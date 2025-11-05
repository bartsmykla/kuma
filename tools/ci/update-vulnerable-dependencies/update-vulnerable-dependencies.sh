#!/usr/bin/env bash

set -Eeuo pipefail

command -v osv-scanner >/dev/null 2>&1 || { echo >&2 "osv-scanner not installed!"; exit 1; }
command -v jq >/dev/null 2>&1 || { echo >&2 "jq not installed!"; exit 1; }

SCRIPT_PATH="${BASH_SOURCE[0]:-$0}";
SCRIPT_DIR="$(dirname -- "$SCRIPT_PATH")"

# Source shared library functions
# shellcheck source=lib.sh
source "$SCRIPT_DIR/lib.sh"

# Get current branch for configuration lookup
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "")

# Arrays to track update results for GitHub step summary
declare -a UPDATED_PACKAGES=()
declare -a SKIPPED_DIRECT=()
declare -a SKIPPED_TRANSITIVE=()
declare -a SKIPPED_FAILED=()

# Parse exclusion configuration from GitHub variable (JSON format)
# Expected format:
# {
#   "branch-name": {
#     "exclude_direct": ["package1", "package2"],
#     "exclude_transitive": ["prefix1/", "prefix2/"]
#   }
# }
EXCLUDE_DIRECT_JSON="[]"
EXCLUDE_TRANSITIVE_JSON="[]"

if [ -n "${SECURITY_UPDATE_CONFIG:-}" ] && [ -n "$CURRENT_BRANCH" ]; then
  echo "Loading exclusion configuration for branch: $CURRENT_BRANCH"
  EXCLUDE_DIRECT_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_direct")
  EXCLUDE_TRANSITIVE_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_transitive")

  if [ "$EXCLUDE_DIRECT_JSON" != "[]" ]; then
    echo "Direct exclusions: $EXCLUDE_DIRECT_JSON"
  fi
  if [ "$EXCLUDE_TRANSITIVE_JSON" != "[]" ]; then
    echo "Transitive exclusions: $EXCLUDE_TRANSITIVE_JSON"
  fi
fi

# Function to check if a package is directly excluded
is_excluded_direct() {
  local package="$1"
  matches_exclusion_pattern "$package" "$EXCLUDE_DIRECT_JSON"
}

# Function to check if an update would pull in excluded transitive dependencies
would_update_excluded_transitive() {
  local package="$1"
  local version="$2"
  local deps_before
  local deps_after
  local changed
  local excluded_found=false
  local dep_line
  local dep_name

  # Get current dependencies snapshot
  deps_before="$(go list -m all 2>/dev/null | sort)"

  # Try the update
  if ! GOTOOLCHAIN=auto go get -d "${package}@v${version}" 2>/dev/null; then
    # Update failed, restore and return error
    git restore go.mod go.sum 2>/dev/null || true
    return 2
  fi

  # Get new dependencies snapshot
  deps_after="$(go list -m all 2>/dev/null | sort)"

  # Find what changed (new or updated dependencies)
  changed="$(comm -13 <(echo "$deps_before") <(echo "$deps_after"))"

  # Check if any changed dependency matches excluded patterns
  while IFS= read -r dep_line; do
    if [ -z "$dep_line" ]; then
      continue
    fi
    # Extract package name (first field)
    dep_name="${dep_line%% *}"

    if matches_exclusion_pattern "$dep_name" "$EXCLUDE_TRANSITIVE_JSON"; then
      echo "  ⚠️  Would transitively update excluded dependency: $dep_line"
      excluded_found=true
    fi
  done <<< "$changed"

  # Revert the changes (always, since this is a dry-run check)
  git restore go.mod go.sum 2>/dev/null || true

  if [ "$excluded_found" = true ]; then
    return 0
  fi

  return 1
}

OSV_FLAGS=(--lockfile=go.mod --json)

