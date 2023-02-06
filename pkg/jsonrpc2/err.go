package jsonrpc2

import (
	"errors"
)

var (
        ErrTransferClosed = errTransferClosed()
)

func errTransferClosed() error {
        return errors.New("transfer is closed by user")
}
