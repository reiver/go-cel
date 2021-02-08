package cel_message_test

import (
	"github.com/reiver/go-cel/message"

	"encoding/json"
	"reflect"

	"testing"
)

func TestType_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected cel_message.Type
	}{
		{
			JSON: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{}
			}`,
			Expected: cel_message.Command("4", "SEND", nil),
		},
		{
			JSON: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{}
			}`,
			Expected: cel_message.Command("4", "SEND", map[string]interface{}{}),
		},
		{
			JSON: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
			Expected: cel_message.Command("4", "SEND", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
		},



		{
			JSON: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{}
			}`,
			Expected: cel_message.Log("5", "TOTALS", nil),
		},
		{
			JSON: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{}
			}`,
			Expected: cel_message.Log("5", "TOTALS", map[string]interface{}{}),
		},
		{
			JSON: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
			Expected: cel_message.Log("5", "TOTALS", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
		},



		{
			JSON: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{}
			}`,
			Expected: cel_message.Event("123", "ACTIVATED", nil),
		},
		{
			JSON: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{}
			}`,
			Expected: cel_message.Event("123", "ACTIVATED", map[string]interface{}{}),
		},
		{
			JSON: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
			Expected: cel_message.Event("123", "ACTIVATED", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
		},
	}

	for testNumber, test := range tests {

		var actual cel_message.Type

		if err := json.Unmarshal([]byte(test.JSON), &actual); nil != err {
			t.Errorf("For test #%d, did not expect an error but actually received one.", testNumber)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %s", err)
			continue
		}

		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, actual value is not what was expected", testNumber)
			t.Logf("JSON: %s", test.JSON)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
