package cliutil

import "fmt"

// ClearPreviousLine clears the previous line on the console.
func ClearPreviousLine() {
	fmt.Print("\033[1A")
	fmt.Print("\033[2K")
}
