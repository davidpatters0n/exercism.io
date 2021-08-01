// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
	case questionAndShouting(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure"
	case isSilent(remark):
		return "Fine. Be that way!"
	case isShouting(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isQuestion(str string) bool {
	return strings.HasSuffix(sanitizeString(str), "?")
}

func questionAndShouting(str string) bool {
	return isQuestion(str) && isShouting(str) && !numbersOnly(str)
}

func isShouting(str string) bool {
	return strings.ToUpper(str) == str && !numbersOnly(str)
}

func isSilent(str string) bool {
	return sanitizeString(str) == ""
}

func sanitizeString(str string) string {
	return strings.Trim(str, " ")
}

func numbersOnly(str string) bool {
	newStr := prettifyString(str)
	return emptyString(sanitizeString(newStr))
}

func emptyString(str string) bool {
	return len(str) == 0
}

func prettifyString(str string) string {
	regex := regexp.MustCompile(`([0-9?:),])+`)
	return regex.ReplaceAllString(str, "")
}
