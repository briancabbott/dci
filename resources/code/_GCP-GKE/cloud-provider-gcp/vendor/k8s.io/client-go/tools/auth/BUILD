load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "auth",
    srcs = ["clientauth.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/tools/auth",
    importpath = "k8s.io/client-go/tools/auth",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/client-go/rest"],
)
