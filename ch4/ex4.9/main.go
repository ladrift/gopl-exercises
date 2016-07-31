// Exericse 4.9:
// Write a program wordfreq to report the frequency of each word in an input
// text file. Call input.Split(bufio.ScanWords) before the first call to Scan
// to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Set up an word-oriented scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	wordFreq := make(map[string]int)
	for scanner.Scan() {
		wordFreq[scanner.Text()]++
	}

	for k, v := range wordFreq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
