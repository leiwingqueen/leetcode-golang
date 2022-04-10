package string

import "strings"

var arr = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
	".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...",
	"-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	mp := make(map[string]struct{})
	for _, word := range words {
		s := strings.Builder{}
		for i := 0; i < len(word); i++ {
			s.WriteString(arr[word[i]-'a'])
		}
		mp[s.String()] = struct{}{}
	}
	return len(mp)
}
