package main

import (
	"fmt"
	"os"
	"text/template"
)

var structTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []Structure */}}
{{typeCheck .}}

package protocol

import (
        "errors"
        "encoding/json"
)

{{range .}}
        {{$name := getName .Name}}
        {{getComment .Documentation}}
        type {{$name}} struct {
                {{if .Extends}}
                // extends
                        {{range .Extends}}
                                {{getName .Name}}
                        {{end}}
                {{end}}
                {{if .Mixins}}
                // mixins
                        {{range .Mixins}}
                                {{getName .Name}}
                        {{end}}
                {{end}}
                {{range .Properties}}
                        {{$current := getName .Name}}
                        {{getComment .Documentation}}
                        {{$current}} {{if .Optional}}*{{end}} {{getTypeName $name $current .Type}} {{getAnno .Name}}
                {{end}}
        }
{{end}}
`

var structureTemplateTypeCheck = func([]Structure) string {
	return ""
}
var structureTemplateGetComment = Comment
var structureTemplateGetName = MethodNameFromString
var structureTemplateGetTypeProps = func(t *Type) string {
	return ""
}

var structureTemplateGetTypeName = typeName

func typeName(prefix string, current string, t *Type) string {
	if t.Kind == "reference" {
		return t.Name
	}
	if t.Kind == "base" {
		return baseTypeName(t.Name)
	}
	if t.Kind == "array" {
		return "[]" + typeName(prefix+"_"+current, "Element", t.Element)
	}
	if t.Kind == "or" {
		// register
		return prefix + "_" + current + "__Or"
	}
	if t.Kind == "map" {
		return "map[" + typeName(
			prefix+"_"+current,
			"Key",
			t.Key,
		) + "]interface{}"
	}
	if t.Kind == "literal" {
		return prefix + "_" + current
	}
	if t.Kind == "stringLiteral" {
		return "string"
	}
	log.Fatalw("unexpected kind",
		"kind", t.Kind,
	)
	panic("")
}

var structureTemplateGetAnno = func(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

func genStruct(metaModel *MetaModel) {
	fileName := "structures_gen.go"
	structuresGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer structuresGenFile.Close()

	funcs := map[string]any{
		"typeCheck":    structureTemplateTypeCheck,
		"getComment":   structureTemplateGetComment,
		"getName":      structureTemplateGetName,
		"getTypeProps": structureTemplateGetTypeProps,
		"getTypeName":  structureTemplateGetTypeName,
		"getAnno":      structureTemplateGetAnno,
	}

	serverTemplate, err := template.New("structures").
		Funcs(funcs).
		Parse(structTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for structs",
			"error", err)
	}

	data := metaModel.Structures

	err = serverTemplate.Execute(structuresGenFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for structs",
			"error", err)
	}

}