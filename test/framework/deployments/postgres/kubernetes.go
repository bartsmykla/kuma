package postgres

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"

	"github.com/kumahq/kuma/test/framework"
)

// NOTE: We intentionally do not use the tag form like
// "oci://ghcr.io/cloudnative-pg/charts/cloudnative-pg:<version>@<digest>"
// for these charts because the tags appear to be mutable and sometimes
// get overridden by the maintainers with new digests. This would cause
// issues in our CI where all branches could suddenly resolve to a
// different digest, leading to inconsistent and flaky runs. To keep CI
// stable and reproducible, we pin the charts by immutable digest only
const (
	cnpgChart    = "oci://ghcr.io/cloudnative-pg/charts/cloudnative-pg@sha256:c5673ca11b45d3bc46ee4078ca491305202145fc3ee9366aa7d09f730313c73c" // v0.26.0
	clusterChart = "oci://ghcr.io/cloudnative-pg/charts/cluster@sha256:bc577ce089834151a538b19f12e624b086e476b22d43864670d4d806fc56e775" // v0.3.1
)

type k8SDeployment struct {
	envVars map[string]string
	options *deployOptions
}

func (t *k8SDeployment) GetEnvVars() map[string]string {
	return t.envVars
}

var _ PostgresDeployment = &k8SDeployment{}

func (t *k8SDeployment) Name() string {
	return t.options.deploymentName
}

func (t *k8SDeployment) Deploy(cluster framework.Cluster) error {
	extraArgs := map[string][]string{"upgrade": {"--create-namespace", "--install", "--wait"}}

	if err := helm.UpgradeE(
		cluster.GetTesting(),
		&helm.Options{
			KubectlOptions: cluster.GetKubectlOptions(t.options.namespace),
			ExtraArgs:      extraArgs,
		},
		cnpgChart,
		"cnpg",
	); err != nil {
		return err
	}

	if err := k8s.KubectlApplyFromStringE(
		cluster.GetTesting(),
		cluster.GetKubectlOptions(t.options.namespace),
		dbSecrets(
			t.options.namespace,
			"postgres",
			t.options.postgresPassword,
			t.options.username,
			t.options.password,
		),
	); err != nil {
		return err
	}

	opts := &helm.Options{
		SetValues: map[string]string{
			"version.postgresql":                       "16.10",
			"cluster.instances":                        "1",
			"cluster.storage.size":                     "100Mi",
			"cluster.initdb.database":                  t.options.database,
			"cluster.initdb.postInitSQL":               t.options.initScript,
			"cluster.initdb.owner":                     t.options.username,
			"cluster.initdb.secret.name":               fmt.Sprintf("db-%s-secret", t.options.username),
			"cluster.superuserSecret":                  "db-postgres-secret",
			"cluster.services.disabledDefaultServices": "{r,ro}",
		},
		KubectlOptions: cluster.GetKubectlOptions(t.options.namespace),
		ExtraArgs:      extraArgs,
	}

	return helm.UpgradeE(cluster.GetTesting(), opts, clusterChart, t.options.primaryName)
}

func (t *k8SDeployment) Delete(framework.Cluster) error {
	// we delete the namespace anyway and helm.DeleteE is flaky here
	return nil
}

func NewK8SDeployment(opts *deployOptions) *k8SDeployment {
	return &k8SDeployment{
		options: opts,
	}
}

func dbSecrets(namespace string, userPassPairs ...string) string {
	var result []string

	if len(userPassPairs)%2 != 0 {
		panic("userPassPairs must be a multiple of 2")
	}

	for pair := range slices.Chunk(userPassPairs, 2) {
		result = append(result, dbSecret(namespace, pair[0], pair[1]))
	}

	return strings.Join(result, "---\n")
}

func dbSecret(namespace, username, password string) string {
	return fmt.Sprintf(`apiVersion: v1
kind: Secret
type: kubernetes.io/basic-auth
metadata:
  name: db-%[2]s-secret
  namespace: %[1]s
stringData:
  username: %[2]s
  password: %[3]s
`, namespace, username, password)
}
