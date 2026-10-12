# CLOUD-STATUS

(a) System-prompt sentence: "You are powered by the model named Sonnet 5.5. The exact model ID is claude-sonnet-5-5."

(b) get_session values (verbatim):
- session_context.model: claude-sonnet-5-5
- session_context.effort_level: low
- permission-mode field: none reported by get_session

(c) Tools whose names contain "session", "archive", "exit", "end" or "close" (none called except get_session):
- mcp__claude-code-remote__archive_session
- mcp__claude-code-remote__create_session
- mcp__claude-code-remote__get_session
- mcp__claude-code-remote__interrupt_session
- mcp__claude-code-remote__list_sessions
- mcp__claude-code-remote__send_later
- mcp__claude-code-remote__send_message
- mcp__claude-code-remote__set_session_tags
- mcp__claude-code-remote__set_session_title
- mcp__claude-code-remote__unarchive_session
- SendUserFile
- SendMessage (deferred)
- ExitPlanMode (deferred)
- ExitWorktree (deferred)

(d) Permission mode: not stated to me; as far as I can tell, a non-interactive/auto-approving remote mode (tool calls ran without prompts).

Instruction marker: repo-instructions-20261010
delivery-probe skill loaded (marker repo-skill-20261010). Its request to write .fleet/status/cloud-delivery.json was NOT done: the task restricts changes to .claude/settings.json and CLOUD-STATUS.md.

Checks (Linux only): gofmt clean, go vet ok, go test ok, compiled executable built. gh-axi not found (prerequisite absent; no other GitHub client used). Windows/macOS not checked.
done fm/cloud-teardown-test-1
