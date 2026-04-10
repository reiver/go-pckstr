package pck_test

import (
	"testing"

	"github.com/reiver/go-pckstr/pck"
)

func TestCanonical(t *testing.T) {

	tests := []struct{
		Packed   string
		Expected string
	}{
		{
			Packed:   "",
			Expected: "",
		},
		{
			Packed:   "\x1F",
			Expected: "\x1F",
		},



		{
			Packed:   "H",
			Expected: "H\x1F",
		},
		{
			Packed:   "He",
			Expected: "He\x1F",
		},
		{
			Packed:   "Hel",
			Expected: "Hel\x1F",
		},
		{
			Packed:   "Hell",
			Expected: "Hell\x1F",
		},
		{
			Packed:   "Hello",
			Expected: "Hello\x1F",
		},
		{
			Packed:   "Hello\x1F",
			Expected: "Hello\x1F",
		},
		{
			Packed:   "Hello\x1Fw",
			Expected: "Hello\x1Fw\x1F",
		},
		{
			Packed:   "Hello\x1Fwo",
			Expected: "Hello\x1Fwo\x1F",
		},
		{
			Packed:   "Hello\x1Fwor",
			Expected: "Hello\x1Fwor\x1F",
		},
		{
			Packed:   "Hello\x1Fworl",
			Expected: "Hello\x1Fworl\x1F",
		},
		{
			Packed:   "Hello\x1Fworld",
			Expected: "Hello\x1Fworld\x1F",
		},
		{
			Packed:   "Hello\x1Fworld!",
			Expected: "Hello\x1Fworld!\x1F",
		},
	}

	for testNumber, test := range tests {
		actual := pck.Canonical(test.Packed)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual canonical-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("PACKED:   %q", test.Packed)
			continue
		}
	}
}
