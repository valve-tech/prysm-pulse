# bazel-in-docker builder, for use with the build-in-docker.sh script and as a base image in pipelines
FROM ubuntu:22.04 as builder

# install build dependencies
RUN apt update && apt install -y cmake git libssl-dev libgmp-dev libtinfo5 unzip curl g++ python3 ca-certificates xz-utils

# install protoc at specific required version
RUN PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# install bazelisk
RUN curl -o bazelisk -L https://github.com/bazelbuild/bazelisk/releases/download/v1.14.0/bazelisk-linux-amd64 && \
    chmod +x bazelisk && \
    mv bazelisk /usr/local/bin && \
    ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel && \
    bazelisk --version

# add the official docker certificates (for installing docker-cli)
RUN install -m 0755 -d /etc/apt/keyrings && \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc && \
    chmod a+r /etc/apt/keyrings/docker.asc

# add the official docker repository to apt sources
RUN echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" \
    | tee /etc/apt/sources.list.d/docker.list > /dev/null

# install docker-cli for use in pipeline builds
RUN apt update && apt install -y docker-ce-cli

# create app placeholder
RUN mkdir /app
WORKDIR /app

# prevent git warnings throughout build process
RUN git config --global --add safe.directory /app

# this command is only run in local builds when using the build-in-docker script
# perform the build then print output directory mappings
CMD bazel build --config=release //cmd/beacon-chain:oci_tarball && \
    bazel build --config=release //cmd/validator:oci_tarball && \
    bazel build --config=release //cmd/prysmctl:oci_tarball && \
    ls -l | grep -o bazel-.* | sed 's/\/root\/.cache/.bazel-in-docker-cache/g' | awk '{print "ln -sf " $3 " " $1}' > .bazel-in-docker-outputs && \
    chmod 666 .bazel-in-docker-outputs
