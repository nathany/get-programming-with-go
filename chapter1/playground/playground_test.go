package main

import (
	"testing"

	"github.com/nathany/learn-go/cmd"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
