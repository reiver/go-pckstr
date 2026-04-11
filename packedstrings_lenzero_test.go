package pckstr_test

import (
	"testing"

	"github.com/reiver/go-pckstr"
)

func TestPackedStrings_LenZero(t *testing.T) {
	tests := []struct{
		PackedStrings pckstr.PackedStrings
		Expected bool
	}{
		{
			Expected: true,
		},
		{
			PackedStrings: pckstr.Nothing(),
			Expected: true,
		},



		{
			PackedStrings: pckstr.NoStrings(),
			Expected: true,
		},



		{
			PackedStrings: pckstr.SomeStrings(),
			Expected: true,
		},



		{
			PackedStrings: pckstr.SomeString("apple"),
			Expected: false,
		},
		{
			PackedStrings: pckstr.SomeString("BANANA"),
			Expected: false,

		},
		{
			PackedStrings: pckstr.SomeString("Cherry"),
			Expected: false,
		},
		{
			PackedStrings: pckstr.SomeString("dAtE"),
			Expected: false,
		},



		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA"),
			Expected: false,
		},
		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA", "Cherry"),
			Expected: false,

		},
		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA", "Cherry", "dAtE"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.PackedStrings.LenZero()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'len-zero' value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("PACKED-STRINGS: %#v", test.PackedStrings)
			continue
		}
	}
}
