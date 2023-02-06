package transfer

import (
	"encoding/json"

	"github.com/black-desk/notels/pkg/jsonrpc2"
)

type LSP struct {
}


func NewLSP() *LSP {
	return nil
}

// Close implements jsonrpc2.Transfer
func (*LSP) Close() {
	panic("unimplemented")
}

// Read implements jsonrpc2.Transfer
func (*LSP) Read() <-chan json.RawMessage {
	panic("unimplemented")
}

// Run implements jsonrpc2.Transfer
func (*LSP) Run() error {
	panic("unimplemented")
}

// Write implements jsonrpc2.Transfer
func (*LSP) Write() chan<- *jsonrpc2.MessageToWrite {
	panic("unimplemented")
}

// Error implements jsonrpc2.Transfer
func (*LSP) Error() <-chan error {
	panic("unimplemented")
}

var _ jsonrpc2.Transfer = &LSP{}
