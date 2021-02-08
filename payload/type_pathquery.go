package cel_payload

import (
	"github.com/reiver/go-mapstringinterface"
)

func (receiver *Type) PathQuery(query ...string) (interface{}, bool) {
	if nil == receiver {
		return nil, false
	}

	var value interface{}
	var found bool

	receiver.lease(func(){
		value, found = mapstringinterface.PathQuery(receiver.value, query...)
	})

	return value, found
}

func (receiver *Type) PathQueryForString(query ...string) (string, bool) {
	if nil == receiver {
		return "", false
	}

	var value string
	var found bool

	receiver.lease(func(){
		value, found = mapstringinterface.PathQueryForString(receiver.value, query...)
	})

	return value, found
}
