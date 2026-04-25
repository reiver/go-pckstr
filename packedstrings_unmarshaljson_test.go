package pckstr_test

import (
	"testing"

	"github.com/reiver/go-pckstr"
)

func TestPackedStrings_UnmarshalJSON(t *testing.T) {
	tests := []struct{
		Bytes    []byte
		Expected pckstr.PackedStrings
	}{
		{
			Bytes: []byte(`null`),
			Expected: pckstr.Nothing(),
		},



		{
			Bytes: []byte(`[]`),
			Expected: pckstr.NoStrings(),
		},



		{
			Bytes: []byte(`[""]`),
			Expected: pckstr.SomeStrings(""),
		},



		{
			Bytes: []byte(`["apple"]`),
			Expected: pckstr.SomeStrings("apple"),
		},
		{
			Bytes: []byte(`["BANANA"]`),
			Expected: pckstr.SomeStrings("BANANA"),
		},
		{
			Bytes: []byte(`["Cherry"]`),
			Expected: pckstr.SomeStrings("Cherry"),
		},
		{
			Bytes: []byte(`["dAtE"]`),
			Expected: pckstr.SomeStrings("dAtE"),
		},



		{
			Bytes: []byte(`["once"]`),
			Expected: pckstr.SomeStrings("once"),
		},
		{
			Bytes: []byte(`["once","twice"]`),
			Expected: pckstr.SomeStrings("once", "twice"),
		},
		{
			Bytes: []byte(`["once","twice","thrice"]`),
			Expected: pckstr.SomeStrings("once", "twice", "thrice"),
		},
		{
			Bytes: []byte(`["once","twice","thrice","fource"]`),
			Expected: pckstr.SomeStrings("once", "twice", "thrice", "fource"),
		},
	}

	for testNumber, test := range tests {

		var actual pckstr.PackedStrings

		err := actual.UnmarshalJSON(test.Bytes)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTES: (len=%d)\n%s", len(test.Bytes), test.Bytes)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
			t.Logf("EXPECTED:\n%#v\n%#v", expected, expected.Strings())
			t.Logf("ACTUAL:\n%#v\n%#v", actual, actual.Strings())
			continue
		}
	}
}
