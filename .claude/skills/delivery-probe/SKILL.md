---
name: delivery-probe
description: Record repository skill loading for the cloud delivery experiment.
---

Record `repo-skill-20261010` in `.fleet/status/cloud-delivery.json`, together with
the UTC timestamp, current commit SHA, OS/architecture, Go version, session ID,
whether `.fleet/probes/session-start.txt` exists and its non-secret content,
and the value of the non-secret `FLEET_SETTINGS_PROBE` environment variable.
Never dump other environment variables or credentials. State which facts were
observed versus inferred. Commit and push the status promptly on the task branch.
