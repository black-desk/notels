package main

import (
	"os"
	"sort"
	"text/template"
)

var LiteralToGenerate = map[string]*StructureLiteral{}

type StructureLiteralWithName struct {
	Name    string
	Literal *StructureLiteral
}

func RegisterLiteral(name string, s *StructureLiteral) {
	if _, ok := LiteralToGenerate[name]; !ok {
		LiteralToGenerate[name] = s
		parseLiteral(name, s)
	}
}

func parseLiteral(name string, s *StructureLiteral) {
	for _, field := range *s.Properties {
		typeName(name, MethodNameFromString(field.Name), &field.Type)
	}
}

var literalTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []StructureLiteralWithName */}}
{{typeCheck .}}

package protocol

import (
        "encoding/json"
        "fmt"
)

var StructureLiteralValidateFailed = func (name string) error {
        return fmt.Errorf( "literal structure \"%s\"validate failed", name)
}

{{range .}}
        {{$name := .Name}}
        type {{$name}} struct {
                {{range .Literal.Properties}}
                        {{$current := getName .Name}}
                        {{getComment .Documentation}}
                        {{$current}} {{getTypeName $name $current .Type}} {{getAnno .Name}}
                {{end}}
        }
        
        {{if hasRequireField .Literal.Properties}}
                func (this *{{$name}}) UnmarshalJSON(data []byte) error {
                        type {{$name}}Unmarshal {{$name}}
                        var tmpUnmarshal {{$name}}Unmarshal
                        err := json.Unmarshal(data, &tmpUnmarshal)
                        if err != nil {
                                return err
                        }
                        {{range .Literal.Properties}}
                                {{$current := getName .Name}}
                                {{if not .Optional}}
                                        if tmpUnmarshal.{{$current}} == nil {
                                                return StructureLiteralValidateFailed("{{$name}}")
                                        }
                                {{end}}
                        {{end}}
                        *this = {{$name}}(tmpUnmarshal)
                        return nil
                }

                func (this *{{$name}}) MarshalJSON() ([]byte, error) {
                        {{range .Literal.Properties}}
                                {{$current := getName .Name}}
                                {{if not .Optional}}
                                        if this.{{$current}} == nil {
                                                return nil, StructureLiteralValidateFailed("{{$name}}")
                                        }
                                {{end}}
                        {{end}}
                        type {{$name}}Marshal {{$name}}
                        tmpMarshal := {{$name}}Marshal(*this)
                        return json.Marshal(&tmpMarshal)
                }
        {{end}}
{{end}}
`

var structureLiteralTemplateTypeCheck = func([]StructureLiteralWithName) string { return "" }

func genLiteral() {
	fileName := "literal_gen.go"
	extraTypeGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer extraTypeGenFile.Close()

	funcs := map[string]any{
		"typeCheck":       structureLiteralTemplateTypeCheck,
		"getName":         structureTemplateGetName,
		"getComment":      Comment,
		"getTypeName":     structureTemplateGetTypeName,
		"getAnno":         structureTemplateGetAnno,
		"hasRequireField": structureTemplateHasRequireField,
	}

	serverTemplate, err := template.New("structureLiteral").
		Funcs(funcs).
		Parse(literalTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for structs",
			"error", err)
	}

	data := []StructureLiteralWithName{}

	keys := []string{}

	for k := range LiteralToGenerate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		data = append(data, StructureLiteralWithName{
			Name:    k,
			Literal: LiteralToGenerate[k],
		})
	}

	err = serverTemplate.Execute(extraTypeGenFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for structs",
			"error", err)
	}
}
