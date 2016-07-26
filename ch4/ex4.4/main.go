// Exercise 4.4:
// Write a version of rotate that operates in a single pass.
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Printf("%v\n", s)
}

// rotate rotate the s slice with n offset.
func rotate(s []int, n int) {
	a := make([]int, n, n)
	for i := 0; i < len(s); i++ {
		if i < n {
			// Reserve first n integers.
			a[i] = s[i]
		}
		if i+n < len(s) {
			// Shift n size.
			s[i] = s[i+n]
		} else {
			// Copy the reserved integers.
			s[i] = a[i+n-len(s)]
		}
	}
}
