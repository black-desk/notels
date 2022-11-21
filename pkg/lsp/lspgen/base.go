package main

func baseTypeName(name string) string {
	switch name {
	case "null":
		log.Fatalw("should never require typename of null")
		panic("")
	case "URI", "DocumentUri", "RegExp", "string":
		return name
	case "integer":
		return "int64"
	case "uinteger":
		return "uint64"
	case "decimal":
		return "float64"
	case "boolean":
		return "bool"
	default:
		log.Fatalw("unexpected base type",
			"name", name)
		panic("")
	}
}
