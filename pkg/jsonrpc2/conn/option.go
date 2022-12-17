package conn

import (
	"context"

	"github.com/black-desk/notels/pkg/jsonrpc2"
)

type option func(*Conn) *Conn

func WithTransfer(transfer jsonrpc2.Transfer) option {
	return func(conn *Conn) *Conn {
		conn.transfer = transfer
		return conn
	}
}

func WithContext(ctx context.Context) option {
	return func(conn *Conn) *Conn {
		conn.ctx, conn.cancel = context.WithCancel(ctx)
		return conn
	}
}

func WithDebug() option {
	return func(conn *Conn) *Conn {
		conn.debug = true
		return conn
	}
}
