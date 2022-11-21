package main

import (
	"fmt"
	"os"
	"text/template"
)

var serverTemplate = fmt.Sprintf(interfaceTemplate, "adaptor", "LspServer")

func genServer(metaModel *MetaModel) {
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

	serverTemplate, err := template.New("client").
		Funcs(interfaceTemplateFuncs).
		Parse(serverTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for lsp client",
			"error", err)
	}

	data := []any{}

	for _, req := range metaModel.Requests {
		if req.MessageDirection != "serverToClient" {
			data = append(data, req)
		}
	}
	for _, notif := range metaModel.Notifications {
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
