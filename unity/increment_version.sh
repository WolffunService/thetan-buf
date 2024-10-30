#!/bin/bash

cd "$(dirname "$0")" || exit

# Read the current version from package.json
current_version=$(jq -r '.version' package.json)

# Increment the version number
new_version=$(echo "$current_version" | awk -F. -v OFS=. '{$NF++;print}')

# Update the version in package.json
jq --arg new_version "$new_version" '.version = $new_version' package.json > tmp.json && mv tmp.json package.json

echo "Version incremented to $new_version"
echo "version=${new_version}" >> $GITHUB_OUTPUT
