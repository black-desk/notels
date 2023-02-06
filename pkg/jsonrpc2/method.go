package jsonrpc2

import (
	"context"
	"encoding/json"

	"github.com/black-desk/notels/pkg/jsonrpc2/message"
)

type Method func(ctx context.Context, param json.RawMessage) (response *message.Response, err error)
