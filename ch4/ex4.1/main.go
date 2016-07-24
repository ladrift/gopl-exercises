// Exercise 4.1:
// Write a function that counts the number of bits that are different in
// two SHA256 hashes.
package main

import "fmt"

import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println(countDiffBits(c1, c2))
	// Output:
	// 125

	// Trivial test case
	h1 := [sha256.Size]byte{1}
	h2 := [sha256.Size]byte{0}
	fmt.Println(countDiffBits(h1, h2))
	// Output:
	// 1
}

// countDiffBits counts the number of bits that are different in two SHA256 hashes.
func countDiffBits(h1, h2 [sha256.Size]byte) int {
	count := 0
	for i, h := range h1 {
		b := h ^ h2[i]
		count += popCount(b)
	}
	return count
}

// popCount return the number of bits that are set in a byte.
func popCount(h byte) int {
	count := 0
	for h > 0 {
		count++
		h &= h - 1
	}
	return count
}
