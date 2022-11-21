package main

func MethodNameFromMessage(input any) string {
	switch v := input.(type) {
	case Request:
		return MethodNameFromString(v.Method)
	case Notification:
		return MethodNameFromString(v.Method)
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
