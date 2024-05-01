package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		return steps, errors.New("too many loops")
	}
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		steps++
	}
	return steps, nil
}
