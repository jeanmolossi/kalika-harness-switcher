## Context

See `proposal.md` for motivation. The repository has no existing implementation. The first release targets Linux and macOS, must preserve native harness behavior, keep sessions alive after UI exit, and support third-party adapters. Config-home separation provides state isolation but is not a security sandbox. Claude Code, Codex, and Copilot have distinct configuration, authentication, state, and precedence behavior.

## Goals / Non-Goals

**Goals:**

- Keep the resolver and session lifecycle independent of vendor logic.
- Make runtime resolution deterministic, inspectable, and testable.
- Keep multiple interactive sessions alive after clients disconnect.
- Give official and third-party adapters the same integration contract.
- Apply logical isolation and honest security diagnostics without overstating guarantees.

**Non-Goals:**

- Reimplement agent loops, tools, conversations, or native session formats.
- Guarantee read-only behavior through config isolation alone.
- Recover attached PTYs after daemon failure or machine reboot in the MVP.
- Provide a strong OS/container sandbox in the MVP.
- Distribute, monetize, review, or automatically update adapters through a marketplace in the MVP.
- Synchronize runtime instances between machines in the MVP.

## Decisions

### Use a local daemon as the persistent session owner

The CLI and TUI connect to a single-user daemon over a permission-restricted Unix domain socket. The daemon owns PTYs, process groups, runtime metadata, bounded output tails, and lifecycle events. Closing a client only detaches it.

Tmux was considered because it already provides persistence and attach/detach. It remains a possible future terminal backend, but making it foundational would expose domain behavior to user tmux configuration, complicate structured lifecycle events, and impede a future Windows backend.

The MVP guarantees survival across client exit, not daemon failure, logout, or reboot. Lost records are reconciled on startup, but a daemon cannot recover PTY file descriptors it no longer owns.

### Use two narrow local protocols

Client-to-daemon IPC supports session commands, events, resize, input, and output streams. Length-prefixed JSON is sufficient for the MVP and avoids an RPC framework.

Daemon-to-adapter communication uses short-lived subprocesses with versioned JSON requests for `describe`, `probe`, `validate`, `plan`, and `command`. Official adapters use this same contract so third-party support is continuously exercised.

### Keep adapters declarative and process lifecycle centralized

Adapters detect tools, describe compatibility, translate resources, and return materialization and command plans. They do not create PTYs, launch harnesses, resolve secrets, or write arbitrary runtime paths. The core validates and executes filesystem operations and owns every child process.

Adapters are discovered only from configured directories through manifests, not arbitrary `PATH` lookup. First execution and content changes require trust based on origin and digest. Planning operations receive a minimal environment and no resolved secrets.

### Separate profile, security context, runtime spec, and runtime instance

A profile requests capabilities. A security context supplies bounded authority and credential references. Resolution produces an immutable runtime spec. Materialization produces a mutable per-session instance. Two otherwise identical sessions receive different vendor homes by default.

Permission requirements never grant authority. Effective read-only claims are classified as service-enforced, vendor-enforced, best-effort, or declared-only. Shell access prevents the MVP from promising strong containment.

### Resolve a versioned registry snapshot

Registry sources may be local directories or Git repositories. A snapshot records source identity, Git commit where applicable, resource digests, and schema versions. Duplicate IDs are errors unless an explicit override selects the source.

Resources prefer native vendor variants, then explicit adapter conversion, then portable fallback. Translation fidelity and omitted semantics are diagnostics, not silent behavior.

### Materialize atomically with late-bound secrets

Adapters return declarative directory, copy, render, link, mode, environment, and command operations. The core validates destinations, executes them under a staging root, writes a redacted manifest, and commits the directory atomically.

The child environment is constructed from an allowlist. Secret references are resolved only immediately before spawn, held in memory, and never written to manifests or diagnostic output.

### Preserve project and harness-owned state

Native project configuration remains in place and is loaded according to vendor rules. The adapter reports identifiable configuration layers and precedence, including managed layers that the control plane cannot override.

The runtime manifest records baseline hashes only for application-managed files. Changes to those files are presented as diffs. Vendor history, authentication, caches, and conversation state remain vendor-owned and are not promoted automatically.

### Require repository consent by canonical local path

Before first operation in a repository, the client shows its canonical path and detected configuration or executable hooks and offers session-only trust, persistent trust, or denial. Trust is local-path based in the MVP; a Git remote alone is not a sufficient identity.

### Keep UI selection immutable after launch

The TUI is a workspace dashboard. Attach temporarily gives the terminal to the remote PTY byte stream; a configurable reserved sequence detaches. A harness's own exit command ends the child naturally and emits an exit event. Changing profile, context, tool, or project creates a new runtime instead of mutating a live one.

## Risks / Trade-offs

- [Daemon failure loses PTY ownership] → State this limit explicitly and defer per-session supervisors or tmux recovery to resilient-supervisor evolution.
- [Third-party adapters execute local code] → Require manifest discovery, trust by digest, minimal environment, timeouts, no planning secrets, and clear origin diagnostics.
- [Vendor upgrades change schemas or precedence] → Record versions, declare tested ranges, warn on unknown versions, and use golden adapter tests.
- [Config-home isolation is mistaken for security] → Classify guarantees and avoid a global "secure readonly" label.
- [Project settings alter generated behavior] → Preserve them intentionally and expose them in effective-configuration diagnostics.
- [Harness mutates generated files] → Track managed baselines and offer explicit diffs without automatic registry writes.
- [Git registry changes reduce reproducibility] → Resolve against locked snapshots and retain source commits and digests.
- [Socket or metadata expose session content] → Use a user-private state directory, restrictive modes, bounded output, and no full transcript persistence by default.

## Migration Plan

This is a greenfield change. Deliver the daemon and fake-harness integration first, then registry/resolver/materializer, adapter protocol, official adapters, and finally CLI/TUI workflows. Each official adapter is enabled only for harness versions it can detect and diagnose. Rollback consists of stopping the daemon and removing generated runtime instances; canonical registry sources and vendor-global configuration are not mutated.

## Future Evolution

### Resilient supervisor

Add survival across daemon failure, logout, or reboot using per-session supervisors, a service manager, or an optional tmux backend. Define recovery semantics and process identity before claiming resumability; this work is not an incomplete MVP task.

### Strong sandbox

Add enforceable filesystem, network, process, and credential boundaries using OS-native mechanisms or containers. Security contexts can later select a sandbox backend. Existing MVP policy diagnostics remain defense in depth rather than enforcement.

### Adapter marketplace

Add discovery, publisher identity, signing, review metadata, checksummed installation, controlled updates, revocation, and compatibility channels. The MVP adapter directory and trust model are the local foundation, but do not perform marketplace distribution.

### Portable runtimes

Add exportable runtime manifests and content-addressed resources while excluding machine-bound paths, keychain identities, and vendor state that cannot be safely transferred.
