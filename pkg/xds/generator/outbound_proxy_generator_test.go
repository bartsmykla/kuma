package generator_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	mesh_core "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	model "github.com/kumahq/kuma/pkg/core/xds"
	. "github.com/kumahq/kuma/pkg/test/matchers"
	test_model "github.com/kumahq/kuma/pkg/test/resources/model"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/generator"
)

var _ = Describe("OutboundProxyGenerator", func() {

	meta := &test_model.ResourceMeta{
		Name: "mesh1",
	}
	plainCtx := xds_context.Context{
		ControlPlane: &xds_context.ControlPlaneContext{},
		Mesh: xds_context.MeshContext{
			Resource: &mesh_core.MeshResource{
				Meta: meta,
				Spec: &mesh_proto.Mesh{},
			},
		},
	}

	mtlsCtx := xds_context.Context{
		ConnectionInfo: xds_context.ConnectionInfo{
			Authority: "kuma-system:5677",
		},
		ControlPlane: &xds_context.ControlPlaneContext{
			SdsTlsCert: []byte("12345"),
		},
		Mesh: xds_context.MeshContext{
			Resource: &mesh_core.MeshResource{
				Spec: &mesh_proto.Mesh{
					Mtls: &mesh_proto.Mesh_Mtls{
						EnabledBackend: "builtin",
						Backends: []*mesh_proto.CertificateAuthorityBackend{
							{
								Name: "builtin",
								Type: "builtin",
							},
						},
					},
				},
				Meta: meta,
			},
		},
	}

	type testCase struct {
		ctx       xds_context.Context
		dataplane string
		expected  string
	}

	DescribeTable("Generate Envoy xDS resources",
		func(given testCase) {
			// setup
			gen := &generator.OutboundProxyGenerator{}

			dataplane := &mesh_proto.Dataplane{}
			Expect(util_proto.FromYAML([]byte(given.dataplane), dataplane)).To(Succeed())

			outboundTargets := model.EndpointMap{
				"api-http": []model.Endpoint{ // notice that all endpoints have tag `kuma.io/protocol: http`
					{
						Target: "192.168.0.4",
						Port:   8084,
						Tags:   map[string]string{"kuma.io/service": "api-http", "kuma.io/protocol": "http", "region": "us"},
						Weight: 1,
					},
					{
						Target: "192.168.0.5",
						Port:   8085,
						Tags:   map[string]string{"kuma.io/service": "api-http", "kuma.io/protocol": "http", "region": "eu"},
						Weight: 1,
					},
				},
				"api-tcp": []model.Endpoint{ // notice that not every endpoint has a `kuma.io/protocol: http` tag
					{
						Target: "192.168.0.6",
						Port:   8086,
						Tags:   map[string]string{"kuma.io/service": "api-tcp", "kuma.io/protocol": "http", "region": "us"},
						Weight: 1,
					},
					{
						Target: "192.168.0.7",
						Port:   8087,
						Tags:   map[string]string{"kuma.io/service": "api-tcp", "region": "eu"},
						Weight: 1,
					},
				},
				"api-http2": []model.Endpoint{ // notice that all endpoints have tag `kuma.io/protocol: http2`
					{
						Target: "192.168.0.4",
						Port:   8088,
						Tags:   map[string]string{"kuma.io/service": "api-http2", "kuma.io/protocol": "http2"},
						Weight: 1,
					},
				},
				"api-grpc": []model.Endpoint{ // notice that all endpoints have tag `kuma.io/protocol: grpc`
					{
						Target: "192.168.0.4",
						Port:   8089,
						Tags:   map[string]string{"kuma.io/service": "api-grpc", "kuma.io/protocol": "grpc"},
						Weight: 1,
					},
				},
				"backend": []model.Endpoint{ // notice that not every endpoint has a tag `kuma.io/protocol: http`
					{
						Target: "192.168.0.1",
						Port:   8081,
						Tags:   map[string]string{"kuma.io/service": "backend", "region": "us"},
						Weight: 1,
					},
					{
						Target: "192.168.0.2",
						Port:   8082,
						Weight: 1,
					},
				},
				"db": []model.Endpoint{
					{
						Target: "192.168.0.3",
						Port:   5432,
						Tags:   map[string]string{"kuma.io/service": "db", "role": "master"},
						Weight: 1,
					},
				},
				"es": []model.Endpoint{
					{
						Target:          "10.0.0.1",
						Port:            10001,
						Tags:            map[string]string{"kuma.io/service": "es", "kuma.io/protocol": "http"},
						Weight:          1,
						ExternalService: &model.ExternalService{TLSEnabled: false},
					},
				},
				"es2": []model.Endpoint{
					{
						Target:          "10.0.0.2",
						Port:            10002,
						Tags:            map[string]string{"kuma.io/service": "es2", "kuma.io/protocol": "http2"},
						Weight:          1,
						ExternalService: &model.ExternalService{TLSEnabled: false},
					},
				},
			}
			proxy := &model.Proxy{
				Id: model.ProxyId{Name: "side-car", Mesh: "default"},
				Dataplane: &mesh_core.DataplaneResource{
					Meta: &test_model.ResourceMeta{
						Name:    "dp-1",
						Mesh:    given.ctx.Mesh.Resource.Meta.GetName(),
						Version: "1",
					},
					Spec: dataplane,
				},
				APIVersion: envoy_common.APIV3,
				Routing: model.Routing{
					TrafficRoutes: model.RouteMap{
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 40001,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      100,
										Destination: mesh_proto.MatchService("api-http"),
									}},
									LoadBalancer: &mesh_proto.TrafficRoute_LoadBalancer{
										LbType: &mesh_proto.TrafficRoute_LoadBalancer_RoundRobin_{},
									},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 40002,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      100,
										Destination: mesh_proto.MatchService("api-tcp"),
									}},
									LoadBalancer: &mesh_proto.TrafficRoute_LoadBalancer{
										LbType: &mesh_proto.TrafficRoute_LoadBalancer_LeastRequest_{
											LeastRequest: &mesh_proto.TrafficRoute_LoadBalancer_LeastRequest{
												ChoiceCount: 4,
											},
										},
									},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 40003,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      100,
										Destination: mesh_proto.MatchService("api-http2"),
									}},
									LoadBalancer: &mesh_proto.TrafficRoute_LoadBalancer{
										LbType: &mesh_proto.TrafficRoute_LoadBalancer_RingHash_{
											RingHash: &mesh_proto.TrafficRoute_LoadBalancer_RingHash{
												HashFunction: "MURMUR_HASH_2",
												MinRingSize:  64,
												MaxRingSize:  1024,
											},
										},
									},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 40004,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      100,
										Destination: mesh_proto.MatchService("api-grpc"),
									}},
									LoadBalancer: &mesh_proto.TrafficRoute_LoadBalancer{
										LbType: &mesh_proto.TrafficRoute_LoadBalancer_Random_{},
									},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 18080,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      100,
										Destination: mesh_proto.MatchService("backend"),
									}},
									LoadBalancer: &mesh_proto.TrafficRoute_LoadBalancer{
										LbType: &mesh_proto.TrafficRoute_LoadBalancer_Maglev_{},
									},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 54321,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      10,
										Destination: mesh_proto.TagSelector{"kuma.io/service": "db", "role": "master"},
									}, {
										Weight:      90,
										Destination: mesh_proto.TagSelector{"kuma.io/service": "db", "role": "replica"},
									}, {
										Weight:      0, // should be excluded from Envoy configuration
										Destination: mesh_proto.TagSelector{"kuma.io/service": "db", "role": "canary"},
									}},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 18081,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      1,
										Destination: mesh_proto.TagSelector{"kuma.io/service": "es", "kuma.io/protocol": "http"},
									}},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 18082,
						}: &mesh_core.TrafficRouteResource{
							Spec: &mesh_proto.TrafficRoute{
								Conf: &mesh_proto.TrafficRoute_Conf{
									Split: []*mesh_proto.TrafficRoute_Split{{
										Weight:      1,
										Destination: mesh_proto.TagSelector{"kuma.io/service": "es2", "kuma.io/protocol": "http2"},
									}},
								},
							},
						},
						mesh_proto.OutboundInterface{
							DataplaneIP:   "127.0.0.1",
							DataplanePort: 4040,
						}: nil,
					},
					OutboundTargets: outboundTargets,
				},
				Policies: model.MatchedPolicies{
					Logs: model.LogMap{
						"api-http": &mesh_proto.LoggingBackend{
							Name: "file",
							Type: mesh_proto.LoggingFileType,
							Conf: util_proto.MustToStruct(&mesh_proto.FileLoggingBackendConfig{
								Path: "/var/log",
							}),
						},
						"api-tcp": &mesh_proto.LoggingBackend{
							Name: "elk",
							Type: mesh_proto.LoggingTcpType,
							Conf: util_proto.MustToStruct(&mesh_proto.TcpLoggingBackendConfig{
								Address: "logstash:1234",
							}),
						},
					},
					CircuitBreakers: model.CircuitBreakerMap{
						"api-http": &mesh_core.CircuitBreakerResource{
							Spec: &mesh_proto.CircuitBreaker{
								Conf: &mesh_proto.CircuitBreaker_Conf{
									Detectors: &mesh_proto.CircuitBreaker_Conf_Detectors{
										TotalErrors: &mesh_proto.CircuitBreaker_Conf_Detectors_Errors{},
									},
								},
							},
						},
					},
				},

				Metadata: &model.DataplaneMetadata{},
			}

			// when
			given.ctx.ControlPlane.CLACache = &dummyCLACache{outboundTargets: outboundTargets}
			rs, err := gen.Generate(given.ctx, proxy)

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
			Expect(actual).To(MatchGoldenYAML(filepath.Join("testdata", "outbound-proxy", given.expected)))
		},
		Entry("01. transparent_proxying=false, mtls=false, outbound=0", testCase{
			ctx:       plainCtx,
			dataplane: ``,
			expected:  "01.envoy.golden.yaml",
		}),
		Entry("02. transparent_proxying=true, mtls=false, outbound=0", testCase{
			ctx: mtlsCtx,
			dataplane: `
            networking:
              transparentProxying:
                redirectPort: 15001
`,
			expected: "02.envoy.golden.yaml",
		}),
		Entry("03. transparent_proxying=false, mtls=false, outbound=4", testCase{
			ctx: plainCtx,
			dataplane: `
            networking:
              address: 10.0.0.1
              gateway:
                tags:
                  kuma.io/service: gateway
              outbound:
              - port: 18080
                service: backend
              - port: 54321
                service: db
              - port: 40001
                service: api-http
              - port: 40002
                service: api-tcp
              - port: 40003
                service: api-http2
              - port: 40004
                service: api-grpc
`,
			expected: "03.envoy.golden.yaml",
		}),
		Entry("04. transparent_proxying=true, mtls=true, outbound=4", testCase{
			ctx: mtlsCtx,
			dataplane: `
            networking:
              address: 10.0.0.1
              inbound:
              - port: 8080
                tags:
                  kuma.io/service: web
              outbound:
              - port: 18080
                service: backend
              - port: 54321
                service: db
              - port: 40001
                service: api-http
              - port: 40002
                service: api-tcp
              transparentProxying:
                redirectPortOutbound: 15001
                redirectPortInbound: 15006
`,
			expected: "04.envoy.golden.yaml",
		}),
		Entry("05. transparent_proxying=true, mtls=true, outbound=1 with ExternalService", testCase{
			ctx: mtlsCtx,
			dataplane: `
            networking:
              address: 10.0.0.1
              inbound:
              - port: 8080
                tags:
                  kuma.io/service: web
              outbound:
              - port: 18081
                tags:
                  kuma.io/service: es
              transparentProxying:
                redirectPortOutbound: 15001
                redirectPortInbound: 15006
`,
			expected: "05.envoy.golden.yaml",
		}),
		Entry("06. transparent_proxying=true, mtls=true, outbound=1 with ExternalService http2", testCase{
			ctx: mtlsCtx,
			dataplane: `
            networking:
              address: 10.0.0.1
              inbound:
              - port: 8080
                tags:
                  kuma.io/service: web
              outbound:
              - port: 18082
                tags:
                  kuma.io/service: es2
              transparentProxying:
                redirectPortOutbound: 15001
                redirectPortInbound: 15006
`,
			expected: "06.envoy.golden.yaml",
		}),
		Entry("07. no TrafficRoute", testCase{
			ctx: mtlsCtx,
			dataplane: `
            networking:
              address: 10.0.0.1
              inbound:
              - port: 8080
                tags:
                  kuma.io/service: web
              outbound:
              - port: 4040
                tags:
                  kuma.io/service: service-without-traffic-route
`,
			expected: "07.envoy.golden.yaml",
		}),
	)

	It("Add sanitized alternative cluster name for stats", func() {
		// setup
		gen := &generator.OutboundProxyGenerator{}
		dp := `
        networking:
          outbound:
          - port: 18080
            service: backend.kuma-system
          - port: 54321
            service: db.kuma-system`

		dataplane := &mesh_proto.Dataplane{}
		Expect(util_proto.FromYAML([]byte(dp), dataplane)).To(Succeed())

		outboundTargets := model.EndpointMap{
			"backend.kuma-system": []model.Endpoint{
				{
					Target: "192.168.0.1",
					Port:   8082,
					Weight: 1,
				},
			},
			"db.kuma-system": []model.Endpoint{
				{
					Target: "192.168.0.2",
					Port:   5432,
					Tags:   map[string]string{"kuma.io/service": "db", "role": "master"},
					Weight: 1,
				},
			},
		}
		proxy := &model.Proxy{
			Id: model.ProxyId{Name: "side-car", Mesh: "default"},
			Dataplane: &mesh_core.DataplaneResource{
				Meta: &test_model.ResourceMeta{
					Version: "1",
				},
				Spec: dataplane,
			},
			APIVersion: envoy_common.APIV3,
			Routing: model.Routing{
				TrafficRoutes: model.RouteMap{
					mesh_proto.OutboundInterface{
						DataplaneIP:   "127.0.0.1",
						DataplanePort: 18080,
					}: &mesh_core.TrafficRouteResource{
						Spec: &mesh_proto.TrafficRoute{
							Conf: &mesh_proto.TrafficRoute_Conf{
								Split: []*mesh_proto.TrafficRoute_Split{{
									Weight:      100,
									Destination: mesh_proto.MatchService("backend.kuma-system"),
								}},
							},
						},
					},
					mesh_proto.OutboundInterface{
						DataplaneIP:   "127.0.0.1",
						DataplanePort: 54321,
					}: &mesh_core.TrafficRouteResource{
						Spec: &mesh_proto.TrafficRoute{
							Conf: &mesh_proto.TrafficRoute_Conf{
								Split: []*mesh_proto.TrafficRoute_Split{{
									Weight:      100,
									Destination: mesh_proto.TagSelector{"kuma.io/service": "db", "version": "3.2.0"},
								}},
							}},
					},
				},
				OutboundTargets: outboundTargets,
			},
			Metadata: &model.DataplaneMetadata{},
		}

		// when
		plainCtx.ControlPlane.CLACache = &dummyCLACache{outboundTargets: outboundTargets}
		rs, err := gen.Generate(plainCtx, proxy)

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
		Expect(actual).To(MatchGoldenYAML(filepath.Join("testdata", "outbound-proxy", "cluster-dots.envoy.golden.yaml")))
	})
})
