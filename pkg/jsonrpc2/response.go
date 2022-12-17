package jsonrpc2

import "encoding/json"

type Error struct {
	// NOTE(black_desk): "code" could be 0
	Code *int `json:"code" validate:"required"`

	// NOTE(black_desk): seems that "message" could be empty string
	Message *string         `json:"message"        validate:"required"`
	Data    json.RawMessage `json:"data,omitepmty"`
}

type Response struct {
	JsonRPC string          `json:"jsonrpc" validate:"required,eq=2.0"`
	Result  json.RawMessage `json:"result"  validate:"excluded_with=Error,required_without=Error"`
	Error   *Error          `json:"error"   validate:"excluded_with=Result,required_without=Result"`
	Id      ID              `json:"id"      validate:"required"`
}
