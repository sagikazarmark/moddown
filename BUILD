github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "4c40fa674130e6d92bcdb4ef9bd17954fdbf3fab",
)

go_binary(
    name = "moddown",
    srcs = ["main.go"],
)

gentest(
    name = "test_download",
    test_cmd = " && ".join([
        "go mod init test",
        "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= emperror.dev/errors@v0.4.2",
    ]),
    data = [":moddown"],
    no_test_output = True,
)

gentest(
    name = "test_file",
    test_cmd = " && ".join([
        "go mod init test",
        "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 > mod.json",
        "./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f mod.json",
    ]),
    data = [":moddown"],
    no_test_output = True,
)

gentest(
    name = "test_stdin",
    test_cmd = " && ".join([
        "go mod init test",
        "go mod download -modcacherw -x -json emperror.dev/errors@v0.4.2 | ./moddown -v -sum h1:snD5ODyv4c9DOBBZh645dy/TziVHZivuFtRRMZP8zK8= -f -",
    ]),
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

subinclude("///pleasings2//misc")

sha256sum(
    name = "checksums.txt",
    srcs = [
        "@linux_amd64//:package",
        "@darwin_amd64//:package",
    ],
    out = "checksums.txt",
    labels = ["dist"],
)

subinclude("///pleasings2//github")

github_release(
    name = "publish",
    assets = [
        "@linux_amd64//:package",
        "@darwin_amd64//:package",
        ":checksums.txt",
    ],
    labels = ["dist"],
)
