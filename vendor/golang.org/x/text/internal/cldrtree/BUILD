load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cldrtree.go",
        "generate.go",
        "option.go",
        "tree.go",
        "type.go",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/golang.org/x/text/internal/cldrtree",
    importpath = "golang.org/x/text/internal/cldrtree",
    visibility = ["//vendor/golang.org/x/text:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/text/internal:go_default_library",
        "//vendor/golang.org/x/text/internal/gen:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/unicode/cldr:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["cldrtree_test.go"],
    embed = [":go_default_library"],
    deps = [
        "//vendor/golang.org/x/text/internal/gen:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/unicode/cldr:go_default_library",
    ],
)
