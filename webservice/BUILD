load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
	name = "go_default_library",
	srcs = [
	     "serve.go",
	],
	visibility = ["//visibility:public"],
	deps = [
	     "//scalestate:go_default_library",
	     "@gorilla_websocket//:go_default_library",
	],
	importpath = "github.com/weirdsoul/scaley_beast/webservice"
)
