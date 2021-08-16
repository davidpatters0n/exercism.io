package isogram

import "strings"

func IsIsogram(word string) bool {
	collectionWord := strings.Split(
		strings.ToLower(strings.Replace(strings.Replace(word, "-", "", -1), " ", "", -1)), "")
	return strings.Join(collectionWord, "") == unique(collectionWord)
}

func unique(arr []string) string {
	m := make(map[string]bool)
	uniqueWords := []string{}
	for _, word := range arr {
		if m[word] {
			continue
		} else {
			m[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}
	return strings.Join(uniqueWords, "")
}
