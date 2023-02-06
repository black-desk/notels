package jsonrpc2

type Adaptor interface {
	Method(string) Method
}
