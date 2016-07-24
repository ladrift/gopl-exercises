// Exercise 4.2:
// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print print the SHA384 or SHA512 hash
// instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var hashType = flag.String("hash", "sha256", "hash type")

func main() {
	flag.Parse()
	var h hash.Hash
	switch *hashType {
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		fmt.Fprintf(os.Stderr,
			"%s: %s hash type not exists.\n", os.Args[0], *hashType)
		os.Exit(1)
	}

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"%s: copy from stdin: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	hash := h.Sum(nil)
	fmt.Printf("The %s hash is %x.\n", *hashType, hash)
}
