github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = "e1e9715ea1a22991fdef2effccee87e05f4b756f",
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

sh_cmd(
    name = "publish",
    cmd = "$(out_location ///pleasings2//tools/misc:hub) release create -a $(out_location @linux_amd64//:package) -a $(out_location @darwin_amd64//:package) $@",
    deps = [
        "///pleasings2//tools/misc:hub",
        "@linux_amd64//:package",
        "@darwin_amd64//:package",
    ],
)
