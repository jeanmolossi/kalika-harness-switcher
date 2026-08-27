package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{name: "no arguments", wantCode: 0, wantStdout: "Usage: khs <command>"},
		{name: "help command", args: []string{"help"}, wantCode: 0, wantStdout: "Usage: khs <command>"},
		{name: "help flag", args: []string{"--help"}, wantCode: 0, wantStdout: "Usage: khs <command>"},
		{name: "short help flag", args: []string{"-h"}, wantCode: 0, wantStdout: "Usage: khs <command>"},
		{name: "version command", args: []string{"version"}, wantCode: 0, wantStdout: "khs dev (commit unknown, built unknown)\n"},
		{name: "version flag", args: []string{"--version"}, wantCode: 0, wantStdout: "khs dev (commit unknown, built unknown)\n"},
		{name: "short version flag", args: []string{"-v"}, wantCode: 0, wantStdout: "khs dev (commit unknown, built unknown)\n"},
		{name: "unknown input", args: []string{"launch"}, wantCode: 2, wantStderr: "khs: unknown command: launch\nRun 'khs help' for usage.\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if got := Run(tt.args, &stdout, &stderr); got != tt.wantCode {
				t.Errorf("Run() code = %d, want %d", got, tt.wantCode)
			}
			if !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, want it to contain %q", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != tt.wantStderr {
				t.Errorf("stderr = %q, want %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}

func TestRunRejectsExtraArguments(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if got := Run([]string{"help", "extra"}, &stdout, &stderr); got != 2 {
		t.Fatalf("Run() code = %d, want 2", got)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q, want empty", stdout.String())
	}
}
