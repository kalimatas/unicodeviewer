package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/kalimatas/unicodeviewer"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: unicodeviewer <unicode char>\n")
	os.Exit(2)
}

func consoleFmt(c unicodeviewer.Codepoint) {
	fmtString := `
Print:   %c
Unicode: %U
Hex:     %s
Binary:  %s
`
	fmt.Printf(fmtString, c.R, c.R, c.Hex(), c.Bin())
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	r, _ := utf8.DecodeRuneInString(os.Args[1])
	if r == utf8.RuneError {
		panic("invalid Unicode character")
	}

	consoleFmt(unicodeviewer.Codepoint{R: r})
}
