## Purpose

Provide consistent CLI and terminal UI operations for creating, understanding, connecting to, and ending isolated harness sessions.

## ADDED Requirements

### Requirement: Provide equivalent session operations
The system SHALL expose create, list, inspect, attach, detach, stop, and delete operations through a CLI and an interactive terminal UI.

#### Scenario: Reopen the UI
- **WHEN** a user starts a new TUI client while the daemon has active sessions
- **THEN** the UI lists the existing sessions and allows attachment

### Requirement: Display session identity
The system SHALL clearly display each session's profile, security context, tool, project, lifecycle state, and attention or exit status.

#### Scenario: Multiple similar sessions
- **WHEN** several sessions use the same tool
- **THEN** the user can distinguish them by profile, context, project, and session identity

### Requirement: Treat runtime selection as immutable
The system SHALL NOT mutate the profile, security context, tool, or project of a running session.

#### Scenario: Change the profile of a running session
- **WHEN** the user requests a different profile for an existing session
- **THEN** the system offers to create or restart as a new runtime rather than modifying the active process

