load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/golang-training
gazelle(name = "gazelle")

go_binary(
    name = "adder",
    srcs =  ["main.go"],
    embed = [":golang-training_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "golang-training_lib",
    srcs = [
        "adder.go"
    ],
    importpath = "github.com/golang-training",
    visibility = ["//visibility:private"],
    deps = ["@com_github_uber_go_tally//:tally"],
)

go_test(
    name = "golang-training_test",
    srcs = ["adder_test.go"],
    embed = [":golang-training_lib"],
)
