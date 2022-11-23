package jsonrpc2

type jsonrpcContext struct {
	ID   MessageID
	Conn *Conn
}
