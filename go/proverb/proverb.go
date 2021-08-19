// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	rhymes := []string{}

	if len(rhyme) == 0 {
		return rhymes
	}
	for idx, word := range rhyme {
		str := ""
		if len(rhyme) > idx+1 {
			str = fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[idx+1])
			rhymes = append(rhymes, str)
		} else {
			str = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			rhymes = append(rhymes, str)
		}
	}
	fmt.Println(rhymes)
	return rhymes
}
