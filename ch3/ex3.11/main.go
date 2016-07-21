// Exercise 3.11:
// Enhance comma so that it deals correctly with floating-point numbers and
// optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	// Handle optional sign
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	dotIdx := strings.IndexByte(s, '.')
	inte := s[:dotIdx]
	frac := s[dotIdx+1:]
	// Write integer part
	prefix := utf8.RuneCountInString(inte) % 3
	if prefix > 0 {
		buf.WriteString(inte[:prefix])
		buf.WriteByte(',')
	}
	for i, r := range inte[prefix:] {
		if i > 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	// Write fractional part
	buf.WriteByte('.')
	for i, r := range frac {
		if i > 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}

	return buf.String()
}
