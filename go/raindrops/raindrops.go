package raindrops

import (
	"fmt"
	"strings"
)

func Convert(input int) string {
	word := []string{}

	if input%3 == 0 {
		word = append(word, "Pling")
	}
	if input%5 == 0 {
		word = append(word, "Plang")
	}
	if input%7 == 0 {
		word = append(word, "Plong")
	}

	finalWord := strings.Join(word, "")

	if len(finalWord) == 0 {
		return fmt.Sprint(input)
	} else {
		return finalWord
	}
}
