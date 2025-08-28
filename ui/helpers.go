package ui

import (
	"bufio"
	"fmt"
	"os"
	"resumecli/utils"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// TerminalWidth holds the width of the terminal (used for layout/word wrap).
var TerminalWidth int

// Initialize terminal width on package load (fallback to 80 if detection fails).
func init() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		TerminalWidth = 80 // Default fallback width
	} else {
		TerminalWidth = width
	}
}

// typeWithColor animates character-by-character printing of `text` using `c` color and delay.
func typeWithColor(c *color.Color, text string, delay time.Duration) {
	for _, char := range text {
		_, err := c.Printf("%c", char)
		if err != nil {
			return
		}
		time.Sleep(delay)
	}
}

// createLink creates a terminal-supported clickable hyperlink (only supported in modern terminals).
func createLink(text, url string) string {
	return fmt.Sprintf("\x1b]8;;%s\x07%s\x1b]8;;\x07", url, text)
}

// printSectionHeader displays a stylized section header with underline.
func printSectionHeader(title string) {
	fmt.Println()
	titleColor := color.New(color.FgHiWhite, color.Bold)
	lineColor := color.New(color.FgHiCyan)

	fmt.Print(" ")
	_, err := titleColor.Println(strings.ToUpper(title))
	if err != nil {
		return
	}
	_, err = lineColor.Println(" " + strings.Repeat("─", len(title)+2))
	if err != nil {
		return
	}
	fmt.Println()
}

// PromptToContinue waits for the user to press Enter before proceeding.
func PromptToContinue() {
	fmt.Println()
	promptColor := color.New(color.FgYellow)
	_, err := promptColor.Print("\nPress Enter to return to menu...")
	if err != nil {
		return
	}
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		return
	}
}

// ShowIntro displays an animated personal introduction message.
func ShowIntro() {
	hiWhiteBold := color.New(color.FgHiWhite, color.Bold)
	yellowBold := color.New(color.FgHiYellow, color.Bold)
	greenBold := color.New(color.FgHiGreen, color.Bold)
	grey := color.New(color.FgWhite)
	typingDelay := 25 * time.Millisecond

	fmt.Println()

	// Animated typing intro
	typeWithColor(hiWhiteBold, "Hey there! ", typingDelay)
	typeWithColor(grey, "I'm ", typingDelay)
	typeWithColor(yellowBold, "Siva Prakash", typingDelay)
	typeWithColor(grey, ", I'm a ", typingDelay)
	typeWithColor(greenBold, "machine learning enthusiast", typingDelay)
	typeWithColor(grey,
		" from India and currently working on self-paced mini-projects in data analytics, software engineering, and machine learning.",
		typingDelay)

	fmt.Println()
}

// ShowExitMessage clears the screen and displays a centered farewell message.
func ShowExitMessage() {
	utils.CallClear()
	cyanBold := color.New(color.FgHiCyan, color.Bold)

	line1 := "Thank you for visiting!"
	line2 := "Have a great day! 👋"

	// Center the messages based on terminal width
	padding1 := (TerminalWidth - len(line1)) / 2
	padding2 := (TerminalWidth - len(line2)) / 2

	fmt.Printf("\n\n")
	fmt.Printf("%s", strings.Repeat(" ", padding1))
	_, err := cyanBold.Println(line1)
	if err != nil {
		return
	}

	fmt.Printf("%s", strings.Repeat(" ", padding2))
	fmt.Println(line2)
	fmt.Printf("\n\n")
}
