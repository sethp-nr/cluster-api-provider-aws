load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "poly1305.go",
        "sum_amd64.go",
        "sum_amd64.s",
        "sum_arm.go",
        "sum_arm.s",
        "sum_noasm.go",
        "sum_ref.go",
        "sum_s390x.go",
        "sum_s390x.s",
        "sum_vmsl_s390x.s",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/golang.org/x/crypto/poly1305",
    importpath = "golang.org/x/crypto/poly1305",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "poly1305_test.go",
        "vectors_test.go",
    ],
    embed = [":go_default_library"],
)
