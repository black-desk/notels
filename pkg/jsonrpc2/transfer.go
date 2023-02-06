package jsonrpc2

import (
	"context"
	"encoding/json"
)

// Transfer is a interface used to handle the underlying JSON transfer details.
// jsonrpc2/conn.Conn need a Transfer to be initialized. There are some common
// transfers in jsonrpc2/transfer. Transfer should be safe for concurrency use.
type Transfer interface {
	// Read read a json message.
	// NOTE: jsonrpc2/conn.Conn do not handle parse error, transfer should
        // take over. Check https://www.jsonrpc.org/specification#error_object
	Read() <-chan json.RawMessage

	// Write write a json message.
	// The messages pass to channel will always be valid JSON-RPC 2.0
	// response.
	Write() chan<- *MessageToWrite

	// Run start this Transfer.
	// After Run called, this Transfer start to handle json reading/writing.
	// Run should never return unless Close called or some uncovered error
	// occur.
	Run() error

	// Close close this Transfer, check documentation of Run. After Close
        // return, the channel returned by Read should be closed
	Close()

        Error() <-chan error
}

type MessageToWrite struct {
        Ctx context.Context
        Msg json.RawMessage
}
