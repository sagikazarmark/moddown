github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "e4cc66bc0cd5b2bc86fc9bd058319a5c864c4261",
)

go_binary(
    name = "moddown",
    srcs = ["main.go"],
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
