package cel_test

import (
	"github.com/reiver/go-cel"
	"github.com/reiver/go-cel/message"

	"testing"
)

func TestMessenger_Message(t *testing.T) {

	var x cel.Messenger = cel_message.Type{} // THIS IS THE PART THAT ACTUALLY MATTERS.

	if nil == x {
		t.Error("This should never happen.")
	}
}
