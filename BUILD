github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "e4cc66bc0cd5b2bc86fc9bd058319a5c864c4261",
)

go_binary(
    name = "moddown",
    srcs = ["main.go"],
)

gentest(
    name = "test_download",
    test_cmd = "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= emperror.dev/errors@v0.4.2",
    data = [":moddown"],
    no_test_output = True,
)

gentest(
    name = "test_file",
    test_cmd = " && ".join([
        "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 > mod.json",
        "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f mod.json",
    ]),
    data = [":moddown"],
    no_test_output = True,
)

gentest(
    name = "test_stdin",
    test_cmd = "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 | ./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f -",
    data = [":moddown"],
    no_test_output = True,
)

tarball(
    name = "package",
    srcs = ["README.md", "LICENSE", ":moddown"],
    out = f"moddown_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    gzip = True,
    labels = ["dist"],
)

subinclude("///pleasings2//github")

github_release(
    name = "publish",
    assets = [
        "@linux_amd64//:package",
        "@darwin_amd64//:package",
    ],
    labels = ["dist"],
)

sh_cmd(
    name = "goroot",
    cmd = "go env GOROOT",
)
