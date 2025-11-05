#!/usr/bin/env bash

# Unit tests for lib.sh functions

set -Eeuo pipefail

# Get script directory and source library
SCRIPT_DIR="$(dirname -- "${BASH_SOURCE[0]}")"
# shellcheck source=lib.sh
source "$SCRIPT_DIR/lib.sh"

echo "Testing lib.sh functions..."
echo

# Test counters
test_count=0
pass_count=0

# Helper function to run tests
run_test() {
  local test_name="$1"
  local expected_result="$2"
  shift 2
  local actual_result

  test_count=$((test_count + 1))

  # Run the command and capture result
  if "$@"; then
    actual_result="pass"
  else
    actual_result="fail"
  fi

  if [ "$actual_result" = "$expected_result" ]; then
    echo "✅ PASS: $test_name"
    pass_count=$((pass_count + 1))
  else
    echo "❌ FAIL: $test_name (expected: $expected_result, got: $actual_result)"
  fi
}

# Helper function to test string output
run_output_test() {
  local test_name="$1"
  local expected_output="$2"
  shift 2
  local actual_output

  test_count=$((test_count + 1))

  # Run the command and capture output
  actual_output=$("$@")

  if [ "$actual_output" = "$expected_output" ]; then
    echo "✅ PASS: $test_name"
    pass_count=$((pass_count + 1))
  else
    echo "❌ FAIL: $test_name"
    echo "  Expected: $expected_output"
    echo "  Got:      $actual_output"
  fi
}

echo "=== Testing matches_exclusion_pattern ==="
echo

# Test exact match
run_test "Exact match - single pattern" "pass" \
  matches_exclusion_pattern "helm.sh/helm/v3" '["helm.sh/helm/v3"]'

# Test prefix match
run_test "Prefix match - k8s.io/" "pass" \
  matches_exclusion_pattern "k8s.io/api" '["k8s.io/"]'

run_test "Prefix match - github.com/example/" "pass" \
  matches_exclusion_pattern "github.com/example/foo" '["github.com/example/"]'

# Test non-match
run_test "Non-match - different package" "fail" \
  matches_exclusion_pattern "github.com/other/package" '["helm.sh/helm/v3"]'

run_test "Non-match - similar but different" "fail" \
  matches_exclusion_pattern "helm.sh/helm/v4" '["helm.sh/helm/v3"]'

# Test empty patterns array
run_test "Empty patterns array" "fail" \
  matches_exclusion_pattern "any.package" '[]'

# Test multiple patterns
run_test "Multiple patterns - matches first" "pass" \
  matches_exclusion_pattern "k8s.io/api" '["k8s.io/", "helm.sh/"]'

run_test "Multiple patterns - matches second" "pass" \
  matches_exclusion_pattern "helm.sh/helm/v3" '["k8s.io/", "helm.sh/helm/v3"]'

run_test "Multiple patterns - matches neither" "fail" \
  matches_exclusion_pattern "github.com/other" '["k8s.io/", "helm.sh/"]'

echo
echo "=== Testing extract_branch_config ==="
echo

TEST_CONFIG='{
  "release-2.7": {
    "exclude_direct": ["pkg1", "pkg2"],
    "exclude_transitive": ["k8s.io/"]
  },
  "master": {
    "exclude_direct": [],
    "exclude_transitive": []
  }
}'

# Test extracting direct exclusions from existing branch
run_output_test "Extract direct exclusions from release-2.7" '["pkg1","pkg2"]' \
  extract_branch_config "$TEST_CONFIG" "release-2.7" "exclude_direct"

# Test extracting transitive exclusions from existing branch
run_output_test "Extract transitive exclusions from release-2.7" '["k8s.io/"]' \
  extract_branch_config "$TEST_CONFIG" "release-2.7" "exclude_transitive"

# Test extracting empty arrays
run_output_test "Extract empty direct exclusions from master" '[]' \
  extract_branch_config "$TEST_CONFIG" "master" "exclude_direct"

run_output_test "Extract empty transitive exclusions from master" '[]' \
  extract_branch_config "$TEST_CONFIG" "master" "exclude_transitive"

# Test non-existent branch (should return empty array)
run_output_test "Extract from non-existent branch" '[]' \
  extract_branch_config "$TEST_CONFIG" "non-existent" "exclude_direct"

# Test non-existent config key (should return empty array)
run_output_test "Extract non-existent config key" '[]' \
  extract_branch_config "$TEST_CONFIG" "release-2.7" "non_existent_key"

# Test empty config
run_output_test "Extract from empty config" '[]' \
  extract_branch_config '{}' "any-branch" "exclude_direct"

# Test malformed JSON (should gracefully return empty array)
run_output_test "Extract from malformed JSON" '[]' \
  extract_branch_config 'invalid json' "any-branch" "exclude_direct"

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
