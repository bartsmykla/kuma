# Naming Envoy resources and stats that result from multiple Kuma policies/resources

* Status: accepted

Technical Story: https://github.com/kumahq/kuma/issues/13886

## Context and Problem Statement

### Background: Kuma's unified naming strategy

Kuma has established a comprehensive naming strategy for Envoy resources and their corresponding stats across a series of MADRs:
- [MADR-070 Resource Identifier](https://github.com/kumahq/kuma/blob/d19b78a4556962f4d9d3cc5921c7bdc73dc93d26/docs/madr/decisions/070-resource-identifier.md)
- [MADR-076 Standardized Naming for system xDS Resources](https://github.com/kumahq/kuma/blob/210d6f7a354bc279b0f2184bb4a4ae9229394b8c/docs/madr/decisions/076-naming-system-envoy-resources.md)
- [MADR-077 Defining and migrating to consistent naming for non-system Envoy resources and stats](https://github.com/kumahq/kuma/blob/dc2dd5564cac3f97f85085a8aa498f4a520ad4ba/docs/madr/decisions/077-migrating-to-consistent-and-well-defined-naming-for-non-system-envoy-resources-and-stats.md)
- [MADR-083 Finalized schemes for unified Envoy resources and stats naming formats](https://github.com/kumahq/kuma/blob/148987f8106e629c9a646c8bdae0f215cf6f1290/docs/madr/decisions/083-finalized-schemes-for-unified-envoy-resources-and-stats-naming-formats.md)

The fundamental principle behind this naming effort was to create a consistent way to identify Envoy resources, their corresponding stats/metrics, and the Kuma policies or resources responsible for their existence. The goal was that the name of an Envoy resource should correspond to the KRI of the policy/resource that created it, enabling users to:

- Trace from a metric back to the Kuma resource that generated it
- Understand which policy is responsible for specific Envoy configuration
- Filter and query metrics by Kuma resource attributes (mesh, zone, namespace, name)
- Debug and troubleshoot issues by correlating Envoy behavior with Kuma configuration

To achieve this traceability, the strategy defines three complementary naming formats:

- **KRI format** (`kri_<resourceType>_<mesh>_<zone>_<namespace>_<resourceName>_<sectionName>`) - for resources that directly map to a single Kuma resource like `MeshService`, `MeshHTTPRoute`, or policy instances
- **Contextual format** (`self_<category>_<scope>_<rest>`) - for resources that are local to a specific proxy (like inbound listeners or transparent proxy passthrough)
- **System format** (`system_...`) - for internal Kuma resources not directly tied to user configuration

This principle works well when there is a **one-to-one relationship** between a Kuma resource and an Envoy resource.

### The gap: Multi-policy resources

**The core problem is that some Envoy resources and their components are the result of multiple Kuma policies/resources being merged or combined together.** 

Envoy's architecture is compositional - resources are built from nested components, many of which have their own name fields (`stat_prefix`, `alt_stat_name`) that affect observability. When multiple Kuma policies contribute to a single Envoy resource or component, there is no single Kuma resource whose KRI can be used for naming. Examples include:

- **Routes** merged from multiple `MeshHTTPRoute` policies targeting the same service
- **RBAC filters** combining rules from multiple `MeshTrafficPermission` policies
- **Default routes** generated when no `MeshHTTPRoute` exists (not tied to any user policy)
- **MeshPassthrough** resources that aggregate configuration from multiple sources

Previous MADRs explicitly marked these multi-policy cases as out of scope:
- [MADR-070](https://github.com/kumahq/kuma/blob/d19b78a4556962f4d9d3cc5921c7bdc73dc93d26/docs/madr/decisions/070-resource-identifier.md#L164) identifies Routes as a multi-policy challenge
- [MADR-076](https://github.com/kumahq/kuma/blob/210d6f7a354bc279b0f2184bb4a4ae9229394b8c/docs/madr/decisions/076-naming-system-envoy-resources.md#L119-L120) marks multi-policy resources out of scope
- [MADR-077](https://github.com/kumahq/kuma/blob/dc2dd5564cac3f97f85085a8aa498f4a520ad4ba/docs/madr/decisions/077-migrating-to-consistent-and-well-defined-naming-for-non-system-envoy-resources-and-stats.md#L59) excludes default routes and other multi-policy cases

**This MADR addresses that gap** by evaluating how to handle multi-policy Envoy resources and components. The evaluation focuses on resources and components that **generate stats and metrics**, where naming directly impacts observability. For resources or components that don't affect stats/metrics, not having an explicit name is a valid option. This MADR will weigh the trade-offs between different approaches - including maintaining current behavior, defining new naming schemes, or explicitly choosing not to name certain components - and document the chosen path forward.

## Scope

This MADR addresses the complete lifecycle of the multi-policy naming challenge, from problem definition through decision-making. The scope includes:

### Problem definition and analysis

Identifies which resources from multiple policies need names and why. Examines affected resources, when naming is necessary (resources with observability impact) versus unnecessary (no stats/metrics), and the trade-offs between traceability and metric cardinality.

### Decision-making process

Evaluates and selects an approach by analyzing current behavior, comparing options (maintaining status quo, new formats, adapting existing patterns, hybrids, or leaving resources unnamed), and weighing trade-offs in impact, value, effort, and feasibility.

### Decision outcomes

Documents the chosen approach, rationale, implementation guidance, and migration considerations.

## Out of scope

This MADR does not cover:

- **Already-solved naming**: KRI-based, contextual, and system naming are established patterns
- **Resources without observability impact**: Resources that don't emit stats/metrics don't need explicit naming
- **Non-xDS resources**: Endpoints, secrets, and other resources handled separately
- **Deprecated patterns**: Legacy `kuma.io/service` tag-based naming

## Why this matters

### Observability and consistency impact

Current ad-hoc naming (hashes, fixed strings, undefined names) creates traceability gaps, breaks consistency with single-policy resources, complicates troubleshooting, and makes dashboards and alerts harder to build.

### Design constraints

Solutions must balance traceability versus cardinality, stability versus completeness. Naively including all policy details risks extremely long names, metric explosion, and configuration churn.

### Tooling and extensibility

Naming affects Kuma GUI, Prometheus, OpenTelemetry exporters, and third-party tools. The strategy must accommodate future policies and merge scenarios without constant revisions.

## Envoy resources that generate stats and metrics

Naming primarily matters for resources that generate stats and metrics impacting observability. This section identifies which Envoy resources and components emit stats - these are the cases requiring evaluation. Resources without observable metrics may not need explicit names.

### Primary resources

The main xDS resource types that generate metrics:

1. **Listeners** - Generate `listener.<name>.*` stats (`downstream_cx_active`, `downstream_cx_total`, `downstream_rq_active`, etc.)

2. **Clusters** - Generate `cluster.<name>.*` stats (`upstream_cx_active`, `upstream_cx_total`, `upstream_rq_total`, `health_check.success`, etc.)

3. **Routes** - Generate stats within HTTP connection manager using the virtual host name. Individual route entries are tracked through tracing and access logs.

### Nested components

Components within Envoy resources that have naming fields affecting stats:

**Network Filters** (applied to listeners):
- **HTTPConnectionManager** - Uses `stat_prefix` for stats like `http.<stat_prefix>.downstream_cx_active`
- **TCPProxy** - Uses `stat_prefix` for stats like `tcp.<stat_prefix>.downstream_cx_total`
- **RBAC** (Network Filter) - Uses `stat_prefix` for stats like `<stat_prefix>.allowed`, `<stat_prefix>.denied`
   - **Key point**: Often merged from multiple `MeshTrafficPermission` policies

**HTTP Filters** (applied within HTTPConnectionManager):
- **RBAC** (HTTP Filter) - Uses `stat_prefix`, similarly merged from multiple policies
- Various other filters with stat prefixes

**Route Configuration**:
- Contains virtual hosts and routes
- Route entries can merge from multiple `MeshHTTPRoute` policies

## Envoy resource composition in Kuma

The diagrams below show Envoy resource hierarchy as generated by Kuma. Each component is annotated with "Affected by:" to indicate which policies influence its configuration, identifying where multiple policies merge into single components - the core naming challenge.

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

Based on analysis of Kuma's policy application and Envoy resource generation, these are the key areas where multiple policies/resources merge into single Envoy resources or components.

**Real-world examples**: The examples in this section come from a live multi-zone Kuma deployment (global control plane `global`, zone control planes `zone-1` and `zone-2`) with multiple services and various policies applied. Examples show actual Envoy configuration fragments from data plane proxy `config_dump` endpoints.

### 1. Routes merged from multiple `MeshHTTPRoute` policies

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

**Naming challenge**: HTTP stats like `http.<stat_prefix>.downstream_rq_timeout` or `http.<stat_prefix>.downstream_rq_2xx` aggregate across all policies without attribution to specific `MeshHTTPRoute`, `MeshTimeout`, or `MeshRetry` policies. Individual routes can be identified in access logs and tracing by policy KRI.

**Stats affected**: Access logs and tracing identify individual routes. HTTP connection manager stats (`http.<stat_prefix>.*`) aggregate without attribution.

### 2. RBAC filters merged from multiple `MeshTrafficPermission` policies

**Problem**: Multiple `MeshTrafficPermission` policies applying to the same inbound/outbound merge into a single `RBAC` filter whose `stat_prefix` must identify the combined configuration.

**Current behavior**: Multiple restrictive `MeshTrafficPermission` policies merge into a single `RBAC` filter (network or HTTP) with a `stat_prefix` field. The prefix uses either contextual format for inbounds (e.g., `self_inbound_dp_<sectionName>`) or target service KRI for outbounds (e.g., `kri_msvc_..._backend-api_httpport`).

**Example**: Three policies (`allow-from-frontend`, `allow-from-gateway`, `allow-admin-access`) merge into one `RBAC` filter with a single `stat_prefix`.

**Naming challenge**: The `stat_prefix` doesn't reflect which specific `MeshTrafficPermission` policies contributed to allow/deny rules. Stats like `<stat_prefix>.allowed` and `<stat_prefix>.denied` aggregate across all merged policies.

**Stats affected**: `<stat_prefix>.allowed`, `<stat_prefix>.denied`, `<stat_prefix>.shadow_allowed`, `<stat_prefix>.shadow_denied`

### 3. Default routes when no `MeshHTTPRoute` is defined

**Problem**: When no `MeshHTTPRoute` is defined for an outbound, Kuma generates a default route not derived from user policy.

**Current behavior**: The `RouteConfiguration` and `HTTPConnectionManager` `stat_prefix` use the service KRI (making stats distinguishable per service), but the default route entry uses the same hash (`9Zuf5Tg79OuZcQITwBbQykxAk2u4fRKrwYn3//AL4Yo=`) across all services.

**Example**: Six services without `MeshHTTPRoute` policies share the same default route hash:

```yaml
Listener:
  name: kri_msvc_default_zone-1_demo-app_payment-service_5678
  filter_chains:
    - filters:
        - name: envoy.filters.network.http_connection_manager
          typed_config:
            stat_prefix: kri_msvc_default_zone-1_demo-app_payment-service_5678
            route_config:
              name: kri_msvc_default_zone-1_demo-app_payment-service_5678
              virtual_hosts:
                - name: kri_msvc_default_zone-1_demo-app_payment-service_5678
                  domains: ["*"]
                  routes:
                    - name: 9Zuf5Tg79OuZcQITwBbQykxAk2u4fRKrwYn3//AL4Yo=
                      match: { prefix: / }
```

**Naming challenge**: The hash makes it impossible to identify which service a default route belongs to from the route name alone, distinguish between default routes in access logs/tracing without context, or follow a consistent naming pattern.

**Stats affected**: Per-route stats, access logs. HTTPConnectionManager stats use the service KRI `stat_prefix` and are distinguishable per service.

### 4. MeshPassthrough resources

**Problem**: `MeshPassthrough` policy generates listeners, clusters, and routes for external traffic, with configuration influenced by multiple policy aspects and mesh configuration.

**Current situation**: Explicitly marked as out of scope in [MADR-077](https://github.com/kumahq/kuma/blob/dc2dd5564cac3f97f85085a8aa498f4a520ad4ba/docs/madr/decisions/077-migrating-to-consistent-and-well-defined-naming-for-non-system-envoy-resources-and-stats.md#L59) - to be addressed separately.

**Stats affected**: Passthrough-related listener and cluster stats, route statistics for external traffic.

### 5. Multi-policy filter chains

**Problem**: Listener configurations with TLS/SNI routing or protocol detection may have filter chains where selection and configuration depend on multiple policies.

**Example**: Multiple `MeshTimeout` or `MeshRetry` policies with different selectors affect the same HTTP connection manager, whose `stat_prefix` represents all merged configurations.

**Stats affected**: `http.<stat_prefix>.*`, `tcp.<stat_prefix>.*`

### 6. Configuration policies that merge without creating named components

**Problem**: Several policy types (`MeshTimeout`, `MeshRetry`, `MeshRateLimit`, `MeshFaultInjection`, `MeshHealthCheck`, `MeshCircuitBreaker`) don't generate separate named filters or components. Instead, they merge configuration directly into existing components like `HTTPConnectionManager` or `Cluster`. When multiple policies apply to the same resource, stats aggregate without attribution.

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
