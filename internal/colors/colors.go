package colors

import (
	"regexp"
	"strings"
)

// ANSI color codes
const (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"
	Grey    = "\x1b[90m"
	Orange  = "\x1b[38;5;214m"
	// Dim Colors
	DimYellow = "\x1b[38;5;144;48;238m"
	DimBlue   = "\x1b[38;5;66;48;237m"
	DimGreen  = "\x1b[38;5;72;48;237m"
	// Reset colors
	Reset = "\x1b[0m"
)

// getLogLevelColors returns ANSI color codes for log levels
func getLogLevelColors(word string) string {
	// Assign a color to each word
	switch strings.ToLower(word) {
	case "error":
		return Red
	case "warn":
		return Yellow
	case "info":
		return Blue
	case "debug":
		return DimBlue
	default:
		return ""
	}
}

// ColorizeLogLevels highlights log levels in a string using ANSI color codes
func ColorizeLogLevels(line string) string {
	var highlightWords = []string{"error", "warn", "info", "debug"}
	// case-insensitive regular expression to find the word
	re := regexp.MustCompile(`(?i)\b(` + strings.Join(highlightWords, "|") + `)\b`)

	// highlight word
	return re.ReplaceAllStringFunc(line, func(match string) string {
		return getLogLevelColors(match) + match + Reset
	})
}

// Highlight highlights a specific word in a string using ANSI color codes
func Highlight(line string, word string) string {
	re := regexp.MustCompile(`(?i)` + word)
	return re.ReplaceAllStringFunc(line, func(match string) string {
		return Orange + match + Reset
	})
}
