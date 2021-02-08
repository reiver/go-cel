package cel_payload

func copymap(src map[string]interface{}) map[string]interface{} {
	var result map[string]interface{} = map[string]interface{}{}

	for k, v := range src {
		switch casted := v.(type) {
		case map[string]interface{}:
			result[k] = copymap(casted)
		default:
			result[k] = casted
		}
	}

	return result
}
