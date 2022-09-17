// returns string in goat latin
package main

import (
	"bytes"
	"strings"
)

func goatLatin(s string) string {
	if len(s) == 0 {
		return ""
	}

	var sb strings.Builder
	// calculate sb's final size -- the sum of 1+2+3+4 is n(n+1)/2
	wordCount := len(bytes.Fields([]byte(s)))
	bufferSize := len(s) + int((wordCount*(wordCount+1))/2) + (wordCount * len("ma"))
	sb.Grow(bufferSize)

	for i, v := range bytes.Fields([]byte(s)) {
		if strings.ContainsAny(string(v[0]), "aeiouAEIOU") {
			// first letter of word is a vowel
			sb.WriteString(string(v) + "ma" + strings.Repeat("a", i+1) + " ")
		} else {
			// first letter of word is a consonant
			if len(v) == 1 {
				sb.WriteString(string(v) + "ma" + strings.Repeat("a", i+1) + " ")
			} else {
				sb.WriteString(string(v[1:]) + string(v[0]) + "ma" + strings.Repeat("a", i+1) + " ")
			}
		}
	}
	return strings.TrimRight(sb.String(), " ")
}

func main() {
		s := "This is an example of goat latin"
		println(goatLatin(s))

}
