load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "generator",
    srcs = [
        "default_generator.go",
        "default_package.go",
        "doc.go",
        "error_tracker.go",
        "execute.go",
        "generator.go",
        "import_tracker.go",
        "snippet_writer.go",
        "transitive_closure.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/gengo/generator",
    importpath = "k8s.io/gengo/generator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/imports",
        "//vendor/k8s.io/gengo/namer",
        "//vendor/k8s.io/gengo/parser",
        "//vendor/k8s.io/gengo/types",
        "//vendor/k8s.io/klog/v2:klog",
    ],
)
