#!/bin/bash

################################################################
# This is a convenience script for devs which runs bazel within 
# docker to build and load local docker images.
# - Removes dependency requirements (other than docker)
# - Uses mounted directories for bazel caching (fast builds)
# - Does NOT push images to remote repository
################################################################

# exit on failures
set -e

# build the prysm-builder image
docker build --platform linux/amd64 -f Dockerfile.builder -t prysm-builder .

# run the build
docker run --platform linux/amd64 -it -v $(pwd):/app -v $(pwd)/.bazel-in-docker-cache:/root/.cache prysm-builder

# source the outputs file to build the correct symlink directories: bazel-*
source .bazel-in-docker-outputs
rm .bazel-in-docker-outputs

# load new images into docker
docker load -i bazel-bin/cmd/beacon-chain/oci_tarball/tarball.tar
docker load -i bazel-bin/cmd/validator/oci_tarball/tarball.tar
docker load -i bazel-bin/cmd/prysmctl/oci_tarball/tarball.tar

# finish up
echo ""
echo "Build Complete! Bazel outputs are available in the following directories:"
ls | grep bazel- | tr '\n' '\t'
echo ""

# print new docker images
echo ""
echo "Docker images loaded:"
docker image ls | grep --color=never -E "REPOSITORY|prysm-pulse.* local"
