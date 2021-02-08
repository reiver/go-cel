package cel_message

import (
	"github.com/reiver/go-cel/kind"

	"fmt"
)

func (receiver Type) GoString() string {
	var version string
	var name    string

	var versionLoaded bool
	var    kindLoaded bool
	var    nameLoaded bool

	version, versionLoaded = receiver.version.String()
	_,          kindLoaded = receiver.kind.String()
	name,       nameLoaded = receiver.name.String()

	switch {
	case !versionLoaded || !kindLoaded || !nameLoaded:
		fallthrough
	case cel_kind.Command() != receiver.kind && cel_kind.Event() != receiver.kind && cel_kind.Log() != receiver.kind:
		fallthrough
	default:
		return fmt.Sprintf(
			"cel_message.Type{version:%#v, kind:%#v, name:%#v, payload:%#v}",
			receiver.version,
			receiver.kind,
			receiver.name,
			&receiver.payload,
		)
	case cel_kind.Command() == receiver.kind:
		return fmt.Sprintf("cel_message.Command(%q, %q, %#v)", version, name, receiver.payload.MapStringInterface())
	case cel_kind.Event() == receiver.kind:
		return fmt.Sprintf("cel_message.Event(%q, %q, %#v)",   version, name, receiver.payload.MapStringInterface())
	case cel_kind.Log() == receiver.kind:
		return fmt.Sprintf("cel_message.Log(%q, %q, %#v)",     version, name, receiver.payload.MapStringInterface())
	}
}
