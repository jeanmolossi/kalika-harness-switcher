# Kalika Harness Switcher

Kalika Harness Switcher (`khs`) is an experimental local control plane for
profiles and isolated sessions of native AI coding harnesses. It configures and
launches those tools while leaving their native behavior in the harnesses
themselves.

The project currently provides only its walking skeleton: a small command-line
interface and the development tooling needed to evolve it safely. It does not
yet implement profiles, sessions, daemon behavior, or harness adapters.

## Requirements

- Linux or macOS
- Go 1.26.0 or newer, with automatic toolchain selection enabled or the
  preferred Go 1.26.7 toolchain installed

The module's `go` and `toolchain` directives are the source of truth for the Go
version.

## Build and test

Build the `khs` executable into `bin/`:

```sh
make build
```

Run formatting verification, `go vet`, and unit tests:

```sh
make check
```

Run the race detector separately (supported on the initial Linux and macOS
platforms):

```sh
make test-race
```

## CLI usage

The implemented interface is intentionally small:

```text
khs help
khs --help
khs version
khs --version
```

Running `khs` without arguments also prints help. Unknown commands print a
diagnostic to standard error and exit with status 2.

## Platform path policy

All Go path construction must use `path/filepath`. Semantic roots must come
from the corresponding native Go resolver—such as `os.UserConfigDir`,
`os.UserCacheDir`, `os.UserHomeDir`, or `os.TempDir`—rather than inferred
environment conventions. A path category without a standard-library resolver
must use explicit, isolated per-OS implementations. Linux and macOS are the
initial supported platforms; Windows support is reserved for future work.
