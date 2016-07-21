// Exercise 3.12:
// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s and %s is anagram.\n",
			os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s and %s is not anagram.\n",
			os.Args[1], os.Args[2])
	}
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	m2 := make(map[rune]int)
	for _, r := range s2 {
		m2[r]++
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}
	return true
}
