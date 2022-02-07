package generator

import (
	"github.com/pkg/errors"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/core"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	model "github.com/kumahq/kuma/pkg/core/xds"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	envoy_clusters "github.com/kumahq/kuma/pkg/xds/envoy/clusters"
	envoy_endpoints "github.com/kumahq/kuma/pkg/xds/envoy/endpoints"
	envoy_listeners "github.com/kumahq/kuma/pkg/xds/envoy/listeners"
	envoy_names "github.com/kumahq/kuma/pkg/xds/envoy/names"
	"github.com/kumahq/kuma/pkg/xds/envoy/tls"
)

const (
	EgressProxy = "egress-proxy"

	// OriginEgress is a marker to indicate by which ProxyGenerator resources
	// were generated.
	OriginEgress = "egress"
)

type EgressGenerator struct {
}

func (g EgressGenerator) Generate(
	_ xds_context.Context,
	proxy *model.Proxy,
) (*model.ResourceSet, error) {
	resources := model.NewResourceSet()

	zoneEgressProxy := proxy.ZoneEgressProxy
	zoneEgressResource := zoneEgressProxy.ZoneEgressResource
	apiVersion := proxy.APIVersion

	// destinations := buildDestinations(zoneEgressProxy)

	determineRoutes(proxy)

	listener, err := g.generateLDS(proxy, zoneEgressResource, apiVersion)
	if err != nil {
		return nil, err
	}

	resources.Add(&model.Resource{
		Name:     listener.GetName(),
		Origin:   OriginEgress,
		Resource: listener,
	})

	cdsResources, err := g.generateCDS(proxy, apiVersion)
	if err != nil {
		return nil, err
	}

	resources.AddSet(cdsResources)

	edsResources, err := generateEgressEDS(proxy, apiVersion)
	if err != nil {
		return nil, err
	}

	resources.AddSet(edsResources)

	return resources, nil
}

// generateLDS generates one Egress Listener
// It assumes that mTLS is on. Using TLSInspector we sniff SNI value.
// SNI value has service name and tag values specified with the following
// format: "backend{cluster=2,version=1}"
func (g EgressGenerator) generateLDS(
	proxy *model.Proxy,
	egress *core_mesh.ZoneEgressResource,
	apiVersion envoy_common.APIVersion,
) (envoy_common.NamedResource, error) {
	networking := egress.Spec.GetNetworking()
	address := networking.GetAddress()
	port := networking.GetPort()

	listenerBuilder := envoy_listeners.NewListenerBuilder(apiVersion).
		Configure(envoy_listeners.InboundListener(
			envoy_names.GetInboundListenerName(address, port),
			address, port,
			model.SocketAddressProtocolTCP,
		)).
		Configure(envoy_listeners.TLSInspector())

	zoneEgressProxy := proxy.ZoneEgressProxy
	// externalServices := zoneEgressProxy.ExternalServices
	meshMap := zoneEgressProxy.MeshMap
	destinationMap := zoneEgressProxy.DestinationMap
	meshEndpointMap := zoneEgressProxy.MeshEndpointMap

	// if len(externalServices) == 0 {
	// 	listenerBuilder.Configure(
	// 		envoy_listeners.FilterChain(
	// 			envoy_listeners.NewFilterChainBuilder(apiVersion),
	// 		),
	// 	)
	// }

	for meshName, endpointMap := range meshEndpointMap {
		for serviceName := range endpointMap {
			mesh := meshMap[meshName]
			destination := destinationMap[serviceName]

			xdsContext := xds_context.Context{
				Mesh: xds_context.MeshContext{
					Resource: mesh,
				},
			}

			core.Log.Info("mesh", "mesh", mesh)

			filterChainBuilder := envoy_listeners.NewFilterChainBuilder(apiVersion)

			for _, tagSelector := range destination {
				tags := envoy_common.Tags(tagSelector).
					WithTags("mesh", meshName)

				tagsWithoutService := tags.
					WithoutTags(mesh_proto.ServiceTag)

				sni := tls.SNIFromTags(tags)

				core.Log.Info("sni", "sni", sni)
				core.Log.Info("tags", "tags", tags)
				core.Log.Info("tagsWithoutService", "tagsWithoutService", tagsWithoutService)

				filterChainBuilder.
					Configure(
						envoy_listeners.ServerSideMTLS(xdsContext),
					).
					Configure(
						envoy_listeners.MatchTransportProtocol("tls"),
						envoy_listeners.MatchServerNames(sni),
					).
					Configure(envoy_listeners.TcpProxyWithMetadata(
						meshName,
						envoy_common.NewCluster(
							envoy_common.WithService(serviceName),
							envoy_common.WithTags(tagsWithoutService),
						),
					))
			}

			listenerBuilder.Configure(
				envoy_listeners.FilterChain(filterChainBuilder),
			)
		}
	}

	// for _, es := range externalServices {
	// 	esMeta := es.GetMeta()
	// 	esName := esMeta.GetName()
	// 	esMeshName := esMeta.GetMesh()
	// 	esMesh := meshMap[esMeshName]
	// 	esDestination := destinationMap[esName]
	//
	// 	xdsContext := xds_context.Context{
	// 		Mesh: xds_context.MeshContext{
	// 			Resource: esMesh,
	// 		},
	// 	}
	//
	// 	filterChainBuilder := envoy_listeners.NewFilterChainBuilder(apiVersion)
	//
	// 	// core.Log.Info("destinationMap", "destinationMap", destinationMap)
	// 	// core.Log.Info("esDestination", "esDestination", esDestination)
	//
	// 	for _, tagSelector := range esDestination {
	// 		tags := envoy_common.Tags(tagSelector).
	// 			WithTags("mesh", esMeshName)
	//
	// 		tagsWithoutService := tags.
	// 			WithoutTags(mesh_proto.ServiceTag)
	//
	// 		sni := tls.SNIFromTags(tags)
	//
	// 		// core.Log.Info("tagSelector", "tagSelector", tagSelector)
	// 		// core.Log.Info("tags", "tags", tags)
	// 		// core.Log.Info("tagsWithoutService", "tagsWithoutService", tagsWithoutService)
	// 		// core.Log.Info("sni", "sni", sni)
	// 		// core.Log.Info("empty line")
	// 		// core.Log.Info("empty line")
	// 		// core.Log.Info("empty line")
	// 		// core.Log.Info("empty line")
	//
	// 		filterChainBuilder.
	// 			Configure(
	// 				envoy_listeners.ServerSideMTLS(xdsContext),
	// 			).
	// 			Configure(
	// 				envoy_listeners.MatchTransportProtocol("tls"),
	// 				envoy_listeners.MatchServerNames(sni),
	// 			).
	// 			Configure(envoy_listeners.TcpProxyWithMetadata(
	// 				esMeshName,
	// 				envoy_common.NewCluster(
	// 					envoy_common.WithService(esName),
	// 					envoy_common.WithTags(tagsWithoutService),
	// 				),
	// 			))
	// 	}
	//
	// 	listenerBuilder.Configure(
	// 		envoy_listeners.FilterChain(filterChainBuilder),
	// 	)
	// }

	return listenerBuilder.Build()
}

