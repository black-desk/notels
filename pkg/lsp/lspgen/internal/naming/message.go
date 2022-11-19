package naming

import (
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
)

func MethodNameFromMessage(input any) string {
	switch v := input.(type) {
	case model.Request:
		return MethodNameFromString(v.Method)
	case model.Notification:
		return MethodNameFromString(v.Method)
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}

func JsonRPCMethodNameFromMessage(input any) string {
	switch v := input.(type) {
	case model.Request:
		return v.Method
	case model.Notification:
		return v.Method
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
