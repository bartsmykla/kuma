package builder

import (
	"github.com/kumahq/kuma/pkg/transparentproxy/config"
	. "github.com/kumahq/kuma/pkg/transparentproxy/iptables/consts"
	. "github.com/kumahq/kuma/pkg/transparentproxy/iptables/parameters"
	"github.com/kumahq/kuma/pkg/transparentproxy/iptables/rules"
	"github.com/kumahq/kuma/pkg/transparentproxy/iptables/tables"
)

// buildRawTable constructs the raw table for iptables with the necessary rules
// for handling DNS traffic and connection tracking zone splitting. This table
// is configured based on the provided configuration, including handling IPv4
// and IPv6 DNS servers.
func buildRawTable(cfg config.InitializedConfig, ipv6 bool) *tables.RawTable {
	// Initialize the raw table.
	raw := tables.Raw()

	// Determine DNS servers and connection tracking zone split
	// based on IP version.
	dnsServers := cfg.Redirect.DNS.ServersIPv4
	conntractZoneSplit := cfg.Redirect.DNS.ConntrackZoneSplitIPv4
	if ipv6 {
		dnsServers = cfg.Redirect.DNS.ServersIPv6
		conntractZoneSplit = cfg.Redirect.DNS.ConntrackZoneSplitIPv6
	}

	if conntractZoneSplit {
		raw.Output().AddRules(
			rules.
				NewRule(
					Protocol(Udp(DestinationPort(DNSPort))),
					Match(Owner(Uid(cfg.Owner.UID))),
					Jump(Ct(Zone("1"))),
				).
				WithCommentf("assign connection tracking zone 1 to DNS traffic from the kuma-dp user (UID %s)", cfg.Owner.UID),
			rules.
				NewRule(
					Protocol(Udp(SourcePort(cfg.Redirect.DNS.Port))),
					Match(Owner(Uid(cfg.Owner.UID))),
					Jump(Ct(Zone("2"))),
				).
				WithComment("assign connection tracking zone 2 to DNS responses from the kuma-dp DNS proxy"),
		)

		if cfg.ShouldCaptureAllDNS() {
			raw.Output().AddRules(
				rules.
					NewRule(
						Protocol(Udp(DestinationPort(DNSPort))),
						Jump(Ct(Zone("2"))),
					).
					WithComment("assign connection tracking zone 2 to all DNS requests"),
			)

			raw.Prerouting().AddRules(
				rules.
					NewRule(
						Protocol(Udp(SourcePort(DNSPort))),
						Jump(Ct(Zone("1"))),
					).
					WithComment("assign connection tracking zone 1 to all DNS responses"),
			)
		} else {
			for _, ip := range dnsServers {
				raw.Output().AddRules(
					rules.
						NewRule(
							Destination(ip),
							Protocol(Udp(DestinationPort(DNSPort))),
							Jump(Ct(Zone("2"))),
						).
						WithCommentf("assign connection tracking zone 2 to DNS requests destined for %s", ip),
				)
				raw.Prerouting().AddRules(
					rules.
						NewRule(
							Destination(ip),
							Protocol(Udp(SourcePort(DNSPort))),
							Jump(Ct(Zone("1"))),
						).
						WithCommentf("assign connection tracking zone 1 to DNS responses from %s", ip),
				)
			}
		}
	}

	return raw
}
