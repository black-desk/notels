package main

func baseTypeName(name string) string {
	switch name {
	case "null":
		return "Null"
	case "URI", "DocumentUri", "RegExp":
		return name
	case "string":
		return "String"
	case "integer":
		return "Integer"
	case "uinteger":
		return "Uinteger"
	case "decimal":
		return "Decimal"
	case "boolean":
		return "Boolean"
	default:
		log.Fatalw("unexpected base type",
			"name", name)
		panic("")
	}
}
