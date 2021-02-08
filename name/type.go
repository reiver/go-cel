package cel_name

// Type represents a CEL message name.
//
// Note that cel_name.Type is an option type.
type Type struct {
	loaded bool
	value string
}

func Nothing() Type {
	return Type{}
}

func Something(value string) Type {
	return Type{
		loaded:true,
		value:value,
	}
}
