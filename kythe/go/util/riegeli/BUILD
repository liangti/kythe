load("//tools:build_rules/shims.bzl", "go_library", "go_test")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "riegeli",
    srcs = [
        "compression.go",
        "reader.go",
        "riegeli.go",
        "transpose.go",
        "transpose_util.go",
        "util.go",
        "writer.go",
    ],
    deps = [
        "//third_party/riegeli:records_metadata_go_proto",
        "@com_github_datadog_zstd//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_minio_highwayhash//:go_default_library",
        "@org_brotli_go//cbrotli",
    ],
)

go_test(
    name = "riegeli_test",
    srcs = ["riegeli_test.go"],
    library = ":riegeli",
)

go_test(
    name = "golden_test",
    srcs = ["golden_test.go"],
    data = [
        "testdata/golden.entries.brotli.riegeli",
        "testdata/golden.entries.brotli_transpose.riegeli",
        "testdata/golden.entries.json",
        "testdata/golden.entries.uncompressed.riegeli",
        "testdata/golden.entries.uncompressed_transpose.riegeli",
        "testdata/golden.entries.zstd.riegeli",
        "testdata/golden.records_metadata.textproto",
    ],
    library = ":riegeli",
    deps = [
        "//kythe/go/storage/stream",
        "//kythe/proto:storage_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_google_go_cmp//cmp:go_default_library",
    ],
)

go_test(
    name = "block_test",
    srcs = ["block_test.go"],
    library = ":riegeli",
)

go_test(
    name = "riegeli_bench_test",
    srcs = ["riegeli_bench_test.go"],
    library = ":riegeli",
)

go_test(
    name = "transpose_test",
    srcs = ["transpose_test.go"],
    library = ":riegeli",
)