# Add additional OSV scanner options if provided
# Check if OSV_SCANNER_ADDITIONAL_OPTS is set and is an array
if [ -n "${OSV_SCANNER_ADDITIONAL_OPTS+set}" ]; then
  for i in "${OSV_SCANNER_ADDITIONAL_OPTS[@]}"; do
    # Skip empty items
    if [ -n "$i" ]; then
      OSV_FLAGS+=("$i")
    fi
  done
fi

while IFS= read -r dep; do
  [ -z "$dep" ] && continue

  fixVersion="$(go run "$SCRIPT_DIR"/main.go <<< "$dep")"

  if [ "$fixVersion" != "null" ]; then
    package="$(jq -r .name <<< "$dep")"

    # Check if package is directly excluded
    if is_excluded_direct "$package"; then
      echo "⏭️  Skipping $package (directly excluded by configuration)"
      SKIPPED_DIRECT+=("$package -> v$fixVersion")
      continue
    fi

    # Check if update would pull in excluded transitive dependencies
    # Only check for transitive exclusions if we have any configured
    if [ "$EXCLUDE_TRANSITIVE_JSON" != "[]" ]; then
      would_update_excluded_transitive "$package" "$fixVersion"
      transitive_check_result=$?

      if [ $transitive_check_result -eq 0 ]; then
        echo "⏭️  Skipping $package -> v$fixVersion (would update excluded transitive dependencies)"
        SKIPPED_TRANSITIVE+=("$package -> v$fixVersion")
        continue
      elif [ $transitive_check_result -eq 2 ]; then
        echo "⚠️  Failed to check transitive dependencies for $package -> v$fixVersion, skipping"
        SKIPPED_FAILED+=("$package -> v$fixVersion")
        continue
      fi
    fi

    echo "✅ Updating $package to v$fixVersion"

    if [[ "$package" == "stdlib" ]]; then
      go mod edit -go="$fixVersion"
      UPDATED_PACKAGES+=("stdlib -> $fixVersion")
    else
      # Always use GOTOOLCHAIN=auto to allow downloading newer Go toolchain
      # when updating dependencies that require it (e.g., helm requiring Go 1.24+)
      GOTOOLCHAIN=auto go get "$package"@v"$fixVersion"
      UPDATED_PACKAGES+=("$package -> v$fixVersion")
    fi
  fi
done < <(osv-scanner "${OSV_FLAGS[@]}" | jq -c '.results[].packages[] | .package.name as $vulnerablePackage | {
  name: $vulnerablePackage,
  current: .package.version,
  fixedVersions: [.vulnerabilities[].affected[] | select(.package.name == $vulnerablePackage) | .ranges[].events |
  map(select(.fixed != null) | .fixed)] | map(select(length > 0)) } | select(.name != "github.com/kumahq/kuma")')

# Always use GOTOOLCHAIN=auto when running `go mod tidy` to allow downloading
# newer Go versions if `go.mod` was updated to require a newer version
# (either explicitly via go mod edit or implicitly via go get)
GOTOOLCHAIN=auto go mod tidy

