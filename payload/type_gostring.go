package cel_payload

import (
	"fmt"
)

func (receiver *Type) GoString() string {
	if nil == receiver {
		return "cel_payload.Nothing()"
	}

	var result string

	receiver.lease(func(){
		if !receiver.loaded {
			result = "cel_payload.Nothing()"
		} else {
			result = fmt.Sprintf("cel_payload.Something(%#v)", receiver.value)
		}
	})

	return result
}
