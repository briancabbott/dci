load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "eventclock",
    srcs = [
        "interface.go",
        "real.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiserver/pkg/util/flowcontrol/fairqueuing/eventclock",
    importpath = "k8s.io/apiserver/pkg/util/flowcontrol/fairqueuing/eventclock",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/utils/clock"],
)
