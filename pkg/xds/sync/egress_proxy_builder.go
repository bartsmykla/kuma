package sync

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/datasource"
	"github.com/kumahq/kuma/pkg/core/dns/lookup"
	"github.com/kumahq/kuma/pkg/core/permissions"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/registry"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/core/xds"
	xds_cache "github.com/kumahq/kuma/pkg/xds/cache/mesh"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	xds_topology "github.com/kumahq/kuma/pkg/xds/topology"
)

type EgressProxyBuilder struct {
	ctx context.Context

	ResManager         manager.ResourceManager
	ReadOnlyResManager manager.ReadOnlyResourceManager
	LookupIP           lookup.LookupIPFunc
	MetadataTracker    DataplaneMetadataTracker
	meshCache          *xds_cache.Cache
	DataSourceLoader   datasource.Loader

	apiVersion envoy_common.APIVersion
}

func NewEgressProxyBuilder(
	ctx context.Context,
	resManager manager.ResourceManager,
	readOnlyResManager manager.ReadOnlyResourceManager,
	lookupIP lookup.LookupIPFunc,
	metadataTracker DataplaneMetadataTracker,
	meshCache *xds_cache.Cache,
	dataSourceLoader datasource.Loader,
	apiVersion envoy_common.APIVersion,
) *EgressProxyBuilder {
	return &EgressProxyBuilder{
		ctx:                ctx,
		ResManager:         resManager,
		ReadOnlyResManager: readOnlyResManager,
		LookupIP:           lookupIP,
		MetadataTracker:    metadataTracker,
		meshCache:          meshCache,
		DataSourceLoader:   dataSourceLoader,
		apiVersion:         apiVersion,
	}
}

func filterMeshesByMTLS(rs core_model.Resource) bool {
	return rs.(*core_mesh.MeshResource).Spec.GetMtls().GetEnabledBackend() != ""
}

func filterDataplanesByZone(zone string) core_store.ListFilterFunc {
	return func(rs core_model.Resource) bool {
		return rs.(*core_mesh.DataplaneResource).Spec.Matches(map[string]string{
			mesh_proto.ZoneTag: zone,
		})
	}
}

func filterOutZoneIngressesByZone(zone string) core_store.ListFilterFunc {
	return func(rs core_model.Resource) bool {
		return rs.(*core_mesh.ZoneIngressResource).Spec.GetZone() != zone
	}
}

var lockedHeaders []string

func lock(headers ...string) {
	lockedHeaders = append(lockedHeaders, headers...)
}

func printHelper(resource interface{}, header string) {
	isLocked := len(lockedHeaders) == 0
	for _, lockedHeader := range lockedHeaders {
		if header == lockedHeader {
			isLocked = true
		}
	}

	if !isLocked {
		return
	}

	if header != "" {
		_, _ = fmt.Fprintln(os.Stderr, header)
	}

	yamlBytes, _ := yaml.Marshal(resource)

	_, _ = fmt.Fprintln(os.Stderr, " ", strings.ReplaceAll(string(yamlBytes), "\n", "\n  "))
}

