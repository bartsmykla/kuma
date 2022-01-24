package sync

import (
	"context"

	"github.com/kumahq/kuma/pkg/core/dns/lookup"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/core/xds"
	"github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/ingress"
	xds_topology "github.com/kumahq/kuma/pkg/xds/topology"
)

type EgressProxyBuilder struct {
	ResManager         manager.ResourceManager
	ReadOnlyResManager manager.ReadOnlyResourceManager
	LookupIP           lookup.LookupIPFunc
	MetadataTracker    DataplaneMetadataTracker

	apiVersion envoy.APIVersion
}

func (p *EgressProxyBuilder) build(key core_model.ResourceKey) (*xds.Proxy, error) {
	ctx := context.Background()

	zoneEgress, err := p.getZoneEgress(key)
	if err != nil {
		return nil, err
	}
	zoneEgress, err = xds_topology.ResolveZoneEgressPublicAddress(p.LookupIP, zoneEgress)
	if err != nil {
		return nil, err
	}

	allMeshDataplanes := &core_mesh.DataplaneResourceList{}
	if err := p.ReadOnlyResManager.List(ctx, allMeshDataplanes); err != nil {
		return nil, err
	}
	allMeshDataplanes.Items = xds_topology.ResolveAddresses(syncLog, p.LookupIP, allMeshDataplanes.Items)

	routing, err := p.resolveRouting(ctx, zoneEgress, allMeshDataplanes)
	if err != nil {
		return nil, err
	}

	proxy := &xds.Proxy{
		Id:         xds.FromResourceKey(key),
		APIVersion: p.apiVersion,
		ZoneEgress: zoneEgress,
		Metadata:   p.MetadataTracker.Metadata(key),
		Routing:    *routing,
	}
	return proxy, nil
}

func (p *EgressProxyBuilder) getZoneEgress(key core_model.ResourceKey) (*core_mesh.ZoneEgressResource, error) {
	ctx := context.Background()

	zoneEgress := core_mesh.NewZoneEgressResource()
	if err := p.ReadOnlyResManager.Get(ctx, zoneEgress, core_store.GetBy(key)); err != nil {
		return nil, err
	}
	// Update Egress' Available Services
	// This was placed as an operation of DataplaneWatchdog out of the convenience.
	// Consider moving to the outside of this component (follow the pattern of updating VIP outbounds)
	if err := p.updateEgress(zoneEgress); err != nil {
		return nil, err
	}
	return zoneEgress, nil
}

func (p *EgressProxyBuilder) resolveRouting(ctx context.Context, zoneEgress *core_mesh.ZoneEgressResource, dataplanes *core_mesh.DataplaneResourceList) (*xds.Routing, error) {
	destinations := ingress.BuildDestinationMap(zoneEgress)
	endpoints := ingress.BuildEndpointMap(destinations, dataplanes.Items)
	routes := &core_mesh.TrafficRouteResourceList{}
	if err := p.ReadOnlyResManager.List(ctx, routes); err != nil {
		return nil, err
	}

	routing := &xds.Routing{
		OutboundTargets:  endpoints,
		TrafficRouteList: routes,
	}
	return routing, nil
}

func (p *EgressProxyBuilder) updateEgress(zoneEgress *core_mesh.ZoneEgressResource) error {
	ctx := context.Background()

	allMeshDataplanes := &core_mesh.DataplaneResourceList{}
	if err := p.ReadOnlyResManager.List(ctx, allMeshDataplanes); err != nil {
		return err
	}
	allMeshDataplanes.Items = xds_topology.ResolveAddresses(syncLog, p.LookupIP, allMeshDataplanes.Items)

	// Update Egress' Available Services
	// This was placed as an operation of DataplaneWatchdog out of the convenience.
	// Consider moving to the outside of this component (follow the pattern of updating VIP outbounds)
	return ingress.UpdateAvailableServices(ctx, p.ResManager, zoneEgress, allMeshDataplanes.Items)
}
