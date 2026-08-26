## Purpose

Define how portable profiles become deterministic, explainable runtime inputs without conflating requested functionality with execution authority.

## ADDED Requirements

### Requirement: Resolve profiles declaratively
The system SHALL expand a selected profile into capabilities and resources while preserving the origin of every resolved item.

#### Scenario: Inspect resolved provenance
- **WHEN** a user inspects a resolved profile
- **THEN** the system reports which profile, capability, resource, and registry source introduced each item

### Requirement: Separate capability from authority
The system SHALL treat capability requirements and security-context authority as separate inputs and SHALL NOT grant authority while resolving a capability.

#### Scenario: Insufficient authority
- **WHEN** a capability requires database read access and the selected security context does not provide it
- **THEN** resolution fails with an explanation of the unmet requirement

### Requirement: Resolve registry snapshots deterministically
The system SHALL resolve resources from local and Git sources using an identified snapshot and SHALL reject undeclared resource conflicts.

#### Scenario: Conflicting resource identifiers
- **WHEN** two active sources provide the same resource identifier without an explicit override
- **THEN** resolution fails instead of selecting a resource by incidental load order

