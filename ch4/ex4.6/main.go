// Exercise 4.6:
// Write an in-place function that squashes each run of adjacent Unicode spaces
// in UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := `啊实打实 
dasd  打死		.`
	b := squashSpaces([]byte(s))
	fmt.Printf("%s\n", string(b))
}

func squashSpaces(bs []byte) []byte {
	var flag bool
	rbs := bs[:0]
	for len(bs) > 0 {
		r, size := utf8.DecodeRune(bs)
		if unicode.IsSpace(r) {
			if !flag {
				flag = true
				rbs = append(rbs, ' ')
			}
		} else {
			if flag {
				flag = false
			}
			rbs = append(rbs, bs[:size]...)
		}
		bs = bs[size:]
	}
	return rbs
}
