package cel_version_test

import (
	"github.com/reiver/go-cel/version"

	"testing"
)

func TestType_String(t *testing.T) {

	tests := []struct{
		Datum cel_version.Type
		ExpectedValue string
		ExpectedReturned bool
	}{
		{
			Datum: cel_version.Nothing(),
			ExpectedValue:               "",
			ExpectedReturned: false,
		},



		{
			Datum: cel_version.Something(""),
			ExpectedValue:               "",
			ExpectedReturned: true,
		},



		{
			Datum: cel_version.Something("apple"),
			ExpectedValue:               "apple",
			ExpectedReturned: true,
		},
		{
			Datum: cel_version.Something("banana"),
			ExpectedValue:               "banana",
			ExpectedReturned: true,
		},
		{
			Datum: cel_version.Something("cherry"),
			ExpectedValue:               "cherry",
			ExpectedReturned: true,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualReturned := test.Datum.String()

		if expected, actual := test.ExpectedReturned, actualReturned; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("DATUM: %#v", test.Datum)
			t.Logf("EXPECTED VALUE: %q", test.ExpectedValue)
			t.Logf("ACTUAL   VALUE: %q", actualValue)
			t.Logf("EXPECTED RETURNED: %t", test.ExpectedReturned)
			t.Logf("ACTUAL   RETURNED: %t", actualReturned)
	/////////////// CONTINUE
			continue
		}

		if expected, actual := test.ExpectedValue, actualValue; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("DATUM: %#v", test.Datum)
			t.Logf("EXPECTED VALUE: %q", test.ExpectedValue)
			t.Logf("ACTUAL   VALUE: %q", actualValue)
			t.Logf("EXPECTED RETURNED: %t", test.ExpectedReturned)
			t.Logf("ACTUAL   RETURNED: %t", actualReturned)
	/////////////// CONTINUE
			continue
		}
	}
}
