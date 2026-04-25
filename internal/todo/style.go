// Package todo provides functionality for the app
package todo

import (
	"fmt"
	"io"
)

const (
	// ColorIndigo used in styling
	ColorIndigo = "\033[38;5;99m"
	// BgIndigo used in styling
	BgIndigo = "\033[48;5;99m"
	// ColorRed used in styling
	ColorRed = "\033[31m"
	// ColorWhiteBold used in styling
	ColorWhiteBold = "\033[1;37m"
	// ColorReset used in styling
	ColorReset = "\033[0m"
	// ColorDim used in styling
	ColorDim = "\033[38;5;245m"
	// BgIndigoLight used in styling
	BgIndigoLight = "\033[48;5;105m" // A slightly lighter indigo for the accent
	// TextIndigo used in styling
	TextIndigo = "\033[38;5;141m" // A bright "glowing" indigo
)

// Dim styles the text. Use this like internal.Dim("text")
func Dim(text string) string {
	return ColorDim + text + ColorReset
}

// Indigo styles the text with indigo color. Use this like internal.Indigo("text")
func Indigo(text string) string {
	return ColorIndigo + text + ColorReset
}

// Red styles the text with red color. Use this like internal.Indigo("text")
func Red(text string) string {
	return ColorRed + text + ColorReset
}

// StyledBar styles the text with indigo background color. Use this like internal.StyledBar("text")
func StyledBar(text string) string {
	return fmt.Sprintf("%s %s %s%s", BgIndigo, ColorWhiteBold+text, ColorReset, "\n")
}

// FancyBar returns a formatted bar
func FancyBar(title string, version string) string {
	iconPart := BgIndigo + ColorWhiteBold + " ⚡ " + ColorReset
	titlePart := BgIndigoLight + ColorWhiteBold + " " + title + " " + ColorReset
	versionPart := "\033[38;5;239m" + " " + version + ColorReset // Dark gray version
	return iconPart + titlePart + versionPart + "\n"
}

// StatsBar returns a formatted string containing the progress info
func StatsBar(total, completed int) string {
	percentage := 0
	if total > 0 {
		percentage = (completed * 100) / total
	}

	// ORDER MATTERS:
	// 1. %s (Indigo Σ)
	// 2. %d (total)
	// 3. %s (Indigo ✓)
	// 4. %d (completed)
	// 5. %d (percentage)
	return fmt.Sprintf("   %s %d  %s %d  %s %d%%",
		Indigo("Σ (Total)"), total,
		Indigo("✓ (Done)"), completed,
		Dim("— Progress:"), percentage)
}

// MegaLogo shows the logo of the app
func megaLogo() string {
	return ColorIndigo + `
 ██████╗  ██████╗       ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗      ██╔══██╗██╔═══██╗
██║  ███╗██║   ██║█████╗██║  ██║██║   ██║
██║   ██║██║   ██║╚════╝██║  ██║██║   ██║
╚██████╔╝╚██████╔╝      ██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝       ╚═════╝  ╚═════╝                        
                                                             ` + ColorReset + "\n"
}

func PrintHeader(w io.Writer) {
	fmt.Fprint(w, megaLogo())
	fmt.Fprintf(w, "\n   %s %s %s\n",
		Indigo("⚡ TASKS"),
		Dim("v1.0.0"),
		Indigo("●")+" "+Dim("Local Storage Active"))
	fmt.Fprintln(w, ColorIndigo+"  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"+ColorReset)
	fmt.Fprintln(w)
}

func PrintProgress(w io.Writer, list *TodoList) {
	total, done := list.GetStats()
	fmt.Fprintf(w, "\n   %s  \n", StyledBar("Progress:"))
	fmt.Fprintln(w, StatsBar(total, done))
	fmt.Fprint(w, "   "+Indigo("● "))
	for i := 0; i < 20; i++ {
		if total > 0 && i < (done*20/total) {
			fmt.Fprint(w, Indigo("■"))
		} else {
			fmt.Fprint(w, Dim("□"))
		}
	}
	fmt.Fprint(w, "\n\n")
}
