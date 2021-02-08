package cel_payload

import (
	"github.com/reiver/go-mapstringinterface"

	"sync"
)

// WARNING: cel_payload.Type contains a sync.Mutex. Do NOT copy it after first use.
type Type struct {
	mutex sync.Mutex
	value map[string]interface{}
}

func Nothing() Type {
	return Type{}
}

func Something(value map[string]interface{}) Type {
	return Type{value:copymap(value)}
}
