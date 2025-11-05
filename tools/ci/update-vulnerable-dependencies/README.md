# Update Vulnerable Dependencies Tool

This tool automatically scans for and updates vulnerable Go dependencies in the Kuma project using [Google's OSV Scanner](https://github.com/google/osv-scanner).

## How It Works

1. **Scan**: Uses `osv-scanner` to identify vulnerable dependencies in `go.mod`
2. **Calculate**: For each vulnerable package, calculates the minimum version that fixes all known CVEs using the "least common version" algorithm
3. **Filter**: Checks if the package or its transitive dependencies are excluded per configuration
4. **Update**: Applies the update using `go get` if not excluded
5. **Tidy**: Runs `go mod tidy` to clean up dependencies

## Least Common Version Algorithm

When a package has multiple CVEs, each with different fixed versions, the tool calculates the smallest version that fixes **all** CVEs:

- For each CVE, find the minimum fixed version greater than the current version
- Take the maximum across all CVEs → this is the least common version

**Example:**
```
Current version: 1.0.0
CVE-1 fixed in: [1.0.3, 1.1.0] → pick 1.0.3
CVE-2 fixed in: [1.0.5, 1.2.0] → pick 1.0.5
CVE-3 fixed in: [1.0.2]        → pick 1.0.2
Result: max(1.0.3, 1.0.5, 1.0.2) = 1.0.5
```

This ensures we get the minimum version that addresses all security issues without unnecessary major version bumps.

## Exclusion Configuration

Sometimes certain dependency updates are too complex to automate (e.g., Helm requiring massive Kubernetes dependency updates). You can configure exclusions using GitHub repository variables.

### Configuration Format

Create a repository variable named `SECURITY_UPDATE_CONFIG` in JSON format:

```json
{
  "branch-name": {
    "exclude_direct": ["package1", "package2/prefix"],
    "exclude_transitive": ["prefix1/", "prefix2/"]
  },
  "another-branch": {
    "exclude_direct": [],
    "exclude_transitive": []
  }
}
```

### Exclusion Types

#### 1. Direct Exclusions (`exclude_direct`)

Packages that should **never** be updated by the automation on this branch.

**Example:**
```json
{
  "release-2.7": {
    "exclude_direct": ["helm.sh/helm/v3"]
  }
}
```

This will skip any security updates for the Helm package on the `release-2.7` branch.

#### 2. Transitive Exclusions (`exclude_transitive`)

Package prefixes that, if updated **transitively** (as a side effect of updating another package), should block the entire update.

**Example:**
```json
{
  "release-2.7": {
    "exclude_transitive": ["k8s.io/"]
  }
}
```

This means: "If updating package X would also update any `k8s.io/*` package, skip updating X entirely."

This is useful when you want to avoid cascading updates that pull in many other dependencies.

### Real-World Example

For the scenario mentioned in [Kong/kong-mesh#8606](https://github.com/Kong/kong-mesh/pull/8606):

```json
{
  "release-2.7": {
    "exclude_direct": [],
    "exclude_transitive": ["k8s.io/"]
  }
}
```

Or to directly exclude Helm:

```json
{
  "release-2.7": {
    "exclude_direct": ["helm.sh/helm/v3"],
    "exclude_transitive": []
  }
}
```

### Setting Up the Configuration

1. Go to your repository on GitHub
2. Navigate to **Settings** → **Secrets and variables** → **Actions**
3. Click on the **Variables** tab
4. Click **New repository variable**
5. Name: `SECURITY_UPDATE_CONFIG`
6. Value: Your JSON configuration
7. Click **Add variable**

### Pattern Matching

Both exclusion types support prefix matching:

- `helm.sh/helm/v3` matches `helm.sh/helm/v3` exactly
- `k8s.io/` matches `k8s.io/api`, `k8s.io/client-go`, `k8s.io/apimachinery`, etc.

## GitHub Workflow

The automation runs via the `.github/workflows/update-insecure-dependencies.yaml` workflow:

- **Trigger**: Daily at 3 AM UTC or manually via `workflow_dispatch`
- **Matrix**: Runs on all active branches
- **Output**: Creates a PR with before/after OSV scan results
- **Labels**: Automatically tags PRs with `dependencies` and the branch name

### Workflow Behavior with Exclusions

When exclusions are configured:

1. The workflow reads `SECURITY_UPDATE_CONFIG` from repository variables
2. The script filters out excluded packages and logs skip messages
3. A **GitHub Actions Step Summary** is generated showing:
   - Exclusion configuration for the branch
   - List of updated packages
   - List of skipped packages (by category: direct, transitive, failed)
   - Summary statistics table
4. The PR is created with only the non-excluded updates
5. Skipped packages appear in both the workflow logs and step summary with clear indicators:
   - `⏭️  Skipping <package> (directly excluded by configuration)`
   - `⏭️  Skipping <package> (would update excluded transitive dependencies)`
   - `⚠️  Would transitively update excluded dependency: <dep>`

#### GitHub Actions Step Summary

The step summary provides a clear, formatted overview:

```markdown
## 🔒 Security Dependency Update Summary

**Branch:** `release-2.7`

### ⚙️ Active Exclusions for This Branch

**Directly excluded packages** (skipped by configuration):

- `helm.sh/helm/v3`

**Transitive exclusion patterns** (updates blocked if they would pull these in):

- `k8s.io/`

### ✅ Updated Packages (3)

| Package | Old Version | New Version |
|---------|-------------|-------------|
| `package1` | `v1.2.2` | `v1.2.3` |
| `package2` | `v1.9.5` | `v2.0.0` |
| `stdlib` | `1.20.0` | `1.21.0` |

### 🔄 Transitive Dependencies Updated (5)

| Package | Old Version | New Version | Updated By |
|---------|-------------|-------------|------------|
| `github.com/dep1` | `v0.5.0` | `v0.6.0` | `package1` |
| `github.com/dep2` | `v1.1.0` | `v1.2.0` | `package1` |
| `github.com/dep3` | `v2.0.0` | `v2.1.0` | `package2` |
| `golang.org/x/net` | `v0.15.0` | `v0.17.0` | `package2` |
| `golang.org/x/sys` | `v0.12.0` | `v0.13.0` | `stdlib` |

### ⏭️ Skipped - Transitive Exclusions (1)

| Package | Current Version | Target Version | Blocked By |
|---------|-----------------|----------------|------------|
| `helm.sh/helm/v3` | `v3.13.0` | `v3.14.0` | `k8s.io/api`, `k8s.io/client-go` |

### 📊 Summary

| Status | Count |
|--------|-------|
| ✅ Updated | 3 |
| ⏭️ Skipped (Direct) | 0 |
| ⏭️ Skipped (Transitive) | 1 |
| ⚠️ Skipped (Failed) | 0 |
| **Total** | **4** |
```

### Benefits

- **Prevents PR Limbo**: No more PRs that can't be merged but keep getting recreated
- **Allows Partial Updates**: Security fixes for simple dependencies can proceed while complex ones are handled manually
- **Per-Branch Control**: Different release branches can have different policies
- **Clear Visibility**: Exclusions are logged in workflow output

## Manual Usage

You can run the tool manually:

```bash
# Without exclusions
make update-vulnerable-dependencies

# With exclusions (inline JSON)
export SECURITY_UPDATE_CONFIG='{"master": {"exclude_direct": ["pkg1"]}}'
make update-vulnerable-dependencies

# With exclusions (fetch from GitHub repo variables using gh CLI)
export SECURITY_UPDATE_CONFIG=$(gh variable get SECURITY_UPDATE_CONFIG)
make update-vulnerable-dependencies
```

**Note:** The `gh variable get` command requires:
- GitHub CLI (`gh`) installed
- Appropriate repository permissions
- The variable to be configured in the repository settings

## Files

**Core Components:**
- `update-vulnerable-dependencies.sh` - Main orchestration script
- `lib.sh` - Shared library with reusable functions (pattern matching, config extraction)

**Version Calculation:**
- `main.go` - CLI wrapper for the least common version calculator
- `leastcommonversion/leastcommonversion.go` - Least common version algorithm implementation
- `leastcommonversion/leastcommonversion_test.go` - Go tests for the algorithm

**Test Suite:**
- `test-lib.sh` - Unit tests for shared library functions (17 tests)
- `test-exclusions.sh` - Integration tests for exclusion logic (11 tests)

## Testing

The tool includes comprehensive test coverage:

### Unit Tests (test-lib.sh)

Tests the core library functions in isolation:

```bash
./tools/ci/update-vulnerable-dependencies/test-lib.sh
```

Coverage:
- Pattern matching (exact matches, prefix matches, multiple patterns)
- Configuration extraction (valid/invalid branches, missing configs, malformed JSON)
- **17 test cases** covering edge cases and error scenarios

### Integration Tests (test-exclusions.sh)

Tests the end-to-end exclusion logic:

```bash
./tools/ci/update-vulnerable-dependencies/test-exclusions.sh
```

Coverage:
- Direct exclusion patterns
- Transitive exclusion patterns
- Multi-branch scenarios
- **11 test cases** covering real-world usage patterns

### Run All Tests

```bash
# Run both test suites
./tools/ci/update-vulnerable-dependencies/test-lib.sh && \
  ./tools/ci/update-vulnerable-dependencies/test-exclusions.sh
```

## Requirements

- `osv-scanner` - Install with: `go install github.com/google/osv-scanner/cmd/osv-scanner@latest`
- `jq` - JSON processor
- `git` - For branch detection and file restoration

## Troubleshooting

### Package still showing in scans after update

If a package appears in the OSV scan but isn't being updated, it might be because:
- No fixed version exists yet
- The package is excluded by configuration
- The update would pull in excluded transitive dependencies

Check the workflow logs for skip messages.

### Configuration not being applied

1. Verify the variable name is exactly `SECURITY_UPDATE_CONFIG`
2. Ensure the JSON is valid (use a JSON validator)
3. Check that the branch name in the config matches exactly (e.g., `release-2.7` not `origin/release-2.7`)
4. Review the workflow logs for configuration parsing errors
