load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["protosize.pb.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/github.com/gogo/protobuf/test/protosize",
    importpath = "github.com/gogo/protobuf/test/protosize",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "protosize_test.go",
        "protosizepb_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/jsonpb:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
    ],
)
