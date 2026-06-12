package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
const (
	recommendationMarker = '❗' // U+2757
	searchMarker         = '🔍' // U+1F50D
	weatherMarker        = '☀' // U+2600
)

func Application(log string) string {
	for _, r := range log {
		switch r {
		case recommendationMarker:
			return "recommendation"

		case searchMarker:
			return "search"

		case weatherMarker:
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result []rune

	for _, r := range log {
		if r == oldRune {
			r = newRune
		}
		result = append(result, r)
	}

	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
