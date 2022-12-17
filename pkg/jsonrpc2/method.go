package jsonrpc2

import (
	"context"

	"github.com/black-desk/notels/pkg/jsonrpc2/message"
)

type Method func(ctx context.Context, request *message.Request) (response *message.Response, err error)
