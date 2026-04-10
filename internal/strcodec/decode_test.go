package strcodec_test

import (
	"testing"

	"github.com/reiver/go-pckstr/internal/strcodec"
)

func TestDecode(t *testing.T) {
	tests := []struct{
		Packed         string
		ExpectedString string
		ExpectedIndex  int
	}{
		{},


		{
			Packed:            "apple",
			ExpectedString:    "apple",
			ExpectedIndex: len("apple"),
		},
		{
			Packed:            "BANANA",
			ExpectedString:    "BANANA",
			ExpectedIndex: len("BANANA"),
		},
		{
			Packed:            "Cherry",
			ExpectedString:    "Cherry",
			ExpectedIndex: len("Cherry"),
		},
		{
			Packed:            "dAtE",
			ExpectedString:    "dAtE",
			ExpectedIndex: len("dAtE"),
		},



		{
			Packed:            "Once",
			ExpectedString:    "Once",
			ExpectedIndex: len("Once"),
		},
		{
			Packed:            "Once\x1FTwice",
			ExpectedString:    "Once",
			ExpectedIndex: len("Once\x1F"),
		},
		{
			Packed:            "Once\x1FTwice\x1FThrice",
			ExpectedString:    "Once",
			ExpectedIndex: len("Once\x1F"),
		},
		{
			Packed:            "Once\x1FTwice\x1FThrice\x1FFource",
			ExpectedString:    "Once",
			ExpectedIndex: len("Once\x1F"),
		},



		{
			Packed:            "\x1F",
			ExpectedString:   "",
			ExpectedIndex: len("\x1F"),
		},
		{
			Packed:            "\x1F\x1F",
			ExpectedString:   "",
			ExpectedIndex: len("\x1F"),
		},
		{
			Packed:            "\x1F\x1F\x1F",
			ExpectedString:   "",
			ExpectedIndex: len("\x1F"),
		},
		{
			Packed:            "\x1F\x1F\x1F\x1F",
			ExpectedString:   "",
			ExpectedIndex: len("\x1F"),
		},



		{
			Packed:            "\x1B\x1F",
			ExpectedString:        "\x1F",
			ExpectedIndex: len("\x1B\x1F"),
		},
		{
			Packed:            "\x1B\x1F\x1F",
			ExpectedString:        "\x1F",
			ExpectedIndex: len("\x1B\x1F\x1F"),
		},
		{
			Packed:            "\x1B\x1F\x1F\x1F",
			ExpectedString:        "\x1F",
			ExpectedIndex: len("\x1B\x1F\x1F"),
		},
		{
			Packed:            "\x1B\x1F\x1F\x1F\x1F",
			ExpectedString:        "\x1F",
			ExpectedIndex: len("\x1B\x1F\x1F"),
		},



		{
			Packed:            "\x1B",
			ExpectedString:   "",
			ExpectedIndex: len("\x1B"),
		},
		{
			Packed:            "\x1B\x1B",
			ExpectedString:        "\x1B",
			ExpectedIndex: len("\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B",
			ExpectedString:        "\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B\x1B\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B\x1B\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B\x1B\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B\x1B\x1B\x1B\x1B"),
		},
		{
			Packed:            "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B\x1B\x1B\x1B",
			ExpectedIndex: len("\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"),
		},

		{
			Packed:            "\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
			ExpectedString:        "\x1B",
			ExpectedIndex: len("\x1B\x1B"+"\x1F"),
		},



		{
			Packed:
				"\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F\x1B\x1F\x1B\x1F",
			ExpectedString:
				    "\x1F",
			ExpectedIndex: len(
				"\x1B\x1F"+
				"\x1F",
			),
		},



		{
			Packed:            "Hello\x1F\x1Fworld!\x1F\x1F",
			ExpectedString:    "Hello",
			ExpectedIndex: len("Hello\x1F"),
		},



		{
			Packed:           "\x1F\x1B\x1F\x1F",
			ExpectedString:  "",
			ExpectedIndex: len("\x1F"),
		},
	}

	for testNumber, test := range tests {

		actualString, actualIndex := strcodec.Decode(test.Packed)

		{
			var (
				actual   = actualString
				expected = test.ExpectedString
			)

			if expected != actual {
				t.Errorf("For test #%d, the actual unpacked-head-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("PACKED:   %q", test.Packed)
				continue
			}
		}

		{
			var (
				actual   = actualIndex
				expected = test.ExpectedIndex
			)

			if expected != actual {
				t.Errorf("For test #%d, the actual unpacked-head-index is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("PACKED:   %q", test.Packed)
				continue
			}
		}
	}
}
