load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "base",
    srcs = ["api_base.pb.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/appengine/internal/base",
    importpath = "google.golang.org/appengine/internal/base",
    visibility = ["//vendor/google.golang.org/appengine:__subpackages__"],
    deps = ["//vendor/github.com/golang/protobuf/proto"],
)
