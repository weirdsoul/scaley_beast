load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_prefix("github.com/weirdsoul/scaley_beast/scalestate")

go_library(
	name = "go_default_library",
	srcs = [
	     "scalestate.go",
	],
	visibility = ["//visibility:public"],
)

go_test(
	name = "scalestate_test",
	srcs = [
	     "scalestate_test.go",
	],
	library = ":go_default_library",
)
