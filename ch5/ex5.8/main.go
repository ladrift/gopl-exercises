// Exercise 5.8:
// Modify forEachNode so that the pre and post functions return
// a boolean result indicating whether to continue the traversal.
// Use it to write a function ElementByID with the following signature
// that finds the first HTML element with the specified id attribute.
// The function should stop the traversal as soon as a match is found.
//	func ElementByID(doc *html.Node, id string) *html.Node
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stderr, "%s: Need a URL and a id as arguments\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetching %s failed: %v\n", url, err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parsing %s failed: %v\n", url, err)
	}
	resp.Body.Close()
	e := ElementByID(doc, id)
	if e == nil {
		fmt.Printf("element whose id is %s not found in %s.\n", id, url)
		os.Exit(1)
	}
	fmt.Printf("Found element: %#v\n", e)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var elmt *html.Node
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					elmt = n
					return false
				}
			}
		}
		return true
	}
	forEachNode(doc, pre, nil)
	return elmt
}
