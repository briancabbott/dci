load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "keepalive",
    srcs = ["keepalive.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/grpc/keepalive",
    importpath = "google.golang.org/grpc/keepalive",
    visibility = ["//visibility:public"],
)
