package raindrops

import "strconv"

type sound struct {
	factor int
	word   string
}

func Convert(number int) string {
	result := ""

	sounds := []sound{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	for _, s := range sounds {
		if number%s.factor == 0 {
			result += s.word
		}
	}

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}