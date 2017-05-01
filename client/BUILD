load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_library")
load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_binary")

closure_js_library(
	name = "instruments_client_main",
	srcs = [
	     "main.js",
	],
	deps = [
	  "@io_bazel_rules_closure//closure/library",
	],
)

closure_js_binary(
	name = "instruments_client",
	deps = [
	     ":instruments_client_main",
	],
	visibility = ["//visibility:public"],
)

filegroup(
	name = "data",
	srcs = [
	     "index.html",
	     ":instruments_client",
	] + glob(["assets/*"]),
	visibility = ["//visibility:public"],
)
