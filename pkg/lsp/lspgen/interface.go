package main

import (
	"fmt"

	"github.com/black-desk/notels/internal/utils/logger"
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
                {{getMethodName .}}{{getMethodArgs .}}{{getMethodReturn .}} //jsonrpc2gen:"{{.Method}}"
        {{end}}
}
`

var interfaceTemplateFuncs = map[string]any{
	"getComment":                interfaceTemplateGetComment,
	"getMethodName":             interfaceTemplateGetMethodName,
	"getMethodArgs":             interfaceTemplateGetMethodArgs,
	"getMethodReturn":           interfaceTemplateGetMethodReturn,
	"getRegistrationMethodName": interfaceTemplateGetRegistrationMethodName,
}

var interfaceTemplateGetComment = Comment
var interfaceTemplateGetMethodName = MethodNameFromMessage
var interfaceTemplateGetMethodArgs = func(input any) string {
	result := "(\nctx context.Context"
	switch v := input.(type) {
	case Request:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			name := fmt.Sprintf("%s_Params",
				MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nparams %s",
				TypeName(&v.Params[0], name))
		}
		result += ",\n)"
		return result
	case Notification:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			name := fmt.Sprintf("%s_Params",
				MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nparams %s",
				TypeName(&v.Params[0], name))
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
	case Request:
		if v.ErrorData != nil {
			name := fmt.Sprintf("%s_ErrorData",
				MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nerrorData %s",
				TypeName(v.ErrorData, name))
		}
		if v.Result != nil &&
			!(v.Result.Kind == "base" && v.Result.Name == "null") {
			name := fmt.Sprintf("%s_Result",
				MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\nresult %s",
				TypeName(v.Result, name))
		}
		if v.PartialResult != nil {
			name := fmt.Sprintf("%s_PartialResult",
				MethodNameFromString(v.Method))
			result += fmt.Sprintf(",\npartialResult %s",
				TypeName(v.PartialResult, name))
		}
		result += ",\n)"
		return result
	case Notification:
		result += ",\n)"
		return result
	default:
		log.Fatalw("unexpected type", "input", input)
		panic("")
	}
}
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
	case Request:
		if v.RegistrationOptions != nil {
			// genType(v.RegistrationOptions)
		}
		return v.RegistrationMethod
	case Notification:
		return v.RegistrationMethod
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
