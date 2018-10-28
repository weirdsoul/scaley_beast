load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.1/rules_go-0.16.1.tar.gz"],
    sha256 = "f5127a8f911468cd0b2d7a141f17253db81177523e4429796e14d429f5444f5f",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()
go_register_toolchains()

git_repository(
     name = "io_bazel_rules_closure",
     remote = "https://github.com/bazelbuild/rules_closure.git",
     commit = "0.8.0",
)

load("@io_bazel_rules_closure//closure:defs.bzl", "closure_repositories")

closure_repositories()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

BARE_BUILD = """
load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    importpath = "github.com/gorilla/websocket",
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
