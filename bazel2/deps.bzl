load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_uber_go_tally",
        importpath = "github.com/uber-go/tally",
        sum = "h1:nFHIuW3VQ22wItiE9kPXic8dEgExWOsVOHwpmoIvsMw=",
        version = "v3.3.17+incompatible",
    )
