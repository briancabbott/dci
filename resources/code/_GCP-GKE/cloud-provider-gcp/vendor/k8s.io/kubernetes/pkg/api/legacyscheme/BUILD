load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "legacyscheme",
    srcs = ["scheme.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/kubernetes/pkg/api/legacyscheme",
    importpath = "k8s.io/kubernetes/pkg/api/legacyscheme",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer",
    ],
)
