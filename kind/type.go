package cel_kind

// Type represents a CEL message kind.
//
// Note that cel_kind.Type is an option type.
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
