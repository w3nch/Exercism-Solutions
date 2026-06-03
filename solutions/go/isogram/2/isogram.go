package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})

	for _, ch := range word {
		if ch == '-' || ch == ' ' {
			continue
		}

		ch = unicode.ToLower(ch)

		if _, exists := seen[ch]; exists {
			return false
		}

		seen[ch] = struct{}{}
	}

	return true
}