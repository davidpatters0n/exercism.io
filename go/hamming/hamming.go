package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot identify hamming difference")
	}

	s1 := strings.Split(a, "")
	s2 := strings.Split(b, "")
	count := 0

	for idx, char := range s1 {
		if char != s2[idx] {
			count += 1
		}
	}
	return count, nil
}
