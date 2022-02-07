package generator_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/core"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
	. "github.com/kumahq/kuma/pkg/test/matchers"
	test_model "github.com/kumahq/kuma/pkg/test/resources/model"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/generator"
)

func init() {
	core.SetLogger(core.NewLogger(2))
}

var _ = FDescribe("EgressGenerator", func() {
	type testCase struct {
		expected      string
		trafficRoutes []*core_mesh.TrafficRouteResource
	}

	DescribeTable("should generate Envoy xDS resources",
		func(given testCase) {
			gen := generator.EgressGenerator{}

			zoneEgressResource := &core_mesh.ZoneEgressResource{
				Meta: &test_model.ResourceMeta{
					Name: "zoneegress-1",
				},
				Spec: &mesh_proto.ZoneEgress{
					Zone: "zone-1",
					Networking: &mesh_proto.ZoneEgress_Networking{
						Address: "10.0.0.1",
						Port:    10002,
					},
				},
			}

			mesh1WithMTLSName := "mesh-1"
			mesh1WithMTLS := &core_mesh.MeshResource{
				Meta: &test_model.ResourceMeta{
					Name: mesh1WithMTLSName,
				},
				Spec: &mesh_proto.Mesh{
					Mtls: &mesh_proto.Mesh_Mtls{
						EnabledBackend: "builtin-1",
						Backends: []*mesh_proto.CertificateAuthorityBackend{
							{
								Name: "builtin-1",
								Type: "builtin",
							},
						},
					},
				},
			}

			// mesh2WithoutMTLSName := "mesh-2-without-mtls"
			// mesh2WithoutMTLS := &core_mesh.MeshResource{
			// 	Meta: &test_model.ResourceMeta{
			// 		Name: mesh2WithoutMTLSName,
			// 	},
			// }

			externalService1Name := "externalservice-1"
			externalService2Name := "externalservice-2"
			externalServices := []*core_mesh.ExternalServiceResource{
				{
					Meta: &test_model.ResourceMeta{
						Name: externalService1Name,
						Mesh: mesh1WithMTLSName,
					},
					Spec: &mesh_proto.ExternalService{
						Networking: &mesh_proto.ExternalService_Networking{
							Address: "konghq.com:80",
						},
						Tags: map[string]string{
							mesh_proto.ServiceTag:  externalService1Name,
							mesh_proto.ProtocolTag: "http",
						},
					},
				},
				{
					Meta: &test_model.ResourceMeta{
						Name: externalService2Name,
						Mesh: mesh1WithMTLSName,
					},
					Spec: &mesh_proto.ExternalService{
						Networking: &mesh_proto.ExternalService_Networking{
							Address: "kuma.io:443",
						},
						Tags: map[string]string{
							mesh_proto.ServiceTag: externalService2Name,
						},
					},
				},
			}

			dataplane1Name := "dp-1"
			someOtherService1Name := "some-other-service"

			gatewayRoutes := []*core_mesh.GatewayRouteResource{
				{
					Meta: &test_model.ResourceMeta{
						Mesh: mesh1WithMTLSName,
						Name: "gatewayroute-1",
					},
					Spec: &mesh_proto.GatewayRoute{
						Selectors: []*mesh_proto.Selector{
							{
								Match: mesh_proto.MatchAnyService(),
							},
						},
						Conf: &mesh_proto.GatewayRoute_Conf{
							Route: &mesh_proto.GatewayRoute_Conf_Http{
								Http: &mesh_proto.GatewayRoute_HttpRoute{
									Hostnames: []string{
										"kuma.io",
									},
									Rules: []*mesh_proto.GatewayRoute_HttpRoute_Rule{
										{
											Backends: []*mesh_proto.GatewayRoute_Backend{
												{
													Weight: 1,
													Destination: map[string]string{
														mesh_proto.ServiceTag: "some-service",
													},
												},
											},
											Matches: []*mesh_proto.GatewayRoute_HttpRoute_Match{
												{
													Method: mesh_proto.HttpMethod_GET,
												},
											},
											Filters: []*mesh_proto.GatewayRoute_HttpRoute_Filter{
												{
													Filter: &mesh_proto.GatewayRoute_HttpRoute_Filter_Redirect_{
														Redirect: &mesh_proto.GatewayRoute_HttpRoute_Filter_Redirect{
															Port: 8080,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}

			proxy := &core_xds.Proxy{
				Id:         *core_xds.BuildProxyId("default", "egress"),
				APIVersion: envoy_common.APIV3,
				ZoneEgressProxy: &core_xds.ZoneEgressProxy{
					ZoneEgressResource: zoneEgressResource,
					MeshRoutingMap: core_xds.MeshRoutingMap{
						mesh1WithMTLSName: &core_xds.MeshRouting{
							DataplaneRoutingMap: core_xds.DataplaneRoutingMap{
								dataplane1Name: &core_xds.DataplaneRouting{
									DestinationMap: core_xds.DestinationMap{
										someOtherService1Name: core_xds.TagSelectorSet{
											map[string]string{
												mesh_proto.ServiceTag: someOtherService1Name,
												mesh_proto.ZoneTag:    "zone-1",
											},
											map[string]string{
												mesh_proto.ServiceTag: someOtherService1Name,
												mesh_proto.ZoneTag:    "zone-2",
											},
										},
										externalService1Name: core_xds.TagSelectorSet{
											map[string]string{
												mesh_proto.ServiceTag: externalService1Name,
											},
										},
									},
									GatewayRoutes: gatewayRoutes,
									IsIPV6:        false,
									RouteMap: core_xds.RouteMap{
										mesh_proto.OutboundInterface{
											DataplaneIP:   "127.0.0.1",
											DataplanePort: 18081,
										}: &core_mesh.TrafficRouteResource{
											Spec: &mesh_proto.TrafficRoute{
												Conf: &mesh_proto.TrafficRoute_Conf{
													Destination: mesh_proto.MatchService(externalService1Name),
												},
											},
										},
										mesh_proto.OutboundInterface{
											DataplaneIP:   "127.0.0.1",
											DataplanePort: 18082,
										}: &core_mesh.TrafficRouteResource{
											Spec: &mesh_proto.TrafficRoute{
												Conf: &mesh_proto.TrafficRoute_Conf{
													Split: []*mesh_proto.TrafficRoute_Split{
														{
															Weight: util_proto.UInt32(1),
															Destination: mesh_proto.TagSelector{
																mesh_proto.ServiceTag: someOtherService1Name,
																mesh_proto.ZoneTag:    "zone-1",
															},
														},
														{
															Weight: util_proto.UInt32(10),
															Destination: mesh_proto.TagSelector{
																mesh_proto.ServiceTag: someOtherService1Name,
																mesh_proto.ZoneTag:    "zone-2",
															},
														},
													},
												},
											},
										},
									},
									EndpointMap: core_xds.EndpointMap{
										externalService1Name: []core_xds.Endpoint{
											{
												Target: "konghq.com",
												Port:   80,
												Tags: map[string]string{
													mesh_proto.ServiceTag: externalService1Name,
												},
												Weight:          1,
												ExternalService: &core_xds.ExternalService{},
											},
										},
										someOtherService1Name: []core_xds.Endpoint{
											{
												Target: "10.0.0.254",
												Port:   10001,
												Tags: map[string]string{
													mesh_proto.ServiceTag: someOtherService1Name,
												},
												Weight: 1,
											},
											{
												Target: "10.0.254.254",
												Port:   10001,
												Tags: map[string]string{
													mesh_proto.ServiceTag: someOtherService1Name,
												},
												Weight: 3,
											},
										},
									},
									MatchedPolicies: &core_xds.MatchedPolicies{
										Timeouts: core_xds.TimeoutMap{},
									},
								},
							},
							MeshResource: mesh1WithMTLS,
						},
					},
					MeshMap: map[string]*core_mesh.MeshResource{
						mesh1WithMTLSName: mesh1WithMTLS,
					},
					ExternalServices: externalServices,
					MeshEndpointMap: map[string]core_xds.EndpointMap{
						mesh1WithMTLSName: {
							externalService1Name: []core_xds.Endpoint{
								{
									Target: "konghq.com",
									Port:   80,
									Tags: map[string]string{
										mesh_proto.ServiceTag: externalService1Name,
									},
									Weight:          1,
									ExternalService: &core_xds.ExternalService{},
								},
							},
							externalService2Name: []core_xds.Endpoint{
								{
									Target: "kuma.io",
									Port:   443,
									Tags: map[string]string{
										mesh_proto.ServiceTag: externalService2Name,
									},
									Weight:          1,
									ExternalService: &core_xds.ExternalService{},
								},
							},
							someOtherService1Name: []core_xds.Endpoint{
								{
									Target: "10.0.0.254",
									Port:   10001,
									Tags: map[string]string{
										mesh_proto.ServiceTag: someOtherService1Name,
									},
									Weight: 1,
								},
								{
									Target: "10.0.254.254",
									Port:   10001,
									Tags: map[string]string{
										mesh_proto.ServiceTag: someOtherService1Name,
									},
									Weight: 3,
								},
							},
						},
					},
					DestinationMap: core_xds.DestinationMap{
						externalService1Name: core_xds.TagSelectorSet{
							map[string]string{
								mesh_proto.ServiceTag: externalService1Name,
							},
						},
						externalService2Name: core_xds.TagSelectorSet{
							map[string]string{
								mesh_proto.ServiceTag: externalService2Name,
							},
						},
						someOtherService1Name: core_xds.TagSelectorSet{
							map[string]string{
								mesh_proto.ServiceTag: someOtherService1Name,
							},
						},
					},
				},
			}

			// when
			rs, err := gen.Generate(xds_context.Context{}, proxy)
			// then
			Expect(err).ToNot(HaveOccurred())

			// when
			resp, err := rs.List().ToDeltaDiscoveryResponse()
			// then
			Expect(err).ToNot(HaveOccurred())
			// when
			actual, err := util_proto.ToYAML(resp)
			// then
			Expect(err).ToNot(HaveOccurred())

			// and output matches golden files
			Expect(actual).To(MatchGoldenYAML(filepath.Join("testdata", "egress", given.expected)))
		},
		Entry("01. default trafficroute, single mesh", testCase{
			expected: "01.envoy.golden.yaml",
			trafficRoutes: []*core_mesh.TrafficRouteResource{
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
			},
		}),
		Entry("02. default trafficroute, empty egress", testCase{
			expected: "02.envoy.golden.yaml",
			trafficRoutes: []*core_mesh.TrafficRouteResource{
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
			},
		}),
		Entry("03. trafficroute with many destinations", testCase{
			expected: "03.envoy.golden.yaml",
			trafficRoutes: []*core_mesh.TrafficRouteResource{
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Split: []*mesh_proto.TrafficRoute_Split{
								{
									Weight: util_proto.UInt32(10),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "backend",
										"version":             "v2",
									},
								},
								{
									Weight: util_proto.UInt32(90),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "backend",
										"region":              "eu",
									},
								},
							},
						},
					},
				},
			},
		}),
		Entry("04. several services in several meshes", testCase{
			expected: "04.envoy.golden.yaml",
			trafficRoutes: []*core_mesh.TrafficRouteResource{
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: map[string]string{
								mesh_proto.ServiceTag: "*",
								"version":             "v2",
							},
						},
					},
				},
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Split: []*mesh_proto.TrafficRoute_Split{
								{
									Weight: util_proto.UInt32(10),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "backend",
										"version":             "v2",
									},
								},
								{
									Weight: util_proto.UInt32(90),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "backend",
										"region":              "eu",
									},
								},
							},
						},
					},
				},
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Split: []*mesh_proto.TrafficRoute_Split{
								{
									Weight: util_proto.UInt32(10),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "frontend",
										"region":              "eu",
										"cloud":               "gke",
									},
								},
								{
									Weight: util_proto.UInt32(90),
									Destination: map[string]string{
										mesh_proto.ServiceTag: "frontend",
										"cloud":               "aks",
									},
								},
							},
						},
					},
				},
			},
		}),
		Entry("05. trafficroute repeated", testCase{
			expected: "05.envoy.golden.yaml",
			trafficRoutes: []*core_mesh.TrafficRouteResource{
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
				{
					Spec: &mesh_proto.TrafficRoute{
						Sources: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchService("foo"),
						}},
						Destinations: []*mesh_proto.Selector{{
							Match: mesh_proto.MatchAnyService(),
						}},
						Conf: &mesh_proto.TrafficRoute_Conf{
							Destination: mesh_proto.MatchAnyService(),
						},
					},
				},
			},
		}),
	)
})
