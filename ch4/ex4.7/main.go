// Exercise 4.7:
// Modify reverse to reverse the characters of a []byte slice that represents
// a UTF-8-encoding string, in place. Can you do it without allocating new
// memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("啊哈哈，我爱菲宝宝～")
	reverseUTF8(b)
	fmt.Printf("%s\n", string(b))
}

func reverseUTF8(b []byte) {
	// Reverse the multi-bytes UTF-8 code point
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	// Reverse whole []byte slice
	reverse(b)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
