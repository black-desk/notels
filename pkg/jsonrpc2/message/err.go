package message

import "errors"

var (
	ErrInvalidID = errInvalidID()
)

func errInvalidID() error {
	return errors.New("Invalid JSON-RPC 2.0 ID")
}