func (p *EgressProxyBuilder) Build(key core_model.ResourceKey) (*xds.Proxy, error) {
	ze := core_mesh.NewZoneEgressResource()
	if err := p.ReadOnlyResManager.Get(p.ctx, ze, core_store.GetBy(key)); err != nil {
		return nil, err
	}

	zeZone := ze.Spec.GetZone()

	var meshes core_mesh.MeshResourceList
	if err := p.ReadOnlyResManager.List(
		p.ctx,
		&meshes,
		core_store.ListByFilterFunc(filterMeshesByMTLS),
	); err != nil {
		return nil, err
	}

	// externalServiceMap := map[string]*core_mesh.ExternalServiceResource{}
	meshMap := map[string]*core_mesh.MeshResource{}
	destinationMap := xds.DestinationMap{}
	endpointMap := map[string]xds.EndpointMap{}
	tlsReadinessMap := map[string]map[string]bool{}

	for _, mesh := range meshes.Items {
		meshName := mesh.GetMeta().GetName()

		meshExternalServiceMap := map[string]*core_mesh.ExternalServiceResource{}

		meshCtx, err := p.meshCache.GetMeshContext(p.ctx, syncLog, meshName)
		if err != nil {
			return nil, err
		}

		tlsReadinessMap[meshName] = meshCtx.ServiceTLSReadiness

		meshResources := meshCtx.Resources
		dataplanes := meshResources.Dataplanes().Items
		trafficPermissions := meshResources.TrafficPermissions()
		trafficRoutes := meshResources.TrafficRoutes().Items
		externalServices := meshResources.ExternalServices().Items
		zoneIngresses := meshResources.ZoneIngresses().Items

		core.Log.Info("dataplanes", "dataplanes", dataplanes)
		core.Log.Info("trafficRoutes", "trafficRoutes", trafficRoutes)

		for dpName, dp := range meshCtx.DataplanesByName {
			core.Log.Info("dpName", "dpName", dpName)
			dpRouteMap := xds_topology.BuildRouteMap(dp, trafficRoutes)
			dpDestinationMap := xds_topology.BuildDestinationMap(dp, dpRouteMap)

			core.Log.Info("dpRouteMap", "dpRouteMap", dpRouteMap)
			core.Log.Info("dpDestinationMap", "dpDestinationMap", dpDestinationMap)

			printHelper(dpDestinationMap, "dpDestinationMap")

			var dpExternalServices []*core_mesh.ExternalServiceResource

			// TODO (bartsmykla): optimization - check if necessary
			for _, es := range externalServices {
				esName := es.GetMeta().GetName()

				if _, ok := dpDestinationMap[esName]; ok {
					dpExternalServices = append(dpExternalServices, es)
				}
			}
			// TODO (bartsmykla): end

			for serviceName, tagSelectorSet := range dpDestinationMap {
				if _, ok := destinationMap[serviceName]; !ok {
					destinationMap[serviceName] = xds.TagSelectorSet{}
				}

				for _, tagSelector := range tagSelectorSet {
					destinationMap[serviceName] = destinationMap[serviceName].Add(tagSelector)
				}
			}

			printHelper(destinationMap, "destinationMap")

			allowedExternalServices, err := permissions.MatchExternalServices(
				dp,
				dpExternalServices,
				trafficPermissions,
			)
			if err != nil {
				return nil, err
			}

			core.Log.Info("allowedExternalServices", "allowedExternalServices", allowedExternalServices)

			for _, es := range allowedExternalServices {
				esName := es.GetMeta().GetName()

				// if _, ok := externalServiceMap[esName]; !ok {
				// 	externalServiceMap[esName] = es
				// }

				if _, ok := meshExternalServiceMap[esName]; !ok {
					meshExternalServiceMap[esName] = es
				}

				// At this point we are sure, that this mesh contains at least
				// one dataplane proxy, which should have access to at least one
				// external service, so we should configure egress for this mesh
				// It's unnecessary to store mesh in this map if none of the
				// dataplanes from it are connecting to any external services
				// and waste a memory
				if _, ok := meshMap[meshName]; !ok {
					meshMap[meshName] = mesh
				}
			}

			if _, ok := meshMap[meshName]; !ok {
				meshMap[meshName] = mesh
			}
		}

		var meshExternalServices []*core_mesh.ExternalServiceResource
		for _, es := range meshExternalServiceMap {
			meshExternalServices = append(meshExternalServices, es)
		}

		meshEndpointMap := xds_topology.BuildRemoteEndpointMap(
			mesh,
			zeZone,
			zoneIngresses,
			meshExternalServices,
			p.DataSourceLoader,
		)

		endpointMap[meshName] = meshEndpointMap
	}

	// var finalExternalServices []*core_mesh.ExternalServiceResource
	// for _, es := range externalServiceMap {
	// 	finalExternalServices = append(finalExternalServices, es)
	// }

	printHelper(endpointMap, "endpointMap")
	// printHelper(finalExternalServices, "finalExternalServices")

	proxy := &xds.Proxy{
		Id:         xds.FromResourceKey(key),
		APIVersion: p.apiVersion,
		ZoneEgressProxy: &xds.ZoneEgressProxy{
			ZoneEgressResource: ze,
			DataSourceLoader:   p.DataSourceLoader,
			// ExternalServices:   finalExternalServices,
			MeshMap:         meshMap,
			DestinationMap:  destinationMap,
			MeshEndpointMap: endpointMap,
			TLSReadinessMap: tlsReadinessMap,
		},
		Metadata: p.MetadataTracker.Metadata(key),
	}

	// bs, _ := json.Marshal(proxy)
	// b64bs := base64.StdEncoding.EncodeToString(bs)
	//
	// core.Log.Info("b64bs", "b64bs", b64bs)

	return proxy, nil
}

