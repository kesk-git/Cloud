# Cloud status: model probe

Instruction marker: `repo-instructions-20261010`

## Model sentence (verbatim from the system prompt)

> You are powered by the model named Fable 5.1. The exact model ID is claude-fable-5-1.

## Session

- Launch label: baseline-noflag
- Session ID: session_01StDephh5Sz8KTeh516Xd94
- Session URL: https://claude.ai/code/session_01StDephh5Sz8KTeh516Xd94
- Recorded at (UTC): 2026-10-12T05:40:34Z
- Probe line appended to `.scout/model-probe.txt`
- Delivery-probe status recorded in `.fleet/status/cloud-delivery.json` (marker `repo-skill-20261010`)

## Checks (linux/amd64 only; no Go files changed)

- gofmt -l: no files listed
- go vet: pass
- go test: pass
- go build + `--smoke`: printed `cloud-stub: ok`, exit 0
- gh-axi: absent on PATH (not needed; only git push was used)

Follow-up (label=settings-fixture) at 2026-10-12T05:45:39Z: push of branch claude/model-probe-settings-fixture (commit 4b839ac, from origin/main 2c59bb8, edits only .claude/settings.json) to origin SUCCEEDED; no error.

done claude/model-probe-baseline-noflag-gqynbu
