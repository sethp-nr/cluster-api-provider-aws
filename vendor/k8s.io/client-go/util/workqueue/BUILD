load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "default_rate_limiters.go",
        "delaying_queue.go",
        "doc.go",
        "metrics.go",
        "parallelizer.go",
        "queue.go",
        "rate_limitting_queue.go",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/k8s.io/client-go/util/workqueue",
    importpath = "k8s.io/client-go/util/workqueue",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/time/rate:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/clock:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "default_rate_limiters_test.go",
        "delaying_queue_test.go",
        "queue_test.go",
        "rate_limitting_queue_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/util/clock:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
    ],
)
