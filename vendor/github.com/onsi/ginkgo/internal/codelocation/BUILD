load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["code_location.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/github.com/onsi/ginkgo/internal/codelocation",
    importpath = "github.com/onsi/ginkgo/internal/codelocation",
    visibility = ["//vendor/github.com/onsi/ginkgo:__subpackages__"],
    deps = ["//vendor/github.com/onsi/ginkgo/types:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "code_location_suite_test.go",
        "code_location_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/ginkgo/types:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
    ],
)
