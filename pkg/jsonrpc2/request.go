package jsonrpc2

import (
	"encoding/json"
)

type Request struct {
	JsonRPC string          `json:"jsonrpc"          validate:"required,eq=2.0"`
	Method  string          `json:"method"           validate:"required"`
	Params  json.RawMessage `json:"params,omitempty" validate:"nullable|json_object|json_array"`
	Id      *ID             `json:"id,omitempty"`
}

func (req *Request) IsNotification() bool { return req.Id == nil }
