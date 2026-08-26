## Purpose

Require informed consent for operation in previously untrusted repositories while retaining and explaining native project-level harness configuration.

## ADDED Requirements

### Requirement: Prompt for unknown repository trust
The system SHALL request consent before launching a harness in a repository whose canonical local path has not been trusted.

#### Scenario: First launch in repository
- **WHEN** a user requests a session in an unknown repository
- **THEN** the system shows the repository identity and relevant detected configuration before offering session-only trust, persistent trust, or denial

### Requirement: Preserve native project configuration
The system SHALL allow the selected harness to read its native project-level configuration without copying, replacing, or silently suppressing it.

#### Scenario: Project contains vendor configuration
- **WHEN** the selected project contains native configuration for the selected harness
- **THEN** the harness sees that configuration according to its own precedence rules

### Requirement: Diagnose effective configuration layers
The system SHALL report generated, project, local-project, user, and managed configuration layers that the adapter can identify, including their expected precedence and limitations.

#### Scenario: Inspect project overrides
- **WHEN** project configuration can override or augment generated runtime configuration
- **THEN** inspection identifies the source and expected effect of that layer

