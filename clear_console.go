// Package ciutil provides utility functions for working with the command line
package cliutil

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearConsole clears the console. It is the equivalent of the 'clear' command in bash.
func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
