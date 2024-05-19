#!/bin/sh

set -e

tmpFile=$(mktemp)

go build -o "$tmpFile" .)

exec "$tmpFile"
