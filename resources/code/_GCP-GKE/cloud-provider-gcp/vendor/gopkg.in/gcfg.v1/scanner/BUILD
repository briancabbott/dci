load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "scanner",
    srcs = [
        "errors.go",
        "scanner.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/gopkg.in/gcfg.v1/scanner",
    importpath = "gopkg.in/gcfg.v1/scanner",
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/gcfg.v1/token"],
)
