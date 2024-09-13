//go:build !windows
// +build !windows

package install

import (
	"fmt"

	"github.com/spf13/cobra"

	tproxy_config "github.com/kumahq/kuma/pkg/transparentproxy/config"
	tproxy_validate "github.com/kumahq/kuma/pkg/transparentproxy/validate"
)

func newInstallTransparentProxyValidator() *cobra.Command {
	var ipFamilyMode tproxy_config.IPFamilyMode
	var serverPort uint16

	cmd := &cobra.Command{
		Use:   "transparent-proxy-validator",
		Short: "Validates if transparent proxy has been set up successfully",
		Long: `Validates the transparent proxy setup by testing if the applied 
iptables rules are working correctly onto the pod.

Follow the following steps to validate:
 1) install the transparent proxy using 'kumactl install transparent-proxy'
 2) run this command

The result will be shown as text in stdout as well as the exit code.
`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return tproxy_validate.RunValidation(cmd.Context(), serverPort, ipFamilyMode)
		},
	}

	cmd.Flags().Var(
		&ipFamilyMode,
		"ip-family-mode",
		fmt.Sprintf(
			"specify the IP family mode for traffic redirection when setting up the transparent proxy; accepted values: %s",
			tproxy_config.AllowedIPFamilyModes(),
		),
	)

	cmd.Flags().Uint16Var(
		&serverPort,
		"validation-server-port",
		serverPort,
		"port number for the validation server to listen on",
	)

	return cmd
}
