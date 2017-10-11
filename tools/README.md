# Automatically update deb_package rules in the WORKSPACE file
Similar to the `gazelle` tool which helps with managing golang bazel rules, it is possible to run this helper program by running `bazel run update_deb_packages`.

Add the following to the `BUILD` or `BUILD.bazel` file in the root directory of your repository:

```bzl
load("@rules_pkg//tools:update_deb_packages.bzl", "update_deb_packages")

update_deb_packages(
    name = "update_deb_packages",
    pgp_keys = ["@rule_name_of_http_file_rule_of_pgp_key//file"],
)
```

The `pgp_keys` list must contain all `http_file` rules that are used in the `pgp_key` portion of the `deb_packages` rules in your `WORKSPACE` file.
Referring to them is necessary, since otherwise these files wouldn't actually be downloaded by Bazel before executing the tool.

This repository also contains the `gazelle` boilerplate in the root `BUILD` file, since the `update_deb_packages` tool is written in go and gazelle helps with automatically generating `BUILD` files for the tool's dependencies.

Then you can run `bazel run update_deb_packages` and it will automatically add missing packages and update hashes and paths of the new and existing ones in your `WORKSPACE` file.

## Behind the scenes

This rule works very similar to the [gazelle](https://github.com/bazelbuild/rules_go/blob/master/go/private/tools/gazelle.bzl) rules ([stable link](https://github.com/bazelbuild/rules_go/blob/ee1fef7ec1379fcf36c002fd3ac0d00d940b147e/go/private/tools/gazelle.bzl)) to execute the `gazelle` tool using `bazel run`.

To escape the sandboxing and have direct access to the actual `WORKSPACE` and `BUILD` files, the small shell script resolves the softlink that Bazel creates into the build environment and operates at the root of the actual repository.

This still creates some challenges, as it is also necessary to have access to the PGP keys, which are back in the sandbox.
Moving them to the repository would be an option, but then it would need some reliable cleanup.

Instead, the tool itself uses the fact that the `bazel-<workspacename>` folder is also linked into the repository for convenience and looks for the key in there instead of the sandbox it came from.

As Bazel's sandboxing gets more sophisticated, it might be necessary to reevaluate this approach.
For now it works.
