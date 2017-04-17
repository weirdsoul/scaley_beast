load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")

go_prefix("github.com/weirdsoul/browser_instruments/webservice")

go_library(
	name = "go_default_library",
	srcs = [
	     "serve.go",
	],
	visibility = ["//visibility:public"],
	deps = [
	     "//planestate:go_default_library",
	],
)
