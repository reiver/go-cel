package cel_payload

func (receiver *Type) lease(fn func()) {
	if nil == receiver {
		return
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	fn()
}
