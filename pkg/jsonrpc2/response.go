package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

const (
	CodeParseError       int = -32700
	CodeInvalidRequest       = -32600
	CodeMethodNotFound       = -32601
	CodeInvalidParams        = -32602
	CodeInternalError        = -32603
	CodeServerErrorStart     = -32000
	CodeServerErrorEnd       = -32099
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
		log.Warnw(
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
		Data:    []byte{},
	}

	*ret.Code = code
	*ret.Message = message
	ret.Data = data

	return ret

}

func newMethodNotFound(method string) *Error {
	return NewError(
		CodeMethodNotFound,
		fmt.Sprintf(`%s [method name="%s"]`,
			ErrMethodNotFound.Error(), method),
		nil,
	)
}

type Response struct {
	JsonRPC string          `json:"jsonrpc" validate:"required,eq=2.0"`
	Result  json.RawMessage `json:"result"  validate:"excluded_with=Error,required_without=Error"`
	Error   *Error          `json:"error"   validate:"excluded_with=Result,required_without=Result"`
	Id      ID              `json:"id"      validate:"required"`
}
