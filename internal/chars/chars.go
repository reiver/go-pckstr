package chars

import (
	"codeberg.org/reiver/go-ascii"
)

const (
	EscapeByte    = byte(ascii.ESC)
	SeparatorByte = byte(ascii.US)
)

const (
	EscapeString    = string(rune(ascii.ESC))
	SeparatorString = string(rune(ascii.US))
)

const (
	EscapedEscapeString    = EscapeString + EscapeString
	EscapedSeparatorString = EscapeString + SeparatorString
)
