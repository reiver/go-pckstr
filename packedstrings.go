package pckstr

import (
	"github.com/reiver/go-opt"

	"github.com/reiver/go-pckstr/pck"
)

type PackedStrings struct {
	optional opt.Optional[string]
}

// Nothing return a 'nothing' optional-type value.
//
// [PackedStrings] is an optional-type.
func Nothing() PackedStrings {
	return PackedStrings{}
}

func NoStrings() PackedStrings {
	return Nothing()
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

func (receiver PackedStrings) IsNothing() bool {
	return receiver.optional.IsNothing()
}

func (receiver PackedStrings) Strings() []string {
	packed, something := receiver.optional.Get()

	if !something {
		return nil
	}

	return pck.Unpack(packed)
}
