package cel_version

import (
	"encoding/json"
	"fmt"
)

func (receiver Type) MarshalJSON() ([]byte, error) {
	if !receiver.loaded {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver *Type) UnmarshalJSON(data []byte) error {
	var dst interface{}

	if err := json.Unmarshal(data, &dst); nil != err {
		return err
	}

	switch {
	case nil == dst:
		*receiver = Nothing()
		return nil
	default:
		switch casted := dst.(type) {
		case string:
			*receiver = Something(casted)
			return nil
		default:
			return fmt.Errorf("unsupported type: cannot unmarshal from type = %T ; value = %#v ; json = %q", dst, dst, string(data))
		}
	}
}
