load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "controller",
    srcs = ["interfaces.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/controller-manager/controller",
    importpath = "k8s.io/controller-manager/controller",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/controller-manager/pkg/healthz"],
)
