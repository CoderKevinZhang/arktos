load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dispatcher.go",
        "interface.go",
    ],
    importpath = "k8s.io/kubernetes/globalscheduler/pkg/apis/dispatcher/client/informers/externalversions/dispatcher/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//globalscheduler/pkg/apis/dispatcher/client/clientset/versioned:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/client/informers/externalversions/internalinterfaces:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/client/listers/dispatcher/v1:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//staging/src/k8s.io/client-go/tools/cache:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
