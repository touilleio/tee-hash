package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	outputFile = flag.String("output", "hash.sha256", "output file")
)

func main() {

	flag.Parse()

	hasher := sha256.New()

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	c := io.TeeReader(r, hasher)
	_, err := io.Copy(w, c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying stdin to stdout, err = %v", err)
		os.Exit(1)
	}

	cksum := hex.EncodeToString(hasher.Sum(nil))
	err = os.WriteFile(*outputFile, []byte(cksum), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got writing hash to file %s, err = %v", *outputFile, err)
		os.Exit(1)
	}
}
