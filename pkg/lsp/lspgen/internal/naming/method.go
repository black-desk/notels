package naming

import (
	"strings"

	"github.com/black-desk/notels/internal/utils/logger"
)

var log = logger.Get("lspgen")

func MethodNameFromString(lspJsonRPCMethodName string) string {
	lspJsonRPCMethodName = strings.Replace(
		lspJsonRPCMethodName,
		"$",
		"Lsp",
		-1,
	)
	tmp := strings.Split(lspJsonRPCMethodName, "/")
	for i, s := range tmp {
		tmp[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	lspJsonRPCMethodName = strings.Join(tmp, "")
	return lspJsonRPCMethodName
}
