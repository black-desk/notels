package interfaces

import (
	"fmt"

	"github.com/black-desk/notels/internal/utils/logger"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/common"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/naming"
)

var log = logger.Get("lspgen")

var interfaceTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "context"
)

//jsonrpc2gen:%s
type %s interface {
        {{range .}}
                {{getComment .Documentation}} {{getRegistrationMethodName .}}
                {{getMethodName .}}{{getMethodArgs .}}{{getMethodReturn .}} //jsonrpc2gen:"{{getJsonRPCMethodName .}}"
        {{end}}
}
`

var interfaceTemplateFuncs = map[string]any{
	"getComment":                interfaceTemplateGetComment,
	"getMethodName":             interfaceTemplateGetMethodName,
	"getMethodArgs":             interfaceTemplateGetMethodArgs,
	"getMethodReturn":           interfaceTemplateGetMethodReturn,
	"getJsonRPCMethodName":      interfaceTemplateGetJsonRPCMethodName,
	"getRegistrationMethodName": interfaceTemplateGetRegistrationMethodName,
}

var interfaceTemplateGetComment = common.Comment
var interfaceTemplateGetMethodName = naming.MethodNameFromMessage
var interfaceTemplateGetMethodArgs = func(input any) string {
	result := "(\nctx context.Context"
	switch v := input.(type) {
	case model.Request:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			name := fmt.Sprintf("%s_Params",
				naming.MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nparams %s",
				naming.TypeName(&v.Params[0], name))
		}
		result += ",\n)"
		return result
	case model.Notification:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			name := fmt.Sprintf("%s_Params",
				naming.MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nparams %s",
				naming.TypeName(&v.Params[0], name))
		}
		result += ",\n)"
		return result
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
var interfaceTemplateGetMethodReturn = func(input any) string {
	result := "(\nerr error"
	switch v := input.(type) {
	case model.Request:
		if v.ErrorData != nil {
			name := fmt.Sprintf("%s_ErrorData",
				naming.MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nerrorData %s",
				naming.TypeName(v.ErrorData, name))
		}
		if v.Result != nil &&
			!(v.Result.Kind == "base" && v.Result.Name == "null") {
			name := fmt.Sprintf("%s_Result",
				naming.MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nresult %s",
				naming.TypeName(v.Result, name))
		}
		if v.PartialResult != nil {
			name := fmt.Sprintf("%s_PartialResult",
				naming.MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\npartialResult %s",
				naming.TypeName(v.PartialResult, name))
		}
		result += ",\n)"
		return result
	case model.Notification:
		result += ",\n)"
		return result
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
var interfaceTemplateGetJsonRPCMethodName = naming.JsonRPCMethodNameFromMessage
var interfaceTemplateGetRegistrationMethodName = func(input any) string {
	methodName := getRegistrationMethodNameFromMessage(input)
	if len(methodName) != 0 {
		return "\n// registration method: " + methodName
	} else {
		return ""
	}
}

func getRegistrationMethodNameFromMessage(input any) string {
	switch v := input.(type) {
	case model.Request:
		return v.RegistrationMethod
	case model.Notification:
		return v.RegistrationMethod
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
