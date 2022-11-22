package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

var TuplesToGenerate = map[string]*Type{}

func RegisterTuple(name string, t *Type) {
	if _, ok := TuplesToGenerate[name]; !ok {
		TuplesToGenerate[name] = t
		parseTuple(name, t)
	}
}

func parseTuple(name string, t *Type) {
	for i, item := range t.Items {
		typeName(name, fmt.Sprint(i), &item)
	}
}

type tupleTypeWithName struct {
	Name string
	T    *Type
}

var tupleTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

{{/* type check to ensure data is []Structure */}}
{{typeCheck .}}

package protocol

import (
        "encoding/json"
        "fmt"
)

var TupleValidateFailed = func (name string) error {
        return fmt.Errorf( "tuple \"%s\"validate failed", name)
}

{{range .}}
        {{$name := .Name}}
        type {{$name}} struct {
                {{range $i, $item := .T.Items}}
                        {{$current:=printf "%v" $i}}
                        TupleItem{{$current}} {{getTypeName $name $current $item}}
                {{end}}
        }
        
        func (this *{{$name}}) UnmarshalJSON(data []byte) error {
                var msg = []json.RawMessage{}
                if err := json.Unmarshal(data, &msg); err != nil{
                        return err
                }

                var tmpUnmarshal {{$name}}
                {{range $i, $item := .T.Items}}
                        {{$current:=printf "%v" $i}}
                        if err:=json.Unmarshal(msg[{{$current}}], &tmpUnmarshal.TupleItem{{$current}});err!=nil{
                                return err
                        }
                {{end}}
                *this = tmpUnmarshal
                return nil
        }

        func (this *{{$name}}) MarshalJSON() ([]byte, error) {
                var msg = []json.RawMessage{}
                {{range $i, $item := .T.Items}}
                        {{$current:=printf "%v" $i}}
                        {
                                bytes, err:=json.Marshal(this.TupleItem{{$current}})
                                if err != nil {
                                        return nil, err
                                }
                                msg = append(msg, bytes)
                        }
                {{end}}
                return json.Marshal(msg)
        }
{{end}}
`

var tupleTemplateTypeCheck = func([]tupleTypeWithName) string { return "" }
var tupleTemplateGetTypeName = typeName

func genTuple() {
	log.Info("generating tuples")

	fileName := "tuples_gen.go"
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
		"typeCheck":   tupleTemplateTypeCheck,
		"getTypeName": tupleTemplateGetTypeName,
	}

	codeTemplate, err := template.New("structures").
		Funcs(funcs).
		Parse(tupleTemplate)
	if err != nil {
		log.Fatalw("failed to parse template for structs",
			"error", err)
	}

	keys := []string{}
	for k := range TuplesToGenerate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	data := []tupleTypeWithName{}
	for _, k := range keys {
		data = append(data, tupleTypeWithName{
			Name: k,
			T:    TuplesToGenerate[k],
		})
	}

	err = codeTemplate.Execute(genFile, data)
	if err != nil {
		log.Fatalw("failed to execute template for tuples",
			"error", err)
	}
}
