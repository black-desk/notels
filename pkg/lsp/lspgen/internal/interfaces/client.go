package interfaces

import (
	"fmt"
	"os"
	"text/template"

	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
)

var clientTemplate = fmt.Sprintf(interfaceTemplate, "proxy", "LspClient")

func GenClient(metaModel *model.MetaModel) {
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

	for _, req := range metaModel.Requests {
		if req.MessageDirection != "clientToServer" {
			data = append(data, req)
		}
	}
	for _, notif := range metaModel.Notifications {
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
