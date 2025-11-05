#!/usr/bin/env bash

# Shared library for update-vulnerable-dependencies scripts
# Contains common functions to avoid duplication

# Check if a package matches any pattern in a JSON array of exclusion patterns
# Args:
#   $1: package name to check
#   $2: JSON array of patterns (e.g., '["k8s.io/", "helm.sh/helm/v3"]')
# Returns:
#   0 if package matches any pattern, 1 otherwise
matches_exclusion_pattern() {
  local package="$1"
  local patterns_json="$2"

  echo "$patterns_json" |
    jq -e --arg pkg "$package" \
      'any(. as $pattern | $pkg == $pattern or ($pkg | startswith($pattern)))' \
      >/dev/null 2>&1
}

# Extract exclusion configuration for a specific branch
# Args:
#   $1: config JSON
#   $2: branch name
#   $3: config key (e.g., "exclude_direct" or "exclude_transitive")
# Outputs: JSON array of exclusions
extract_branch_config() {
  local config_json="$1"
  local branch="$2"
  local key="$3"

  echo "$config_json" |
    jq -c --arg branch "$branch" --arg key "$key" \
      '.[$branch][$key] // []' 2>/dev/null ||
    echo "[]"
}

