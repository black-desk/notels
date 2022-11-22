package main

import (
	"fmt"
	"os"
	"text/template"
)

var structTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []Structure */}}
{{typeCheck .}}

package lsp

import (
        "encoding/json"
        "fmt"
)

var StructureValidateFailed = func (name string) error {
        return fmt.Errorf( "structure \"%s\"validate failed", name)
}

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
                        {{$current}} {{getTypeName $name $current .Type}} {{getAnno .Name}}
                {{end}}
        }
        
        {{if hasRequireField .Properties}}
                func (this *{{$name}}) UnmarshalJSON(data []byte) error {
                        type {{$name}}Unmarshal {{$name}}
                        var tmpUnmarshal {{$name}}Unmarshal
                        err := json.Unmarshal(data, &tmpUnmarshal)
                        if err != nil {
                                return err
                        }
                        {{range .Properties}}
                                {{$current := getName .Name}}
                                {{if not .Optional}}
                                        if tmpUnmarshal.{{$current}} == nil {
                                                return StructureValidateFailed("{{$name}}")
                                        }
                                {{end}}
                        {{end}}
                        *this = {{$name}}(tmpUnmarshal)
                        return nil
                }

                func (this *{{$name}}) MarshalJSON() ([]byte, error) {
                        {{range .Properties}}
                                {{$current := getName .Name}}
                                {{if not .Optional}}
                                        if this.{{$current}} == nil {
                                                return nil, StructureValidateFailed("{{$name}}")
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

var structureTemplateTypeCheck = func([]Structure) string {
	return ""
}
var structureTemplateGetComment = Comment
var structureTemplateGetName = MethodNameFromString
var structureTemplateGetTypeProps = func(t *Type) string {
	return ""
}

var structureTemplateGetTypeName = func(prefix string, current string, t *Type) string {
	ret := typeName(prefix, current, t)
	if ret[0:1] == "[" || ret[0:3] == "map" {
		return ret
	}
	return "*" + ret
}
var structureTemplateGetAnno = func(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}
var structureTemplateHasRequireField = func(p []Property) bool {
	for _, prop := range p {
		if !prop.Optional {
			return true
		}
	}
	return false

}

func genStruct(metaModel *MetaModel) {
	log.Info("generating structures")

	fileName := "structures_gen.go"
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
		"typeCheck":       structureTemplateTypeCheck,
		"getComment":      structureTemplateGetComment,
		"getName":         structureTemplateGetName,
		"getTypeProps":    structureTemplateGetTypeProps,
		"getTypeName":     structureTemplateGetTypeName,
		"getAnno":         structureTemplateGetAnno,
		"hasRequireField": structureTemplateHasRequireField,
	}

	codeTemplate, err := template.New("structures").
		Funcs(funcs).
		Parse(structTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for structures",
			"error", err)
	}

	data := metaModel.Structures

	err = codeTemplate.Execute(codeFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for structures",
			"error", err)
	}
}
