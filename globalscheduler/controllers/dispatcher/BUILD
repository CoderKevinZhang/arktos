load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dispatcher_controller.go",
        "dispatcher_crd.go",
        "dispatcher_process.go",
        "eventType.go",
    ],
    importpath = "k8s.io/kubernetes/globalscheduler/controllers/dispatcher",
    visibility = ["//visibility:public"],
    deps = [
        "//globalscheduler/cmd/conf:go_default_library",
        "//globalscheduler/controllers:go_default_library",
        "//globalscheduler/controllers/util:go_default_library",
        "//globalscheduler/controllers/util/openstack:go_default_library",
        "//globalscheduler/pkg/apis/cluster/client/clientset/versioned:go_default_library",
        "//globalscheduler/pkg/apis/cluster/client/informers/externalversions/cluster/v1:go_default_library",
        "//globalscheduler/pkg/apis/cluster/v1:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/client/clientset/versioned:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/client/clientset/versioned/scheme:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/client/informers/externalversions/dispatcher/v1:go_default_library",
        "//globalscheduler/pkg/apis/dispatcher/v1:go_default_library",
        "//globalscheduler/pkg/scheduler/utils/runtime:go_default_library",
        "//staging/src/k8s.io/api/core/v1:go_default_library",
        "//staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1:go_default_library",
        "//staging/src/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/fields:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/scheme:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
        "//staging/src/k8s.io/client-go/tools/cache:go_default_library",
        "//vendor/k8s.io/klog:go_default_library",
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
