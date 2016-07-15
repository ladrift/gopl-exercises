/*
Exercise 1.4:
Modify dup2 to print the names of all files in which each duplicated line occurs.

Example usage:
$ cd ch1/ex1.4
$ go build
$ ./ex1.4 dup_text/*
3       you     dup_text/dup2.txt
5       dup     dup_text/dup1.txt, dup_text/dup2.txt
6       love    dup_text/dup2.txt, dup_text/dup3.txt
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineCountFiles struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]lineCountFiles)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr,
					"%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(f, filename, counts)
			f.Close()
		}
	}
	for line, lcf := range counts {
		if lcf.count > 1 {
			fmt.Printf("%d\t%s\t%s", lcf.count, line, lcf.files[0])
			for _, file := range lcf.files[1:] {
				fmt.Printf(", %s", file)
			}
			fmt.Println("")
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]lineCountFiles) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count := counts[input.Text()].count
		files := counts[input.Text()].files
		counts[input.Text()] = lineCountFiles{count + 1, appendFiles(files, filename)}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func appendFiles(files []string, file string) []string {
	for _, f := range files {
		if f == file {
			return files
		}
	}
	return append(files, file)
}
