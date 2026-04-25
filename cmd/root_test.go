package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"Go-Do/internal/todo"
)

func resetState(t *testing.T) {
	t.Helper()
	myList = todo.TodoList{}
	showDone = false
	t.Chdir(t.TempDir())
}

func runCmd(t *testing.T, args ...string) (string, error) {
	t.Helper()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return buf.String(), err
}

func TestRootCmd_Metadata(t *testing.T) {
	if rootCmd.Use != "go-do" {
		t.Errorf("expected use 'go-do', got %q", rootCmd.Use)
	}
	if rootCmd.Version == "" {
		t.Error("expected non-empty version")
	}
}

func TestRootCmd_Run(t *testing.T) {
	resetState(t)
	out, err := runCmd(t)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !strings.Contains(out, "OPEN TASKS") {
		t.Errorf("expected output to contain OPEN TASKS, got: %s", out)
	}
}

func TestExecute_Success(t *testing.T) {
	resetState(t)
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})
	if err := Execute(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestExecute_InvalidFlag(t *testing.T) {
	resetState(t)
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--non-existent-flag"})
	if err := Execute(); err == nil {
		t.Error("expected error for invalid flag, got nil")
	}
}

func TestAddCmd(t *testing.T) {
	resetState(t)
	out, err := runCmd(t, "add", "hello", "world")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !strings.Contains(out, "hello world") {
		t.Errorf("expected output to contain title, got: %s", out)
	}
	if len(myList.Tasks) != 1 || myList.Tasks[0].Title != "hello world" {
		t.Errorf("expected one task 'hello world', got %+v", myList.Tasks)
	}
	if _, err := os.Stat("tasks.json"); err != nil {
		t.Errorf("expected tasks.json to exist: %v", err)
	}
}

func TestAddCmd_SaveError(t *testing.T) {
	resetState(t)
	// Put a directory where the file should go so SaveToFile fails.
	if err := os.Mkdir("tasks.json", 0755); err != nil {
		t.Fatal(err)
	}
	if _, err := runCmd(t, "add", "will-fail"); err == nil {
		t.Error("expected save error when tasks.json is a directory")
	}
}

func TestCompleteCmd(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "add", "first"); err != nil {
		t.Fatal(err)
	}
	out, err := runCmd(t, "complete", "1")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !strings.Contains(out, "Task completed") {
		t.Errorf("expected completion message, got: %s", out)
	}
	if !myList.Tasks[0].IsDone {
		t.Error("expected task 1 to be marked done")
	}
}

func TestCompleteCmd_InvalidID(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "complete", "not-a-number"); err == nil {
		t.Error("expected error for non-numeric ID")
	}
}

func TestCompleteCmd_NotFound(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "complete", "999"); err == nil {
		t.Error("expected error for missing task")
	}
}

func TestDeleteCmd(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "add", "to-delete"); err != nil {
		t.Fatal(err)
	}
	out, err := runCmd(t, "delete", "1")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !strings.Contains(out, "Task deleted") {
		t.Errorf("expected deletion message, got: %s", out)
	}
	if len(myList.Tasks) != 0 {
		t.Errorf("expected no tasks after delete, got %+v", myList.Tasks)
	}
}

func TestDeleteCmd_InvalidID(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "delete", "xyz"); err == nil {
		t.Error("expected error for non-numeric ID")
	}
}

func TestDeleteCmd_NotFound(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "delete", "42"); err == nil {
		t.Error("expected error for missing task")
	}
}

func TestListCmd_Default(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "add", "task-a"); err != nil {
		t.Fatal(err)
	}
	out, err := runCmd(t, "list")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "OPEN TASKS") || !strings.Contains(out, "task-a") {
		t.Errorf("expected open tasks with task-a, got: %s", out)
	}
}

func TestListCmd_Done(t *testing.T) {
	resetState(t)
	if _, err := runCmd(t, "add", "task-done"); err != nil {
		t.Fatal(err)
	}
	if _, err := runCmd(t, "complete", "1"); err != nil {
		t.Fatal(err)
	}
	out, err := runCmd(t, "list", "--done")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "COMPLETED TASKS") {
		t.Errorf("expected completed section, got: %s", out)
	}
}
