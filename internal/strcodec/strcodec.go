package strcodec

import (
	"github.com/reiver/go-pckstr/internal/chars"
)

// AppendEncode encodes a string such that, any byte with value 0x1B (i.e., ESC) gets turned into 0x1B 0x1B (i.e., ESC ESC),
// and any byte with value 0x1F (i.e, US) gets turned into 0x1B 0x1F (i.e., ESC US).
//
// For example:
//
//	var str string = "Hello \x1B[7mworld\x1b[0m!"
//
//	p = strcodec.AppendEncode(p, str)
//
// See also:
//
//	• [Encode]
func AppendEncode(p []byte, str string) []byte {
	if len(str) <= 0 {
		return p
	}

	for index:=0; index<len(str); index++ {
		var b byte = str[index]

		switch b {
		case chars.EscapeByte, chars.SeparatorByte:
			p = append(p, chars.EscapeByte)
			p = append(p, b)
		default:
			p = append(p, b)
		}
	}

	return p
}

// Encode encodes a string such that, any byte with value 0x1B (i.e., ESC) gets turned into 0x1B 0x1B (i.e., ESC ESC),
// and any byte with value 0x1F (i.e, US) gets turned into 0x1B 0x1F (i.e., ESC US).
//
// For example:
//
//	var str string = "Hello \x1B[7mworld\x1b[0m!"
//
//	encoded := strcodec.Encode(str)
//	// encoded == "Hello \x1B\x1B[7mworld\x1B\x1B[0m!"
//	//                   ^^^^            ^^^^
//
// See also:
//
//	• [AppendEncode]
func Encode(str string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = AppendEncode(p, str)
	return string(p)
}

// Decode accepts a packed-strings and return the (decoded) head of the packed-strings
// and how many bytes to skip in the packed-strings to get the rest (after the head).
//
// For example:
//
//	var packed string = "Once\x1FTwice\x1FThrice\x1FFource"
//	
//	head, skip := strcodec.Decode(packed)
//	// head == "Once"
//	// skip == len("Once\x1F")
//	
//	rest := packed[skip:]
//	// rest == "Twice\x1FThrice\x1FFource"
func Decode(packed string) (string, int) {
	if len(packed) <= 0 {
		return "", 0
	}

	var encoded []byte

	var skip int // = 0

	var index int // = 0
	loop: for ; index<len(packed); index++ {
		var b byte = packed[index]

		switch b {
		case chars.EscapeByte:
			skip++
			index++
			if len(packed) <= index {
				break
			}

			b = packed[index]

			skip++
			encoded = append(encoded, b)
		case chars.SeparatorByte:
			skip++
			break loop
		default:
			skip++
			encoded = append(encoded, b)
		}
	}

	return string(encoded), skip
}
