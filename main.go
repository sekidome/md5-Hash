package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagFile = flag.String(
		"file",
		"",
		"wenn gesetzt wird direkt dieses File verwendet")
	flagURL = flag.String(
		"url",
		"",
		"wenn gesetzt wird direkt von diesem URL geladen")
)

func main() {
	flag.Parse()
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	switch {
	case *flagFile != "":
		fd, err := os.Open(*flagFile)
		if err != nil {
			fmt.Fprintln(output, "Fehler beim Öffnen der File")
			os.Exit(1)
		}
		defer fd.Close()
		input = fd
	case *flagURL != "":
		resp, err := http.Get(*flagURL)
		if err != nil {
			fmt.Fprintln(output, "Fehler beim Laden der URL")
			os.Exit(1)
		}
		defer resp.Body.Close()
		input = resp.Body
	}
	printMD5(input, os.Stdout)
}

func printMD5(r io.Reader, w io.Writer) {
	h := md5.New()
	_, err := io.Copy(h, r)
	if err != nil {
		fmt.Println("Fehler beim Lesen", err)
		return
	}
	fmt.Fprintf(w, "%x\n", h.Sum(nil))
}
