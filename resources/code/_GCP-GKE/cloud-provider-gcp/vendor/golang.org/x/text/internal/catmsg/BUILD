load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "catmsg",
    srcs = [
        "catmsg.go",
        "codec.go",
        "varint.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/text/internal/catmsg",
    importpath = "golang.org/x/text/internal/catmsg",
    visibility = ["//vendor/golang.org/x/text:__subpackages__"],
    deps = ["//vendor/golang.org/x/text/language"],
)
