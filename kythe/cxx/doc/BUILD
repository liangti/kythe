package(default_visibility = ["//kythe:default_visibility"])

licenses(["notice"])

cc_library(
    name = "markup_handler",
    srcs = [
        "html_markup_handler.cc",
        "javadoxygen_markup_handler.cc",
        "markup_handler.cc",
    ],
    hdrs = [
        "html_markup_handler.h",
        "javadoxygen_markup_handler.h",
        "markup_handler.h",
    ],
    deps = [
        "//kythe/proto:common_cc_proto",
        "//kythe/proto:xref_cc_proto",
        "@com_github_google_glog//:glog",
    ],
)

cc_library(
    name = "markup_handler_testlib",
    testonly = 1,
    srcs = [
        "markup_handler_test.cc",
    ],
    deps = [
        ":markup_handler",
        "//third_party:gtest",
    ],
)

cc_test(
    name = "markup_handler_test",
    size = "small",
    deps = [
        ":markup_handler_testlib",
    ],
)

cc_library(
    name = "html_renderer",
    srcs = [
        "html_renderer.cc",
    ],
    hdrs = [
        "html_renderer.h",
    ],
    deps = [
        ":markup_handler",
        "//kythe/cxx/common:lib",
        "//kythe/proto:common_cc_proto",
        "//kythe/proto:xref_cc_proto",
        "@com_github_google_glog//:glog",
    ],
)

cc_library(
    name = "html_renderer_testlib",
    testonly = 1,
    srcs = [
        "html_renderer_test.cc",
    ],
    deps = [
        ":html_renderer",
        ":markup_handler",
        "//kythe/proto:common_cc_proto",
        "//third_party:gtest",
    ],
)

cc_test(
    name = "html_renderer_test",
    size = "small",
    deps = [
        ":html_renderer_testlib",
    ],
)

cc_library(
    name = "doclib",
    srcs = [
        "doc.cc",
    ],
    deps = [
        ":html_renderer",
        ":markup_handler",
        "//kythe/cxx/common:lib",
        "//kythe/cxx/common:net_client",
        "//kythe/cxx/common/schema:edges",
        "//kythe/cxx/common/schema:facts",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_absl//absl/memory",
    ],
)

cc_binary(
    name = "doc",
    deps = [
        ":doclib",
    ],
)
