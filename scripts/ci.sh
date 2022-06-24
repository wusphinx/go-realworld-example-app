#!/bin/bash
# This file modified from https://github.com/uptrace/opentelemetry-go-extra/blob/main/scripts/test.sh

set -e

PACKAGE_DIRS=$(find . -mindepth 2 -type f -name 'go.mod' -exec dirname {} \; \
  | sed 's/^\.\///' \
  | sort)

for dir in $PACKAGE_DIRS
do
    (
        echo "lint ${dir}..."
        cd $dir
        golangci-lint run 
        go test -v ./...
    )
done