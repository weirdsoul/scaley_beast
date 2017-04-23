git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.3",
)

git_repository(
     name = "io_bazel_rules_closure",
     remote = "https://github.com/bazelbuild/rules_closure.git",
     commit = "9507f68",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")
load("@io_bazel_rules_closure//closure:defs.bzl", "closure_repositories")

go_repositories()
closure_repositories()

BARE_BUILD = """
load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_library")

go_prefix("github.com/gorilla/websocket")

go_library(
    name = "go_default_library",
    srcs = [
    	 "client.go",
	 "client_clone.go",
	 "client_clone_legacy.go",
	 "compression.go",
	 "conn.go",
	 "conn_read.go",
	 "conn_read_legacy.go",
	 "doc.go",
	 "json.go",
	 "mask.go",
	 "mask_safe.go",
	 "prepared.go",
	 "server.go",
	 "util.go",
	 ],
    visibility = ["//visibility:public"],
)

"""

new_git_repository(
    name = "gorilla_websocket",
    remote = "https://github.com/gorilla/websocket.git",
    build_file_content = BARE_BUILD,
    commit = "a91eba7",
)
