load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "bipartitegraph",
    srcs = [
        "bipartitegraph.go",
        "bipartitegraphmatching.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/onsi/gomega/matchers/support/goraph/bipartitegraph",
    importpath = "github.com/onsi/gomega/matchers/support/goraph/bipartitegraph",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/onsi/gomega/matchers/support/goraph/edge",
        "//vendor/github.com/onsi/gomega/matchers/support/goraph/node",
        "//vendor/github.com/onsi/gomega/matchers/support/goraph/util",
    ],
)
