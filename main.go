// Package sha512_256 calculates the sha512_256 hash for STDIN or first file argument.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"crypto/sha512"
)

func getReader() io.Reader {
	if len(flag.Args()) > 0 {
		r, err := os.Open(flag.Arg(0))
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
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage:", os.Args[0], "[file]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) > 1 {
		flag.Usage()
		os.Exit(64) // command line usage error
	}

	fmt.Printf("%x\n", sum512_256(getReader()))
}
