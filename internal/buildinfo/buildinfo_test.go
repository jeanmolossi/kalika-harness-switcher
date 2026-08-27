package buildinfo

import "testing"

func TestDevelopmentDefaults(t *testing.T) {
	got := Current()
	want := Info{Version: "dev", Commit: "unknown", BuildDate: "unknown"}
	if got != want {
		t.Fatalf("Current() = %#v, want %#v", got, want)
	}
}
