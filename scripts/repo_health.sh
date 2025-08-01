#!/bin/bash

set -e

desired_golang_version() {
    local top_level
    top_level=$(git rev-parse --show-toplevel)
    cat "$top_level/GOLANG_VERSION"
}

export GOPROXY="https://proxy.golang.org,direct" 
export GOSUMDB="sum.golang.org"

hab pkg install -b core/git core/ruby/3.0.6/20240108025751 core/jq-static core/shellcheck/0.8.0/20240108154129 core/cacerts

git config --global --add safe.directory /go/src/github.com/chef/automate
git config --global --add safe.directory '*'
hab pkg install -b "core/go1_24/1.24.4"

echo "Checking Go Dependencies And Vendored Protos"
go mod verify
scripts/revendor.sh # Revendor extracts vendored protos to protovendor/
git diff --exit-code --ignore-submodules=all # fail if anything's been changed

echo "Checking automate-deployment binding data"
(
  cd ./components/automate-deployment
  make check-bindings
)

yml2json() {
  ruby -ryaml -rjson -e 'puts JSON.pretty_generate(YAML.load(ARGF))' "$1"
}

echo "Checking if Golang license fallbacks/exceptions are needed"

# Convert the YAML to JSON
license_scout_json=$(yml2json .license_scout.yml)

# Get the list of Golang packages marked as exceptions
exceptions=$(echo "$license_scout_json" | jq -r '.exceptions.golang // [] | .[].name')

# Loop through the fallbacks and exceptions
for d in $(echo "$license_scout_json" | jq -r '.fallbacks.golang // [] | .[].name'); do
    # Check if the package is in go.sum in any relevant folders
    if ! grep -q "$d" go.sum && ! grep -q "$d" protovendor/go.sum && ! grep -q "$d" api/external/go.sum; then
        # Check if it's an exception
        if echo "$exceptions" | grep -q "$d"; then
            echo "Skipping exception for dependency \"$d\""
        else
            echo "License_scout exception for dependency \"$d\" not required anymore"
            exit 1
        fi
    fi
done

echo "Checking for up-to-date bldr config"
go run ./tools/bldr-config-gen
if ! git diff --exit-code --ignore-submodules=all; then
    echo "The bldr config appears to be out of date!"
    echo "To fix this, run:"
    echo ""
    echo "   hab studio run \"source .studiorc && generate_bldr_config\""
    echo ""
    echo "Inspect and commit the resulting changes if they look reasonable"
    exit 1
fi

echo "Shellchecking!"
shellcheck -s bash -ax \
  .expeditor/*.sh \
  .expeditor/**/*.sh \
  .buildkite/hooks/* \
  scripts/*.sh \
  integration/**/*.sh

# Applying shellcheck to the studio scripts is still in-progress.  To
# help, choose one of the violations excluded here and fix all
# instances of it.
shellcheck -s bash -ax \
  -e SC2012,SC2086,SC2128,SC2164,SC2181,SC2207 \
  .studiorc .studio/*

echo "Checking for possible credentials in the source code"
go run ./tools/credscan