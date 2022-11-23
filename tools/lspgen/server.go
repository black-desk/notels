package main

import (
	"fmt"
	"os"
	"text/template"
)

var serverTemplate = fmt.Sprintf(interfaceTemplate, "LspServer")

func genServer(metaModel *MetaModel) {
	log.Info("generating lsp server interface")

	fileName := "server_gen.go"
	genFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer genFile.Close()

	codeTemplate, err := template.New("client").
		Funcs(interfaceTemplateFuncs).
		Parse(serverTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for lsp client",
			"error", err)
	}

	data := []any{}

	for i := range metaModel.Requests {
		if metaModel.Requests[i].MessageDirection != "serverToClient" {
			data = append(data, &metaModel.Requests[i])
		}
	}
	for i := range metaModel.Notifications {
		if metaModel.Notifications[i].MessageDirection != "serverToClient" {
			data = append(data, &metaModel.Notifications[i])
		}
	}

	err = codeTemplate.Execute(genFile, data)

	if err != nil {
		log.Fatalw("failed to execute template for lsp server",
			"error", err)
	}
}
