load("@rules_oci//oci:defs.bzl", "oci_image", "oci_image_index", "oci_push", "oci_tarball")
load("@rules_pkg//:pkg.bzl", "pkg_tar")
load("//tools:multi_arch.bzl", "multi_arch")

def prysm_image_upload(
        name,
        binary,
        entrypoint,
        symlinks,
        repository):
    pkg_tar(
        name = "binary_tar",
        srcs = [binary],
        symlinks = symlinks,
    )

    oci_image(
        name = "oci_image",
        base = "@linux_debian11_multiarch_base",
        entrypoint = entrypoint,
        tars = [
            "//tools:passwd_tar",
        ] + select({
          "@platforms//cpu:x86_64": [
            "@amd64_debian11_bash",
            "@amd64_debian11_libtinfo6",
            "@amd64_debian11_coreutils",
            "@amd64_debian11_libacl1",
            "@amd64_debian11_libattr1",
            "@amd64_debian11_libselinux",
            "@amd64_debian11_libpcre2",
          ],
          "@platforms//cpu:arm64": [
            "@arm64_debian11_bash",
            "@arm64_debian11_libtinfo6",
            "@arm64_debian11_coreutils",
            "@arm64_debian11_libacl1",
            "@arm64_debian11_libattr1",
            "@arm64_debian11_libselinux",
            "@arm64_debian11_libpcre2",
          ],
        }) + [
            ":binary_tar",
        ],
        labels = {
          "org.opencontainers.image.source": "https://gitlab.com/pulsechaincom/prysm-pulse",
        },
    )

    multi_arch(
        name = "oci_multiarch",
        image = ":oci_image",
        platforms = [
            "@io_bazel_rules_go//go/toolchain:linux_amd64_cgo",
            "@io_bazel_rules_go//go/toolchain:linux_arm64_cgo",
        ],
    )

    oci_image_index(
        name = "oci_image_index",
        images = [
            ":oci_multiarch",
        ],
    )
    
    # oci_push rule is used to push the multi-platform image to the registry
    # must pass a docker image --tag
    oci_push(
        name = name,
        image = ":oci_image_index",
        repository = repository,
    )

    # oci_tarball rule is used to create an image that can be loaded by the local docker runtime
    oci_tarball (
        name = "oci_tarball",
        image = ":oci_image",
        repo_tags = [
            repository + ":local"
        ],
    )
