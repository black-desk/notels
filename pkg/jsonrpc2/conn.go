package jsonrpc2

import (
	"io"
	"sync"
)

type Conn struct {
	conn             io.ReadWriteCloser
	protocol         Protocol
	waiting          sync.Map
	requestsToRead   chan *RequestMessage
	requestsToWrite  chan *RequestMessage
	responsesToRead  chan *ResponseMessage
	responsesToWrite chan *ResponseMessage
}

func (c *Conn) RequestsToRead() <-chan *RequestMessage {
	return c.requestsToRead
}

func (c *Conn) RequestsWrite() chan<- *RequestMessage {
	return c.requestsToWrite
}

func (c *Conn) ResponsesRead() <-chan *ResponseMessage {
	return c.responsesToRead
}

func (c *Conn) ResponsesToWrite() chan<- *ResponseMessage {
	return c.responsesToWrite
}

func (c *Conn) startLoop() {
	go func() {
		for resp := range c.responsesToWrite {
			err := c.protocol.WriteResponse(resp)
			if err != nil {
				log.Errorw("Failed to write response",
					"error", err)
			}
		}
	}()

	go func() {
		for req := range c.requestsToWrite {
			err := c.protocol.WriteRequest(req)
			if err != nil {
				log.Errorw("Failed to write request",
					"error", err)
			}
		}
	}()

	go func() {
		for {
			req, resp, err := c.protocol.Read()
			if err != nil {
				log.Errorw("Failed to read jsonrpc 2.0 message",
					"error", err)
			}
			if req != nil {
				c.requestsToRead <- req
			}
			if resp != nil {
				c.responsesToRead <- resp
			}
		}
	}()
}

func (c *Conn) registerID(ID MessageID) <-chan *ResponseMessage {
	if ID == nil {
		log.Fatalw("should not register nil ID")
	}
	any, ok := c.waiting.LoadOrStore(ID, make(chan *ResponseMessage))
	if !ok {
		log.Errorw("duplicate message ID detected", "ID", ID)
	}

	respChan, ok := any.(<-chan *ResponseMessage)
	if !ok {
		log.Fatal("unexpected value type", "any", any)
	}

	return respChan
}

func (c *Conn) SendMessage(req *RequestMessage) (*ResponseMessage, error) {
	var respChan <-chan *ResponseMessage = nil
	if req.ID != nil {
		respChan = c.registerID(req.ID)
	}

	c.RequestsWrite() <- req
	if respChan == nil {
		return nil, nil
	}

	return <-respChan, nil
}

func NewConn(io io.ReadWriter, protocol Protocol) *Conn {
	ret := &Conn{
		protocol:         protocol.SetIO(io),
		requestsToRead:   make(chan *RequestMessage),
		requestsToWrite:  make(chan *RequestMessage),
		responsesToRead:  make(chan *ResponseMessage),
		responsesToWrite: make(chan *ResponseMessage),
	}
	ret.startLoop()
	return ret
}
