package cel_payload

func (receiver *Type) MapStringInterface() map[string]interface{} {
	if nil == receiver {
		return nil
	}

	var value map[string]interface{}

	receiver.lease(func(){
		value = copymap(receiver.value)
	})

	return value
}
