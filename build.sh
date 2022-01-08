#!/bin/bash

set -ex

KV=`curl -sS https://api.github.com/repos/kubernetes-sigs/kustomize/releases | grep https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize/ | head -n 1 | awk -F '"' '{print $4}' | awk -F '/' '{print $9}'`
docker build -t chaunceyshannon/kustomize:$KV .
# docker push chaunceyshannon/kustomize:$KV