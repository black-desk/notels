package main

import (
	"fmt"
	"os"
	"text/template"
)

var clientTemplate = fmt.Sprintf(interfaceTemplate, "proxy", "LspClient")

func genClient(metaModel *MetaModel) {
	log.Info("generating lsp client interface")

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

	clientTemplate, err := template.New("client").
		Funcs(interfaceTemplateFuncs).
		Parse(clientTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for lsp client",
			"error", err)
	}

	data := []any{}

	for i := range metaModel.Requests {
		if metaModel.Requests[i].MessageDirection != "clientToServer" {
			data = append(data, &metaModel.Requests[i])
		}
	}
	for i := range metaModel.Notifications {
		if metaModel.Notifications[i].MessageDirection != "clientToServer" {
			data = append(data, &metaModel.Notifications[i])
		}
	}

	err = clientTemplate.Execute(clientGenFile, data)

	if err != nil {
		log.Fatalw("failed to execute template for lsp client",
			"error", err)
	}
}
