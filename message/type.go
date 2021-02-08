package cel_message

import (
	"github.com/reiver/go-cel/kind"
	"github.com/reiver/go-cel/name"
	"github.com/reiver/go-cel/payload"
	"github.com/reiver/go-cel/version"
)

type Type struct {
	version cel_version.Type
	kind    cel_kind.Type
	name    cel_name.Type
	payload cel_payload.Type
}

func (receiver Type) Version() cel_version.Type {
	return receiver.version
}

func (receiver Type) Kind() cel_kind.Type {
	return receiver.kind
}

func (receiver Type) Name() cel_name.Type {
	return receiver.name
}

func (receiver Type) Payload() cel_payload.Payloader {
	return &receiver.payload
}
