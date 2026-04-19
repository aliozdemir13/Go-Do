package todo

import (
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

func TestMegaLogo(t *testing.T) {
	got := MegaLogo()
	if !strings.Contains(got, "██") {
		t.Error("MegaLogo() should contain ASCII art blocks")
	}
	if !strings.Contains(got, ColorIndigo) {
		t.Error("MegaLogo() should be colored indigo")
	}
}
