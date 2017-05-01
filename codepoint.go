package unicodeviewer

import (
	"fmt"

	"github.com/fatih/color"
)

// Codepoint reprensent a rune
type Codepoint struct {
	R rune
}

// Hex returns a base 16 string of the R rune.
func (c *Codepoint) Hex() (res string) {
	for _, b := range c.bytes() {
		res += fmt.Sprintf("% 8X ", b)
	}
	return
}

// Bin returns a base 2 string of the R rune
// with colored prefix and data bits.
func (c *Codepoint) Bin() (res string) {
	isASCII := false

	// ASCII characters start with "0" and don't
	// have prefix.
	if c.bytes()[:1][0]>>7 == 0 {
		isASCII = true
	}

	for _, b := range c.bytes() {
		var isPrefix bool
		if !isASCII {
			isPrefix = true
		}

		for i := 0; i < 8; i++ {
			bit := (b << uint(i)) >> 7

			if isPrefix {
				res += color.RedString("%d", bit)
			} else {
				res += color.GreenString("%d", bit)
			}

			if bit == 0 && !isASCII {
				isPrefix = false
			}
		}
		res += " "
	}

	return
}

func (c *Codepoint) bytes() []byte {
	return []byte(string(c.R))
}
