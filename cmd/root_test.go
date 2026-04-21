package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestRootCmd(t *testing.T) {
	// Verify metadata (Covers the variable initialization)
	if rootCmd.Use != "go-do" {
		t.Errorf("expected go-do, got %s", rootCmd.Use)
	}

	// Capture output to prevent clutter
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)

	// Execute the Run anonymous function (Covers the Run: func block)
	if rootCmd.Run != nil {
		rootCmd.Run(rootCmd, []string{})
	}
}

func TestExecute_Success(t *testing.T) {
	// Use --help to ensure Execute() finishes without an error
	rootCmd.SetArgs([]string{"--help"})

	// Prevent output from printing during test
	rootCmd.SetOut(new(bytes.Buffer))

	// This covers the success branch of the if err := ... block
	Execute()
}

func TestExecute_ExitPath(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		// Provide an invalid flag to force an error
		rootCmd.SetArgs([]string{"--non-existent-flag"})
		// This will call Execute(), hit fmt.Println, and then os.Exit(1)
		Execute()
		return
	}

	// Re-run the current test function but with the BE_CRASHER environment variable
	cmd := exec.Command(os.Args[0], "-test.run=TestExecute_ExitPath")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()

	// Verify that the process actually exited with an error (status 1)
	e, ok := err.(*exec.ExitError)
	if !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1. Coverage for os.Exit(1) depends on this.", err)
	}
}
