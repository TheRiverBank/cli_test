git add/commit/push
git tag -a v0.0.0 -m "release"
git push origin v0.0.0
goreleaser release --clean

winget install --manifest .\manifests\t\TheRiverBank\cli_test\0.7.0\