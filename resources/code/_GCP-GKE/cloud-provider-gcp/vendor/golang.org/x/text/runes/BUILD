load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "runes",
    srcs = [
        "cond.go",
        "runes.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/text/runes",
    importpath = "golang.org/x/text/runes",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/text/transform"],
)
