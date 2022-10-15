package main

var interfaceTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
        "context"
)

//jsonrpc2gen:%s
type %s interface {
{{range .}}
{{comment .}}
{{methodName .}}{{methodArgs .}}{{methodReturn .}} //jsonrpc2gen:"{{jsonName .}}"
{{end}}
}
`