func (g EgressGenerator) generateCDS(
	proxy *model.Proxy,
	apiVersion envoy_common.APIVersion,
) (*model.ResourceSet, error) {
	resources := model.NewResourceSet()

	egressProxy := proxy.ZoneEgressProxy
	meshEndpointMap := egressProxy.MeshEndpointMap
	meshTLSReadinessMap := egressProxy.TLSReadinessMap
	destinationMap := egressProxy.DestinationMap

	for meshName, endpointMap := range meshEndpointMap {
		tlsReadinessMap := meshTLSReadinessMap[meshName]

		mesh := egressProxy.MeshMap[meshName]

		xdsContext := xds_context.Context{
			Mesh: xds_context.MeshContext{
				Resource: mesh,
			},
		}

		for serviceName, endpoints := range endpointMap {
			isTLSReady := tlsReadinessMap[serviceName]
			destination := destinationMap[serviceName]

			var isExternalService bool
			if len(endpoints) > 0 {
				isExternalService = endpoints[0].IsExternalService()
			}

			if isExternalService {
				cluster, err := envoy_clusters.NewClusterBuilder(apiVersion).
					Configure(envoy_clusters.ProvidedEndpointCluster(
						serviceName,
						false, // TODO (bartsmykla)
						endpoints...,
					)).
					Configure(envoy_clusters.ClientSideTLS(endpoints)).
					Build()
				if err != nil {
					return nil, err
				}

				resources.Add(&model.Resource{
					Name:     serviceName,
					Origin:   OriginEgress,
					Resource: cluster,
				})

				continue
			}

			for _, tagSelector := range destination {
				tags := envoy_common.Tags(tagSelector).
					WithTags("mesh", meshName)

				clusterBuilder := envoy_clusters.NewClusterBuilder(apiVersion).
					Configure(envoy_clusters.EdsCluster(serviceName)).
					Configure(envoy_clusters.ClientSideMTLS(
						xdsContext,
						serviceName,
						isTLSReady,
						[]envoy_common.Tags{tags},
					)).
					Configure(envoy_clusters.Http2())

				cluster, err := clusterBuilder.Build()
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"build CDS for cluster %s failed",
						serviceName,
					)
				}

				resources.Add(&model.Resource{
					Name:     serviceName,
					Origin:   OriginEgress,
					Resource: cluster,
				})
			}
		}
	}

	return resources, nil
}

