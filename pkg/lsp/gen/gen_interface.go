package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/black-desk/notels/pkg/lsp/gen/internal/components"
	"github.com/black-desk/notels/pkg/lsp/gen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/gen/internal/naming"
)

func methodName(input any) string {
	switch v := input.(type) {
	case model.Request:
		return naming.MethodName(v.Method)
	case model.Notification:
		return naming.MethodName(v.Method)
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}

func methodArgs(input any) string {
	result := "(\nctx context.Context"
	switch v := input.(type) {
	case model.Request:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("---")
			}
			result += fmt.Sprintf(",\nparams %s", v.Params[0].GoName())
		}
		result += ",\n)"
		return result
	case model.Notification:
		if v.Params != nil {
			if len(v.Params) != 1 {
				panic("---")
			}
			result += fmt.Sprintf(",\nparams %s", v.Params[0].GoName())
		}
		result += ",\n)"
		return result
	default:
		log.Fatalw("unexpected type",
			"input", input)
		panic("")
	}
}

func jsonName(input any) string {
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

func jsonRegistrationMethodName(input any) string {
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

func methodReturn(input any) string {
	result := "(\nerr error"
	switch v := input.(type) {
	case model.Request:
		if v.ErrorData != nil {
			result += fmt.Sprintf(",\nerrorData %s", naming.MethodName(v.Method)+"_ErrorData")
		}
		if v.Result != nil {
			result += fmt.Sprintf(",\nresult %s", naming.MethodName(v.Method)+"_Result")
		}
		if v.PartialResult != nil {
			result += fmt.Sprintf(",\npartialResult %s", naming.MethodName(v.Method)+"_PartialResult")
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

func genClient(model *model.MetaModel) {
	fileName := "client_gen.go"
	clientGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer clientGenFile.Close()

	clientTemplate, err := template.New("client").Funcs(map[string]any{
		"comment":      components.Comment,
		"methodName":   methodName,
		"methodArgs":   methodArgs,
		"methodReturn": methodReturn,
		"jsonName":     jsonName,
	}).Parse(fmt.Sprintf(interfaceTemplate, "proxy", "LspClient"))
	if err != nil {
		log.Fatalw("failed to parse template for lsp client",
			"error", err)
	}

	data := []any{}

	for _, req := range model.Requests {
		if req.MessageDirection != "clientToServer" {
			data = append(data, req)
		}
	}
	for _, notif := range model.Notifications {
		if notif.MessageDirection != "clientToServer" {
			data = append(data, notif)
		}
	}

	err = clientTemplate.Execute(clientGenFile, data)

	if err != nil {
		log.Fatalw("failed to execute template for lsp client",
			"error", err)
	}

}

func genServer(model *model.MetaModel) {

	fileName := "server_gen.go"
	serverGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer serverGenFile.Close()

	serverTemplate, err := template.New("server").Funcs(map[string]any{
		"comment":      components.Comment,
		"methodName":   methodName,
		"methodArgs":   methodArgs,
		"methodReturn": methodReturn,
		"jsonName":     jsonName,
	}).Parse(fmt.Sprintf(interfaceTemplate, "adaptor", "LspServer"))
	if err != nil {
		log.Fatalw("failed to parse template for lsp server",
			"error", err)
	}

	data := []any{}

	for _, req := range model.Requests {
		if req.MessageDirection != "serverToClient" {
			data = append(data, req)
		}
	}
	for _, notif := range model.Notifications {
		if notif.MessageDirection != "serverToClient" {
			data = append(data, notif)
		}
	}

	err = serverTemplate.Execute(serverGenFile, data)

	if err != nil {
		log.Fatalw("failed to execute template for lsp server",
			"error", err)
	}
}
