package cel_message

import (
	"github.com/reiver/go-cel/kind"
	"github.com/reiver/go-cel/name"
	"github.com/reiver/go-cel/payload"
	"github.com/reiver/go-cel/version"

	"encoding/json"
	"fmt"
)

func (receiver Type) MarshalJSON() ([]byte, error) {

	var version interface{} = nil
	var kind    interface{} = nil
	var name    interface{} = nil

	{
		value, loaded := receiver.version.String()
		if loaded {
			version = value
		}
	}
	{
		value, loaded := receiver.kind.String()
		if loaded {
			kind = value
		}
	}
	{
		value, loaded := receiver.name.String()
		if loaded {
			name = value
		}
	}

	var src map[string]interface{} = map[string]interface{}{
		"magic"   : "CEL/1",
		"version" : version,
		"kind"    : kind,
		"name"    : name,
		"payload" : receiver.payload.MapStringInterface(),
	}

	return json.Marshal(src)
}

func (receiver *Type) UnmarshalJSON(data []byte) error {

	var dst map[string]interface{} = map[string]interface{}{}

	err := json.Unmarshal(data, &dst)
	if nil != err {
		return fmt.Errorf("could not unmarshal json data: %s", err)
	}

	{
		const key = "magic"

		valueInterface, found := dst[key]
		if !found {
			return fmt.Errorf("could not unmarshal json data: missing ‘%s’", key)
		}

		value, casted := valueInterface.(string)
		if !casted {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected type ‘string’ but is actually ‘%T’", key, valueInterface)
		}

		const expected = "CEL/1"

		if actual := value; expected != actual {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected value to be %q but actually was %q", key, expected, actual)
		}
	}

	{
		const key = "version"

		valueInterface, found := dst[key]
		if !found {
			return fmt.Errorf("could not unmarshal json data: missing ‘%s’", key)
		}

		value, casted := valueInterface.(string)
		if !casted {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected type ‘string’ but is actually ‘%T’", key, valueInterface)
		}

		receiver.version = cel_version.Something(value)
	}
	{
		const key = "kind"

		valueInterface, found := dst[key]
		if !found {
			return fmt.Errorf("could not unmarshal json data: missing ‘%s’", key)
		}

		value, casted := valueInterface.(string)
		if !casted {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected type ‘string’ but is actually ‘%T’", key, valueInterface)
		}

		receiver.kind = cel_kind.Something(value)
	}
	{
		const key = "name"

		valueInterface, found := dst[key]
		if !found {
			return fmt.Errorf("could not unmarshal json data: missing ‘%s’", key)
		}

		value, casted := valueInterface.(string)
		if !casted {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected type ‘string’ but is actually ‘%T’", key, valueInterface)
		}

		receiver.name = cel_name.Something(value)
	}
	{
		const key = "payload"

		valueInterface, found := dst[key]
		if !found {
			return fmt.Errorf("could not unmarshal json data: missing ‘%s’", key)
		}

		value, casted := valueInterface.(map[string]interface{})
		if !casted {
			return fmt.Errorf("could not unmarshal json data: bad ‘%s’ — expected type ‘map[string]interface{}’ but is actually ‘%T’", key, valueInterface)
		}

		receiver.payload = cel_payload.Something(value)
	}

	return nil
}
