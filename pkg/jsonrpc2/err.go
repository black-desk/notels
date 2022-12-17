package jsonrpc2

import "errors"

var (
	ErrNoMessageToRead            = errNoMessageToRead()
	ErrInvalidID                  = errInvalidID()
	ErrUnsupportedProtocolVersion = errUnsupportedProtocolVersion()
	ErrInvalidMessage             = errInvalidMessage()
)

func errNoMessageToRead() error {
	return errors.New("No message to read on this JSON RPC 2 connection")
}

func errInvalidID() error {
	return errors.New("Invalid JSON-RPC 2.0 ID")
}

func errUnsupportedProtocolVersion() error {
	return errors.New("Unsupported JSON-RPC protocol version")
}

func errInvalidMessage() error {
	return errors.New(
		"Invalid JSON-RPC 2.0 message, check your JsonRPCReadWriteCloser implementation",
	)
}
