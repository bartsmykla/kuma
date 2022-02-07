package sync

import (
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/datasource"
	"github.com/kumahq/kuma/pkg/core/permissions"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/core/xds"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	"github.com/kumahq/kuma/pkg/xds/envoy"
	xds_topology "github.com/kumahq/kuma/pkg/xds/topology"
)

var syncLog = core.Log.WithName("sync")

type DataplaneProxyBuilder struct {
	DataSourceLoader datasource.Loader
	MetadataTracker  DataplaneMetadataTracker

	Zone       string
	APIVersion envoy.APIVersion
}

func (p *DataplaneProxyBuilder) Build(key core_model.ResourceKey, meshContext xds_context.MeshContext) (*xds.Proxy, error) {
	dp, found := meshContext.DataplanesByName[key.Name]
	if !found {
		return nil, core_store.ErrorResourceNotFound(core_mesh.DataplaneType, key.Name, key.Mesh)
	}

	routing, destinations, err := p.resolveRouting(meshContext, dp)
	if err != nil {
		return nil, err
	}

	matchedPolicies, err := matchPolicies(meshContext, dp, destinations)
	if err != nil {
		return nil, err
	}

	matchedPolicies.TrafficRoutes = routing.TrafficRoutes

	proxy := &xds.Proxy{
		Id:         xds.FromResourceKey(key),
		APIVersion: p.APIVersion,
		Dataplane:  dp,
		Metadata:   p.MetadataTracker.Metadata(key),
		Routing:    *routing,
		Policies:   *matchedPolicies,
	}
	return proxy, nil
}

func (p *DataplaneProxyBuilder) resolveRouting(
	meshContext xds_context.MeshContext,
	dataplane *core_mesh.DataplaneResource,
) (*xds.Routing, xds.DestinationMap, error) {
	resources := meshContext.Resources

	matchedExternalServices, err := permissions.MatchExternalServices(
		dataplane,
		resources.ExternalServices().Items,
		resources.TrafficPermissions(),
	)
	if err != nil {
		return nil, nil, err
	}

	p.resolveVIPOutbounds(meshContext, dataplane)

	trafficRoutes := resources.TrafficRoutes().Items
	dataplanes := resources.Dataplanes().Items
	zoneIngresses := resources.ZoneIngresses().Items
	zoneEgresses := resources.ZoneEgresses().Items
	mesh := meshContext.Resource

	// pick a single the most specific route for each outbound interface
	routeMap := xds_topology.BuildRouteMap(dataplane, trafficRoutes)

	// create a map of selectors to match other dataplanes reachable via given
	// routes
	destinations := xds_topology.BuildDestinationMap(dataplane, routeMap)

	// resolve all endpoints that match given selectors
	endpointMap := xds_topology.BuildEndpointMap(
		mesh,
		p.Zone,
		dataplanes,
		zoneIngresses,
		zoneEgresses,
		matchedExternalServices,
		p.DataSourceLoader,
	)

	routing := &xds.Routing{
		TrafficRoutes:   routeMap,
		OutboundTargets: endpointMap,
	}

	return routing, destinations, nil
}

func (p *DataplaneProxyBuilder) resolveVIPOutbounds(meshContext xds_context.MeshContext, dataplane *core_mesh.DataplaneResource) {
	if dataplane.Spec.Networking.GetTransparentProxying() == nil {
		return
	}
	// Update the outbound of the dataplane with the generatedVips
	generatedVips := map[string]bool{}
	for _, ob := range meshContext.VIPOutbounds {
		generatedVips[ob.Address] = true
	}
	outbounds := meshContext.VIPOutbounds
	for _, outbound := range dataplane.Spec.Networking.GetOutbound() {
		if generatedVips[outbound.Address] { // Useful while we still have resources with computed vip outbounds
			continue
		}
		outbounds = append(outbounds, outbound)
	}
	dataplane.Spec.Networking.Outbound = outbounds
}
