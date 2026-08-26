## 1. Foundation

- [ ] 1.1 Initialize the Go module, command entry point, and internal package boundaries and verify `go test ./...` succeeds
- [ ] 1.2 Define domain types for profiles, capabilities, resources, security contexts, runtime specs, runtime instances, and sessions and verify serialization and validation unit tests
- [ ] 1.3 Implement private application-directory discovery and secure file modes for Linux and macOS and verify filesystem tests reject permissive or escaping paths

## 2. Registry and Resolution

- [ ] 2.1 Define versioned YAML schemas and loaders for profiles, capabilities, resources, security contexts, and source configuration and verify valid and invalid fixtures
- [ ] 2.2 Implement local and Git registry source snapshots with commit and content digests and verify repeat resolution uses the locked snapshot
- [ ] 2.3 Implement graph expansion, cycle detection, explicit overrides, and conflict diagnostics and verify resolver table tests
- [ ] 2.4 Implement separation of capability requirements from security-context authority and verify insufficient authority never becomes an implicit grant
- [ ] 2.5 Record provenance and translation candidates in immutable runtime specs and verify inspect fixtures show the complete resolution chain

## 3. Adapter Protocol and Trust

- [ ] 3.1 Define the versioned adapter manifest and JSON protocol for describe, probe, validate, plan, and command and verify protocol round-trip and compatibility tests
- [ ] 3.2 Implement adapter discovery from configured directories without implicit PATH discovery and verify duplicate and invalid manifests are rejected
- [ ] 3.3 Implement adapter trust by origin and executable digest and verify new or changed adapters require consent before execution
- [ ] 3.4 Implement the constrained adapter subprocess host with minimal environment, timeout, bounded output, and no resolved secrets and verify adversarial fake-adapter tests
- [ ] 3.5 Implement fidelity diagnostics for native, converted, degraded, and unsupported resources and verify unsupported semantics are never silently omitted

## 4. Runtime Materialization and Credentials

- [ ] 4.1 Define declarative materialization operations and implement staging-root validation, secure modes, atomic commit, and rollback and verify traversal and symlink attack tests
- [ ] 4.2 Implement distinct runtime instances and vendor homes per session and verify identical launch requests do not share mutable paths
- [ ] 4.3 Implement redacted runtime manifests with tool, adapter, registry, configuration-layer, and resource provenance data and verify no fixture leaks secret values
- [ ] 4.4 Implement allowlisted child-environment construction and secret-reference providers for explicit environment and external command sources and verify unrelated daemon credentials are removed
- [ ] 4.5 Implement managed-file baseline hashing and diff generation while classifying vendor-owned state separately and verify modified configuration produces a reviewable diff

## 5. Persistent Daemon and PTYs

- [ ] 5.1 Implement single-user daemon startup, lock ownership, Unix socket permissions, version handshake, and idle connection handling and verify client reconnection tests
- [ ] 5.2 Implement length-prefixed client protocol messages for session commands, lifecycle events, terminal input/output, and resize and verify codec and malformed-message tests
- [ ] 5.3 Implement PTY creation, process groups, resize, wait, exit status, and bounded output tails on Linux and macOS and verify behavior with a fake harness
- [ ] 5.4 Implement the session state machine and persistent metadata store and verify invalid transitions and startup reconciliation are diagnosed
- [ ] 5.5 Implement attach with a single controlling writer and configurable detach sequence and verify detach preserves the harness process and a later client can reattach
- [ ] 5.6 Implement graceful process-group stop followed by forced termination after a deadline and verify child subprocesses do not remain alive
- [ ] 5.7 Verify closing or crashing a client leaves multiple daemon-owned harness sessions running and attachable

## 6. Repository Trust and Effective Configuration

- [ ] 6.1 Implement canonical repository identification and session-only, persistent, and denied trust decisions and verify an unknown path cannot launch without consent
- [ ] 6.2 Detect relevant native project configuration and executable hooks before trust approval and verify the consent summary reports detected sources
- [ ] 6.3 Carry native project configuration layers into runtime diagnostics without copying or modifying project files and verify adapter precedence output

## 7. Official Adapters

- [ ] 7.1 Implement the Claude Code adapter using the external protocol and verify golden tests for detection, config-home materialization, resources, environment, command, and project layers
- [ ] 7.2 Implement the OpenAI Codex adapter using the external protocol and verify golden tests for detection, config-home materialization, resources, environment, command, and project layers
- [ ] 7.3 Implement the GitHub Copilot CLI adapter using the external protocol and verify golden tests for detection, config-home materialization, resources, environment, command, and project layers
- [ ] 7.4 Verify all official adapters report native and best-effort resource fidelity consistently using shared conformance fixtures
- [ ] 7.5 Build a third-party sample adapter outside the core adapter packages and verify it can be installed and used without changing resolver, materializer, or session-manager code

## 8. CLI and Terminal UI

- [ ] 8.1 Implement CLI commands for daemon status and session create, list, inspect, attach, stop, and delete and verify command-level integration tests
- [ ] 8.2 Implement the TUI session dashboard with profile, security context, tool, project, lifecycle, attention, and exit status and verify view-model tests
- [ ] 8.3 Implement the new-session workflow with profile, context, tool, project, arguments, trust confirmation, and pre-launch diagnostics and verify a complete fake-harness flow
- [ ] 8.4 Implement terminal attach/detach mode and natural harness-exit handling and verify both the reserved key and the harness `exit` path return control correctly
- [ ] 8.5 Prevent mutation of runtime selection for running sessions and verify profile or tool changes create a new-runtime workflow

## 9. MVP Verification and Documentation

- [ ] 9.1 Run unit, protocol, materializer, resolver, daemon, PTY, adapter conformance, and CLI/TUI integration tests on Linux and macOS and document the commands and results
- [ ] 9.2 Perform a manual concurrent-session acceptance test with Claude Code, Codex, and Copilot, including UI closure, reattach, natural exit, forced stop, and configuration diff
- [ ] 9.3 Document security guarantees and limitations, especially config isolation, best-effort read-only policies, shell escape risk, child environment filtering, and adapter trust
- [ ] 9.4 Document explicitly deferred evolution for resilient supervision across daemon failure/reboot, strong sandbox enforcement, adapter marketplace, and portable runtimes without adding them to MVP completion criteria
