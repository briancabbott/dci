load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "names",
    srcs = ["generate.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiserver/pkg/storage/names",
    importpath = "k8s.io/apiserver/pkg/storage/names",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/apimachinery/pkg/util/rand"],
)
