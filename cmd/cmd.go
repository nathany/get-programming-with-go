// Package cmd provides helper functions for testing main packages.
package cmd

import (
	"os/exec"
	"testing"
)

// CheckOutput from running a Go program against expectations.
func CheckOutput(t *testing.T, file, expected string) {
	actual := Run(t, file)
	if expected != actual {
		t.Errorf("Expected output %q, got %q.", expected, actual)
	}
}

// Run a Go program and capture the output.
func Run(t *testing.T, file string) string {
	out, err := exec.Command("go", "run", file).Output()
	if e, ok := err.(*exec.ExitError); ok {
		t.Fatalf("%s", e.Stderr)
	} else if err != nil {
		t.Fatalf("%s", err)
	}
	return string(out)
}
