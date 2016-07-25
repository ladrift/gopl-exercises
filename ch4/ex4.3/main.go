// Exercise 4.3:
// Rewrite reverse to use an array pointer instead of a slice.
package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&a)
	fmt.Printf("%v\n", a)
}

// reverse reverses the 10-size integer array.
// Because the size of an array is the part of its type,
// so we can use slice to get more flexible with variadic size.
func reverse(p *[10]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		tmp := (*p)[i]
		(*p)[i] = (*p)[j]
		(*p)[j] = tmp
	}
}
