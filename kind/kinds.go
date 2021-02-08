package cel_kind

const (
	kindCommand = "COMMAND"
	kindEvent   = "EVENT"
	kindLog     = "LOG"
)

func Command() Type {
	return Something(kindCommand)
}

func Event() Type {
	return Something(kindEvent)
}

func Log() Type {
	return Something(kindLog)
}