# Generate GitHub Actions step summary if running in CI
if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
  {
    echo "## 🔒 Security Dependency Update Summary"
    echo ""
    echo "**Branch:** \`$CURRENT_BRANCH\`"
    echo ""

    # Show exclusion configuration if any
    if [ "$EXCLUDE_DIRECT_JSON" != "[]" ] || [ "$EXCLUDE_TRANSITIVE_JSON" != "[]" ]; then
      echo "### ⚙️ Active Exclusions for This Branch"
      echo ""
      if [ "$EXCLUDE_DIRECT_JSON" != "[]" ]; then
        echo "**Directly excluded packages** (skipped by configuration):"
        echo ""
        echo "$EXCLUDE_DIRECT_JSON" | jq -r '.[] | "- \u0060" + . + "\u0060"'
        echo ""
      fi
      if [ "$EXCLUDE_TRANSITIVE_JSON" != "[]" ]; then
        echo "**Transitive exclusion patterns** (updates blocked if they would pull these in):"
        echo ""
        echo "$EXCLUDE_TRANSITIVE_JSON" | jq -r '.[] | "- \u0060" + . + "\u0060"'
        echo ""
      fi
    fi

    # Show updated packages
    if [ ${#UPDATED_PACKAGES[@]} -gt 0 ]; then
      echo "### ✅ Updated Packages (${#UPDATED_PACKAGES[@]})"
      echo ""
      for pkg in "${UPDATED_PACKAGES[@]}"; do
        echo "- $pkg"
      done
      echo ""
    else
      echo "### ✅ Updated Packages"
      echo ""
      echo "_No packages were updated_"
      echo ""
    fi

    # Show skipped packages - direct exclusions
    if [ ${#SKIPPED_DIRECT[@]} -gt 0 ]; then
      echo "### ⏭️ Skipped - Direct Exclusions (${#SKIPPED_DIRECT[@]})"
      echo ""
      echo "These packages are directly excluded by configuration and were not updated:"
      echo ""
      for pkg in "${SKIPPED_DIRECT[@]}"; do
        echo "- $pkg"
      done
      echo ""
    fi

    # Show skipped packages - transitive exclusions
    if [ ${#SKIPPED_TRANSITIVE[@]} -gt 0 ]; then
      echo "### ⏭️ Skipped - Transitive Exclusions (${#SKIPPED_TRANSITIVE[@]})"
      echo ""
      echo "These packages were skipped because updating them would transitively update excluded dependencies:"
      echo ""
      for pkg in "${SKIPPED_TRANSITIVE[@]}"; do
        echo "- $pkg"
      done
      echo ""
      echo "> **Note:** Check the workflow logs for details on which transitive dependencies caused the exclusion."
      echo ""
    fi

    # Show failed checks
    if [ ${#SKIPPED_FAILED[@]} -gt 0 ]; then
      echo "### ⚠️ Skipped - Check Failed (${#SKIPPED_FAILED[@]})"
      echo ""
      echo "These packages were skipped because the transitive dependency check failed:"
      echo ""
      for pkg in "${SKIPPED_FAILED[@]}"; do
        echo "- $pkg"
      done
      echo ""
    fi

    # Summary statistics
    total_processed=$((${#UPDATED_PACKAGES[@]} + ${#SKIPPED_DIRECT[@]} + ${#SKIPPED_TRANSITIVE[@]} + ${#SKIPPED_FAILED[@]}))
    echo "### 📊 Summary"
    echo ""
    echo "| Status | Count |"
    echo "|--------|-------|"
    echo "| ✅ Updated | ${#UPDATED_PACKAGES[@]} |"
    echo "| ⏭️ Skipped (Direct) | ${#SKIPPED_DIRECT[@]} |"
    echo "| ⏭️ Skipped (Transitive) | ${#SKIPPED_TRANSITIVE[@]} |"
    echo "| ⚠️ Skipped (Failed) | ${#SKIPPED_FAILED[@]} |"
    echo "| **Total** | **$total_processed** |"
    echo ""

    if [ ${#SKIPPED_DIRECT[@]} -gt 0 ] || [ ${#SKIPPED_TRANSITIVE[@]} -gt 0 ]; then
      echo "---"
      echo ""
      echo "💡 **Tip:** Excluded packages need to be updated manually. Review the exclusion configuration if these packages should be auto-updated."
    fi
  } >> "$GITHUB_STEP_SUMMARY"
fi

echo ""
echo "Update process completed!"
echo "  Updated: ${#UPDATED_PACKAGES[@]} packages"
echo "  Skipped (direct): ${#SKIPPED_DIRECT[@]} packages"
echo "  Skipped (transitive): ${#SKIPPED_TRANSITIVE[@]} packages"
echo "  Skipped (failed): ${#SKIPPED_FAILED[@]} packages"
