package cel_name

func (receiver Type) String() (string, bool) {
	if !receiver.loaded {
		return "", false
	}

	return receiver.value, true
}
