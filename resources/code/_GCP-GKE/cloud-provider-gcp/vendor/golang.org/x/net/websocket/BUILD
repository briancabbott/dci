load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "websocket",
    srcs = [
        "client.go",
        "dial.go",
        "hybi.go",
        "server.go",
        "websocket.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/net/websocket",
    importpath = "golang.org/x/net/websocket",
    visibility = ["//visibility:public"],
)
