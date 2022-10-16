package main

var interfaceTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "context"
)

//jsonrpc2gen:%s
type %s interface {
{{range .}}
{{comment .Documentation}} {{methodName .}}{{methodArgs .}}{{methodReturn .}} //jsonrpc2gen:"{{jsonName .}}"
{{end}}
}
`

var enumerationsTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "errors"
)

var EnumerationValidateFailed error = errors.New("enumeration validate failed")

{{range .}}
{{comment .Documentation}} {{$type := goEnumType .Type.Name}} {{$name := goMethodName .Name}} type {{$name}} {{$type}}

const (
        {{range .Values}}
        {{comment .Documentation}} {{goEnumName $name .Name}} = {{goEnumValue .Value}} {{end}}
)

var {{$name}}ValidateSlice = []{{$name}}{ {{range .Values}}
        {{goEnumName $name .Name}}, {{end}}
}

func (this *{{$name}}) Validate() error {
        for _,x := range {{$name}}ValidateSlice {
                if *this == x {
                        return nil
                }
        }
        return EnumerationValidateFailed
}
{{end}}
`
