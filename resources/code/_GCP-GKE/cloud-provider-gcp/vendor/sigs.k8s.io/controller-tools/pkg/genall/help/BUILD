load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "help",
    srcs = [
        "doc.go",
        "sort.go",
        "types.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/sigs.k8s.io/controller-tools/pkg/genall/help",
    importpath = "sigs.k8s.io/controller-tools/pkg/genall/help",
    visibility = ["//visibility:public"],
    deps = ["//vendor/sigs.k8s.io/controller-tools/pkg/markers"],
)
