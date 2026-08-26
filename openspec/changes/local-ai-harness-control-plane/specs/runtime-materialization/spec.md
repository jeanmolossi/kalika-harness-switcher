## Purpose

Create isolated, inspectable runtime instances from adapter plans while protecting secrets and tracking changes to application-managed configuration.

## ADDED Requirements

### Requirement: Materialize one runtime instance per session
The system SHALL create a distinct mutable runtime instance for every session and SHALL NOT make concurrent sessions write to the same vendor home by default.

#### Scenario: Start identical sessions
- **WHEN** two sessions use the same profile, security context, tool, and project
- **THEN** each receives a distinct runtime instance and vendor home

### Requirement: Control materialization operations
The system SHALL validate and execute declarative adapter operations atomically and SHALL reject destinations that escape the runtime staging root.

#### Scenario: Adapter requests path traversal
- **WHEN** a materialization operation targets a path outside the staging root
- **THEN** the operation is rejected and no partial runtime is committed

### Requirement: Filter child environments
The system SHALL construct child environments from an allowlist plus explicitly resolved values rather than inheriting the complete daemon environment.

#### Scenario: Unrequested credential exists in daemon
- **WHEN** the daemon environment contains a credential not selected by the security context
- **THEN** that credential is absent from the harness environment

### Requirement: Resolve secrets late
The system SHALL persist secret references rather than secret values and SHALL resolve their values only for process launch.

#### Scenario: Inspect runtime
- **WHEN** a user inspects a runtime containing credential references
- **THEN** the output identifies the references and resolution status without exposing secret values

### Requirement: Detect managed configuration changes
The system SHALL compare application-managed files against their materialized baseline and offer changes as a diff without automatically updating a registry source.

#### Scenario: Harness modifies generated configuration
- **WHEN** a harness changes a managed file in its runtime home
- **THEN** the user can inspect the diff separately from vendor-owned history and cache files

