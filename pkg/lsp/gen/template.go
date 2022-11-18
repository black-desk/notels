package main

var interfaceTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "context"
)

//jsonrpc2gen:%s
type %s interface {
{{range .}}
{{comment .Documentation}}
{{methodName .}}{{methodArgs .}}{{methodReturn .}} //jsonrpc2gen:"{{jsonName .}}"
{{end}}
}
`

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
