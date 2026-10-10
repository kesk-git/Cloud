# Cloud delivery worker task

Task ID: cloud-delivery. This is a public throwaway delivery experiment.

1. Read CLAUDE.md and run the delivery-probe skill. Record bootstrap evidence in
   `.fleet/status/cloud-delivery.json` and promptly commit/push it to your task
   branch so the Mac can see startup. Include status, sequence, UTC timestamp,
   source SHA and session ID. Do not push main or merge. Do not create subagents.
2. Add optional `--name NAME` to `--smoke`. Default output stays `cloud-stub: ok`.
   A supplied nonempty name produces `cloud-stub: ok (NAME)`. Reject an empty
   name or malformed arguments. Add focused tests and README examples.
3. Expect one steering message from the Mac while you work. Record its exact
   marker, receipt UTC time and implementation result in the status JSON.
4. Run gofmt, go vet, go test, build, and the executable smoke tests. Commit and
   push all work, updating status with test evidence and current state. Use a
   cloud-created task branch if the platform requires one; record its name.
5. Open a pull request to main in this repository with gh-axi, including the
   full cloud session URL and validation. If gh-axi is absent, report that in
   status and stop before opening it; do not install or authenticate anything.
   Report the full PR URL. Never merge or enable auto-merge.

Do not change the Actions gate to make a failure disappear. Do not use any
VibeHub source or files outside this clone. Do not add a GUI. No deployments,
paid services, tokens, login/setup flows, or account changes. Record blockers
instead of circumventing them. Keep this task small.
