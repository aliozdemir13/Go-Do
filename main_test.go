package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// We need to save the original arguments and restore them after the test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// We set the command line arguments to "lumina --help"
	// This ensures cmd.Execute() runs successfully and returns nil
	os.Args = []string{"lumina", "--help"}

	// Calling main() will now:
	// 1. Run internal.PrintHeader() -> (Covered!)
	// 2. Run cmd.Execute() with help args -> (Covered!)
	main()
}
