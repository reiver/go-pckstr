package pckstr_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-pckstr"
)

func TestPackedStrings_MarshalJSON(t *testing.T) {
	tests := []struct{
		PackedStrings pckstr.PackedStrings
		Expected      string
	}{
		{
			Expected: `null`,
		},
		{
			PackedStrings: pckstr.Nothing(),
			Expected: `null`,
		},



		{
			PackedStrings: pckstr.NoStrings(),
			Expected: `[]`,
		},



		{
			PackedStrings: pckstr.SomeStrings(""),
			Expected: `[""]`,
		},



		{
			PackedStrings: pckstr.SomeStrings("apple"),
			Expected: `["apple"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("BANANA"),
			Expected: `["BANANA"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("Cherry"),
			Expected: `["Cherry"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("dAtE"),
			Expected: `["dAtE"]`,
		},



		{
			PackedStrings: pckstr.SomeStrings("once"),
			Expected: `["once"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("once", "twice"),
			Expected: `["once","twice"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("once", "twice", "thrice"),
			Expected: `["once","twice","thrice"]`,
		},
		{
			PackedStrings: pckstr.SomeStrings("once", "twice", "thrice", "fource"),
			Expected: `["once","twice","thrice","fource"]`,
		},
	}

	for testNumber, test := range tests {

		actual, err := test.PackedStrings.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("PACKED-STRINGS: %#v", test.PackedStrings)
			t.Logf("PACKED-STRINGS: %#v", test.PackedStrings.Strings())
			continue
		}

		expected := []byte(test.Expected)

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
			t.Logf("EXPECTED: (len=%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (len=%d)\n%s", len(actual), actual)
			continue
		}
	}
}
