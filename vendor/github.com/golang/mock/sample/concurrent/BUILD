load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["concurrent.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/github.com/golang/mock/sample/concurrent",
    importpath = "github.com/golang/mock/sample/concurrent",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["concurrent_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/github.com/golang/mock/sample/concurrent/mock:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
