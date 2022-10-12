package main

var interfaceTemplate string = `// Code generated from metaModel.json by "lspgen". DO NOT EDIT
package protocol
import "context"
type %s interface {
{{range .}}
{{comment .}}
{{methodName .}}{{methodArgs .}}{{methodReturn .}} //jsonrpc2gen:"{{jsonName .}}"
{{if shouldGen .RegistrationMethod}}{{registrationMethodName .}}{{registrationMethodArgs .}}{{registrationMethodReturn .}} //jsonrpc2gen:"{{jsonRegistrationMethodName .}}" {{end}}
{{end}}
}
`
