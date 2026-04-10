package strcodec_test

import (
	"testing"

	"github.com/reiver/go-pckstr/internal/strcodec"
)

func TestDecode_ignoreSkip(t *testing.T) {

	tests := []struct {
		Encoded  string
		Expected string
	}{
		{},



		{
			Encoded:  "apple",
			Expected: "apple",
		},
		{
			Encoded:  "apple\x1F",
			Expected: "apple",
		},
		{
			Encoded:  "BANANA",
			Expected: "BANANA",
		},
		{
			Encoded:  "BANANA\x1F",
			Expected: "BANANA",
		},
		{
			Encoded:  "Cherry",
			Expected: "Cherry",
		},
		{
			Encoded:  "Cherry\x1F",
			Expected: "Cherry",
		},
		{
			Encoded:  "dAtE",
			Expected: "dAtE",
		},
		{
			Encoded:  "dAtE\x1F",
			Expected: "dAtE",
		},



		{
			Encoded:  "\x1B",
			Expected: "",
		},
		{
			Encoded:  "\x1B\x1B",
			Expected: "\x1B",
		},



		{
			Encoded:  "\x1F",
			Expected: "",
		},
		{
			Encoded:  "\x1B\x1F",
			Expected: "\x1F",
		},



		{
			Encoded:  "Once",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1F",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice\x1F",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice\x1FThrice",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice\x1FThrice\x1F",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice\x1FThrice\x1FFource",
			Expected: "Once",
		},
		{
			Encoded:  "Once\x1FTwice\x1FThrice\x1FFource\x1F",
			Expected: "Once",
		},
	}

	for testNumber, test := range tests {

		actual, _ := strcodec.Decode(test.Encoded)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual decoded-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("ENCODED: %q", test.Encoded)
			continue
		}
	}
}
