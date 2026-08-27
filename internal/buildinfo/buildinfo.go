// Package buildinfo provides metadata embedded in the khs executable.
package buildinfo

// These values are variables so release builds can replace them with -ldflags
// -X options. Their defaults are deterministic for local development and tests.
var (
	Version   = "dev"
	Commit    = "unknown"
	BuildDate = "unknown"
)

// Info is the build metadata reported by the CLI.
type Info struct {
	Version   string
	Commit    string
	BuildDate string
}

// Current returns the metadata embedded in this executable.
func Current() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		BuildDate: BuildDate,
	}
}
