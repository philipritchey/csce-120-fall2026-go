package mostimportantword

import "strings"

type Set map[string]struct{}

// MostImportantWord finds the most frequent word that is not ignored
func MostImportantWord(s string, ignore []string) string {
	count := WordCounts(s)
	ignored := MakeSet(ignore)
	// find the most frequent and not ignored word
	maxCnt := 0
	var miw string
	for word, cnt := range count {
		_, ignoreWord := ignored[word]
		if cnt > maxCnt && !ignoreWord {
			maxCnt = cnt
			miw = word
		}
	}
	return miw
}

func WordCounts(s string) map[string]int {
	count := make(map[string]int)
	for word := range strings.FieldsSeq(s) {
		count[word]++
	}
	return count
}

func MakeSet(s []string) Set {
	set := make(Set)
	for _, word := range s {
		set[word] = struct{}{}
	}
	return set
}
