load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "authhandler",
    srcs = ["authhandler.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/oauth2/authhandler",
    importpath = "golang.org/x/oauth2/authhandler",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/oauth2"],
)
