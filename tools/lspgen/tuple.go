package main

import "fmt"

var TuplesTogenerate = map[string]*Type{}

func RegisterTuple(name string, t *Type) {
	if _, ok := TuplesTogenerate[name]; !ok {
		TuplesTogenerate[name] = t
		parseTuple(name, t)
	}
}

func parseTuple(name string, t *Type) {
	for i, item := range t.Items {
		typeName(name, fmt.Sprint(i), &item)
	}
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

func genTuple() {

}
