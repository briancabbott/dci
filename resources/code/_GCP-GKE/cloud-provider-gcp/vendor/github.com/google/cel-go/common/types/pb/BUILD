load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "pb",
    srcs = [
        "checked.go",
        "enum.go",
        "equal.go",
        "file.go",
        "pb.go",
        "type.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/google/cel-go/common/types/pb",
    importpath = "github.com/google/cel-go/common/types/pb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/genproto/googleapis/api/expr/v1alpha1",
        "//vendor/google.golang.org/protobuf/encoding/protowire",
        "//vendor/google.golang.org/protobuf/proto",
        "//vendor/google.golang.org/protobuf/reflect/protoreflect",
        "//vendor/google.golang.org/protobuf/reflect/protoregistry",
        "//vendor/google.golang.org/protobuf/types/dynamicpb",
        "//vendor/google.golang.org/protobuf/types/known/anypb",
        "//vendor/google.golang.org/protobuf/types/known/durationpb",
        "//vendor/google.golang.org/protobuf/types/known/emptypb",
        "//vendor/google.golang.org/protobuf/types/known/structpb",
        "//vendor/google.golang.org/protobuf/types/known/timestamppb",
        "//vendor/google.golang.org/protobuf/types/known/wrapperspb",
    ],
)
