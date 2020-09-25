# Go Module downloader

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sagikazarmark/moddown/CI?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/moddown?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/moddown)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/sagikazarmark/moddown)](https://pkg.go.dev/mod/github.com/sagikazarmark/moddown)

`moddown` is a simplified version of [fetch_repo](https://github.com/bazelbuild/bazel-gazelle/tree/5c00b77/cmd/fetch_repo).
It focuses on downloading a module using `go mod download`.
Unlike `fetch_repo`, this tool [does not create a dummy module](https://github.com/golang/go/issues/29522)
and uses the [-modcacherw](https://golang.org/doc/go1.14#go-flags) flag to make the cache writable (removable),
so it requires at least Go 1.14.

## Usage

The basic functionality of `moddown` includes:

- downloading a module (by calling `go mod download` under the hood)
- checking the module sum against a known value
- copying files to a destination directory

```bash
./moddown -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -dest ./emperror emperror.dev/errors@v0.4.2
```

`moddown` can also use existing module information and delegate downloading to the Go tool:

```bash
go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 > mod.json
./moddown -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f mod.json

# OR pipe from stdin directly

go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 | ./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f -
```

Last, but not least: you can enable verbose logging with `-v`.


## Development

Contributions are welcome! :)

1. Clone the repository
1. Make changes on a new branch
1. Run the test suite:
    ```bash
    ./pleasew build
    ./pleasew test
    ./pleasew lint
    ```
1. Commit, push and open a PR


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
