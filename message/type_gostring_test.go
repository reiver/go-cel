package cel_message_test

import (
	"github.com/reiver/go-cel/message"

	"fmt"

	"testing"
)

func TestType_GoString(t *testing.T) {

	tests := []struct{
		Datum cel_message.Type
		Expected string
	}{
		{
			Datum:     cel_message.Type{},
			Expected: `cel_message.Type{version:cel_version.Nothing(), kind:cel_kind.Nothing(), name:cel_name.Nothing(), payload:cel_payload.Nothing()}`,
		},



		{
			Datum:     cel_message.Command("123", "SEND", nil),
			Expected: `cel_message.Command("123", "SEND", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Command("123", "SEND", map[string]interface{}(nil)),
			Expected: `cel_message.Command("123", "SEND", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Command("123", "SEND", map[string]interface{}{}),
			Expected: `cel_message.Command("123", "SEND", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Command("123", "SEND", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `cel_message.Command("123", "SEND", map[string]interface {}{"ONE":"apple", "THREE":"cherry", "TWO":"banana"})`,
		},



		{
			Datum:     cel_message.Event("123", "ACTIVATED", nil),
			Expected: `cel_message.Event("123", "ACTIVATED", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Event("123", "ACTIVATED", map[string]interface{}(nil)),
			Expected: `cel_message.Event("123", "ACTIVATED", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Event("123", "ACTIVATED", map[string]interface{}{}),
			Expected: `cel_message.Event("123", "ACTIVATED", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Event("123", "ACTIVATED", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `cel_message.Event("123", "ACTIVATED", map[string]interface {}{"ONE":"apple", "THREE":"cherry", "TWO":"banana"})`,
		},



		{
			Datum:     cel_message.Log("123", "TOTALS", nil),
			Expected: `cel_message.Log("123", "TOTALS", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Log("123", "TOTALS", map[string]interface{}(nil)),
			Expected: `cel_message.Log("123", "TOTALS", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Log("123", "TOTALS", map[string]interface{}{}),
			Expected: `cel_message.Log("123", "TOTALS", map[string]interface {}{})`,
		},
		{
			Datum:     cel_message.Log("123", "TOTALS", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `cel_message.Log("123", "TOTALS", map[string]interface {}{"ONE":"apple", "THREE":"cherry", "TWO":"banana"})`,
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf("%#v", test.Datum)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
	/////////////// CONTINUE
			continue
		}
	}
}
