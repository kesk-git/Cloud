# CLOUD-STATUS

Task: add a `--version` flag to cloud-stub that prints `cloud-stub 0.2.0` and exits 0.
Instruction marker: `repo-instructions-20261010`. Skill marker: `repo-skill-20261010`.
Session: https://claude.ai/code/session_01DcLAfCQuonLoebowjdNskx
Branch: claude/cloud-stub-version-flag-kev28v

Changes: `main.go` (version constant and `--version` check in `run`), `main_test.go` (`TestVersion`), `README.md` (one sentence). Usage string, `--smoke`, `--name`, and all rejection rules are unchanged.

Checks run on linux/amd64 (go1.24.7): gofmt clean, `go vet ./...` pass, `go test ./...` pass (6 tests), `go build -o dist/cloud-stub .` pass, executable `--version` prints `cloud-stub 0.2.0` exit 0, `--smoke` prints `cloud-stub: ok`, golden fixture diff exit 0, `--version extra` and `--smoke --version` rejected with the usage error. Windows and macOS were not checked here.

Prerequisite note: gh-axi is not installed in this session; no GitHub API client was substituted. Only ordinary git push was used. No pull request was opened.

done claude/cloud-stub-version-flag-kev28v
