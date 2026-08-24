package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}

	number := n
	steps := 0

	for number != 1 {
		if number%2 == 0 {
			number /= 2
		} else {
			number = number*3 + 1
		}

		steps++
	}

	return steps, nil
}