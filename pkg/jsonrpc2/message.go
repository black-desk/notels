package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Version string `json:"jsonrpc"` // always be "2.0"
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type MessageUnmarshal Message
	var tmp MessageUnmarshal
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp.Version != "2.0" {
		return fmt.Errorf(
			"Unexpected jsonrpc version \"%v\", should be 2.0",
			tmp.Version,
		)
	}
	*m = Message(tmp)
	return nil
}

var message = Message{
	Version: "2.0",
}

// Integer or String
type MessageID interface {
	isMessageID()
}

func (*Integer) isMessageID() {}
func (*String) isMessageID()  {}

type RequestMessage struct {
	Message
	ID     MessageID       `json:"id,omitempty"` // when id is nil this message is a notification
	Params json.RawMessage `json:"params"`       // slice or object
	Method string          `json:"method"`
}

func (r *RequestMessage) UnmarshalJSON(data []byte) error {
	type RequestMessageUnmarshal struct {
		Message
		ID     json.RawMessage `json:"id"`
		Params json.RawMessage `json:"params"`
		Method string          `json:"method"`
	}
	var tmp RequestMessageUnmarshal
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var id MessageID
	for {
		{
			var tmp_ Integer
			if err := json.Unmarshal(tmp.ID, &tmp_); err == nil {
				id = &tmp_
				break
			}
		}
		{
			var tmp_ String
			if err := json.Unmarshal(tmp.ID, &tmp_); err == nil {
				id = &tmp_
				break
			}
		}
		return fmt.Errorf(
			"Failed to unmarshal message ID, it should be a number or string",
		)
	}

	r.Version = "2.0"
	r.ID = id
	r.Params = tmp.Params
	r.Method = tmp.Method
	return nil
}

type ResponseMessage struct {
	Message
	ID     MessageID             `json:"id"`
	Result json.RawMessage       `json:"result,omitempty"`
	Error  *ResponseMessageError `json:"error,omitempty"`
}

type ResponseMessageError struct {
	Code    Code           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (r *ResponseMessage) UnmarshalJSON(data []byte) error {
	type ResponseMessageUnmarshal struct {
		Message
		ID     json.RawMessage       `json:"id"`
		Result json.RawMessage       `json:"result,omitempty"`
		Error  *ResponseMessageError `json:"error,omitempty"`
	}
	var tmp ResponseMessageUnmarshal
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var id MessageID
	for {
		{
			var tmp_ Integer
			if err := json.Unmarshal(tmp.ID, &tmp_); err == nil {
				id = &tmp_
				break
			}
		}
		{
			var tmp_ String
			if err := json.Unmarshal(tmp.ID, &tmp_); err == nil {
				id = &tmp_
				break
			}
		}
		return fmt.Errorf(
			"Failed to unmarshal message ID, it should be a number or string",
		)
	}

	r.Version = "2.0"
	r.ID = id
	r.Result = tmp.Result
	r.Error = tmp.Error
	return nil
}