func (p *EgressProxyBuilder) BuildOld(key core_model.ResourceKey) (*xds.Proxy, error) {
	lock(
		// "routeMap",
		"dpName",
		// "endpointMap",
		"meshName",
		// "policies",
		"destinationMap",
		// "timeouts",
		// "isLocal",
		// "allowedExternalServices",
		"destinations",
	)

	ze := core_mesh.NewZoneEgressResource()
	if err := p.ReadOnlyResManager.Get(p.ctx, ze, core_store.GetBy(key)); err != nil {
		return nil, err
	}

	var meshes core_mesh.MeshResourceList
	if err := p.ReadOnlyResManager.List(
		p.ctx,
		&meshes,
		core_store.ListByFilterFunc(filterMeshesByMTLS),
	); err != nil {
		return nil, err
	}

	meshRouting := xds.MeshRoutingMap{}

	for _, mesh := range meshes.Items {
		meshName := mesh.GetMeta().GetName()
		zeZone := ze.Spec.GetZone()

		printHelper(meshName, "meshName")

		meshRouting[meshName] = &xds.MeshRouting{
			MeshResource:        mesh,
			DataplaneRoutingMap: xds.DataplaneRoutingMap{},
		}

		var dataplanes core_mesh.DataplaneResourceList
		if err := p.ReadOnlyResManager.List(
			p.ctx,
			&dataplanes,
			core_store.ListByMesh(meshName),
			core_store.ListByFilterFunc(filterDataplanesByZone(zeZone)),
		); err != nil {
			return nil, err
		}

		meshCtx, err := p.meshCache.GetMeshContext(p.ctx, syncLog, meshName)
		if err != nil {
			message := fmt.Sprintf("getting mesh context for %+q failed", meshName)

			return nil, errors.Wrap(err, message)
		}

		meshResources := meshCtx.Resources

		printHelper(meshResources.Timeouts().Items, "timeouts")

		externalServices := meshResources.ExternalServices()
		trafficPermissions := meshResources.TrafficPermissions()
		trafficRoutes := meshResources.TrafficRoutes().Items
		zoneIngresses := meshResources.ZoneIngresses().Items

		var gatewayRoutes []*core_mesh.GatewayRouteResource
		if _, err := registry.Global().DescriptorFor(core_mesh.GatewayRouteType); err == nil {
			gatewayRoutes = meshResources.GatewayRoutes().Items
		}

		destinations := buildDestinations(trafficRoutes, gatewayRoutes)
		printHelper(destinations, "destinations")

		for _, dp := range dataplanes.Items {
			dpName := dp.GetMeta().GetName()

			printHelper(dpName, "dpName")
			printHelper(trafficPermissions, "trafficPermissions")

			allowedExternalServices, err := permissions.MatchExternalServices(
				dp,
				externalServices.Items,
				trafficPermissions,
			)
			if err != nil {
				return nil, err
			}

			printHelper(allowedExternalServices, "allowedExternalServices")

			routeMap := xds_topology.BuildRouteMap(dp, trafficRoutes)

			printHelper(routeMap, "routeMap")

			// TODO (bartsmykla)
			endpointMap := xds_topology.BuildRemoteEndpointMap(
				mesh,
				zeZone,
				zoneIngresses,
				allowedExternalServices,
				p.DataSourceLoader,
			)

			if len(routeMap) > 0 {
				printHelper(endpointMap, "endpointMap")
			}

			// TODO (bartsmykla): check what for are used destination maps,
			//  and how they are being generated
			destinationMap := xds_topology.BuildDestinationMap(dp, routeMap)
			if len(destinationMap) > 0 {
				printHelper(destinationMap, "destinationMap")
			}

			matchedPolicies, err := matchPolicies(meshCtx, dp, destinationMap)
			if err != nil {
				return nil, err
			}
			printHelper(matchedPolicies, "matchedPolicies")

			meshRouting[meshName].DataplaneRoutingMap[dpName] = &xds.DataplaneRouting{
				IsIPV6:          dp.IsIPv6(),
				RouteMap:        routeMap,
				EndpointMap:     endpointMap,
				MatchedPolicies: matchedPolicies,
				GatewayRoutes:   gatewayRoutes,
				DestinationMap:  destinationMap,
			}
		}
	}

	proxy := &xds.Proxy{
		Id:         xds.FromResourceKey(key),
		APIVersion: p.apiVersion,
		ZoneEgressProxy: &xds.ZoneEgressProxy{
			ZoneEgressResource: ze,
			DataSourceLoader:   p.DataSourceLoader,
			MeshRoutingMap:     meshRouting,
		},
		Metadata: p.MetadataTracker.Metadata(key),
	}

	return proxy, nil
}

func buildDestinations(
	trafficRoutes []*core_mesh.TrafficRouteResource,
	gatewayRoutes []*core_mesh.GatewayRouteResource,
) map[string][]envoy_common.Tags {
	// _, _ = fmt.Fprintln(os.Stderr, ".")
	destinations := map[string][]envoy_common.Tags{}

	for _, route := range trafficRoutes {
		for _, split := range route.Spec.Conf.GetSplitWithDestination() {
			service := split.Destination[mesh_proto.ServiceTag]
			destinations[service] = append(destinations[service], split.Destination)
		}

		for _, http := range route.Spec.Conf.Http {
			for _, split := range http.GetSplitWithDestination() {
				service := split.Destination[mesh_proto.ServiceTag]
				destinations[service] = append(destinations[service], split.Destination)
			}
		}
	}

	var backends []*mesh_proto.GatewayRoute_Backend

	for _, route := range gatewayRoutes {
		for _, rule := range route.Spec.GetConf().GetHttp().GetRules() {
			backends = append(backends, rule.Backends...)
		}
		for _, rule := range route.Spec.GetConf().GetTcp().GetRules() {
			backends = append(backends, rule.Backends...)
		}
		for _, rule := range route.Spec.GetConf().GetTls().GetRules() {
			backends = append(backends, rule.Backends...)
		}
		for _, rule := range route.Spec.GetConf().GetUdp().GetRules() {
			backends = append(backends, rule.Backends...)
		}
	}

	printHelper(backends, "destinations")

	for _, backend := range backends {
		serviceName := backend.Destination[mesh_proto.ServiceTag]
		destinations[serviceName] = append(destinations[serviceName], backend.Destination)
	}

	return destinations
}
