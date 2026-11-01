#!/usr/bin/env bash

set -e -u -o pipefail

rm -rf box3d-c
git clone https://github.com/erincatto/box3d.git box3d-c
git -C box3d-c checkout -q 336263320469e5e9febfe18697ab562488f7e97a

git -C box3d-c apply < box3d.patch
