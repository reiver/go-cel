package cel_message

import (
	"github.com/reiver/go-cel/kind"
	"github.com/reiver/go-cel/name"
	"github.com/reiver/go-cel/payload"
	"github.com/reiver/go-cel/version"
)

func Event(version string, name string, payload map[string]interface{}) Type {
	return Type{
		version: cel_version.Something(version),
		kind:    cel_kind.Event(),
		name:    cel_name.Something(name),
		payload: cel_payload.Something(payload),
	}
}
