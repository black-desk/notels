package main

import (
	"os"
	"sort"
	"text/template"
)

var StringLiteralToGenerate = map[string]string{}

func RegisterStringLiteral(name string, s *StringLiteral) {
	if _, ok := StringLiteralToGenerate[name]; !ok {
		StringLiteralToGenerate[name] = string(*s)
	}
}

type StringLiteralWithName struct {
	Name  string
	Value string
}

var stringLiteralTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []*StringLiteralWithName */}}
{{typeCheck .}}

package lsp

import (
        "encoding/json"
        "fmt"
)

var StringLiteralValidateFailed = func (name string) error {
        return fmt.Errorf( "literal structure \"%s\"validate failed", name)
}

{{range .}}
        type {{ .Name }} struct {
                V String // {{ .Value }}
        }
        
        func (this *{{.Name}}) UnmarshalJSON(data []byte) error {
                var tmp string
                err := json.Unmarshal(data, &tmp)
                if err != nil {
                        return err
                }
                if tmp != "{{ .Value }}" {
                        return StringLiteralValidateFailed("{{.Name}}")
                }
                this.V = String(tmp)
                return nil
        }

        func (this *{{.Name}}) MarshalJSON() ([]byte, error) {
                if this.V != "{{.Value}}" {
                        return nil, StringLiteralValidateFailed("{{.Name}}")
                }
                return json.Marshal(&this.V)
        }
{{end}}
`

var stringLiteralTemplateTypeCheck = func([]*StringLiteralWithName) string { return "" }

func genStringLiteral() {
	fileName := "string_literal_gen.go"
	genFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer genFile.Close()

	funcs := map[string]any{
		"typeCheck": stringLiteralTemplateTypeCheck,
	}

	codeTemplate, err := template.New("stringLiteral").
		Funcs(funcs).
		Parse(stringLiteralTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for string literal",
			"error", err)
	}

	data := []*StringLiteralWithName{}

	keys := []string{}

	for k := range StringLiteralToGenerate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		data = append(data, &StringLiteralWithName{
			Name:  k,
			Value: StringLiteralToGenerate[k],
		})
	}

	err = codeTemplate.Execute(genFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for string literal",
			"error", err)
	}
}
