package otel

import (
	"math"
	net_url "net/url"
	"strconv"
	"strings"

	"github.com/kumahq/kuma/v2/pkg/core/validators"
	core_xds "github.com/kumahq/kuma/v2/pkg/core/xds"
)

// ValidateEndpoint validates an OpenTelemetry endpoint that can be either
// an HTTP/HTTPS URL or a gRPC host:port address. For HTTP/HTTPS URLs, it
// validates the scheme, port, and hostname. Returns isURL=true if the endpoint
// was validated as a URL, false if it's a plain gRPC host:port.
func ValidateEndpoint(endpointPath validators.PathBuilder, endpoint string) (bool, validators.ValidationError) {
	var verr validators.ValidationError
	parsedURL, err := net_url.ParseRequestURI(endpoint)
	switch {
	case err == nil && parsedURL.Host != "":
		if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			verr.AddViolationAt(endpointPath, "URL scheme must be http or https")
		}
		if portStr := parsedURL.Port(); portStr != "" {
			port, err := strconv.Atoi(portStr)
			if err != nil || port < 1 || port > math.MaxUint16 {
				verr.AddViolationAt(endpointPath, "port must be valid (1-65535)")
			}
		}
		if parsedURL.Hostname() == "" {
			verr.AddViolationAt(endpointPath, "hostname must be defined")
		}
		return true, verr
	case strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://"):
		verr.AddViolationAt(endpointPath, "must be a valid URL")
		return true, verr
	case err != nil && strings.Contains(endpoint, "://"):
		verr.AddViolationAt(endpointPath, "must be a valid URL")
		return true, verr
	default:
		return false, verr
	}
}

// ParseEndpoint parses an OpenTelemetry endpoint string into an xds.Endpoint.
// For HTTP/HTTPS URLs, it extracts hostname, port, and TLS config.
// For gRPC host:port endpoints, it uses the provided default port if none specified.
func ParseEndpoint(endpoint string, defaultGrpcPort uint32) *core_xds.Endpoint {
	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		// Error is ignored because the endpoint was already validated by the validator.
		url, _ := net_url.ParseRequestURI(endpoint)
		port := uint32(80)
		if url.Scheme == "https" {
			port = 443
		}
		if portStr := url.Port(); portStr != "" {
			if val, err := strconv.ParseInt(portStr, 10, 32); err == nil {
				port = uint32(val)
			}
		}
		return &core_xds.Endpoint{
			Target: url.Hostname(),
			Port:   port,
			ExternalService: &core_xds.ExternalService{
				TLSEnabled:         url.Scheme == "https",
				AllowRenegotiation: true,
			},
		}
	}

	// gRPC endpoint (host:port format)
	target := strings.Split(endpoint, ":")
	port := defaultGrpcPort
	if len(target) > 1 {
		if val, err := strconv.ParseInt(target[1], 10, 32); err == nil && val > 0 && val <= 65535 {
			port = uint32(val)
		}
	}
	return &core_xds.Endpoint{
		Target: target[0],
		Port:   port,
	}
}

// IsHTTP returns true if the endpoint uses HTTP or HTTPS scheme.
func IsHTTP(endpoint string) bool {
	return strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://")
}
