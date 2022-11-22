package main

import (
	"os"
	"text/template"
)

var typeAliasTemplate = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

{{range .}}
        {{$name := getName .Name}}
        type {{$name}} = {{getTypeName "" $name .Type}}
{{end}}
`

var typeAliasTemplateFuncs = map[string]any{
	"getName":     MethodNameFromString,
	"getTypeName": typeName,
}

func genTypeAliases(metaModel *MetaModel) {
	log.Info("generating type aliases")

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
		Funcs(typeAliasTemplateFuncs).
		Parse(typeAliasTemplate)
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
