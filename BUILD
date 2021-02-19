github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "f644350aab5e0090ab69022beeb2feadfcdc223e",
)

go_binary(
    name = "moddown",
    srcs = ["main.go"],
    deps = [],
)

gentest(
    name = "test_download",
    data = [":moddown"],
    no_test_output = True,
    test_cmd = " && ".join([
        "go mod init test",
        "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= emperror.dev/errors@v0.4.2",
    ]),
)

gentest(
    name = "test_file",
    data = [":moddown"],
    no_test_output = True,
    test_cmd = " && ".join([
        "go mod init test",
        "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 > mod.json",
        "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f mod.json",
    ]),
)

gentest(
    name = "test_stdin",
    data = [":moddown"],
    no_test_output = True,
    test_cmd = " && ".join([
        "go mod init test",
        "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 | ./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f -",
    ]),
)

tarball(
    name = "artifact",
    srcs = [
        "LICENSE",
        "README.md",
        ":moddown",
    ],
    out = f"moddown_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    gzip = True,
    labels = ["manual"],
)

subinclude("///pleasings2//misc")

build_artifacts(
    name = "artifacts",
    artifacts = [
        "@linux_amd64//:artifact",
        "@darwin_amd64//:artifact",
    ],
    labels = ["manual"],
)

subinclude("///pleasings2//github")

github_release(
    name = "publish",
    assets = [":artifacts"],
    labels = ["manual"],
)
