package cliutil

import (
	"os"
	"testing"
)

func TestClearPreviousLine(t *testing.T) {
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

	ClearPreviousLine()

	w.Close()

	output := make([]byte, 100)
	n, err := r.Read(output)
	if err != nil {
		t.Fatalf("Error reading from pipe: %v", err)
	}

	expectedOutput := "\033[1A\033[2K"

	if string(output[:n]) != expectedOutput {
		t.Errorf("ClearConsole() output %q, expected %q", string(output[:n]), expectedOutput)
	}
}
