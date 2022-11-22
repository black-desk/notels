package main

import (
	"os"
	"text/template"
)

var typeAliasTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

{{range .}}
        {{$name := getName .Name}}
        type {{$name}} = {{getTypeName "" $name .Type}}
{{end}}
`

func genTypeAliases(metaModel *MetaModel) {
	log.Info("generating type aliases")

	fileName := "alias_gen.go"
	codeFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer codeFile.Close()

	funcs := map[string]any{
		"getName":     MethodNameFromString,
		"getTypeName": typeName,
	}

	codeTemplate, err := template.New("alias").
		Funcs(funcs).
		Parse(typeAliasTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for type aliases",
			"error", err)
	}

	data := metaModel.TypeAliases

	err = codeTemplate.Execute(codeFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for type aliases",
			"error", err)
	}
}
