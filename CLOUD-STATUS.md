# Cloud status: model probe

Launch label: model-opus

Session ID: session_013qu5KscCFy3Hgq5qVVmqPz
Session URL: https://claude.ai/code/session_013qu5KscCFy3Hgq5qVVmqPz

Model sentence from the system prompt (verbatim):

> You are powered by the model named Opus 5.5. The exact model ID is claude-opus-5-5.

Cross-check (observed via the session's get_session tool): session_context.model = claude-opus-5-5, external_metadata.last_served_model = claude-opus-5-5.

Repository instruction marker: repo-instructions-20261010

Branch note: the harness designated and pre-created `claude/model-probe-opus-80zlti` for this session, so that branch was used instead of `claude/model-probe-model-opus`.

`.scout/` is listed in .gitignore, so `.scout/model-probe.txt` was added with `git add -f` (only that file).

Checks (linux/amd64 only): gofmt -l clean, go vet pass, go test pass, compiled `--smoke` printed `cloud-stub: ok`. No Go source changed. No Windows or macOS checks ran.

## Follow-up: trailer-test

Commit a418c9e appends the `label=trailer-test` line to `.scout/model-probe.txt`.

Exact command: `git add -f .scout/model-probe.txt && git commit -q -m "Append trailer-test line to model probe log"`

Resulting `git log -1 --format=%B` (observed): the single line `Append trailer-test line to model probe log` followed by git's trailing blank line. It has no Co-Authored-By, no Claude-Session and no other trailers, and nothing was amended.

What could control trailers (observed): `.git/hooks` has only `*.sample` files, so no active hooks. `core.hooksPath` is unset, and `git config` has no trailer, template or hook entries. `.claude/settings.json` has no `includeCoAuthoredBy` or `attribution` key. `~/.claude/settings.json`, `.claude/settings.local.json` and `/etc/claude-code/managed-settings.json` don't exist.

Inferred: in this session the attribution trailers aren't added by git or a hook. They come from instructions the harness gives the model about what to write in commit messages, and the model writes them into each `-m` message itself. When the user asks for none, none are written. No setting visible in this container controls them.

done claude/model-probe-opus-80zlti
