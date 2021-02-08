package cel_version

import (
	"fmt"
)

func (receiver Type) GoString() string {
	if !receiver.loaded {
		return "cel_version.Nothing()"
	}

	return fmt.Sprintf("cel_version.Something(%q)", receiver.value)
}
