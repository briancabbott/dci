load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "number",
    srcs = [
        "common.go",
        "decimal.go",
        "format.go",
        "number.go",
        "pattern.go",
        "roundingmode_string.go",
        "tables.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/text/internal/number",
    importpath = "golang.org/x/text/internal/number",
    visibility = ["//vendor/golang.org/x/text:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/text/internal/language/compact",
        "//vendor/golang.org/x/text/internal/stringset",
        "//vendor/golang.org/x/text/language",
    ],
)
