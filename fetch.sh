#!/usr/bin/env bash

set -e -u -o pipefail

rm -rf box3d-c
git clone https://github.com/erincatto/box3d.git box3d-c
git -C box3d-c checkout -q f36d061

git -C box3d-c apply < box3d.patch