func generateEgressEDS(proxy *model.Proxy, version envoy_common.APIVersion) (*model.ResourceSet, error) {
	resources := model.NewResourceSet()

	apiVersion := proxy.APIVersion

	for _, endpointMap := range proxy.ZoneEgressProxy.MeshEndpointMap {
		for serviceName, endpoints := range endpointMap {
			var isExternalService bool
			if len(endpoints) > 0 {
				isExternalService = endpoints[0].IsExternalService()
			}

			if isExternalService {
				continue
			}

			cla, err := envoy_endpoints.CreateClusterLoadAssignment(
				serviceName,
				endpoints,
				apiVersion,
			)
			if err != nil {
				return nil, err
			}

			resources.Add(&model.Resource{
				Name:     serviceName,
				Resource: cla,
			})
		}
	}

	return resources, nil
}

func clustersFromSplit(
	splits []*mesh_proto.TrafficRoute_Split,
	endpointMap model.EndpointMap,
	timeoutConf *core_mesh.TimeoutResource,
	loadBalancer *mesh_proto.TrafficRoute_LoadBalancer,
	sc *splitCounter,
) []envoy_common.Cluster {
	var clusters []envoy_common.Cluster

	for _, split := range splits {
		splitDestination := split.GetDestination()
		splitWeight := split.GetWeight().GetValue()
		serviceName := splitDestination[mesh_proto.ServiceTag]
		endpoints := endpointMap[serviceName]

		if splitWeight == 0 {
			// 0 assumes no traffic is passed there. Envoy
			// doesn't support 0 weight, so instead of
			// passing it to Envoy we just skip such cluster.
			continue
		}

		// TODO (bartsmykla): confirm if we can safely skip the cluster creation
		//  when there is no endpoint for the service
		if len(endpoints) == 0 {
			continue
		}

		// We assume that all the targets are either
		// ExternalServices or not therefore we check only
		// the first one
		isExternalService := endpoints[0].IsExternalService()

		clusterName := serviceName
		if len(splitDestination) > 1 {
			clusterName = envoy_names.GetSplitClusterName(
				serviceName,
				sc.getAndIncrement(),
			)
		}

		cluster := envoy_common.NewCluster(
			envoy_common.WithService(serviceName),
			envoy_common.WithName(clusterName),
			envoy_common.WithWeight(splitWeight),
			envoy_common.WithTags(splitDestination),
			envoy_common.WithTimeout(timeoutConf),
			envoy_common.WithLB(loadBalancer),
			envoy_common.WithExternalService(isExternalService),
		)

		clusters = append(clusters, cluster)
	}

	return clusters
}

func determineRoutes(proxy *model.Proxy) envoy_common.Routes {
	zoneEgressProxy := proxy.ZoneEgressProxy
	meshRouting := zoneEgressProxy.MeshRoutingMap
	sc := splitCounter{}

	var routes envoy_common.Routes

	// meshName
	for _, dpRouting := range meshRouting {
		// dataplaneName
		for _, routing := range dpRouting.DataplaneRoutingMap {
			for outboundInterface, trafficRoute := range routing.RouteMap {
				trSpec := trafficRoute.Spec
				trConf := trSpec.GetConf()
				loadBalancer := trConf.GetLoadBalancer()
				timeoutConf := routing.MatchedPolicies.Timeouts[outboundInterface]
				endpointMap := routing.EndpointMap

				for _, http := range trConf.GetHttp() {
					clusters := clustersFromSplit(
						http.GetSplitWithDestination(),
						endpointMap,
						timeoutConf,
						loadBalancer,
						&sc,
					)

					// TODO (bartsmykla): move outside of here (appendRoute func?)
					if len(clusters) > 0 {
						// TODO (bartsmykla): add rate limits
						route := envoy_common.Route{
							Match:     http.GetMatch(),
							Modify:    http.GetModify(),
							RateLimit: nil,
							Clusters:  clusters,
						}

						routes = append(routes, route)
					}
				}

				if defaultDestination := trConf.GetSplitWithDestination(); len(defaultDestination) != 0 {
					clusters := clustersFromSplit(
						defaultDestination,
						endpointMap,
						timeoutConf,
						loadBalancer,
						&sc,
					)

					// TODO (bartsmykla): add rate limits
					route := envoy_common.Route{
						Match:     nil,
						Modify:    nil,
						RateLimit: nil,
						Clusters:  clusters,
					}

					routes = append(routes, route)
				}
			}
		}
	}

	return routes
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

	for _, backend := range backends {
		service := backend.Destination[mesh_proto.ServiceTag]
		destinations[service] = append(destinations[service], backend.Destination)
	}

	return destinations
}

// mesh-1:
//  dp-1:
//   o7777: externalservice-1
//   o8888: externalservice-2
//  dp-2:
//  dp-3:
