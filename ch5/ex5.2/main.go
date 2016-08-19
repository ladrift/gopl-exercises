// Exercise 5.2:
// Write a function to populate a mapping from element names p, div, span, and so on
// to the number of elements with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	m := make(map[string]int)
	visit(m, doc)
	fmt.Println(m)
}

// function visit populates the map from tag name to count number
func visit(m map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	visit(m, n.NextSibling)
	visit(m, n.FirstChild)
}
