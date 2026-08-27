// Package cli implements the khs command-line interface.
package cli

import (
	"fmt"
	"io"

	"github.com/jeanmolossi/kalika-harness-switcher/internal/buildinfo"
)

const usage = `Usage: khs <command>

Commands:
  help       Show this help
  version    Show build information

Options:
  -h, --help       Show this help
  -v, --version    Show build information
`

// Run executes the CLI with explicit streams and returns a process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return 0
	}

	if len(args) == 1 {
		switch args[0] {
		case "help", "-h", "--help":
			fmt.Fprint(stdout, usage)
			return 0
		case "version", "-v", "--version":
			info := buildinfo.Current()
			fmt.Fprintf(stdout, "khs %s (commit %s, built %s)\n", info.Version, info.Commit, info.BuildDate)
			return 0
		}
	}

	fmt.Fprintf(stderr, "khs: unknown command: %s\nRun 'khs help' for usage.\n", args[0])
	return 2
}
