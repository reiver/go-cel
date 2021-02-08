package cel

import (
	"github.com/reiver/go-cel/kind"
	"github.com/reiver/go-cel/name"
	"github.com/reiver/go-cel/payload"
	"github.com/reiver/go-cel/version"
)

// Messenger represents a message.
type Messenger interface {
	Version()  cel_version.Type
	Kind()     cel_kind.Type
	Name()     cel_name.Type
	Payload()  cel_payload.Payloader
}
