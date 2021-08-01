// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreivate santizes the string
// splits the string by an empty space
// grabs the first char from the string
// and appends it to the slice and lastly
// we'll convert all strings to an uppercase and
// join it togehter.
func Abbreviate(s string) string {
	str := SanitizeString(s)

	wordArr := []string{}
	for _, word := range strings.Split(str, " ") {
		wordArr = append(wordArr, string(word[0]))
	}
	return strings.ToUpper(strings.Join(wordArr, ""))
}

// Sanitize will match on `-_` and replace it with an
// empty string we'll remove all surrounding white space
// and re-join the string
func SanitizeString(str string) string {
	regex := regexp.MustCompile(`([-_])`)
	str = regex.ReplaceAllString(str, " ")
	return strings.Join(strings.Fields(str), " ")
}
