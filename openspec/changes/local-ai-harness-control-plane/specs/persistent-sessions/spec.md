## Purpose

Keep multiple native harness processes interactive and alive independently of the lifecycle of any individual CLI or terminal UI client.

## ADDED Requirements

### Requirement: Keep sessions after client exit
The system SHALL run session processes under a persistent local daemon so closing or crashing a CLI/TUI client does not terminate running sessions.

#### Scenario: Close the terminal UI
- **WHEN** the user exits the TUI while sessions are running
- **THEN** the daemon keeps those processes and PTYs active

### Requirement: Support concurrent sessions
The system SHALL supervise multiple sessions concurrently with independent runtime instances, PTYs, process groups, status, and exit results.

#### Scenario: Run different harnesses simultaneously
- **WHEN** the user starts Codex, Claude Code, and Copilot sessions
- **THEN** all sessions can remain running independently

### Requirement: Attach and detach
The system SHALL allow a client to attach to a running session and detach using a configurable reserved key sequence without terminating the harness.

#### Scenario: Detach from harness
- **WHEN** the user enters the detach sequence while attached
- **THEN** control returns to the client and the harness remains running

### Requirement: Observe natural process exit
The system SHALL mark a session exited when its harness process terminates, including when the user invokes the harness's own exit command.

#### Scenario: Exit inside harness
- **WHEN** the harness process exits normally from its native interface
- **THEN** the daemon records its exit result and emits a session-exited event

### Requirement: Stop process groups
The system SHALL stop the session process group with a graceful period before forced termination.

#### Scenario: Harness does not stop gracefully
- **WHEN** a stopped session remains alive after the configured grace period
- **THEN** the daemon forcibly terminates its process group and records the outcome

