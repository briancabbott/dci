load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "semver",
    srcs = [
        "semver.go",
        "sort.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/coreos/go-semver/semver",
    importpath = "github.com/coreos/go-semver/semver",
    visibility = ["//visibility:public"],
)
