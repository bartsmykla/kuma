package iptables

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/transparentproxy/config"
	"github.com/kumahq/kuma/pkg/transparentproxy/iptables/builder"
	"github.com/kumahq/kuma/pkg/transparentproxy/iptables/consts"
)

func Setup(ctx context.Context, cfg config.InitializedConfig) (string, error) {
	if cfg.DryRun {
		return dryRun(cfg), nil
	}

	if err := Cleanup(ctx, cfg); err != nil {
		return "", err
	}

	return builder.RestoreIPTables(ctx, cfg)
}

// Cleanup removes iptables rules and chains related to the transparent proxy,
// ensuring that only the relevant rules and chains are removed based on the
// presence of iptables comments. It verifies the new rules after cleanup and
// restores them if they are valid.
//
// Args:
//   - ctx (context.Context): The context for command execution.
//   - cfg (config.InitializedConfig): The configuration containing the
//     iptables settings, including comments and redirect information.
//
// Returns:
//   - error: An error if the cleanup process or verification fails.
func Cleanup(ctx context.Context, cfg config.InitializedConfig) error {
	for _, c := range []config.InitializedConfigIPvX{cfg.IPv4, cfg.IPv6} {
		// Execute iptables-save to retrieve current rules.
		stdout, _, err := c.Executables.IptablesSave.Exec(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to execute iptables-save command")
		}

		output := stdout.String()
		containsTProxyRules := strings.Contains(output, c.Redirect.NamePrefix)
		containsTProxyComments := strings.Contains(output, c.Comment.Prefix)

		switch {
		case !containsTProxyRules && !containsTProxyComments:
			// If there are no transparent proxy rules or chains, there is
			// nothing to do.
			c.Logger.Infof(
				"no transparent proxy %s rules detected. No cleanup necessary",
				consts.IptablesCommandByFamily[c.IPv6],
			)
			continue
		case containsTProxyRules && !containsTProxyComments:
			return errors.New(
				"transparent proxy iptables rules detected, but expected comments are missing. Cleanup cannot proceed safely without comments to identify rules",
			)
		}

		// Split the output into lines and remove lines related to transparent
		// proxy and comments
		lines := strings.Split(output, "\n")
		linesCleaned := slices.DeleteFunc(
			lines,
			func(line string) bool {
				isComment := strings.HasPrefix(line, "#")
				isTProxyRule := strings.Contains(line, c.Comment.Prefix)
				isTProxyChain := strings.HasPrefix(
					line,
					fmt.Sprintf(":%s_", c.Redirect.NamePrefix),
				)

				return isComment || isTProxyRule || isTProxyChain
			},
		)
		newRules := strings.Join(linesCleaned, "\n")

		// Verify if the new rules after cleanup are correct.
		if _, err := c.Executables.RestoreTest(ctx, newRules); err != nil {
			return errors.Wrap(
				err,
				"verification if new rules after cleanup are correct failed",
			)
		}

		if c.DryRun {
			c.Logger.Infof(
				"dry run mode: %s rules after cleanup:",
				consts.IptablesCommandByFamily[c.IPv6],
			)
			c.Logger.InfoWithoutPrefix(newRules)
			continue
		}

		// Restore the new rules with flushing.
		if _, err := c.Executables.RestoreWithFlush(ctx, newRules); err != nil {
			return errors.Wrap(
				err,
				"failed to restore rules with flush after cleanup",
			)
		}
	}

	return nil
}

// dryRun simulates the setup of iptables rules for both IPv4 and IPv6
// configurations based on the provided config.InitializedConfig. It does not
// apply any changes to the system but generates and returns a string
// representation of what would be executed.
//
// The function operates as follows:
//  1. It defines a helper function, ipvxRun, which:
//     - Builds the iptables-restore content for either IPv4 or IPv6 depending
//     on the input.
//     - Prepends a header (### IPv4 ### or ### IPv6 ###) to distinguish between
//     the IP versions.
//     - Returns the formatted iptables rules or an error if the building
//     process fails.
//  2. Executes ipvxRun for IPv4 and, if enabled in the configuration, for IPv6.
//  3. Concatenates the results from IPv4 and IPv6 runs, separating them with
//     newlines for clarity.
//  4. Logs the final combined output using the configured logger without
//     prefixing, to ensure that the output is clear and unmodified, suitable
//     for review or documentation purposes.
//
// Args:
//
//   - cfg (config.InitializedConfig): Configuration settings that include flags
//     for dry run, logging, and IP version preferences.
//
// Returns:
//
//   - string: A combined string of formatted iptables commands for both IPv4
//     and IPv6.
//   - error: An error if there is a failure in generating the iptables commands
//     for any version.
func dryRun(cfg config.InitializedConfig) string {
	output := strings.Join(
		slices.Concat(
			dryRunIPvX(cfg.IPv4, false),
			dryRunIPvX(cfg.IPv6, true),
		),
		"\n\n",
	)

	cfg.Logger.InfoWithoutPrefix(output)

	return output
}

// dryRunIPvX generates iptables rules for either IPv4 or IPv6 based on the
// provided configuration. It returns a slice with a header indicating the
// IP version and the generated rules as a single string.
//
// Args:
//   - cfg (config.InitializedConfigIPvX): Configuration settings for IPv4 or
//     IPv6.
//   - ipv6 (bool): Indicates if the configuration is for IPv6.
//
// Returns:
//   - []string: A slice containing the header and the iptables rules for the
//     specified IP version.
func dryRunIPvX(cfg config.InitializedConfigIPvX, ipv6 bool) []string {
	if !cfg.Enabled() {
		return nil
	}

	return []string{
		fmt.Sprintf("### %s ###", consts.IPTypeMap[ipv6]),
		strings.TrimSpace(builder.BuildIPTablesForRestore(cfg)),
	}
}
