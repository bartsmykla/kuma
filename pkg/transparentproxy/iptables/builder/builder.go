package builder

import (
	"context"
	"slices"
	"strings"

	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/transparentproxy/config"
	"github.com/kumahq/kuma/pkg/transparentproxy/iptables/tables"
)

// BuildIPTablesForRestore constructs the complete set of iptables rules for
// restoring, based on the provided configuration and IP version.
//
// This function performs the following steps:
//  1. Builds the NAT table rules by calling `buildNatTable`. If this step
//     fails, it returns an error indicating the failure to build the NAT table.
//  2. Builds the rules for the raw, NAT, and mangle tables using the
//     `BuildRulesForRestore` function.
//  3. Filters out any empty rule sets from the results.
//  4. Joins the remaining rules with a separator based on the verbosity setting
//     from the configuration.
//  5. Returns the joined rules as a single string.
//
// Args:
//
//   - cfg (config.InitializedConfigIPvX): The configuration used to initialize
//     the iptables rules.
//   - ipv6 (bool): A boolean indicating whether to build rules for IPv6.
//
// Returns:
//
//   - string: The complete set of iptables rules as a single string.
//   - error: An error if the NAT table rules cannot be built.
func BuildIPTablesForRestore(
	l config.Logger,
	cfg config.InitializedConfigIPvX,
) string {
	// Build the rules for raw, NAT, and mangle tables, filtering out any empty
	// sets.
	result := slices.DeleteFunc([]string{
		tables.BuildRulesForRestore(cfg, buildRawTable(cfg)),
		tables.BuildRulesForRestore(cfg, buildNatTable(l, cfg)),
		tables.BuildRulesForRestore(cfg, buildMangleTable(cfg)),
	}, func(s string) bool { return s == "" })

	// Determine the separator based on verbosity setting.
	separator := "\n"
	if cfg.Verbose {
		separator = "\n\n"
	}

	// Join the rules with the determined separator and return.
	return strings.Join(result, separator) + "\n"
}

// TODO(bartsmykla): Probably BuildIPTablesForRestore calls could be moved somewhere else
func RestoreIPTables(ctx context.Context, cfg config.InitializedConfig) (string, error) {
	cfg.Logger.Info("kumactl is about to apply the iptables rules that " +
		"will enable transparent proxying on the machine. The SSH connection " +
		"may drop. If that happens, just reconnect again.")

	output, err := cfg.IPv4.Executables.Restore(
		ctx,
		BuildIPTablesForRestore(cfg.Logger, cfg.IPv4),
	)
	if err != nil {
		return "", errors.Wrap(err, "unable to restore iptables rules")
	}

	if cfg.IPv6.Enabled() {
		ipv6Output, err := cfg.IPv6.Executables.Restore(
			ctx,
			BuildIPTablesForRestore(cfg.Logger, cfg.IPv6),
		)
		if err != nil {
			return "", errors.Wrap(err, "unable to restore ip6tables rules")
		}

		output += ipv6Output
	}

	cfg.Logger.Info("iptables set to divert the traffic to Envoy")

	return output, nil
}
