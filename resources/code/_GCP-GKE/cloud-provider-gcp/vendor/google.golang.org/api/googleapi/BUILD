load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "googleapi",
    srcs = [
        "googleapi.go",
        "types.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/api/googleapi",
    importpath = "google.golang.org/api/googleapi",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/api/internal/third_party/uritemplates"],
)
