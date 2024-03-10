package cliutil

import (
	"os"
	"runtime"
	"testing"
)

func TestClearConsole(t *testing.T) {
	originalStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer func() {
		os.Stdout = originalStdout
		r.Close()
		w.Close()
	}()

	os.Stdout = w

	ClearConsole()

	w.Close()

	output := make([]byte, 100)
	n, err := r.Read(output)
	if err != nil {
		t.Fatalf("Error reading from pipe: %v", err)
	}

	expectedClearCommand := "\x1b[H\x1b[2J\x1b[3J"
	if runtime.GOOS == "windows" {
		expectedClearCommand = "cmd /c cls"
	}

	if string(output[:n]) != expectedClearCommand {
		t.Errorf("ClearConsole() output %q, expected %q", string(output[:n]), expectedClearCommand)
	}
}
