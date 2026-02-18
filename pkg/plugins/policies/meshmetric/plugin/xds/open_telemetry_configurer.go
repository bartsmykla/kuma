package xds

import (
	core_xds "github.com/kumahq/kuma/v2/pkg/core/xds"
	envoy_common "github.com/kumahq/kuma/v2/pkg/xds/envoy"
	envoy_clusters "github.com/kumahq/kuma/v2/pkg/xds/envoy/clusters"
	envoy_listeners "github.com/kumahq/kuma/v2/pkg/xds/envoy/listeners"
)

type OpenTelemetryConfigurer struct {
	Endpoint     *core_xds.Endpoint
	ListenerName string
	ClusterName  string
	SocketName   string
	StatPrefix   string
	ApiVersion   core_xds.APIVersion
	IPv6Enabled  bool
}

func (oc *OpenTelemetryConfigurer) ConfigureCluster(isIPv6 bool) (envoy_common.NamedResource, error) {
	builder := envoy_clusters.NewClusterBuilder(oc.ApiVersion, oc.ClusterName)
	// Only force HTTP/2 for gRPC endpoints. HTTP endpoints use default HTTP/1.1.
	if oc.Endpoint.ExternalService == nil {
		builder.Configure(envoy_clusters.Http2())
	}
	return builder.
		Configure(envoy_clusters.ProvidedEndpointCluster(isIPv6, *oc.Endpoint)).
		Configure(envoy_clusters.ClientSideTLS([]core_xds.Endpoint{*oc.Endpoint})).
		Configure(envoy_clusters.DefaultTimeout()).
		Build()
}

func (oc *OpenTelemetryConfigurer) ConfigureListener() (envoy_common.NamedResource, error) {
	filterChainBuilder := envoy_listeners.NewFilterChainBuilder(oc.ApiVersion, envoy_common.AnonymousResource).
		Configure(envoy_listeners.StaticEndpoints(
			oc.IPv6Enabled,
			oc.ListenerName,
			[]*envoy_common.StaticEndpointPath{{
				ClusterName: oc.ClusterName,
				Path:        "/",
			}},
		))
	// Only add gRPC stats filter for gRPC endpoints.
	if oc.Endpoint.ExternalService == nil {
		filterChainBuilder.Configure(envoy_listeners.GrpcStats())
	}
	return envoy_listeners.NewListenerBuilder(oc.ApiVersion, oc.ListenerName).
		Configure(envoy_listeners.PipeListener(oc.SocketName)).
		Configure(envoy_listeners.StatPrefix(oc.StatPrefix)).
		Configure(envoy_listeners.FilterChain(filterChainBuilder)).
		Build()
}
