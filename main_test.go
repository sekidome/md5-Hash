package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("go")
	buf := &bytes.Buffer{}
	printMD5(in, buf)
	want := "34d1f91fb2e514b8576fab1a75a89a6b\n"
	got := buf.String()
	fmt.Println(got, "Hi")
	fmt.Println(want, "Hi2")
	var x bool = got == want
	fmt.Println(x)
	if got != want {
		t.Errorf("printMD5()=%s\nwant:%s\n", got, want)
	}
}
