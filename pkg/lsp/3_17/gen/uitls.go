package main

import (
	"strings"

        "encoding/json"
)

func goMethodName(json string) string {
	json = strings.Replace(json, "$", "Lsp", -1)
	tmp := strings.Split(json, "/")
	for i, s := range tmp {
		tmp[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	json = strings.Join(tmp, "")
	return json
}

func goEnumType(s string) string {
	switch s {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "uinteger":
		return "uint"
	default:
		log.Fatalw("unexpected type",
			"s", s)
		panic("")
	}

}

func goEnumName(t string, name string) string {
	return t + strings.ToUpper(name[0:1]) + name[1:]
}

func goEnumValue(value json.RawMessage) string {
	return string(value)
}

