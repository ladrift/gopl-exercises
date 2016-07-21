// Exercise 3.10:
// Write a non-recusive version of comma, using bytes.Buffer instead of
// string concatenation.
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	count := utf8.RuneCountInString(s)
	prefixLen := count % 3
	if prefixLen > 0 {
		buf.WriteString(s[:prefixLen])
		buf.WriteByte(',')
	}

	for i, r := range s[prefixLen:] {
		if i > 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	return buf.String()
}
