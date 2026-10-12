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
