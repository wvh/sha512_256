// Package sha512_265 calculates the sha512_265 hash for STDIN or first file argument.
package main

import (
	"fmt"
	"io"
	"os"

	"crypto/sha512"
)

func getReader() io.Reader {
	if len(os.Args) > 1 {
		r, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("error opening file:", err)
			os.Exit(66) // cannot open input
		}
		return r
	}

	return os.Stdin
}

func sum512_256(data io.Reader) []byte {
	hash := sha512.New512_256()
	if _, err := io.Copy(hash, data); err != nil {
		fmt.Println("error creating checksum:", err)
		os.Exit(74) // i/o error
	}
	return hash.Sum(nil)
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("usage:", os.Args[0], "[file]")
		os.Exit(64) // command line usage error
	}

	fmt.Printf("%x\n", sum512_256(getReader()))
}
