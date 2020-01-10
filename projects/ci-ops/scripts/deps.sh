#!/usr/bin/env sh
ROOT="$(cd "$(dirname "$BASH_SOURCE[0]}")"; pwd -P)"
BIN_DIR=$PWD/bin

VERSION_FLY=5.7.2
VERSION_JQ=1.6
VERSION_KUBECTL=1.15.1
# waiting for release of https://github.com/kubernetes-sigs/kustomize/pull/1316
# hash is latest commit from this branch:
# https://github.com/tkellen/kustomize/tree/allinone
VERSION_KUSTOMIZE=e7fc0b8eda765f9f4a8627ed97ec77471155383e
VERSION_SPRUCE=1.21.0

URL_FLY=https://github.com/concourse/concourse/releases/download/v${VERSION_FLY}/fly-${VERSION_FLY}-linux-amd64.tgz
URL_JQ=https://github.com/stedolan/jq/releases/download/jq-${VERSION_JQ}/jq-linux64
URL_KUBECTL=https://storage.googleapis.com/kubernetes-release/release/v${VERSION_KUBECTL}/bin/linux/amd64/kubectl
URL_KUSTOMIZE=https://github.com/kubernetes-sigs/kustomize/releases/download/v${VERSION_KUSTOMIZE}/kustomize_${VERSION_KUSTOMIZE}_linux_amd64
URL_SPRUCE=https://github.com/geofffranks/spruce/releases/download/v${VERSION_SPRUCE}/spruce-linux-amd64

mkdir -p $BIN_DIR

# Create temporary directory to explode files for modification into.
TEMP_DIR=$(mktemp -d)
if [[ ! -e $TEMP_DIR ]]; then
  >&2 echo "Failed to create temp directory ($TEMP_DIR)."
  exit 1
fi

# Ensure we cleanup after ourselves.
trap "exit 1" HUP INT PIPE QUIT TERM
trap 'rm -rf "$TEMP_DIR"' EXIT

(
  cd $TEMP_DIR
  curl -sSL $URL_FLY | tar xzf -
  curl -sSL $URL_JQ > jq
  curl -sSL $URL_KUBECTL > kubectl
  #curl -sSL $URL_KUSTOMIZE > kustomize
  (
    git clone https://github.com/tkellen/kustomize k
    cd k
    git checkout ${VERSION_KUSTOMIZE}
    go build cmd/kustomize/main.go
  )
  mv k/main kustomize
  rm -rf k
  curl -sSL $URL_SPRUCE > spruce
  chmod -R +x *
  mv * $BIN_DIR
)
