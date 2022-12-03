package main

import (
	"os"
	"strings"
	"text/template"
)

var enumTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []model.Enumeration */}}
{{typeCheck .}}

package lsp

import (
        "fmt"
        "encoding/json"
)

var EnumerationValidateFailed = func (name string) error {
        return fmt.Errorf( "enumeration \"%s\"validate failed", name)
}


{{range .}}
        {{$type := getType .Type.Name}}
        {{$name := getName .Name}}

        {{getComment .Documentation}}
        type {{$name}} {{$type}}

        const (
                {{range .Values}}
                        {{getComment .Documentation}}
                        {{getEnumName $name .Name}} = {{getEnumValue .Value}}
                {{end}}
        )

        var _{{$name}} = []{{$name}}{ {{range .Values}}
                {{getEnumName $name .Name}}, {{end}}
        }

        func (this *{{$name}}) UnmarshalJSON(data []byte) error {
                type {{$name}}Unmarshal {{$name}}
                var tmpUnmarshal {{$name}}Unmarshal
                err := json.Unmarshal(data, &tmpUnmarshal)
                if err != nil {
                        return err
                }
                tmp := {{$name}}(tmpUnmarshal)
                for _, x := range _{{$name}} {
                        if tmp == x {
                                this = &tmp
                                return nil
                        }
                }
                return EnumerationValidateFailed("{{$name}}")
        }

        func (this *{{$name}}) MarshalJSON() ([]byte, error) {
                if err := func() error {
                        for _, x := range _{{$name}} {
                                if *this == x {
                                        return nil
                                }
                        }
                        return EnumerationValidateFailed("{{$name}}")
                }(); err != nil {
                        return nil, err
                }

                type {{$name}}Marshal {{$name}}
                tmpMarshal := {{$name}}Marshal(*this)
                return json.Marshal(tmpMarshal)
        }
{{end}}
`

var enumTemplateTypeCheck = func([]Enumeration) string {
	return ""
}
var enumTemplateGetType = baseTypeName
var enumTemplateGetName = func(name string) (ret string) {
	return MethodNameFromString(name)
}
var enumTemplateGetComment = Comment
var enumTemplateGetEnumName = func(t string, name string) string {
	return t + strings.ToUpper(name[0:1]) + name[1:]
}
var enumTemplateGetEnumValue = func(value []byte) string {
	return string(value)
}

func genEnumerations(metaModel *MetaModel) {
	log.Info("generating enumerations")

	fileName := "enumerations_gen.go"
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
		"typeCheck":    enumTemplateTypeCheck,
		"getType":      enumTemplateGetType,
		"getName":      enumTemplateGetName,
		"getComment":   enumTemplateGetComment,
		"getEnumName":  enumTemplateGetEnumName,
		"getEnumValue": enumTemplateGetEnumValue,
	}

	codeTemplate, err := template.New("enumerations").
		Funcs(funcs).
		Parse(enumTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for enumerations",
			"error", err)
	}

	data := metaModel.Enumerations

	err = codeTemplate.Execute(codeFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for enumerations",
			"error", err)
	}
}
