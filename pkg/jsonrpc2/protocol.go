package jsonrpc2

import "io"

// the transfer protocol wrap on the top of jsonrpc 2.0
type Protocol interface {
	SetIO(io.ReadWriter) Protocol
	Read() (*RequestMessage, *ResponseMessage, error)
	WriteRequest(req *RequestMessage) error
	WriteResponse(resp *ResponseMessage) error
}
