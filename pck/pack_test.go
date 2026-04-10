package pck_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-pckstr/pck"
)

func TestPack(t *testing.T) {
	tests := []struct{
		Strings  []string
		Expected string
	}{
		{},



		{
			Strings: []string(nil),
		},
		{
			Strings: []string{},
		},



		{
			Strings: []string{""},
			Expected: "\x1F",
		},
		{
			Strings: []string{
				"",
				"",
			},
			Expected: "\x1F"+"\x1F",
		},
		{
			Strings: []string{
				"",
				"",
				"",
			},
			Expected: "\x1F"+"\x1F"+"\x1F",
		},
		{
			Strings: []string{
				"",
				"",
				"",
				"",
			},
			Expected: "\x1F"+"\x1F"+"\x1F"+"\x1F",
		},


		{
			Strings: []string{
				"apple",
			},
			Expected: "apple"+"\x1F",
		},
		{
			Strings: []string{
				"BANANA",
			},
			Expected: "BANANA"+"\x1F",
		},
		{
			Strings: []string{
				"Cherry",
			},
			Expected: "Cherry"+"\x1F",
		},
		{
			Strings: []string{
				"dAtE",
			},
			Expected: "dAtE"+"\x1F",
		},



		{
			Strings: []string{
				"Once",
			},
			Expected: "Once"+"\x1F",
		},
		{
			Strings: []string{
				"Once",
				"Twice",
			},
			Expected: "Once"+"\x1F"+"Twice"+"\x1F",
		},
		{
			Strings: []string{
				"Once",
				"Twice",
				"Thrice",
			},
			Expected: "Once"+"\x1F"+"Twice"+"\x1F"+"Thrice"+"\x1F",
		},
		{
			Strings: []string{
				"Once",
				"Twice",
				"Thrice",
				"Fource",
			},
			Expected: "Once"+"\x1F"+"Twice"+"\x1F"+"Thrice"+"\x1F"+"Fource"+"\x1F",
		},



		{
			Strings: []string{
				"\x1B",
			},
			Expected: "\x1B\x1B"+"\x1F",
		},
		{
			Strings: []string{
				"\x1B\x1B",
			},
			Expected: "\x1B\x1B\x1B\x1B"+"\x1F",
		},
		{
			Strings: []string{
				"\x1B\x1B\x1B",
			},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F",
		},
		{
			Strings: []string{
				"\x1B\x1B\x1B\x1B",
			},
			Expected: "\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F",
		},

		{
			Strings: []string{
				"\x1B",
				"\x1B\x1B",
				"\x1B\x1B\x1B",
				"\x1B\x1B\x1B\x1B",
			},
			Expected: "\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F"+"\x1B\x1B\x1B\x1B\x1B\x1B\x1B\x1B"+"\x1F",
		},



		{
			Strings: []string{
				"\x1F",
			},
			Expected: "\x1B\x1F"+"\x1F",
		},
		{
			Strings: []string{
				"\x1F\x1F",
			},
			Expected: "\x1B\x1F\x1B\x1F"+"\x1F",
		},
		{
			Strings: []string{
				"\x1F\x1F\x1F",
			},
			Expected: "\x1B\x1F\x1B\x1F\x1B\x1F"+"\x1F",
		},
		{
			Strings: []string{
				"\x1F\x1F\x1F\x1F",
			},
			Expected: "\x1B\x1F\x1B\x1F\x1B\x1F\x1B\x1F"+"\x1F",
		},

		{
			Strings: []string{
				"\x1F",
				"\x1F\x1F",
				"\x1F\x1F\x1F",
				"\x1F\x1F\x1F\x1F",
			},
			Expected:
				"\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F\x1B\x1F"+
				"\x1F"+
				"\x1B\x1F\x1B\x1F\x1B\x1F\x1B\x1F"+
				"\x1F",
		},
	}

	for testNumber, test := range tests {

		actual := pck.Pack(test.Strings...)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual packed-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRINGS: %#v", test.Strings)
				continue
			}
		}

		{
			var expected []string = append([]string(nil), test.Strings...)
			if nil == expected {
				expected = []string{}
			}

			unpacked := pck.Unpack(actual)
			if !reflect.DeepEqual(expected, unpacked) {
				t.Errorf("For test #%d, the actual packed-then-unpacked value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("UNDONE:   %#v", unpacked)
				t.Logf("STRINGS: %#v", test.Strings)
				continue
			}
		}
	}
}
