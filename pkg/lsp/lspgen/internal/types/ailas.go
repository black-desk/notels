package types

import (
	"os"
	"text/template"

	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/naming"
)

var typeTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

{{range .}}
        type {{getName .Name}} = {{getTypeName .Type}}
{{end}}
`

var typeTemplateFuncs = map[string]any{
	"getName":     naming.MethodNameFromString,
	"getTypeName": typeTemplateGetTypeName,
}

var typeTemplateGetTypeName = func(t *model.Type) string {
	name := naming.TypeName(t, "")
	return name
}

func GenAlias(metaModel *model.MetaModel) {
	fileName := "alias_gen.go"
	aliasGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer aliasGenFile.Close()

	codeTemplate, err := template.New("alias").
		Funcs(typeTemplateFuncs).
		Parse(typeTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for enum",
			"error", err)
	}

	data := metaModel.TypeAliases

	err = codeTemplate.Execute(aliasGenFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for enum",
			"error", err)
	}
}
