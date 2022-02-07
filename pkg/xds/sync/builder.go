package sync

import (
	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/core/faultinjections"
	"github.com/kumahq/kuma/pkg/core/logs"
	manager_dataplane "github.com/kumahq/kuma/pkg/core/managers/apis/dataplane"
	"github.com/kumahq/kuma/pkg/core/permissions"
	"github.com/kumahq/kuma/pkg/core/ratelimits"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/xds"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	xds_topology "github.com/kumahq/kuma/pkg/xds/topology"
)

// TODO (bartsmykla)
func matchEgressRelevantPolicies() {}

func matchPolicies(
	meshContext xds_context.MeshContext,
	dp *core_mesh.DataplaneResource,
	outboundSelectors xds.DestinationMap,
) (*xds.MatchedPolicies, error) {
	mesh := meshContext.Resource
	resources := meshContext.Resources

	additionalInbounds, err := manager_dataplane.AdditionalInbounds(dp, mesh)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"could not fetch additional inbounds",
		)
	}

	dpInbound := dp.Spec.GetNetworking().GetInbound()
	inbounds := append(dpInbound, additionalInbounds...)

	rateLimits := resources.RateLimits().Items
	rateLimitMap := ratelimits.BuildRateLimitMap(dp, inbounds, rateLimits)
	trafficPermissions := resources.TrafficPermissions().Items
	trafficLogs := resources.TrafficLogs().Items
	healthChecks := resources.HealthChecks().Items
	circuitBreakers := resources.CircuitBreakers().Items
	trafficTraces := resources.TrafficTraces().Items
	faultInjections := resources.FaultInjections().Items
	retries := resources.Retries().Items
	timeouts := resources.Timeouts().Items

	matchedPolicies := &xds.MatchedPolicies{
		TrafficPermissions: permissions.BuildTrafficPermissionMap(dp, inbounds, trafficPermissions),
		TrafficLogs:        logs.BuildTrafficLogMap(dp, trafficLogs),
		HealthChecks:       xds_topology.BuildHealthCheckMap(dp, outboundSelectors, healthChecks),
		CircuitBreakers:    xds_topology.BuildCircuitBreakerMap(dp, outboundSelectors, circuitBreakers),
		TrafficTrace:       xds_topology.SelectTrafficTrace(dp, trafficTraces),
		FaultInjections:    faultinjections.BuildFaultInjectionMap(dp, inbounds, faultInjections),
		Retries:            xds_topology.BuildRetryMap(dp, retries, outboundSelectors),
		Timeouts:           xds_topology.BuildTimeoutMap(dp, timeouts),
		RateLimitsInbound:  rateLimitMap.Inbound,
		RateLimitsOutbound: rateLimitMap.Outbound,
	}

	return matchedPolicies, nil
}
