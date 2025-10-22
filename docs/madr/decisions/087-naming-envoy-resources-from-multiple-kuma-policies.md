# Naming Envoy resources and stats that result from multiple Kuma policies/resources

* Status: accepted

Technical Story: https://github.com/kumahq/kuma/issues/13886

## Context and Problem Statement

### Background: Kuma's unified naming strategy

Kuma has established a unified naming strategy for Envoy resources and stats across four foundational MADRs: [MADR-070 Resource Identifier](https://github.com/kumahq/kuma/blob/d19b78a4556962f4d9d3cc5921c7bdc73dc93d26/docs/madr/decisions/070-resource-identifier.md), [MADR-076 Standardized Naming for system xDS Resources](https://github.com/kumahq/kuma/blob/210d6f7a354bc279b0f2184bb4a4ae9229394b8c/docs/madr/decisions/076-naming-system-envoy-resources.md), [MADR-077 Defining and migrating to consistent naming for non-system Envoy resources and stats](https://github.com/kumahq/kuma/blob/dc2dd5564cac3f97f85085a8aa498f4a520ad4ba/docs/madr/decisions/077-migrating-to-consistent-and-well-defined-naming-for-non-system-envoy-resources-and-stats.md), and [MADR-083 Finalized schemes for unified Envoy resources and stats naming formats](https://github.com/kumahq/kuma/blob/148987f8106e629c9a646c8bdae0f215cf6f1290/docs/madr/decisions/083-finalized-schemes-for-unified-envoy-resources-and-stats-naming-formats.md).

The fundamental principle is straightforward: **the name of an Envoy resource should correspond to the KRI of the Kuma policy or resource that created it**. This enables users to:
- Trace metrics back to the Kuma resource that generated them
- Understand which policy is responsible for specific Envoy configuration
- Filter and query metrics by Kuma resource attributes (mesh, zone, namespace, name)
- Debug and troubleshoot issues by correlating Envoy behavior with Kuma configuration

To achieve this traceability, the strategy defines three complementary naming formats:
- **KRI format** - for resources mapping directly to a single Kuma resource
- **Contextual format** (`self_...`) - for resources local to a specific proxy
- **System format** (`system_...`) - for internal Kuma resources

### The gap: Multi-policy resources

This principle works well when there is a **one-to-one relationship** between a Kuma resource and an Envoy resource. However, **some Envoy resources and components are the result of multiple Kuma policies being merged together**. When multiple policies contribute to a single Envoy component, there is no single KRI to use for naming.

**Common scenarios where multiple policies merge:**
- Routes from multiple `MeshHTTPRoute` policies targeting the same service
- RBAC filters combining rules from multiple `MeshTrafficPermission` policies
- HTTP connection managers where multiple `MeshTimeout`, `MeshRetry`, or `MeshRateLimit` policies apply
- Clusters where multiple `MeshHealthCheck` or `MeshCircuitBreaker` policies merge
- Passthrough filter chains and clusters from multiple `MeshPassthrough` policies targeting the same proxy

These multi-policy scenarios were explicitly marked as out of scope in [previous MADRs](https://github.com/kumahq/kuma/blob/210d6f7a354bc279b0f2184bb4a4ae9229394b8c/docs/madr/decisions/076-naming-system-envoy-resources.md#L119-L120). **This MADR addresses that gap** by evaluating approaches to handle multi-policy Envoy resources and components. The evaluation focuses on resources and components that **generate stats and metrics**, where naming directly impacts observability. For resources without observability impact, not having an explicit name is acceptable.

## Scope

### Problem definition and analysis

Identifies which resources from multiple policies need names and why. Examines affected resources, when naming is necessary (resources with observability impact) versus unnecessary (no stats/metrics), and the trade-offs between traceability and metric cardinality.

### Decision-making process

Evaluates and selects an approach by analyzing current behavior, comparing options (maintaining status quo, new formats, adapting existing patterns, hybrids, or leaving resources unnamed), and weighing trade-offs in impact, value, effort, and feasibility.

### Decision outcomes

Documents the chosen approach, rationale, implementation guidance, and migration considerations.

## Out of scope

- **Already-solved naming**: KRI-based, contextual, and system naming are established patterns
- **Resources without observability impact**: Resources that don't emit stats/metrics don't need explicit naming
- **Non-xDS resources**: Endpoints, secrets, and other resources handled separately
- **Deprecated patterns**: Legacy `kuma.io/service` tag-based naming
- **Default routes**: Routes generated when no `MeshHTTPRoute` is defined will be addressed in a separate MADR

## Why this matters

**Observability impact**: Current ad-hoc naming (hashes, fixed strings, undefined names) creates traceability gaps, breaks consistency with single-policy resources, complicates troubleshooting, and makes dashboards and alerts harder to build.

**Design constraints**: Solutions must balance traceability versus cardinality, stability versus completeness. Naively including all policy details risks extremely long names, metric explosion, and configuration churn.

**Tooling impact**: Naming affects Kuma GUI, Prometheus, OpenTelemetry exporters, and third-party tools. The strategy must accommodate future policies and merge scenarios without constant revisions.

## Envoy resources that generate stats and metrics

Naming primarily matters for resources generating stats and metrics. This section identifies which Envoy resources and components emit stats - the cases requiring evaluation. Resources without observable metrics may not need explicit names.

### Primary resources

The main xDS resource types that generate metrics:

1. **Listeners** - Generate `listener.<name>.*` stats (`downstream_cx_active`, `downstream_cx_total`, `downstream_rq_active`)

2. **Clusters** - Generate `cluster.<name>.*` stats (`upstream_cx_active`, `upstream_cx_total`, `upstream_rq_total`, `health_check.success`)

3. **Routes** - Generate stats within HTTP connection manager using the virtual host name. Individual route entries are tracked through tracing and access logs.

### Nested components

**Network Filters** (applied to listeners):
- **HTTPConnectionManager** - Uses `stat_prefix` for stats like `http.<stat_prefix>.downstream_cx_active`
- **TCPProxy** - Uses `stat_prefix` for stats like `tcp.<stat_prefix>.downstream_cx_total`
- **RBAC** - Uses `stat_prefix` for stats like `<stat_prefix>.allowed`, `<stat_prefix>.denied`; often merged from multiple `MeshTrafficPermission` policies

**HTTP Filters** (applied within HTTPConnectionManager):
- **RBAC** - Uses `stat_prefix`, similarly merged from multiple policies
- Various other filters with stat prefixes

**Route Configuration**:
- Contains virtual hosts and routes
- Route entries can merge from multiple `MeshHTTPRoute` policies

## Envoy resource composition in Kuma

The diagrams below show Envoy resource hierarchy. Each component is annotated with "Affected by:" to indicate which policies influence its configuration, identifying where multiple policies merge - the core naming challenge.

### Listener composition

```
Listener
├── name
│   └── Affected by: MeshGateway (gateway), Dataplane (sidecar)
├── address/port
│   └── Affected by: Dataplane, MeshGateway, networking.transparentProxying
└── filter_chains → filters
    ├── HTTPConnectionManager (HTTP traffic)
    │   ├── stat_prefix
    │   │   └── Affected by: MeshService/MeshExternalService/MeshMultiZoneService (outbounds)
    │   ├── route_config or RDS
    │   │   └── Affected by: MeshHTTPRoute (multiple policies merge)
    │   └── http_filters
    │       ├── RBAC filter
    │       │   └── Affected by: MeshTrafficPermission (multiple policies merge)
    │       ├── Timeout filter
    │       │   └── Affected by: MeshTimeout (multiple policies)
    │       ├── Retry filter
    │       │   └── Affected by: MeshRetry (multiple policies)
    │       ├── RateLimit filter
    │       │   └── Affected by: MeshRateLimit (multiple policies)
    │       ├── FaultInjection filter
    │       │   └── Affected by: MeshFaultInjection (multiple policies)
    │       ├── Tracing/AccessLog
    │       │   └── Affected by: MeshTrace, MeshAccessLog
    │       └── Router filter
    └── TCPProxy (TCP traffic)
        ├── stat_prefix
        │   └── Affected by: MeshService/MeshExternalService/MeshMultiZoneService (outbounds)
        └── cluster reference
            └── Affected by: MeshTCPRoute, target service
```

### Cluster composition

```
Cluster
├── name
│   └── Affected by: MeshService, MeshExternalService, MeshMultiZoneService
├── alt_stat_name (overrides cluster name in stats)
│   └── Affected by: same as name (should match for consistency)
├── type (EDS, STATIC, etc.)
│   └── Affected by: service type, MeshExternalService configuration
├── load_balancing_policy
│   └── Affected by: MeshLoadBalancingStrategy
├── health_checks
│   └── Affected by: MeshHealthCheck (multiple policies)
├── circuit_breakers
│   └── Affected by: MeshCircuitBreaker (multiple policies)
├── outlier_detection
│   └── Affected by: MeshCircuitBreaker
└── endpoints or EDS reference
    └── Affected by: service discovery, zone configuration
```

### Route configuration composition

```
RouteConfiguration
├── name
│   └── Affected by: target MeshService/MeshExternalService/MeshMultiZoneService, listener context
└── virtual_hosts
    ├── name (appears in vhost stats, derived by Envoy)
    │   └── Derived from: hostname/domain configuration
    ├── domains
    │   └── Affected by: MeshHTTPRoute match, MeshGateway hostname
    └── routes (individual route entries)
        ├── match criteria
        │   └── Affected by: MeshHTTPRoute (multiple policies contribute routes)
        ├── route action
        │   └── Affected by: MeshHTTPRoute backendRefs, target service
        ├── per-route timeout/retry/rate limit/fault injection
        │   └── Affected by: MeshTimeout/MeshRetry/MeshRateLimit/MeshFaultInjection (multiple policies)
        └── per-route filters
            └── Affected by: various HTTP filter policies at route level
```

## Where multiple Kuma policies affect Envoy resource naming

Based on analysis of Kuma's policy application and Envoy resource generation, these are the key areas where multiple policies merge into single Envoy resources or components.

**Real-world examples**: Examples in this section come from a live multi-zone Kuma deployment (global control plane `global`, zone control planes `zone-1` and `zone-2`) with multiple services and various policies applied. Examples show actual Envoy configuration fragments from data plane proxy `config_dump` endpoints.

### 1. Routes merged from multiple MeshHTTPRoute policies

**Problem**: Multiple `MeshHTTPRoute` policies targeting the same outbound service merge into a single Envoy `RouteConfiguration` with individual route entries from different policies.

**Current behavior**: The `RouteConfiguration`, `VirtualHost`, `HTTPConnectionManager` stat_prefix, `Listener`, and `Cluster` all use the target service's KRI (e.g., `kri_msvc_default_zone-1_demo-app_backend-api_httpport`). Each route entry uses its source policy's KRI + `rule_<index>` suffix (e.g., `kri_mhttpr_default_zone-1_demo-app_backend-routing_rule_0`).

**Example**: Listener for `backend-api` with three merged `MeshHTTPRoute` policies:

```yaml
Listener:
  name: kri_msvc_default_zone-1_demo-app_backend-api_httpport
  filter_chains:
    - filters:
        - name: envoy.filters.network.http_connection_manager
          typed_config:
            stat_prefix: kri_msvc_default_zone-1_demo-app_backend-api_httpport
            route_config:
              name: kri_msvc_default_zone-1_demo-app_backend-api_httpport
              virtual_hosts:
                - name: kri_msvc_default_zone-1_demo-app_backend-api_httpport
                  domains: ["*"]
                  routes:
                    - name: kri_mhttpr_default_zone-1_demo-app_backend-routing_rule_0
                      match: { prefix: / }
                    - name: kri_mhttpr_default__kuma-system_client-to-backend_rule_0
                      match: { prefix: /multizone }
                    - name: kri_mhttpr_default__kuma-system_client-to-backend_rule_0
                      match: { prefix: /zone1 }
                    - name: kri_mhttpr_default_zone-1_demo-app_backend-with-timeout_rule_0
                      match: { prefix: /api }
                    - name: kri_mhttpr_default_zone-1_demo-app_backend-routing_rule_0
                      match: { prefix: /v1 }

Cluster:
  name: kri_msvc_default_zone-1_demo-app_backend-api_httpport
```

Three policies merged: `backend-routing`, `client-to-backend`, `backend-with-timeout`.

**Naming challenge**: HTTP stats like `http.<stat_prefix>.downstream_rq_timeout` or `http.<stat_prefix>.downstream_rq_2xx` aggregate across all policies without attribution. Users cannot determine which specific `MeshHTTPRoute`, `MeshTimeout`, or `MeshRetry` policy contributed to observed behavior. Individual routes are identifiable in access logs and tracing by policy KRI.

**Stats affected**: Access logs and tracing identify individual routes. HTTP connection manager stats (`http.<stat_prefix>.*`) aggregate without attribution.

### 2. RBAC filters merged from multiple MeshTrafficPermission policies

**Problem**: Multiple `MeshTrafficPermission` policies applying to the same inbound or outbound merge into a single `RBAC` filter whose `stat_prefix` must identify the combined configuration.

**Current behavior**: Multiple restrictive `MeshTrafficPermission` policies merge into a single `RBAC` filter (network or HTTP) with a `stat_prefix` field. The prefix uses either contextual format for inbounds (e.g., `self_inbound_dp_<sectionName>`) or target service KRI for outbounds (e.g., `kri_msvc_..._backend-api_httpport`).

**Example**: Three policies (`allow-from-frontend`, `allow-from-gateway`, `allow-admin-access`) merge into one `RBAC` filter with a single `stat_prefix`.

**Naming challenge**: The `stat_prefix` doesn't reflect which specific `MeshTrafficPermission` policies contributed to allow/deny rules. Stats like `<stat_prefix>.allowed` and `<stat_prefix>.denied` aggregate across all merged policies, making it impossible to attribute permission decisions to individual policies.

**Stats affected**: `<stat_prefix>.allowed`, `<stat_prefix>.denied`, `<stat_prefix>.shadow_allowed`, `<stat_prefix>.shadow_denied`

### 3. Multi-policy filter chains

**Problem**: Listener configurations with TLS/SNI routing or protocol detection may have filter chains where selection and configuration depend on multiple policies.

**Example**: Multiple `MeshTimeout` or `MeshRetry` policies with different selectors affect the same HTTP connection manager, whose `stat_prefix` represents all merged configurations.

**Stats affected**: `http.<stat_prefix>.*`, `tcp.<stat_prefix>.*`

### 4. Configuration policies that merge without creating named components

**Problem**: Several policy types (`MeshTimeout`, `MeshRetry`, `MeshRateLimit`, `MeshFaultInjection`, `MeshHealthCheck`, `MeshCircuitBreaker`) don't generate separate named filters or components. They merge configuration directly into existing components like `HTTPConnectionManager` or `Cluster`. When multiple policies apply to the same resource, stats aggregate without attribution.

**Current behavior**: These policies merge configuration at various levels:
- **HTTP connection manager level**: Timeout, retry, rate limit, and fault injection settings merge into `HTTPConnectionManager` or individual route entries
- **Cluster level**: Health check and circuit breaker settings merge into `Cluster` configuration

The `stat_prefix` for `HTTPConnectionManager` and cluster names use the target service KRI without indicating which specific policies contributed.

**Example**: Listener with merged policy configurations:

```yaml
Listener:
  name: kri_msvc_default_zone-1_demo-app_backend-api_httpport
  filter_chains:
    - filters:
        - name: envoy.filters.network.http_connection_manager
          typed_config:
            stat_prefix: kri_msvc_default_zone-1_demo-app_backend-api_httpport
            # Multiple MeshTimeout, MeshRetry, MeshRateLimit, MeshFaultInjection
            # policies merge here without individual attribution
```

HTTP stats use this `stat_prefix`: `http.kri_msvc_default_zone-1_demo-app_backend-api_httpport.downstream_rq_timeout`

**Naming challenges by policy type**:

| Policy Type | Challenge | Stats Affected |
|-------------|-----------|----------------|
| **HTTP connection manager level** |
| `MeshTimeout` | Stats like `http.<stat_prefix>.downstream_rq_timeout` aggregate across all timeout policies applying different values at different levels | `http.<stat_prefix>.downstream_rq_timeout`, route-specific timeout stats |
| `MeshRetry` | Stats like `http.<stat_prefix>.retry.*` aggregate across all retry policies applying to different routes/conditions | `http.<stat_prefix>.retry.*`, per-route retry overflow and success metrics |
| `MeshRateLimit` | Stats like `http.<stat_prefix>.ratelimit.*` aggregate across policies applying at different levels and dimensions | `http.<stat_prefix>.ratelimit.ok`, `http.<stat_prefix>.ratelimit.over_limit`, per-route rate limit stats |
| `MeshFaultInjection` | Stats like `http.<stat_prefix>.fault.*` aggregate across policies injecting different fault types on different routes | `http.<stat_prefix>.fault.delays_injected`, `http.<stat_prefix>.fault.aborts_injected` |
| **Cluster level** |
| `MeshHealthCheck` | Cluster health check stats don't reflect which policies configured the checks | `cluster.<name>.health_check.attempt`, `cluster.<name>.health_check.success`, `cluster.<name>.health_check.failure`, outlier detection stats |
| `MeshCircuitBreaker` | Cluster circuit breaker stats don't reflect which policies set thresholds | `cluster.<name>.circuit_breakers.*`, `cluster.<name>.upstream_rq_pending_overflow`, outlier detection ejection stats |

**Common pattern**: All these policy types merge into existing Envoy components without creating separately-named filters or resources, making it impossible to attribute specific stats or behaviors to individual policies.

### 5. Passthrough filter chains and clusters from MeshPassthrough policies

**Problem**: `MeshPassthrough` policies configure filter chains on a common, shared `outbound:passthrough:ipv4` listener. Filter chains are grouped by protocol and port (e.g., `meshpassthrough_http_80`), not by individual policy. This means the naming scheme provides no attribution to source policies - whether it's a single policy with multiple domains or multiple policies contributing to the same protocol/port combination. Each match entry creates separate clusters, but filter chain names reflect only the protocol/port grouping.

**Current behavior**: `MeshPassthrough` policies configure a common, shared listener for all passthrough traffic. The naming scheme is based on protocol/port grouping, not policy identity:
- **Listener name**: `outbound:passthrough:ipv4` (fixed legacy format, common to all passthrough traffic regardless of which policies configure it)
- **Filter chain names**: `meshpassthrough_<protocol>_<port>` (e.g., `meshpassthrough_http_80`) - grouped by protocol/port, not by policy
- **HTTPConnectionManager stat_prefix**: `meshpassthrough_<protocol>_<port>` (same as filter chain)
- **RouteConfiguration names**: `meshpassthrough_<protocol>_<port>` (same as filter chain)
- **Cluster names**: `meshpassthrough_<protocol>_<domain>_<port>` (e.g., `meshpassthrough_http_example1.com_80`)
- **altStatName**: Sanitized version with dots/wildcards replaced by underscores (e.g., `meshpassthrough_http_example1_com_80`)

All `MeshPassthrough` policies targeting the same proxy configure this shared listener. Filter chains are organized by protocol/port grouping - whether domains come from one policy or multiple policies, they share the same filter chain for a given protocol/port combination. Each domain gets its own virtual host and cluster, but there's no attribution to the source policy.

**Example**: Two `MeshPassthrough` policies (`passthrough-aws-services` and `passthrough-google-apis`) targeting the same namespace with multiple HTTP domains on port 80:

```yaml
Listener:
  name: outbound:passthrough:ipv4
  filter_chains:
    - name: meshpassthrough_http_80
      filter_chain_match:
        destination_port: 80
        application_protocols: ["http/1.1", "h2c"]
      filters:
        - name: envoy.filters.network.http_connection_manager
          typed_config:
            stat_prefix: meshpassthrough_http_80
            route_config:
              name: meshpassthrough_http_80
              virtual_hosts:
                - name: s3.amazonaws.com
                  domains: ["s3.amazonaws.com", "s3.amazonaws.com:80"]
                  routes:
                    - match: { prefix: / }
                      route: { cluster: meshpassthrough_http_s3.amazonaws.com_80 }
                - name: dynamodb.us-east-1.amazonaws.com
                  domains: ["dynamodb.us-east-1.amazonaws.com", "dynamodb.us-east-1.amazonaws.com:80"]
                  routes:
                    - match: { prefix: / }
                      route: { cluster: meshpassthrough_http_dynamodb.us-east-1.amazonaws.com_80 }
                - name: storage.googleapis.com
                  domains: ["storage.googleapis.com", "storage.googleapis.com:80"]
                  routes:
                    - match: { prefix: / }
                      route: { cluster: meshpassthrough_http_storage.googleapis.com_80 }

Clusters:
  - name: meshpassthrough_http_s3.amazonaws.com_80
    alt_stat_name: meshpassthrough_http_s3_amazonaws_com_80
  - name: meshpassthrough_http_dynamodb.us-east-1.amazonaws.com_80
    alt_stat_name: meshpassthrough_http_dynamodb_us-east-1_amazonaws_com_80
  - name: meshpassthrough_http_storage.googleapis.com_80
    alt_stat_name: meshpassthrough_http_storage_googleapis_com_80
```

This example shows multiple policies (`passthrough-aws-services` and `passthrough-google-apis`) contributing domains to the same filter chain `meshpassthrough_http_80`. However, the same naming challenge exists even with a single policy that configures multiple domains on the same protocol/port - all domains share the common filter chain name based on protocol/port grouping.

**Naming challenge**: The root cause is the shared listener architecture and protocol/port-based grouping. Filter chain, route configuration, and `stat_prefix` names are derived from protocol/port (`meshpassthrough_http_80`), not from policy identity. This means stats like `http.meshpassthrough_http_80.downstream_rq_total` aggregate across all domains on that protocol/port without any policy attribution - whether those domains come from one policy or multiple policies. Individual clusters are named by domain, but there's no indication of which `MeshPassthrough` policy configured each domain. Virtual host stats (`vhost.<domain>.*`) can identify individual domains but not the source policies.

**Stats affected**: `http.meshpassthrough_<protocol>_<port>.*` (aggregates all domains/policies), `cluster.meshpassthrough_<protocol>_<domain>_<port>.*` (per domain, no policy attribution), `listener.outbound:passthrough:ipv4.*` (all passthrough traffic)
