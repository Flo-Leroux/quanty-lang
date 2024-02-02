#!/usr/bin/env bash
set -eu

cd "$(dirname "$0")/.."

known_failures="$(cat scripts/known-failures.txt)"

# shellcheck disable=2046
tree-sitter parse -q \
	'examples/**/*.qy'
	$(for failure in $known_failures; do echo "!${failure}"; done)

example_count=$(find examples -name '*.qy' | wc -l)
failure_count=$(wc -w <<<"$known_failures")

success_count=$((example_count - failure_count))
success_percent=$(bc -l <<<"100*${success_count}/${example_count}")

printf \
	"Successfully parsed %d of %d example files (%.2f)\n" \
	$success_count "$example_count" "${success_percent/./,}"