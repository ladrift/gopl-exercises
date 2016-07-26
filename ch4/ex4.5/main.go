// Exercise 4.5:
// Write an in-place function to eliminate adjacent duplicates
// in a []string slice.
package main

import "fmt"

func main() {
	s := []string{"abc", "abc",
		"i love feifei",
		"haha", "haha", "haha",
		"aha"}
	s = dedup(s)
	fmt.Printf("%q\n", s)
}

func dedup(ss []string) []string {
	ds := ss[:1]
	for _, s := range ss[1:] {
		if ds[len(ds)-1] != s {
			ds = append(ds, s)
		}
	}
	return ds
}
