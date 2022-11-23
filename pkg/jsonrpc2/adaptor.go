package jsonrpc2

import (
	"context"
	"encoding/json"
)

type Adaptor interface {
	Call(
		ctx context.Context,
		method string,
		params json.RawMessage,
	) (result json.RawMessage, err Error)

	Notify(
		ctx context.Context,
		method string,
		params json.RawMessage,
	) error
}
