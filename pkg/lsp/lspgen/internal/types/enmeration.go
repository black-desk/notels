package types

import (
	"encoding/json"
	"os"
	"strings"
	"text/template"

	"github.com/black-desk/notels/internal/utils/logger"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/common"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/extra"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/naming"
)

var log = logger.Get("lspgen")

var enumTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []model.Enumeration */}}
{{typeCheck .}}

package protocol

import (
        "errors"
        "encoding/json"
)

var EnumerationValidateFailed error = errors.New(
        "enumeration validate failed",
)

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

        var _{{$name}} = []{{$name}}{
                {{range .Values}}
                        {{getEnumName $name .Name}},
                {{end}}
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
                return EnumerationValidateFailed
        }
{{end}}
`

func GenEnumerations(metaModel *model.MetaModel) {

	fileName := "enumerations_gen.go"
	enumerationGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer enumerationGenFile.Close()

	funcs := map[string]any{
		"typeCheck":    enumTemplateTypeCheck,
		"getType":      enumTemplateGetType,
		"getName":      enumTemplateGetName,
		"getComment":   enumTemplateGetComment,
		"getEnumName":  enumTemplateGetEnumName,
		"getEnumValue": enumTemplateGetEnumValue,
	}

	serverTemplate, err := template.New("enumerations").
		Funcs(funcs).
		Parse(enumTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for enum",
			"error", err)
	}

	data := metaModel.Enumerations

	err = serverTemplate.Execute(enumerationGenFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for enum",
			"error", err)
	}
}

var enumTemplateTypeCheck = func([]model.Enumeration) string {
	return ""
}
var enumTemplateGetType = func(s string) string {
	switch s {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "uinteger":
		return "uint"
	default:
		log.Fatalw("unexpected type",
			"s", s)
		panic("")
	}
}
var enumTemplateGetName = func(name string) (ret string) {
	defer extra.RemoveType(ret)
	return naming.MethodNameFromString(name)
}
var enumTemplateGetComment = common.Comment
var enumTemplateGetEnumName = func(t string, name string) string {
	return t + strings.ToUpper(name[0:1]) + name[1:]
}
var enumTemplateGetEnumValue = func(value json.RawMessage) string {
	return string(value)
}
