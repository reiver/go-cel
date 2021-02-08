package cel_name

import (
	"fmt"
)

func (receiver Type) GoString() string {
	if !receiver.loaded {
		return "cel_name.Nothing()"
	}

	return fmt.Sprintf("cel_name.Something(%q)", receiver.value)
}
