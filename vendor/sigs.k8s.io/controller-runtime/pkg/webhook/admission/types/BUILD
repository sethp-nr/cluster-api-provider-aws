load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["types.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/sigs.k8s.io/controller-runtime/pkg/webhook/admission/types",
    importpath = "sigs.k8s.io/controller-runtime/pkg/webhook/admission/types",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mattbaird/jsonpatch:go_default_library",
        "//vendor/k8s.io/api/admission/v1beta1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
    ],
)
