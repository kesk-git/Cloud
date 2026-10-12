# Cloud status

Model sentence (verbatim from system prompt):

> You are powered by the model named Opus 5.5. The exact model ID is claude-opus-5-5.

- Launch label: opus-effort-low
- Session ID: session_018XH89Y2bdnS3KGbF7L44uZ
- Session URL: https://claude.ai/code/session_018XH89Y2bdnS3KGbF7L44uZ
- Instruction marker: repo-instructions-20261010
- gh-axi: not available (prerequisite absent; not needed, plain git push used)


## Follow-up: after-slash-model

- What happened to "/model claude-sonnet-5-5": it executed as a real slash command, run directly by the harness. I got it as a local-command record with stdout "Set model to `Sonnet 5.5` for this session only", not as a request to act on.
- System prompt sentence now (verbatim): "You are powered by the model named Sonnet 5.5. The exact model ID is claude-sonnet-5-5."
- get_session: session_context.model = claude-sonnet-5-5; external_metadata.last_served_model = claude-sonnet-5-5; external_metadata.model = claude-sonnet-5-5; configured_model = claude-opus-5-5 (creation-time value, unchanged).
- Effort: get_session shows session_context.effort_level = "low". The system prompt names no effort level. This matches the --effort low launch.

done claude/model-probe-opus-effort-low-52lihw
