package cel_payload

type Payloader interface {
	MapStringInterface() map[string]interface{}
	PathQuery(...string) (interface{}, bool)
	PathQueryForString(...string) (string, bool)
}
