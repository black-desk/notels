package message

import (
	"encoding/json"
	"fmt"

	. "github.com/black-desk/lib/go/errwrap"
	. "github.com/black-desk/notels/pkg/jsonrpc2/internal/log"
)

const (
	CodeParseError        int = -32700
	CodeInvalidRequest        = -32600
	CodeMethodNotFound        = -32601
	CodeInvalidParams         = -32602
	CodeInternalError         = -32603
	CodeServerErrorStart      = -32000
	CodeServerErrorEnd        = -32099
	MessageInvalidRequest     = "Invalid Request"
	MessageMethodNotFound     = "Method not found"
	MessageInternalError      = "Internal error"
)

type Error struct {
	// NOTE(black_desk): "code" could be 0
	Code *int `json:"code" validate:"required"`

	// NOTE(black_desk): seems that "message" could be empty string
	Message *string         `json:"message"        validate:"required"`
	Data    json.RawMessage `json:"data,omitepmty"`
}

func NewError(code int, message string, data json.RawMessage) *Error {
	if CodeServerErrorStart <= code && code <= CodeServerErrorEnd {
		Log.Warnw(
			fmt.Sprintf(
				"Error code range [%d, %d] is reserved for server implementation you SHOULD NOT use them",
				CodeServerErrorStart,
				CodeServerErrorEnd,
			),
			"code", code,
		)
	}
	return newError(code, message, data)
}

func newError(code int, message string, data json.RawMessage) *Error {
	ret := &Error{
		Code:    new(int),
		Message: new(string),
		Data:    nil,
	}

	*ret.Code = code
	*ret.Message = message
	ret.Data = data

	return ret

}

type Response struct {
	JsonRPC string          `json:"jsonrpc" validate:"required,eq=2.0"`
	Result  json.RawMessage `json:"result"  validate:"excluded_with=Error,required_without=Error"`
	Error   *Error          `json:"error"   validate:"excluded_with=Result,required_without=Result"`
	Id      *ID             `json:"id"      validate:"required"`
}

func (resp *Response) UnmarshalJSON(raw []byte) (err error) {
	type ResponseUnmarshal Response
	tmp := &ResponseUnmarshal{}

	err = Trace(json.Unmarshal(raw, tmp))
	if err != nil {
		return
	}

	err = Trace(validate.Struct(tmp))
	if err != nil {
		return
	}

	*resp = Response(*tmp)
	return
}

func (resp *Response) MarshalJSON() (raw []byte, err error) {
	type ResponseMarshal Response
	tmp := (*ResponseMarshal)(resp)

	raw, err = json.Marshal(tmp)
	if err != nil {
		err = Trace(err)
		return
	}

	err = Trace(validate.Struct(tmp))
	return
}
