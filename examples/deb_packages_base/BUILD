package(default_visibility = ["//visibility:public"])

# Extracting the ca-certificates deb package
sh_binary(
    name = "extract_certs",
    srcs = ["extract.sh"],
)

# For convenience, rename busybox-x86_64 to busybox.
genrule(
    name = "busybox_tar",
    srcs = ["@busybox//file"],
    outs = ["busybox.tar"],
    cmd = "mkdir busybox; cp $(<) busybox/busybox; cd busybox; \
           for cmd in $$(./busybox --list); do \
             ln -s busybox $$cmd; \
           done; \
           cd ..; \
           tar -cf $@ busybox",
)

load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build")
load("@io_bazel_rules_docker//contrib:passwd.bzl", "passwd_file")
load("@runtimes_common//structure_tests:tests.bzl", "structure_test")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

# Create a default passwd_file rule.

passwd_file(
    name = "base_passwd",
    gid = 0,
    info = "root",
    uid = 0,
    username = "root",
)

load("//examples/deb_packages_base:cacerts.bzl", "cacerts")

cacerts(
    name = "cacerts",
)

load("@debian_jessie_amd64//debs:deb_packages.bzl", "debian_jessie_amd64")

# Create /tmp, too many things assume it exists.
# tmp.tar has a /tmp with the correct permissions 01777
# A tar is needed because at the moment there is no way to create a diretory with specific permissions

docker_build(
    name = "base",
    debs = [
        debian_jessie_amd64["libc6"],
        debian_jessie_amd64["libssl1.0.0"],
        debian_jessie_amd64["openssl"],
        debian_jessie_amd64["netbase"],
        debian_jessie_amd64["tzdata"],
    ],
    env = {"PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
    tars = [
        ":base_passwd.passwd.tar",
        ":cacerts.tar",
        ":tmp.tar",
        # os-release files are not implemented in deb_packages
        #"@debian_jessie//file:os_release.tar",
    ],
)

# A debug image with busybox available.
docker_build(
    name = "debug",
    base = ":base",
    directory = "/",
    entrypoint = ["/busybox/sh"],
    env = {"PATH": "$PATH:/busybox"},
    tars = [":busybox_tar"],
)

structure_test(
    name = "debug_test",
    config = "testdata/debug.yaml",
    image = ":debug",
)

structure_test(
    name = "base_test",
    config = "testdata/base.yaml",
    image = ":base",
)

go_binary(
    name = "check_certs",
    srcs = ["testdata/check_certs.go"],
)

docker_build(
    name = "check_certs_image",
    base = "//examples/deb_packages_base:base",
    files = [":check_certs"],
    visibility = ["//visibility:private"],
)

structure_test(
    name = "certs_test",
    config = "testdata/certs.yaml",
    image = ":check_certs_image",
)
