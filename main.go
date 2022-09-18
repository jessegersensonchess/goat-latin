// returns string in goat latin
package main

import (
	"strings"
)

func goatLatin(s string) string {

	if len(s) == 0 {
		return ""
	}

	var sb strings.Builder
	var a strings.Builder
	// calculate sb's final size -- the sum of 1+2+3+4 is n(n+1)/2
	wordCount := strings.Count(s, " ")
	bufferSize := len(s) + int((wordCount*(wordCount+1))/2) + (wordCount * len("ma"))
	sb.Grow(bufferSize)

	var prev rune
	var first rune
	for i, v := range s {
		// first letter of a word
		if v != ' ' {
			if i == 0 || prev == ' ' {
				first = v
				a.WriteRune('a')
				if strings.ContainsAny(string(first), "aeiouAEIOU") {
					sb.WriteRune(v)
				}
			} else {
				sb.WriteRune(v)
			}
		}

		// first space after a word. append ending
		if v == ' ' && prev != ' ' {
			if strings.ContainsAny(string(first), "aeiouAEIOU") {
				sb.WriteString("ma" + a.String() + " ")
			} else {
				sb.WriteString(string(first) + "ma" + a.String() + " ")
			}
		}
		prev = v

	}
	return strings.TrimRight(sb.String(), " ")

}

func main() {
	s := "asdlfkjasdf aslkdfjasdflkj asdflkjasdf lkjasdflkj asdflkjasdflkjasdflkjasdflkj asdlfkjasdflkjasdflkjasdfklkasjdflkjasdflkjasd asdflkjasdflkjasdflkjasdf asdflkjasdflkjasdflkjaas asdflkjasdflkjasdflkjasdfasdlfkjasdf aslkdfjasdflkj asdflkjasdf lkjasdflkj asdflkjasdflkjasdflkjasdflkj asdlfkjasdflkjasdflk"

	println(len(goatLatin(s)))

}
