package cel_version

// Type represents a CEL message version.
//
// Note that cel_version.Type is an option type.
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
