package pck

import (
	"strings"

	"github.com/reiver/go-pckstr/internal/chars"
	"github.com/reiver/go-pckstr/internal/strcodec"
)

// Canonical returns the canonical version of a raw packed-strings.
//
// It basically makes sure there is an unescaped '\x1F' at the
// end of a non-empty packed-strings if there isn't one there already.
func Canonical(packed string) string {
	if "" == packed {
		return ""
	}

	if strings.HasSuffix(packed, chars.EscapedSeparatorString) {
		return packed + chars.SeparatorString
	}

	if !strings.HasSuffix(packed, chars.SeparatorString) {
		return packed + chars.SeparatorString
	}

	return packed
}

// Pack returns a raw packed-strings.
//
// Pack does the opposite of [Unpack].
func Pack(strs ...string) string {
	if len(strs) <= 0 {
		return ""
	}

	var buffer [256]byte
	var packed []byte = buffer[0:0]

	for _, str := range strs {
		packed = strcodec.AppendEncode(packed, str)
		packed = append(packed, chars.SeparatorByte)
	}

	return string(packed)
}

// Unpack returns the []string from a packed-strings.
//
// Unpack does the opposite of [Pack].
func Unpack(packed string) []string {
	if len(packed) <= 0 {
		return []string{}
	}

	rest := packed

	var unpacked []string
	for {
		head, skip := strcodec.Decode(rest)
		rest = rest[skip:]

		unpacked = append(unpacked, head)

		if len(rest) <= 0 {
/////// BREAK
			break
		}
	}

	return unpacked
}
