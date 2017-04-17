load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_prefix("github.com/weirdsoul/browser_instruments/planestate")

go_library(
	name = "go_default_library",
	srcs = [
	     "planestate.go",
	],
	visibility = ["//visibility:public"],
)

go_test(
	name = "planestate_test",
	srcs = [
	     "planestate_test.go",
	],
	library = ":go_default_library",
)
