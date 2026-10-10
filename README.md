# cloud

Throwaway public repository for the Claude cloud delivery scout: a cross-platform stub of VibeHub built by Claude cloud workers, with supervision and merge authority on the operator's Mac.

This dependency-free Go command is a delivery experiment, not a port of VibeHub
or a desktop UI. It uses no VibeHub source. Linux, Windows, and macOS each compile
and run their own native executable in GitHub Actions.

Run `go test ./...`, `go vet ./...`, and `go build -o dist/cloud-stub .`.
On Windows use `dist/cloud-stub.exe` as the output path. Run the resulting
executable with `--smoke`; it prints `cloud-stub: ok`.

Cloud workers may push task branches and open pull requests in this repository.
Only the Mac-side supervisor may merge. No deployments or paid services are used.
