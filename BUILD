load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "somefunc",
    srcs = ["somefunc.go"],
    importpath = "some.import.path",
    deps = ["@org_golang_x_sys//unix"],
    linkmode = "c-archive",
    cgo = True,
)


cc_binary(
    name = "caller",
    srcs = ["caller.cc"],
    deps = [":somefunc.cc"]
)
