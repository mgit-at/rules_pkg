# Documentation for `gcr.io/distroless/base`

## Image Contents

This image contains a minimal Linux, glibc-based system. It is intended for use directly by "mostly-statically compiled" languages like Go, Rust or D.

The image contains:

* glibc
* libssl
* openssl
* ca-certificates
* A /etc/passwd entry for a root user
* A /tmp directory

## Usage

Users are expected to include their compiled application and set the correct cmd in their image.



## Origin

This image contains the same stuff as the distroless base image at https://github.com/GoogleCloudPlatform/distroless/commit/494d25e5f55ed8d14fbc66934e53a0a9cb2309b9
