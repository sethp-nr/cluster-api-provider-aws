load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["message.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/golang.org/x/net/dns/dnsmessage",
    importpath = "golang.org/x/net/dns/dnsmessage",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "example_test.go",
        "message_test.go",
    ],
    embed = [":go_default_library"],
)
