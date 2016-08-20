package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "%s: need provide a url as command-line arguments.\n", os.Args[0])
		os.Exit(1)
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("words =", words, "images =", images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
			return
		}
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		if n.Type == html.TextNode {
			words += countWords(n.Data)
		}
		visit(n.NextSibling)
		visit(n.FirstChild)
	}
	visit(n)
	return
}

func countWords(text string) int {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}
