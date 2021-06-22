workspace(name="rules_go_empty_header")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Buggy version
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

# Working version, patch from https://github.com/CodeIntelligenceTesting/odict/commit/860723ea166539334fd9f3df059c833c70b1fb07
#http_archive(
#    name = "io_bazel_rules_go",
#    sha256 = "e3a80175a7553bb338987b57edb9bc0167a226c61dead15b4707f6ffcdb86dc4",
#    strip_prefix = "rules_go-06ef45c0dbdf73c034e527300a02234848a1347b",
#    urls = [
#        "https://github.com/bazelbuild/rules_go/archive/06ef45c0dbdf73c034e527300a02234848a1347b.zip",
#    ],
#)



load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.14")
