package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

var OrToGenerate = map[string]*Type{}

type OrTypeWithName struct {
	Name  string
	Items []Type
}

func RegisterOr(name string, t *Type) {
	if name == "NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or" {
		log.Info(OrToGenerate[name])
	}
	if t.Items == nil || len(t.Items) < 2 {
		log.Fatalw("Or type should have Items field")
		panic("")
	}
	if _, ok := OrToGenerate[name]; !ok {
		OrToGenerate[name] = t
		parseOr(name, t)
	}
	if name == "NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or" {
		log.Info(OrToGenerate[name])
	}
}

func parseOr(name string, t *Type) {
	for i, item := range t.Items {
		typeName(name, fmt.Sprint(i), &item)
	}
}

var OrTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []Structure */}}
{{typeCheck .}}

package protocol

import (
        "encoding/json"
        "fmt"
)

var OrValidateFailed = func (name string) error {
        return fmt.Errorf( "or type \"%s\"validate failed", name)
}

{{range .}}
        {{$name := .Name}}
        type {{$name}} struct {
                {{range $i, $item := .Items}}{{$texti:=printf "%v" $i}}
                // {{getTypeName $name $texti $item}} {{end}}
                V interface{}
        }
        
        func (this *{{$name}}) UnmarshalJSON(data []byte) error {
                {{range $i, $item := .Items}}
                        {
                                {{$texti:=printf "%v" $i}}
                                var tmp {{getTypeName $name $texti $item}}
                                if err:=json.Unmarshal(data, &tmp);err==nil{
                                        this.V = tmp
                                        return nil
                                }
                        }
                {{end}}
                return OrValidateFailed("{{$name}}")
        }

        func (this *{{$name}}) MarshalJSON() ([]byte, error) {
                for {
                        {{range $i, $item := .Items}}
                                {{$texti:=printf "%v" $i}}
                                if _,ok:=this.V.({{getTypeName $name $texti $item}});ok {
                                        break;
                                }
                        {{end}}
                        return nil, OrValidateFailed("{{$name}}")
                }
                return json.Marshal(this.V)
        }
{{end}}
`
var orTemplateTypeCheck = func([]OrTypeWithName) string { return "" }

func genOr() {
	fileName := "or_gen.go"
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
		"typeCheck":       orTemplateTypeCheck,
		"getTypeProps":    structureTemplateGetTypeProps,
		"getTypeName":     structureTemplateGetTypeName,
		"getAnno":         structureTemplateGetAnno,
		"hasRequireField": structureTemplateHasRequireField,
	}

	serverTemplate, err := template.New("structures").
		Funcs(funcs).
		Parse(OrTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for structs",
			"error", err)
	}

	data := []OrTypeWithName{}

	keys := []string{}
	for k := range OrToGenerate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == "NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or" {
			log.Info(OrToGenerate[k])
		}
		data = append(data, OrTypeWithName{
			Name:  k,
			Items: OrToGenerate[k].Items,
		})
	}

	err = serverTemplate.Execute(extraTypeGenFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for structs",
			"error", err)
	}
}
