load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "stats",
    srcs = [
        "handlers.go",
        "stats.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/grpc/stats",
    importpath = "google.golang.org/grpc/stats",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/grpc/metadata"],
)
