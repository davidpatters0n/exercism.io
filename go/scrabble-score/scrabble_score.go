package scrabble

import (
	"strings"
)

var scrabbleLetters = map[int][]string{
	1:  []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  []string{"D", "G"},
	3:  []string{"B", "C", "M", "P"},
	4:  []string{"F", "H", "V", "W", "Y"},
	5:  []string{"K"},
	8:  []string{"J", "X"},
	10: []string{"Q", "Z"},
}

func Score(input string) int {
	total := 0
	wordToArray := strings.Split(input, "")
	for _, word := range wordToArray {
		for score, scrabbleWords := range scrabbleLetters {
			if contains(scrabbleWords, word) {
				total += score
			}
		}
	}
	return total
}

func contains(collection []string, word string) bool {
	for _, v := range collection {
		if strings.ToLower(v) == strings.ToLower(word) {
			return true
		}
	}
	return false
}
