package cel_message_test

import (
	"github.com/reiver/go-cel/message"

	"encoding/json"
	"reflect"

	"testing"
)

func TestType_MarshalJSON(t *testing.T) {

	tests := []struct{
		Message cel_message.Type
		Expected string
	}{
		{
			Message: cel_message.Command("4", "SEND", nil),
			Expected: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Command("4", "SEND", map[string]interface{}{}),
			Expected: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Command("4", "SEND", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `{
				"version":"4",
				"kind":"COMMAND",
				"name":"SEND",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
		},



		{
			Message: cel_message.Log("5", "TOTALS", nil),
			Expected: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Log("5", "TOTALS", map[string]interface{}{}),
			Expected: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Log("5", "TOTALS", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `{
				"version":"5",
				"kind":"LOG",
				"name":"TOTALS",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
		},



		{
			Message: cel_message.Event("123", "ACTIVATED", nil),
			Expected: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Event("123", "ACTIVATED", map[string]interface{}{}),
			Expected: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{}
			}`,
		},
		{
			Message: cel_message.Event("123", "ACTIVATED", map[string]interface{}{"ONE":"apple","TWO":"banana","THREE":"cherry"}),
			Expected: `{
				"version":"123",
				"kind":"EVENT",
				"name":"ACTIVATED",
				"payload":{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry"
				}
			}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.Message)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually received one.", testNumber)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %s", err)
			continue
		}

		var actual map[string]interface{}
		{
			if err := json.Unmarshal(actualBytes, &actual); nil != err {
				t.Errorf("For test #%d, did not expect an error but actually received one.", testNumber)
				t.Logf("JSON: %s", string(actualBytes))
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %s", err)
				continue
			}
		}

		var expected map[string]interface{}
		{
			if err := json.Unmarshal([]byte(test.Expected), &expected); nil != err {
				t.Errorf("For test #%d, did not expect an error but actually received one.", testNumber)
				t.Logf("JSON: %s", test.Expected)
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %s", err)
				continue
			}
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, actual value is not what was expected", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
