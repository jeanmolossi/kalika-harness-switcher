## 1. Go Module and Command Skeleton

- [x] 1.1 Initialize `github.com/jeanmolossi/kalika-harness-switcher` with `go 1.26.0` and `toolchain go1.26.7`, and verify `go env GOMOD` resolves to the repository module
- [x] 1.2 Create the `cmd/khs` entry point delegating arguments and standard streams to a testable CLI runner, and verify `go build -o bin/khs ./cmd/khs` succeeds
- [x] 1.3 Implement standard-library help, version, and unknown-input handling without placeholder future commands, and verify unit tests cover output streams and exit codes
- [x] 1.4 Establish and document the path rule that all Go path construction MUST use `filepath`, semantic roots MUST come from native resolvers such as `os.UserConfigDir`, `os.UserCacheDir`, `os.UserHomeDir`, and `os.TempDir`, and categories without a standard resolver MUST use explicit per-OS implementations; verify a source audit finds no inferred environment convention, string-concatenated path, or hardcoded Linux, macOS, or Windows directory

## 2. Build Information and Tests

- [x] 2.1 Add internal build information with deterministic development defaults for version, commit, and build date, and verify unit tests do not depend on Git or wall-clock state
- [x] 2.2 Support build-information replacement through linker flags, and verify a built `bin/khs --version` reports injected fixture values
- [x] 2.3 Add focused tests for `khs help`, `khs --help`, `khs version`, `khs --version`, and invalid input, and verify `go test ./...` passes without external Go dependencies

## 3. Developer Workflow

- [x] 3.1 Add `bin/` and Go-generated local artifacts to `.gitignore` without ignoring source or OpenSpec files, and verify a local build leaves no untracked binary
- [x] 3.2 Add Makefile targets for `build`, `test`, `test-race`, `vet`, `fmt-check`, and `check`, and verify `make check` plus `make build` succeed
- [x] 3.3 Update the README with the experimental project purpose, native-harness boundary, Go requirement, build, test, existing CLI usage, and the rule for `filepath`, native path resolvers, and isolated OS-specific behavior, and verify every documented command is implemented and the platform guidance matches the OpenSpec project context

## 4. Continuous Integration

- [x] 4.1 Add a GitHub Actions workflow using immutable action SHAs with version comments and a Linux/macOS matrix, and verify its configuration selects the module-declared Go toolchain
- [ ] 4.2 Configure CI to run formatting verification, vet, tests, Linux race tests, and builds without release, coverage-service, or third-party linter jobs, and verify all jobs pass on the change branch
- [x] 4.3 Run `go test ./...`, `go vet ./...`, `make check`, and a clean `make build`, then verify the repository contains no external Go dependencies or speculative domain packages
