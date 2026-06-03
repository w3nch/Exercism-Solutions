package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i := range word {
		if word[i] == '-' || word[i] == ' ' {
			continue
		}

		for j := range word {
			if word[j] == '-' || word[j] == ' ' {
				continue
			}

			if i != j && word[i] == word[j] {
				return false
			}
		}
	}

	return true
}