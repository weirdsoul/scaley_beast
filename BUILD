load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_prefix("github.com/weirdsoul/browser_instruments")

go_binary(
	name = "instruments_server",
	srcs = [
	     "main.go",
	     "udpreader.go",
	],
	deps = [
	     "//planestate:go_default_library",
	     "//webservice:go_default_library",
	],
)
