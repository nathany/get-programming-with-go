package main

import (
	"testing"

	"github.com/nathany/get-programming-with-go/cmd"
)

func TestPlayground(t *testing.T) {
	cmd.CheckOutput(t, "playground.go", "Hello, playground\n")
}
