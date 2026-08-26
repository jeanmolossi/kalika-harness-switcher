## Purpose

Allow official and third-party harness integrations to translate resolved runtimes without embedding vendor conditionals in the core application.

## ADDED Requirements

### Requirement: Discover adapters by manifest
The system SHALL discover adapters from explicitly configured directories using manifests that identify the adapter, executable, implementation version, and protocol version.

#### Scenario: Install a third-party adapter
- **WHEN** a valid third-party adapter manifest is added to a configured adapter directory
- **THEN** the tool becomes available without modifying the resolver or session manager

### Requirement: Require adapter trust
The system SHALL require explicit trust before first executing an adapter whose identity or content digest has not previously been approved.

#### Scenario: Unknown adapter
- **WHEN** an untrusted adapter is selected
- **THEN** the system shows its origin, executable, and digest and does not execute it without consent

### Requirement: Use a versioned external protocol
The system SHALL invoke adapters as subprocesses through a versioned protocol for description, detection, validation, planning, and command construction.

#### Scenario: Incompatible protocol
- **WHEN** an adapter declares an unsupported protocol version
- **THEN** the system rejects it with a compatibility diagnostic

### Requirement: Report translation fidelity
The adapter SHALL identify each translated resource as native, converted, degraded, or unsupported and SHALL NOT silently discard unsupported semantics.

#### Scenario: Best-effort fallback
- **WHEN** no native variant exists but a portable representation can be used
- **THEN** the runtime plan includes the fallback and reports any lost semantics

