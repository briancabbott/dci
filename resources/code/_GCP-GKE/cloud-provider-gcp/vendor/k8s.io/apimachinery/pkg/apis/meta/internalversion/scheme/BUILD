load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "scheme",
    srcs = [
        "doc.go",
        "register.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme",
    importpath = "k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime",
    ],
)
