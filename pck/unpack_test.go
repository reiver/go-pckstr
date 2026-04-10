package pck_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-pckstr/pck"
)

func TestUnpack(t *testing.T) {
	tests := []struct{
		Packed   string
		Expected []string
	}{
		{
			Packed: "",
			Expected: []string{},
		},
		{
			Packed: "\x1F",
			Expected: []string{""},
		},


		{
			Packed: "apple",
			Expected: []string{
				"apple",
			},
		},
		{
			Packed: "apple\x1F",
			Expected: []string{
				"apple",
			},
		},
		{
			Packed: "BANANA",
			Expected: []string{
				"BANANA",
			},
		},
		{
			Packed: "BANANA\x1F",
			Expected: []string{
				"BANANA",
			},
		},
		{
			Packed: "Cherry",
			Expected: []string{
				"Cherry",
			},
		},
		{
			Packed: "Cherry\x1F",
			Expected: []string{
				"Cherry",
			},
		},
		{
			Packed: "dAtE",
			Expected: []string{
				"dAtE",
			},
		},
		{
			Packed: "dAtE\x1F",
			Expected: []string{
				"dAtE",
			},
		},



		{
			Packed: "Once",
			Expected: []string{
				"Once",
			},
		},
		{
			Packed: "Once\x1F",
			Expected: []string{
				"Once",
			},
		},
		{
			Packed: "Once\x1FTwice",
			Expected: []string{
				"Once",
				"Twice",
			},
		},
		{
			Packed: "Once\x1FTwice\x1F",
			Expected: []string{
				"Once",
				"Twice",
			},
		},
		{
			Packed: "Once\x1FTwice\x1FThrice",
			Expected: []string{
				"Once",
				"Twice",
				"Thrice",
			},
		},
		{
			Packed: "Once\x1FTwice\x1FThrice\x1F",
			Expected: []string{
				"Once",
				"Twice",
				"Thrice",
			},
		},
		{
			Packed: "Once\x1FTwice\x1FThrice\x1FFource",
			Expected: []string{
				"Once",
				"Twice",
				"Thrice",
				"Fource",
			},
		},
		{
			Packed: "Once\x1FTwice\x1FThrice\x1FFource\x1F",
			Expected: []string{
				"Once",
				"Twice",
				"Thrice",
				"Fource",
			},
		},



		{
			Packed: "\x1FHello",
			Expected: []string{
				"",
				"Hello",
			},
		},
		{
			Packed: "\x1FHello\x1F",
			Expected: []string{
				"",
				"Hello",
			},
		},
		{
			Packed: "\x1FHello\x1FHow",
			Expected: []string{
				"",
				"Hello",
				"How",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1F",
			Expected: []string{
				"",
				"Hello",
				"How",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1Fdo",
			Expected: []string{
				"",
				"Hello",
				"How",
				"do",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1Fdo\x1F",
			Expected: []string{
				"",
				"Hello",
				"How",
				"do",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1Fdo\x1Fyou",
			Expected: []string{
				"",
				"Hello",
				"How",
				"do",
				"you",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1Fdo\x1Fyou\x1Fdo",
			Expected: []string{
				"",
				"Hello",
				"How",
				"do",
				"you",
				"do",
			},
		},
		{
			Packed: "\x1FHello\x1FHow\x1Fdo\x1Fyou\x1Fdo\x1F",
			Expected: []string{
				"",
				"Hello",
				"How",
				"do",
				"you",
				"do",
			},
		},



		{
			Packed: "\x1B\x1B",
			Expected: []string{
				"\x1B",
			},
		},
		{
			Packed: "\x1B\x1B\x1B\x1B",
			Expected: []string{
				"\x1B\x1B",
			},
		},
		{
			Packed: "\x1B\x1B\x1B\x1B\x1B\x1B",
			Expected: []string{
				"\x1B\x1B\x1B",
			},
		},
		{
			Packed: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
			Expected: []string{
				"\x1B\x1B\x1B\x1B",
			},
		},

		{
			Packed: "\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B",
			Expected: []string{
				"\x1B",
				"\x1B\x1B",
				"\x1B\x1B\x1B",
				"\x1B\x1B\x1B\x1B",
			},
		},



		{
			Packed: "\x1B\x1F",
			Expected: []string{
				"\x1F",
			},
		},
		{
			Packed: "\x1B\x1F\x1B\x1F",
			Expected: []string{
				"\x1F\x1F",
			},
		},
		{
			Packed: "\x1B\x1F\x1B\x1F\x1B\x1F",
			Expected: []string{
				"\x1F\x1F\x1F",
			},
		},
		{
			Packed: "\x1B\x1F\x1B\x1F\x1B\x1F\x1B\x1F",
			Expected: []string{
				"\x1F\x1F\x1F\x1F",
			},
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
			Expected: []string{
				"\x1F",
				"\x1F\x1F",
				"\x1F\x1F\x1F",
				"\x1F\x1F\x1F\x1F",
			},
		},



		{
			Packed: "\x1F",
			Expected: []string{
				"",
			},
		},
		{
			Packed: "\x1F"+"\x1F",
			Expected: []string{
				"",
				"",
			},
		},
		{
			Packed: "\x1F"+"\x1F"+"\x1F",
			Expected: []string{
				"",
				"",
				"",
			},
		},
		{
			Packed: "\x1F"+"\x1F"+"\x1F"+"\x1F",
			Expected: []string{
				"",
				"",
				"",
				"",
			},
		},



		{
			Packed: "Hello"+"\x1F"+"\x1F"+"world!"+"\x1F",
			Expected: []string{
				"Hello",
				"",
				"world!",
			},
		},
		{
			Packed: "Hello"+"\x1F"+"\x1F"+"world!"+"\x1F"+"\x1F",
			Expected: []string{
				"Hello",
				"",
				"world!",
				"",
			},
		},



		{
			Packed: "\x1F"+"\x1B\x1F",
			Expected: []string{
				"",
				"\x1F",
			},
		},
		{
			Packed: "\x1F"+"\x1B\x1F"+"\x1F",
			Expected: []string{
				"",
				"\x1F",
			},
		},
	}

	for testNumber, test := range tests {

		actual := pck.Unpack(test.Packed)

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual unpacked-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("PACKED:   %q", test.Packed)
				continue
			}
		}

		{
			packed := pck.Pack(actual...)

			expected := pck.Canonical(test.Packed)

			if expected != packed {
				t.Errorf("For test #%d, the actual unpacked-then-packed-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("PACKED:   %q", packed)
				t.Logf("ORIGINAL: %q", test.Packed)
				continue
			}
		}
	}
}
