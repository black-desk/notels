package components

import "strings"

func Comment(s string) string {
	if s != "" {
		return "// " + strings.TrimLeft(
			strings.ReplaceAll(s, "\n", " "),
			" ",
		)
	} else {
		return ""
	}
}
