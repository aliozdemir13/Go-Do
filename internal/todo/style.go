package todo

import "fmt"

const (
	ColorIndigo    = "\033[38;5;99m"
	BgIndigo       = "\033[48;5;99m"
	ColorWhiteBold = "\033[1;37m"
	ColorReset     = "\033[0m"
	ColorDim       = "\033[38;5;245m"
	BgIndigoLight  = "\033[48;5;105m" // A slightly lighter indigo for the accent
	TextIndigo     = "\033[38;5;141m" // A bright "glowing" indigo
)

// Function: Use this like internal.Dim("text")
func Dim(text string) string {
	return ColorDim + text + ColorReset
}

// Function: Use this like internal.Indigo("text")
func Indigo(text string) string {
	return ColorIndigo + text + ColorReset
}

func StyledBar(text string) string {
	return fmt.Sprintf("%s %s %s%s", BgIndigo, ColorWhiteBold+text, ColorReset, "\n")
}

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

func MegaLogo() string {
	return ColorIndigo + `
 ██████╗  ██████╗       ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗      ██╔══██╗██╔═══██╗
██║  ███╗██║   ██║█████╗██║  ██║██║   ██║
██║   ██║██║   ██║╚════╝██║  ██║██║   ██║
╚██████╔╝╚██████╔╝      ██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝       ╚═════╝  ╚═════╝                        
                                                             ` + ColorReset + "\n"
}
