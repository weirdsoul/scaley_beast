load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")

go_prefix("github.com/weirdsoul/browser_instruments")

go_binary(
	name = "instruments_server",
	srcs = [
	     "main.go",
	     "planestate.go",
	     "udpreader.go",
	],
)
