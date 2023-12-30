#!/usr/bin/env bash
set -e

cd $1

echo -n "Running tests "
function testrun {
    sudo -E bash -c "umask 0; PATH=$PATH go test $@"
}
if [ ! -z "${COVERALLS:-""}" ]; then
    # coverage profile only works per-package
    PKGS="$(go list ./... | xargs echo)"
    echo "with coverage profile generation..."
    i=0
    for t in ${PKGS}; do
        testrun "-covermode set -coverprofile ${i}.coverprofile ${t}"
        i=$((i+1))
    done
else
    echo "without coverage profile generation..."
    testrun "./..."
fi

