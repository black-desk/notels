package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

type andWithName struct {
	Name string
	T    *Type
}

var andToGenerate = map[string]andWithName{}

func RegisterAnd(name string, t *Type) {
	if t.Items == nil || len(t.Items) < 2 {
		log.Fatalw("And type should have Items field")
		panic("")
	}
	if _, ok := andToGenerate[name]; !ok {
		andToGenerate[name] = andWithName{
			Name: name,
			T:    t,
		}
		for i, item := range t.Items {
			typeName(name, fmt.Sprint(i), &item)
		}
	} else if t != andToGenerate[name].T {
		log.Fatalw("unexpected situation",
			"name", name,
			"t", t,
			"andToGenerate[name].T", andToGenerate[name].T)
		panic("")
	}
}

var andTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []Structure */}}
{{typeCheck .}}

package lsp

{{range .}}
        type {{getName .Name}} struct {
                {{range .T.Items}}
                        {{getName .Name}}
                {{end}}
        }
{{end}}
`

var andTemplateTypeCheck = func([]andWithName) string {
	return ""
}
var andTemplateGetName = MethodNameFromString

func genAnd() {
	log.Info("generating and type")

	fileName := "and_gen.go"
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
		"typeCheck": andTemplateTypeCheck,
		"getName":   andTemplateGetName,
	}

	codeTemplate, err := template.New("and").
		Funcs(funcs).
		Parse(andTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for and types",
			"error", err)
	}

	data := []andWithName{}

	keys := []string{}
	for key := range andToGenerate {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		data = append(data, andToGenerate[key])
	}

	err = codeTemplate.Execute(codeFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for and types",
			"error", err)
	}
}
