package conn

import (
	"errors"
)

var (
	ErrCanceled          = errCanceled()
	ErrRegisterToRunning = errRegisterToFreezed()
	ErrDoubleFreeze      = errDoubleFreeze()
	ErrUnknowID          = errUnknowID()
	ErrNotACall          = errNotACall()
	ErrNotANotification  = errNotANotification()

	ErrInvalidMessage = errInvalidMessage()
	ErrMethodNotFound = errMethodNotFound()
	ErrDuplicateID    = errDuplicateID()
)

func errCanceled() error {
	return errors.New("context canceled")
}

func errRegisterToFreezed() error {
	return errors.New("can not register adaptor to freezed JSON-RPC 2.0 connection")
}

func errDoubleFreeze() error {
	return errors.New("can not freeze a freezed JSON-RPC 2.0 connection")
}

func errUnknowID() error {
	return errors.New("unknown JSON-RPC 2.0 ID")
}

func errNotACall() error {
	return errors.New("request is not a method call")
}

func errNotANotification() error {
	return errors.New("request is not a notification")
}

func errInvalidMessage() error {
	return errors.New("invalid JSON-RPC 2.0 message")
}

func errMethodNotFound() error {
	return errors.New("No such method or method not available now")
}

func errDuplicateID() error {
	return errors.New("duplicate JSON-RPC 2.0 message ID")
}
