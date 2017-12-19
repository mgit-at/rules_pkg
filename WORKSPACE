workspace(name = "rules_pkg")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "90bb270d0a92ed5c83558b2797346917c46547f6f7103e648941ecdb6b9d0e72",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.1/rules_go-0.8.1.tar.gz",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()

go_register_toolchains()

# Go dependencies of the update_deb_packages helper tool

# "golang.org/x/crypto/openpgp"
go_repository(
    name = "org_golang_x_crypto",
    commit = "d585fd2cc9195196078f516b69daff6744ef5e84",
    importpath = "golang.org/x/crypto",
)

# "github.com/knqyf263/go-deb-version"
go_repository(
    name = "com_github_knqyf263_go_deb_version",
    commit = "9865fe14d09b1c729188ac810466dde90f897ee3",
    importpath = "github.com/knqyf263/go-deb-version",
)

# "github.com/stapelberg/godebiancontrol"
go_repository(
    name = "com_github_stapelberg_godebiancontrol",
    commit = "4376b22fb2c4dfda546c972f686310af907819b2",
    importpath = "github.com/stapelberg/godebiancontrol",
)

# Example for using the deb_packages ruleset
load("//deb_packages:deb_packages.bzl", "deb_packages")

# The Debian jessie archive signing key
# Source: https://ftp-master.debian.org/keys.html
# Full fingerprint: 126C 0D24 BD8A 2942 CC7D F8AC 7638 D044 2B90 D010
http_file(
    name = "jessie_archive_key",
    # It is highly recommended to use the sha256 hash of the key file to make sure it is untampered
    sha256 = "e42141a829b9fde8392ea2c0e329321bb29e5c0453b0b48e33c9f88bdc4873c5",
    urls = ["https://ftp-master.debian.org/keys/archive-key-8.asc"],
)

# The Debian jessie security archive signing key
# Source: https://ftp-master.debian.org/keys.html
# Full fingerprint: D211 6914 1CEC D440 F2EB 8DDA 9D6D 8F6B C857 C906
http_file(
    name = "jessie_security_archive_key",
    # It is highly recommended to use the sha256 hash of the key file to make sure it is untampered
    sha256 = "d05815c66deb71a595279b750aaf06370b6ad8c3b373651473c1c4b3d7da8f3c",
    urls = ["https://ftp-master.debian.org/keys/archive-key-8-security.asc"],
)

# The Debian stretch archive signing key
# Source: https://ftp-master.debian.org/keys.html
# Full fingerprint: E1CF 20DD FFE4 B89E 8026 58F1 E0B1 1894 F66A EC98
http_file(
    name = "stretch_archive_key",
    # It is highly recommended to use the sha256 hash of the key file to make sure it is untampered
    sha256 = "33b6a997460e177804cc44c7049a19350c11034719219390b22887471f0a2b5e",
    urls = ["https://ftp-master.debian.org/keys/archive-key-9.asc"],
)

# The Debian stretch security archive signing key
# Source: https://ftp-master.debian.org/keys.html
# Full fingerprint: 6ED6 F5CB 5FA6 FB2F 460A E88E EDA0 D238 8AE2 2BA9
http_file(
    name = "stretch_security_archive_key",
    # It is highly recommended to use the sha256 hash of the key file to make sure it is untampered
    sha256 = "4adecda0885f192b82c19fde129ca9d991f937437835a058da355b352a97e7dc",
    urls = ["https://ftp-master.debian.org/keys/archive-key-9-security.asc"],
)

deb_packages(
    name = "debian_jessie_amd64",
    arch = "amd64",
    distro = "jessie",
    distro_type = "debian",
    mirrors = [
        "http://deb.debian.org/debian",
        # This ensures old states of this repository will build as long as the snapshot mirror works:
        "http://snapshot.debian.org/archive/debian/20171219T131415Z",
    ],
    packages = {
        "ca-certificates": "pool/main/c/ca-certificates/ca-certificates_20141019+deb8u3_all.deb",
        "libc6": "pool/main/g/glibc/libc6_2.19-18+deb8u10_amd64.deb",
        "libssl1.0.0": "pool/main/o/openssl/libssl1.0.0_1.0.1t-1+deb8u7_amd64.deb",
        "netbase": "pool/main/n/netbase/netbase_5.3_all.deb",
        "openssl": "pool/main/o/openssl/openssl_1.0.1t-1+deb8u7_amd64.deb",
        "tzdata": "pool/main/t/tzdata/tzdata_2017c-0+deb8u1_all.deb",
    },
    packages_sha256 = {
        "ca-certificates": "bd799f47f5ae3260b6402b1fe19fe2c37f2f4125afcd19327bf69a9cf436aeff",
        "libc6": "0a95ee1c5bff7f73c1279b2b78f32d40da9025a76f93cb67c03f2867a7133e61",
        "libssl1.0.0": "d99de2cdca54484d23badc5683c7211b3a191977272d9e5281837af863dcdd56",
        "netbase": "3979bdd40c5666ef9bf71a5391ba01ad38e264f2ec96d289993f2a0805616dd3",
        "openssl": "d0e1464148bb2d682ccdb6f433b27a6848e4d012e8bb8a61ed9f6ad708017640",
        "tzdata": "f53b963b533100380127a20922b4265412ca4cf8f8b21c66e07c4645b7845002",
    },
    pgp_key = "jessie_archive_key",
)

# For the debug image
http_file(
    name = "busybox",
    executable = True,
    sha256 = "b51b9328eb4e60748912e1c1867954a5cf7e9d5294781cae59ce225ed110523c",
    urls = ["https://busybox.net/downloads/binaries/1.27.1-i686/busybox"],
)

# Docker rules.
git_repository(
    name = "io_bazel_rules_docker",
    commit = "cdd259b3ba67fd4ef814c88070a2ebc7bec28dc5",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

# used for testing the examples
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

git_repository(
    name = "runtimes_common",
    remote = "https://github.com/GoogleCloudPlatform/runtimes-common.git",
    tag = "v0.1.0",
)
