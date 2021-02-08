package cel_version_test

import (
	"github.com/reiver/go-cel/version"

	"errors"

	"testing"
)

func TestType_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		Datum []byte
		ExpectedValue cel_version.Type
		ExpectedErr error
	}{
		{
			Datum: []byte(nil),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: errors.New("unexpected end of JSON input"),
		},



		{
			Datum: []byte("null"),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: nil,
		},



		{
			Datum:                       []byte(`"apple"`),
			ExpectedValue: cel_version.Something("apple"),
			ExpectedErr: nil,
		},
		{
			Datum:                       []byte(`"banana"`),
			ExpectedValue: cel_version.Something("banana"),
			ExpectedErr: nil,
		},
		{
			Datum:                       []byte(`"cherry"`),
			ExpectedValue: cel_version.Something("cherry"),
			ExpectedErr: nil,
		},



		{
			Datum: []byte(""),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: errors.New("unexpected end of JSON input"),
		},
		{
			Datum: []byte("5"),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: errors.New(`unsupported type: cannot unmarshal from type = float64 ; value = 5 ; json = "5"`),
		},
		{
			Datum: []byte(`["ONE","TWO","THREE"]`),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: errors.New(`unsupported type: cannot unmarshal from type = []interface {} ; value = []interface {}{"ONE", "TWO", "THREE"} ; json = "[\"ONE\",\"TWO\",\"THREE\"]"`),
		},
		{
			Datum: []byte(`{"apple":"ONE","banana":"TWO","cherry":"THREE"}`),
			ExpectedValue: cel_version.Nothing(),
			ExpectedErr: errors.New(`unsupported type: cannot unmarshal from type = map[string]interface {} ; value = map[string]interface {}{"apple":"ONE", "banana":"TWO", "cherry":"THREE"} ; json = "{\"apple\":\"ONE\",\"banana\":\"TWO\",\"cherry\":\"THREE\"}"`),
		},
	}

	for testNumber, test := range tests {

		var actualValue cel_version.Type

		actualErr := actualValue.UnmarshalJSON(test.Datum)

		switch {
		case nil == test.ExpectedErr && nil == actualErr:
			// Nothing here.
		case nil == test.ExpectedErr && nil != actualErr:
			fallthrough
		case nil != test.ExpectedErr && nil == actualErr:
			fallthrough
		case test.ExpectedErr.Error() != actualErr.Error():
			t.Errorf("For test #%d, the actual error was not what was expected." , testNumber)
			t.Logf("DATUM []byte: %#v", test.Datum)
			t.Logf("DATUM string: %#v", string(test.Datum))
			t.Logf("EXPECTED VALUE: %#v", test.ExpectedValue)
			t.Logf("ACTUAL   VALUE: %#v", actualValue)
			t.Logf("EXPECTED ERR: (%T) %q", test.ExpectedErr, test.ExpectedErr)
			t.Logf("ACTUAL   ERR: (%T) %q", actualErr, actualErr)
	/////////////// CONTINUE
			continue
		}

		if expected, actual := test.ExpectedValue, actualValue; expected != actual {
			t.Errorf("For test #%d, the actual value was not what was expected." , testNumber)
			t.Logf("DATUM []byte: %#v", test.Datum)
			t.Logf("DATUM string: %#v", string(test.Datum))
			t.Logf("EXPECTED VALUE: %#v", test.ExpectedValue)
			t.Logf("ACTUAL   VALUE: %#v", actualValue)
			t.Logf("EXPECTED ERR: (%T) %q", test.ExpectedErr, test.ExpectedErr)
			t.Logf("ACTUAL   ERR: (%T) %q", actualErr, actualErr)
	/////////////// CONTINUE
			continue
		}
	}
}
