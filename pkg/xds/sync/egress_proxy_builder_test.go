package sync_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"

	system_proto "github.com/kumahq/kuma/api/system/v1alpha1"
	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/config/manager"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_manager "github.com/kumahq/kuma/pkg/core/resources/manager"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/registry"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
	"github.com/kumahq/kuma/pkg/dns/vips"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	"github.com/kumahq/kuma/pkg/plugins/resources/memory"
	test_model "github.com/kumahq/kuma/pkg/test/resources/model"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	xds_cache "github.com/kumahq/kuma/pkg/xds/cache/mesh"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/sync"
)

type fakeResourceManager struct {
	core_manager.ResourceManager
}

type fakeMetadataTracker struct{}

func (m fakeMetadataTracker) Metadata(
	core_model.ResourceKey,
) *core_xds.DataplaneMetadata {
	return nil
}

type fakeDataSourceLoader struct {
}

func (f fakeDataSourceLoader) Load(
	context.Context,
	string,
	*system_proto.DataSource,
) ([]byte, error) {
	return []byte("secret"), nil
}

func init() {
	core.SetLogger(core.NewLogger(0))
}

var _ = Describe("EgressProxyBuilder", func() {
	type testCase struct {
		expected          string
		resourcesToCreate []string
	}

	DescribeTable("should generate Envoy xDS resources",
		func(given testCase) {
			Expect(registry.Global().
				RegisterType(core_mesh.GatewayRouteResourceTypeDescriptor)).
				To(Succeed())
			s := memory.NewStore()
			resourceManager := core_manager.NewResourceManager(s)
			metadataTracker := fakeMetadataTracker{}
			metrics, err := core_metrics.NewMetrics("zone-1")
			Expect(err).ToNot(HaveOccurred())

			lookupIPFunc := func(s string) ([]net.IP, error) {
				return []net.IP{net.ParseIP(s)}, nil
			}

			ctx := context.Background()

			for _, fileName := range given.resourcesToCreate {
				resourcePath := filepath.Join(
					"testdata", "resources",
					fileName,
				)

				resourceBytes, err := os.ReadFile(resourcePath)
				Expect(err).ToNot(HaveOccurred())

				resourceReader := bytes.NewReader(resourceBytes)
				yamlDecoder := yaml.NewDecoder(resourceReader)

				var parsedResource map[string]interface{}

				for yamlDecoder.Decode(&parsedResource) == nil {
					var mesh string
					kind := parsedResource["type"].(string)
					name := parsedResource["name"].(string)
					delete(parsedResource, "type")
					delete(parsedResource, "name")
					if m, ok := parsedResource["mesh"].(string); ok {
						mesh = m
						delete(parsedResource, "mesh")
					}

					specBytes, err := yaml.Marshal(parsedResource)
					Expect(err).To(BeNil())

					object, err := registry.Global().
						NewObject(core_model.ResourceType(kind))
					Expect(err).To(BeNil())

					meta := &test_model.ResourceMeta{
						Name: name,
						Mesh: mesh,
					}
					object.SetMeta(meta)

					Expect(util_proto.FromYAML(specBytes, object.GetSpec())).To(Succeed())

					Expect(resourceManager.Create(
						ctx, object,
						core_store.CreateByKey(name, mesh),
					)).To(Succeed())
				}
			}

			meshContextBuilder := xds_context.NewMeshContextBuilder(
				resourceManager,
				registry.Global().ObjectTypes(),
				lookupIPFunc,
				"zone-1",
				vips.NewPersistence(resourceManager, manager.NewConfigManager(s)),
				"mesh",
			)

			meshCache, err := xds_cache.NewCache(0, meshContextBuilder, metrics)
			Expect(err).ToNot(HaveOccurred())

			builder := sync.NewEgressProxyBuilder(
				ctx,
				resourceManager,
				resourceManager,
				lookupIPFunc,
				metadataTracker,
				meshCache,
				fakeDataSourceLoader{},
				envoy_common.APIV3,
			)

			proxy, err := builder.Build(core_model.WithoutMesh("zoneegress-1"))
			Expect(err).To(BeNil())

			f, err := os.OpenFile(
				"/Users/bartsmykla/Projects/github.com/kumahq/kuma/tmp/egress/generated_preview.json",
				os.O_CREATE|os.O_RDWR|os.O_TRUNC,
				0664,
			)
			Expect(err).To(BeNil())
			defer func(f *os.File) {
				_ = f.Close()
			}(f)

			encoder := json.NewEncoder(f)

			encoder.SetIndent("", "  ")

			Expect(encoder.Encode(proxy)).To(Succeed())

			// for outboundInterface, resource := range proxy.ZoneEgressProxy.MeshRoutingMap["mesh-with-mtls-1"]["dataplane-1"].RouteMap {
			// 	Expect(resource.Spec.GetConf().GetSplitWithDestination()).ToNot(BeNil())
			// 	Expect(outboundInterface).ToNot(BeNil())
			// }
		},
		FEntry("hello", testCase{
			expected: "01.envoy.golden.yaml",
			resourcesToCreate: []string{
				"dataplanes.yaml",
			},
		}),
	)
})
