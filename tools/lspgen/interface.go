package main

import (
	"fmt"
)

var interfaceTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

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
var interfaceTemplateGetMethodName = func(input any) string {
	switch v := input.(type) {
	case *Request:
		return MethodNameFromString(v.Method)
	case *Notification:
		return MethodNameFromString(v.Method)
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}

var interfaceTemplateGetMethodArgs = func(input any) string {
	result := "(\nctx context.Context"
	switch v := (input).(type) {
	case *Request:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			result += fmt.Sprintf(",\nparams *%s",
				typeName(MethodNameFromString(v.Method),
					"Params", &v.Params[0]))
		}
		result += ",\n)"
		return result
	case *Notification:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("should not happen")
			}
			result += fmt.Sprintf(",\nparams *%s",
				typeName(MethodNameFromString(v.Method),
					"Params", &v.Params[0]))
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
	case *Request:
		if v.ErrorData != nil {
			result += fmt.Sprintf(",\nerrorData %s",
				typeName(MethodNameFromString(v.Method),
					"ErrorData", v.ErrorData))
		}
		if v.Result != nil &&
			!(v.Result.Kind == "base" && v.Result.Name == "null") {
			result += fmt.Sprintf(",\nresult %s",
				typeName(MethodNameFromString(v.Method),
					"Result", v.Result))
		}
		if v.PartialResult != nil {
			result += fmt.Sprintf(",\npartialResult %s",
				typeName(MethodNameFromString(v.Method),
					"PartialResult", v.PartialResult))
		}
		result += ",\n)"
		return result
	case *Notification:
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
	case *Request:
		if v.RegistrationOptions != nil {
			// genType(v.RegistrationOptions)
		}
		return v.RegistrationMethod
	case *Notification:
		return v.RegistrationMethod
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
