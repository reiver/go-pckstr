package pckstr_test

import (
	"testing"

	"github.com/reiver/go-pckstr"
)

func TestPackedStrings_GoString(t *testing.T) {
	tests := []struct{
		PackedStrings pckstr.PackedStrings
		Expected string
	}{
		{
			Expected: "pckstr.Nothing()",
		},
		{
			PackedStrings: pckstr.Nothing(),
			Expected:     "pckstr.Nothing()",
		},



		{
			PackedStrings: pckstr.NoStrings(),
			Expected:     "pckstr.NoStrings()",
		},



		{
			PackedStrings: pckstr.SomeString("apple"),
			Expected:     `pckstr.SomeString("apple")`,
		},
		{
			PackedStrings: pckstr.SomeString("BANANA"),
			Expected:     `pckstr.SomeString("BANANA")`,
		},
		{
			PackedStrings: pckstr.SomeString("Cherry"),
			Expected:     `pckstr.SomeString("Cherry")`,
		},
		{
			PackedStrings: pckstr.SomeString("dAtE"),
			Expected:     `pckstr.SomeString("dAtE")`,
		},



		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA"),
			Expected:     `pckstr.SomeStrings("apple", "BANANA")`,
		},
		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA", "Cherry"),
			Expected:     `pckstr.SomeStrings("apple", "BANANA", "Cherry")`,
		},
		{
			PackedStrings: pckstr.SomeStrings("apple", "BANANA", "Cherry", "dAtE"),
			Expected:     `pckstr.SomeStrings("apple", "BANANA", "Cherry", "dAtE")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.PackedStrings.GoString()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'go-string' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			continue
		}
	}
}
