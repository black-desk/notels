package main

import (
	"fmt"
)

var interfaceTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
        "context"
        "github.com/black-desk/notels/pkg/jsonrpc2"
)

//jsonrpc2gen: %[1]s adaptor
//jsonrpc2gen: %[1]s proxy
type %[1]s interface {
        {{range .}}
                {{getComment .Documentation}} {{getRegistrationMethodName .}}
                {{getMethodName .}}{{getMethodArgs .}}{{getMethodReturn .}} // jsonrpc2gen:{{getMessageType .}}"{{.Method}}"
        {{end}}
}

`

var interfaceTemplateFuncs = map[string]any{
	"getComment":                interfaceTemplateGetComment,
	"getMethodName":             interfaceTemplateGetMethodName,
	"getMethodArgs":             interfaceTemplateGetMethodArgs,
	"getMethodReturn":           interfaceTemplateGetMethodReturn,
	"getRegistrationMethodName": interfaceTemplateGetRegistrationMethodName,
	"getMessageType":            interfaceTemplateGetMessageType,
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
	result := "(\n"
	switch v := input.(type) {
	case *Request:
		if v.PartialResult != nil {
			typeName(MethodNameFromString(v.Method),
				"PartialResult", v.PartialResult)
		}
		if v.Result != nil {
			name := typeName(MethodNameFromString(v.Method),
				"Result", v.Result)
			if name[0:1] != "[" {
				name = "*" + name
			}
			result += fmt.Sprintf("result %s,\n", name)
		}
		if v.ErrorData != nil {
			name := typeName(MethodNameFromString(v.Method),
				"ErrorData", v.ErrorData)
			if name[0:1] != "[" {
				name = "*" + name
			}
			result += fmt.Sprintf("errorData %s,\n", name)
		}
		result += "code jsonrpc2.Code,\n"
		result += "err error,\n)"
		return result
	case *Notification:
		result += "err error,\n)"
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
			typeName(MethodNameFromString(v.Method),
				"RegistrationOptions", v.RegistrationOptions)
		}
		return v.RegistrationMethod
	case *Notification:
		if v.RegistrationOptions != nil {
			typeName(MethodNameFromString(v.Method),
				"RegistrationOptions", v.RegistrationOptions)
		}
		return v.RegistrationMethod
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}

var interfaceTemplateGetMessageType = func(input any) string {
	switch input.(type) {
	case *Request:
		return "request"
	case *Notification:
		return "notification"
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}
