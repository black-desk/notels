package main

var typeAliasesTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "encoding/json"
)

{{range .}}
{{comment .Documentation}} type {{goTypeAliasesName .}} {{goTypeAliasesBody .}}

{{goTypeAliasesUnmarshalMethod .}}
{{end}}
`

var interfaceUnmarshalMethodTemplate string = `
{{$name := goTypeAliasesName .}}
func lspgenUnmarshalJSONto{{$name}}(
        raw json.RawMessage,
) (
        ret {{$name}},
        err error,
){
        unmarshalError := interfaceUnmarshalJSONError{}
        {{range .Items}}
        {
                var tmp {{goTypeName .}}
                if jsonErr := json.Unmarshal(raw, &tmp); jsonErr == nil {
                        ret = &tmp
                        return
                } else {
                        unmarshalError = append(unmarshalError, jsonErr)
                }
                
        }
        {{end}}

        err = &unmarshalError

        return
}
`
