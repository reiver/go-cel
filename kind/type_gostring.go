package cel_kind

import (
	"fmt"
)

func (receiver Type) GoString() string {
	if !receiver.loaded {
		return "cel_kind.Nothing()"
	}

	return fmt.Sprintf("cel_kind.Something(%q)", receiver.value)
}
