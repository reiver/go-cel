package cel_payload_test

import (
	"github.com/reiver/go-cel/payload"

	"fmt"

	"testing"
)

func TestType_GoString(t *testing.T) {

	tests := []struct{
		Datum cel_payload.Type
		Expected string
	}{
		{
			Datum: cel_payload.Type{},
			Expected: "cel_payload.Nothing()",
		},



		{
			Datum: cel_payload.Nothing(),
			Expected: "cel_payload.Nothing()",
		},



		{
			Datum: cel_payload.Something(nil),
			Expected: "cel_payload.Something(map[string]interface {}{})",
		},
		{
			Datum: cel_payload.Something(map[string]interface{}(nil)),
			Expected: "cel_payload.Something(map[string]interface {}{})",
		},
		{
			Datum: cel_payload.Something(map[string]interface{}{}),
			Expected: "cel_payload.Something(map[string]interface {}{})",
		},



		{
			Datum: cel_payload.Something(map[string]interface{}{"ONE":"apple",}),
			Expected: `cel_payload.Something(map[string]interface {}{"ONE":"apple"})`,
		},
		{
			Datum: cel_payload.Something(map[string]interface{}{"ONE":"apple","TWO":"banana"}),
			Expected: `cel_payload.Something(map[string]interface {}{"ONE":"apple", "TWO":"banana"})`,
		},
		{
			Datum: cel_payload.Something(map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `cel_payload.Something(map[string]interface {}{"ONE":"apple", "THREE":"cherry", "TWO":"banana"})`,
		},

		{
			Datum: cel_payload.Something(map[string]interface{}{
				"ONE":"apple",
				"TWO":"banana",
				"THREE":"cherry",
				"FOUR": map[string]interface{}{
					"once":true,
					"twice":false,
					"thrice":4,
					"fource":5,
				},
			}),
			Expected: `cel_payload.Something(map[string]interface {}{"FOUR":map[string]interface {}{"fource":5, "once":true, "thrice":4, "twice":false}, "ONE":"apple", "THREE":"cherry", "TWO":"banana"})`,
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf("%#v", &test.Datum)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
	/////////////// CONTINUE
			continue
		}

	}
}
