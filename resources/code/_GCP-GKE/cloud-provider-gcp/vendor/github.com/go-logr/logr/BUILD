load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "logr",
    srcs = [
        "discard.go",
        "logr.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/go-logr/logr",
    importpath = "github.com/go-logr/logr",
    visibility = ["//visibility:public"],
)
