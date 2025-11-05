#!/usr/bin/env bash

# Test script for exclusion configuration logic
# This tests the jq-based pattern matching without requiring osv-scanner or actual vulnerabilities

set -Eeuo pipefail

# Get script directory and source library
SCRIPT_DIR="$(dirname -- "${BASH_SOURCE[0]}")"
# shellcheck source=lib.sh
source "$SCRIPT_DIR/lib.sh"

echo "Testing exclusion configuration logic..."
echo

# Test configuration
export SECURITY_UPDATE_CONFIG='{
  "release-2.7": {
    "exclude_direct": ["helm.sh/helm/v3", "github.com/example/"],
    "exclude_transitive": ["k8s.io/"]
  },
  "master": {
    "exclude_direct": [],
    "exclude_transitive": []
  }
}'

# Simulate current branch
export CURRENT_BRANCH="release-2.7"

# Parse configuration using library functions
EXCLUDE_DIRECT_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_direct")
EXCLUDE_TRANSITIVE_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_transitive")

echo "Branch: $CURRENT_BRANCH"
echo "Direct exclusions: $EXCLUDE_DIRECT_JSON"
echo "Transitive exclusions: $EXCLUDE_TRANSITIVE_JSON"
echo

# Run tests
test_count=0
pass_count=0

# Unified test function that works with any exclusion pattern JSON
# Args:
#   $1: test name
#   $2: package to test
#   $3: expected result ("excluded" or "allowed")
#   $4: exclusion patterns JSON
run_exclusion_test() {
  local test_name="$1"
  local package="$2"
  local expected="$3"
  local patterns_json="$4"
  local result

  test_count=$((test_count + 1))

  if matches_exclusion_pattern "$package" "$patterns_json"; then
    result="excluded"
  else
    result="allowed"
  fi

  if [ "$result" = "$expected" ]; then
    echo "✅ PASS: $test_name"
    pass_count=$((pass_count + 1))
  else
    echo "❌ FAIL: $test_name (expected: $expected, got: $result)"
  fi
}

echo "Running direct exclusion tests:"
echo

run_exclusion_test "Exact match - helm" "helm.sh/helm/v3" "excluded" "$EXCLUDE_DIRECT_JSON"
run_exclusion_test "Prefix match - example package" "github.com/example/foo" "excluded" "$EXCLUDE_DIRECT_JSON"
run_exclusion_test "Non-excluded package" "github.com/other/package" "allowed" "$EXCLUDE_DIRECT_JSON"
run_exclusion_test "Similar but not matching" "helm.sh/helm/v4" "allowed" "$EXCLUDE_DIRECT_JSON"

echo
echo "Running transitive exclusion pattern tests:"
echo

run_exclusion_test "k8s.io/api" "k8s.io/api" "excluded" "$EXCLUDE_TRANSITIVE_JSON"
run_exclusion_test "k8s.io/client-go" "k8s.io/client-go" "excluded" "$EXCLUDE_TRANSITIVE_JSON"
run_exclusion_test "k8s.io/apimachinery" "k8s.io/apimachinery" "excluded" "$EXCLUDE_TRANSITIVE_JSON"
run_exclusion_test "github.com/kubernetes/kubernetes" "github.com/kubernetes/kubernetes" "allowed" "$EXCLUDE_TRANSITIVE_JSON"

echo
echo "Testing with branch that has no exclusions:"
echo

CURRENT_BRANCH="master"
EXCLUDE_DIRECT_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_direct")
EXCLUDE_TRANSITIVE_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_transitive")

echo "Branch: $CURRENT_BRANCH"
echo "Direct exclusions: $EXCLUDE_DIRECT_JSON"
echo "Transitive exclusions: $EXCLUDE_TRANSITIVE_JSON"
echo

run_exclusion_test "helm.sh on master" "helm.sh/helm/v3" "allowed" "$EXCLUDE_DIRECT_JSON"
run_exclusion_test "k8s.io on master" "k8s.io/api" "allowed" "$EXCLUDE_TRANSITIVE_JSON"

echo
echo "Testing with non-existent branch:"
echo

CURRENT_BRANCH="non-existent"
EXCLUDE_DIRECT_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_direct")
EXCLUDE_TRANSITIVE_JSON=$(extract_branch_config "$SECURITY_UPDATE_CONFIG" "$CURRENT_BRANCH" "exclude_transitive")

echo "Branch: $CURRENT_BRANCH"
echo "Direct exclusions: $EXCLUDE_DIRECT_JSON"
echo "Transitive exclusions: $EXCLUDE_TRANSITIVE_JSON"
echo

run_exclusion_test "helm.sh on non-existent branch" "helm.sh/helm/v3" "allowed" "$EXCLUDE_DIRECT_JSON"

echo
echo "================================"
echo "Test Results: $pass_count/$test_count passed"
echo "================================"

if [ $pass_count -eq $test_count ]; then
  echo "✅ All tests passed!"
  exit 0
else
  echo "❌ Some tests failed"
  exit 1
fi
