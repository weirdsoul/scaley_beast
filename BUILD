load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

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
             "@tarm_serial//:go_default_library",
	],
	importpath = "github.com/weirdsoul/scaley_beast",
)
