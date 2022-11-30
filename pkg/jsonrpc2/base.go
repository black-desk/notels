package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

type Code int64

const (
	ParseErrorCode     Code = -32700 // Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.
	InvalidRequestCode      = -32600 // The JSON sent is not a valid Request object.
	MethodNotFoundCode      = -32601 // The method does not exist / is not available.
	InvalidParamsCode       = -32602 // Invalid method parameter(s).
	InternalErrorCode       = -32603 // Internal JSON-RPC error.

	ServerErrorStart = -32000
	ServerErrorEnd   = -32099
)

type Integer int64

type String string

type Boolean bool

type Null struct{}

func (this *Null) UnmarshalJSON(data []byte) error {
	var tmp *int
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp != nil {
		return fmt.Errorf("Null should always be null")
	}
	return nil
}

func (this *Null) MarshalJSON() ([]byte, error) {
	var tmp *int = nil
	return json.Marshal(tmp)
}
