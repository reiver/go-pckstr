package pckstr

import (
	gobytes "bytes"
	"encoding/json"
	"strconv"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-pckstr/pck"
)

var (
	null []byte = []byte{'n','u','l','l'}
)

type PackedStrings struct {
	optional opt.Optional[string]
}

// Nothing return a 'nothing' optional-type value.
//
// [PackedStrings] is an optional-type.
//
// Nothing should not be confused with [NoStrings].
func Nothing() PackedStrings {
	return PackedStrings{}
}

func NoStrings() PackedStrings {
	return PackedStrings{
		optional: opt.Something(pck.Pack()),
	}
}

func SomeString(str string) PackedStrings {
	return PackedStrings{
		optional: opt.Something(pck.Pack(str)),
	}
}

func SomeStrings(strs ...string) PackedStrings {
	return PackedStrings{
		optional: opt.Something(pck.Pack(strs...)),
	}
}

func (receiver PackedStrings) GoString() string {
	switch {
	case Nothing() == receiver:
		return "pckstr.Nothing()"
	case NoStrings() == receiver:
		return "pckstr.NoStrings()"
	}

	strings := receiver.Strings()

	switch len(strings) {
	case 1:
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = append(p, "pckstr.SomeString("...)
		p = strconv.AppendQuote(p, strings[0])
		p = append(p, ')')

		return string(p)
	default:
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = append(p, "pckstr.SomeStrings("...)
		for i, str := range strings {
			if 0 < i {
				p = append(p, ", "...)
			}

			p = strconv.AppendQuote(p, str)
		}
		p = append(p, ')')

		return string(p)
	}
}

func (receiver PackedStrings) IsNothing() bool {
	return receiver.optional.IsNothing()
}

func (receiver PackedStrings) LenZero() bool {
	packed, something := receiver.optional.Get()

	switch {
	case !something:
		return true
	case "" == packed:
		return true
	case "\x1F" == packed:
		return true
	default:
		return false
	}
}

func (receiver PackedStrings) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.Strings())
}

func (receiver PackedStrings) Strings() []string {
	packed, something := receiver.optional.Get()

	if !something {
		return nil
	}

	return pck.Unpack(packed)
}

func (receiver *PackedStrings) UnmarshalJSON(bytes []byte) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	if gobytes.Equal(null, bytes) {
		*receiver = Nothing()
		return nil
	}

	var strs []string
	err := json.Unmarshal(bytes, &strs)
	if nil != err {
		return err
	}

	*receiver = SomeStrings(strs...)
	return nil
}
