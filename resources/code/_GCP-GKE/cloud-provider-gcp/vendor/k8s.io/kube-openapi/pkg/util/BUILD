load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "trie.go",
        "util.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/kube-openapi/pkg/util",
    importpath = "k8s.io/kube-openapi/pkg/util",
    visibility = ["//visibility:public"],
)
