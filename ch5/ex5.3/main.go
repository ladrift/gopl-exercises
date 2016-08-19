// Exercise 5.3:
// Write a function to print the contents of all text nodes in an HTML document tree.
// Do not descend into <script> or <style> elements, since their contents are not visible in a web browser.
package main

import (
	"bytes"
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
	text := gatherText(doc)
	fmt.Println(text)
}

func gatherText(n *html.Node) string {
	var f func(text *bytes.Buffer, n *html.Node)
	f = func(text *bytes.Buffer, n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.TextNode {
			text.WriteString(n.Data)
		}
		f(text, n.NextSibling)
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return
		}
		f(text, n.FirstChild)
	}
	text := new(bytes.Buffer)
	f(text, n)
	return text.String()
}
