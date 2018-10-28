load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_prefix("github.com/weirdsoul/scaley_beast")

go_binary(
	name = "scaley_beast",
	srcs = [
	     "main.go",
	     "serialreader.go",
	],
	data = [
	     "//client:data",
	],
	deps = [
	     "//scalestate:go_default_library",
	     "//webservice:go_default_library",
	],
)
