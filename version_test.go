package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestVersionBuild(t *testing.T) {
	command := exec.Command("go", "run", ".", "--version")
	out, err := command.Output()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "0.2.0") {
		t.Fatalf("bad version: %s", out)
	}
}
