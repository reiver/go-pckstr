package strcodec_test

import (
	"testing"

	"github.com/reiver/go-pckstr/internal/strcodec"
)

func TestEncode(t *testing.T) {

	tests := []struct {
		String   string
		Expected string
	}{
		{},



		{
			String:   "",
			Expected: "",
		},



		{
			String:   "apple",
			Expected: "apple",
		},
		{
			String:   "BANANA",
			Expected: "BANANA",
		},
		{
			String:   "Cherry",
			Expected: "Cherry",
		},
		{
			String:   "dAtE",
			Expected: "dAtE",
		},



		{
			String:   "\x1B",
			Expected: "\x1B\x1B",
		},
		{
			String:   "\x1B\x1B",
			Expected: "\x1B\x1B\x1B\x1B",
		},



		{
			String:   "\x1F",
			Expected: "\x1B\x1F",
		},
		{
			String:   "\x1B\x1F",
			Expected: "\x1B\x1B\x1B\x1F",
		},



		{
			String:   "Once\x1FTwice\x1FThrice\x1FFource",
			Expected: "Once\x1B\x1FTwice\x1B\x1FThrice\x1B\x1FFource",
		},
	}

	for testNumber, test := range tests {

		actual := strcodec.Encode(test.String)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual encoded-value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("STRING:   %q", test.String)
			continue
		}
	}
}
