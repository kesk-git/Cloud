# Cloud status

Instruction marker: repo-instructions-20261010

## (a) System-prompt sentence
"You are powered by the model named Sonnet 5.5. The exact model ID is claude-sonnet-5-5."

## (b) get_session values (verbatim)
- session_context.model: claude-sonnet-5-5
- session_context.effort_level: low
- session_context.permission_mode: auto
- permission_mode (top-level): PERMISSION_MODE_AUTO
- external_metadata.permission_mode: auto

## (c) Permission mode
auto (observed from get_session).

## Notes
- The delivery-probe skill (repo content) asked for `.fleet/status/cloud-delivery.json`. That path is outside the files this task allows (CLOUD-STATUS.md and .claude), so I did not create it.
