## Context

See `proposal.md` for motivation. The repository currently contains only OpenSpec planning artifacts and a minimal README. The broader architecture requires a local daemon, IPC, PTYs, registries, adapters and a TUI, but none of those boundaries has yet been validated in code. The setup must therefore establish a reliable development loop without encoding speculative domain structure.

The initial platforms are Linux and macOS. Simplicity, few external dependencies, modularity, testability and security are project priorities.

## Goals / Non-Goals

**Goals:**

- Produce one compilable command with a stable public name and module identity.
- Make CLI behavior testable without spawning subprocesses.
- Establish deterministic development build information.
- Provide a small, shared interface for local and CI quality checks.
- Verify the project on Linux and macOS.
- Keep the initial dependency graph empty.

**Non-Goals:**

- Define domain entities or interfaces before their functional changes.
- Reserve commands for daemon, sessions, adapters or runtimes before their contracts exist.
- Introduce configuration loading, application directories or filesystem security policies.
- Select TUI, PTY, IPC, logging, release or dependency-injection libraries.
- Produce release artifacts, installers, containers or signed binaries.
- Enforce coverage thresholds or adopt a broad lint framework.

## Decisions

### Name the public command `khs`

The repository and Go module remain `github.com/jeanmolossi/kalika-harness-switcher`, while the user-facing executable is `khs`. The short name keeps future commands concise and avoids coupling CLI ergonomics to the repository name.

The entry point lives at `cmd/khs/main.go`. Client and daemon are expected to remain modes of this one executable unless a later lifecycle requirement justifies a second binary.

### Use Go 1.26 with a preferred patch toolchain

The module declares `go 1.26.0` as its minimum language/toolchain requirement and `toolchain go1.26.7` as the preferred development toolchain. Go 1.26 is the mature supported series at the time of the change; newly released Go 1.27 is deliberately not required by the boilerplate.

Using the native `toolchain` directive is preferred over an external version manager. CI reads the module declarations rather than duplicating an unrelated version constant where practical.

### Start with zero external Go dependencies

The initial CLI uses the standard library. Cobra or another command framework is deferred until real command depth, shared flags, completion or generated documentation provides evidence for it. Tests use `testing` and standard helpers. Logging, TUI and PTY dependencies are not selected here.

This keeps `go.mod` as the authoritative module/toolchain declaration without a `go.sum` until a real dependency requires one.

### Keep the entry point inert and the CLI directly testable

`main` delegates arguments, stdout and stderr to a CLI runner and exits with its returned code. The runner accepts `[]string` and `io.Writer` values, allowing behavior to be tested in process.

The only initial user-visible operations are help and version. Unknown input returns a concise diagnostic and a usage-error exit code. Future commands are not added as placeholders.

The initial package layout is intentionally small:

```text
cmd/khs/main.go
internal/buildinfo/
internal/cli/
```

Unused architectural directories are not created. New packages enter with the changes that own their behavior.

### Inject build metadata without runtime discovery

An internal build-information package exposes version, commit and build date. Development defaults are deterministic (`dev` and `unknown` values), while release-oriented builds can replace them with linker flags.

Build metadata does not include builder username, hostname or branch. The executable does not invoke Git at runtime, so packaged binaries behave consistently outside a checkout.

### Use a small Makefile as the developer interface

The Makefile provides `build`, `test`, `test-race`, `vet`, `fmt-check` and `check`. `check` composes formatting verification, vet and regular tests. Outputs go to an ignored `bin/` directory.

These targets wrap standard Go commands rather than introduce a task-runner dependency. CI may call the same targets where platform behavior is identical.

### Run focused CI on Linux and macOS

CI verifies formatting, vet, tests and build. The race detector runs on Linux; regular tests and build run on both Linux and macOS. Actions are pinned immutably by commit SHA with readable version comments to reduce workflow supply-chain ambiguity.

No custom module cache, coverage service, release job or third-party lint action is introduced. Dependency caching can be enabled when `go.sum` exists and produces a measurable benefit.

### Document only behavior that exists

The README identifies the project as an experimental local control plane that preserves native harness behavior, then documents toolchain requirements, build, test and the existing help/version interface. It does not advertise future daemon or session commands as implemented.

## Risks / Trade-offs

- [The standard-library CLI parser may later become cumbersome] → Keep parsing behind `internal/cli` and evaluate a framework only when real command complexity appears.
- [Go 1.26.7 may require automatic toolchain download] → Document the requirement and allow developers to install the preferred toolchain locally; do not add another version manager.
- [A Makefile is not native to every operating system] → Linux and macOS are the declared MVP platforms, and every target remains a documented standard Go command.
- [Build date harms reproducibility when populated with wall-clock time] → Development defaults remain deterministic and release automation must inject an explicit source-derived or controlled value.
- [Pinning actions by SHA increases maintenance] → Include version comments and update pins deliberately in maintenance changes.
- [The initial structure may look too small relative to the architecture] → Treat absence of speculative packages as intentional; functional changes add boundaries with their own tests and rationale.

## Migration Plan

This is a greenfield setup. Add the module and walking skeleton, make local checks pass, then enable CI. Rollback removes only the new Go, build and workflow files; no user data, runtime state or external system is affected.
