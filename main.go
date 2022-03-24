package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	h := md5.New()
	if _, err := io.Copy(h, os.Stdin); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
