package message

import (
	"encoding/json"

	. "github.com/black-desk/lib/go/errwrap"
)

type Request struct {
	JsonRPC string          `json:"jsonrpc"          validate:"required,eq=2.0"`
	Method  string          `json:"method"           validate:"required"`
	Params  json.RawMessage `json:"params,omitempty" validate:"nullable|json_object|json_array"`
	Id      *ID             `json:"id,omitempty"`
}

func (req *Request) IsNotification() bool { return req.Id == nil }

func (req *Request) UnmarshalJSON(raw []byte) (err error) {
	type RequestUnmarshal Request
	tmp := &RequestUnmarshal{}

	err = Trace(json.Unmarshal(raw, tmp))
	if err != nil {
		return
	}

	err = Trace(validate.Struct(tmp))
	if err != nil {
		return
	}

	*req = Request(*tmp)
	return
}

func (req *Request) MarshalJSON() (raw []byte, err error) {
	type RequestMarshal Request
	tmp := (*RequestMarshal)(req)

	raw, err = json.Marshal(tmp)
	if err != nil {
		err = Trace(err)
		return
	}

	err = Trace(validate.Struct(tmp))
	return
}
