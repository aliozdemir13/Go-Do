package todo

import (
	"bytes"
	"strings"
	"testing"
)

func TestColorWrappers(t *testing.T) {
	// Test Dim
	input := "hello"
	gotDim := Dim(input)
	if !strings.Contains(gotDim, ColorDim) || !strings.Contains(gotDim, ColorReset) {
		t.Error("Dim() should wrap text in Dim and Reset colors")
	}

	// Test Indigo
	gotIndigo := Indigo(input)
	if !strings.Contains(gotIndigo, ColorIndigo) || !strings.Contains(gotIndigo, ColorReset) {
		t.Error("Indigo() should wrap text in Indigo and Reset colors")
	}

	// Test Red
	gotRed := Red(input)
	if !strings.Contains(gotRed, ColorRed) || !strings.Contains(gotRed, ColorReset) {
		t.Error("Red() should wrap text in Red and Reset colors")
	}
}

func TestStyledBar(t *testing.T) {
	text := "MY TASKS"
	got := StyledBar(text)

	if !strings.Contains(got, text) {
		t.Errorf("StyledBar() missing text: %s", text)
	}
	if !strings.HasSuffix(got, "\n") {
		t.Error("StyledBar() should end with a newline")
	}
}

func TestFancyBar(t *testing.T) {
	title := "TODO"
	version := "v1.0"
	got := FancyBar(title, version)

	if !strings.Contains(got, title) || !strings.Contains(got, version) {
		t.Error("FancyBar() should contain both title and version")
	}
	if !strings.Contains(got, "⚡") {
		t.Error("FancyBar() should contain the lightning icon")
	}
}

func TestStatsBar(t *testing.T) {
	tests := []struct {
		name      string
		total     int
		completed int
		wantPerc  string
	}{
		{"Empty list", 0, 0, "0%"},
		{"Half done", 10, 5, "50%"},
		{"All done", 4, 4, "100%"},
		{"None done", 5, 0, "0%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StatsBar(tt.total, tt.completed)
			if !strings.Contains(got, tt.wantPerc) {
				t.Errorf("StatsBar(%d, %d) = %v; want percentage %s", tt.total, tt.completed, got, tt.wantPerc)
			}

			// Verify it includes the labels
			if !strings.Contains(got, "Total") || !strings.Contains(got, "Done") {
				t.Error("StatsBar should contain Total and Done labels")
			}
		})
	}
}

func TestPrintHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	PrintHeader(buf)
	got := buf.String()

	tests := []string{
		"TASKS",
		"v1.0.0",
		"Local Storage Active",
		"━━━━━━",
	}

	for _, want := range tests {
		if !strings.Contains(got, want) {
			t.Errorf("PrintHeader() output missing expected string: %q", want)
		}
	}

	if buf.Len() == 0 {
		t.Error("PrintHeader() produced no output")
	}
}

func TestPrintProgress(t *testing.T) {
	tests := []struct {
		name          string
		total         int
		done          int
		wantDoneIcons int // How many ■ we expect
		wantTodoIcons int // How many □ we expect
		list          *TodoList
	}{
		{
			name:          "Zero tasks (Empty State)",
			total:         0,
			done:          0,
			wantDoneIcons: 0,
			wantTodoIcons: 20,
			list: &TodoList{
				Tasks: []Task{
					{Title: "Learn Go", IsDone: false},
					{Title: "Write Tests", IsDone: false},
				},
			},
		},
		{
			name:          "Half completed",
			total:         10,
			done:          5,
			wantDoneIcons: 10, // (5 * 20 / 10) = 10
			wantTodoIcons: 10,
			list: &TodoList{
				Tasks: []Task{
					{Title: "Learn Go", IsDone: false},
					{Title: "Write Tests", IsDone: true},
				},
			},
		},
		{
			name:          "Fully completed",
			total:         10,
			done:          10,
			wantDoneIcons: 20, // (10 * 20 / 10) = 20
			wantTodoIcons: 0,
			list: &TodoList{
				Tasks: []Task{
					{Title: "Learn Go", IsDone: true},
					{Title: "Write Tests", IsDone: true},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)

			PrintProgress(buf, tt.list)
			got := buf.String()

			if !strings.Contains(got, "Progress:") {
				t.Errorf("Missing label 'Progress:'")
			}

			doneCount := strings.Count(got, "■")
			todoCount := strings.Count(got, "□")

			if doneCount != tt.wantDoneIcons {
				t.Errorf("Progress bar math wrong: got %d '■', want %d", doneCount, tt.wantDoneIcons)
			}
			if todoCount != tt.wantTodoIcons {
				t.Errorf("Progress bar math wrong: got %d '□', want %d", todoCount, tt.wantTodoIcons)
			}
		})
	}
}
