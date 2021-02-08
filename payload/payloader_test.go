package cel_payload_test

import (
	"github.com/reiver/go-cel/payload"

	"testing"
)

func TestPayloader_Payload(t *testing.T) {

	var x cel_payload.Payloader = &cel_payload.Type{} // THIS IS THE PART THAT ACTUALLY MATTERS.

	if nil == x {
		t.Error("This should never happen.")
	}
}
