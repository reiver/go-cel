package cel_payload

import (
	"sync"
)

// WARNING: cel_payload.Type contains a sync.Mutex. Do NOT copy it after first use.
type Type struct {
	mutex sync.Mutex
	loaded bool
	value map[string]interface{}
}

func Nothing() Type {
	return Type{}
}

func Something(value map[string]interface{}) Type {
	return Type{
		loaded:true,
		value:copymap(value),
	}
}
