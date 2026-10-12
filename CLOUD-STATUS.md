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

done claude/model-probe-opus-80zlti
