load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_library(
	name = "go_default_library",
	srcs = [
	     "scalestate.go",
	],
	visibility = ["//visibility:public"],
	importpath = "github.com/weirdsoul/scaley_beast/scalestate",
)

go_test(
	name = "scalestate_test",
	srcs = [
	     "scalestate_test.go",
	],
	importpath = "github.com/weirdsoul/scaley_beast/scalestate",
	embed = [":go_default_library"],
)
