load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "parse.go",
        "readfile.go",
        "sysreadfile.go",
        "sysreadfile_compat.go",
        "valueparser.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/prometheus/procfs/internal/util",
    importpath = "github.com/prometheus/procfs/internal/util",
    visibility = ["//vendor/github.com/prometheus/procfs:__subpackages__"],
)
