# Cloud delivery experiment

Work only in this public repository. Do not merge, deploy, change account settings,
add secrets, or start another agent/session. Push only a task branch to this repo.
The Mac supervisor retains merge authority. Stop on credential or paid-tier needs.

Before implementation, use the `delivery-probe` skill and copy this instruction
marker into the task status: `repo-instructions-20261010`.

Run gofmt, go vet, go test and a compiled executable smoke test for changes.
Use gh-axi for GitHub operations if available. If absent, report that prerequisite;
do not substitute a different GitHub API client. Ordinary git push is allowed.

Do not claim Windows/macOS checks passed from a Linux build. Report real checks.
