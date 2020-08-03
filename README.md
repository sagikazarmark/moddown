# Go Module downloader

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/moddown/CI?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/moddown?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/moddown)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/sagikazarmark/moddown)

`moddown` is a simplified version of [fetch_repo](https://github.com/bazelbuild/bazel-gazelle/tree/5c00b77/cmd/fetch_repo).
It focuses on downloading a module using `go mod download`.
Unlike `fetch_repo`, this tool [does not create a dummy module](https://github.com/golang/go/issues/29522)
and uses the [-modcacherw](https://golang.org/doc/go1.14#go-flags) flag to make the cache writable (removable),
so it requires at least Go 1.14.


## Development

1. Clone the repository
1. Make changes on a new branch
1. Run `./pleasew build`
1. Commit, push and open PR


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
