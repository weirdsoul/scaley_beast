load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_library")
load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_binary")

closure_js_library(
	name = "scaley_beast_client_main",
	srcs = [
	     "main.js",
	     "analog_gauge.js",
	],
	deps = [
	  "@io_bazel_rules_closure//closure/library",
	],
)

closure_js_binary(
	name = "scaley_beast_client",
	deps = [
	     ":scaley_beast_client_main",
	],
        entry_points = ["goog:browser_instruments.AssignFromScale"],
	visibility = ["//visibility:public"],
)

filegroup(
	name = "data",
	srcs = [
	     "index.html",
	     "success.html",
	     "failure.html",
	     ":scaley_beast_client",
	] + glob(["assets/*"]),
	visibility = ["//visibility:public"],
)
