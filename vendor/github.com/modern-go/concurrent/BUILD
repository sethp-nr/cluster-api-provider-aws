load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "executor.go",
        "go_above_19.go",
        "go_below_19.go",
        "log.go",
        "unbounded_executor.go",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/github.com/modern-go/concurrent",
    importpath = "github.com/modern-go/concurrent",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "map_test.go",
        "unbounded_executor_test.go",
    ],
    embed = [":go_default_library"],
)
