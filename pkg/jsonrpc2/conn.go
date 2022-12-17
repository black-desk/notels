package jsonrpc2

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	. "github.com/black-desk/lib/go/errwrap"
)

type JsonRPCReadWriteCloser interface {
	Read(context.Context) (json.RawMessage, error)
	Write(context.Context, json.RawMessage) error
	Close() error
}

type Conn struct {
	jsonRPCReadWriteCloser JsonRPCReadWriteCloser

	lastErr      error
	lastErrMutex sync.Mutex

	Ctx    context.Context
	Cancel context.CancelFunc

	messageToRead  chan json.RawMessage
	messageToWrite chan json.RawMessage

	requests  chan *Request
	responses chan *Response
}

func NewConn(jsonRPCReadWriteCloser JsonRPCReadWriteCloser) *Conn {
	ctx, cancel := context.WithCancel(context.Background())
	return &Conn{
		jsonRPCReadWriteCloser: jsonRPCReadWriteCloser,
		lastErr:                nil,
		lastErrMutex:           sync.Mutex{},
		Ctx:                    ctx,
		Cancel:                 cancel,
		messageToRead:          make(chan json.RawMessage),
		messageToWrite:         make(chan json.RawMessage),
		requests:               make(chan *Request),
		responses:              make(chan *Response),
	}
}

func (c *Conn) Requests() <-chan *Request {
	return c.requests
}

func (c *Conn) Responses() <-chan *Response {
	return c.responses
}

func (c *Conn) WriteRequest(req *Request) error {
	if err := validate.Struct(req); err != nil {
		return Trace(err)
	}

	msg, err := json.Marshal(req)
	if err != nil {
		return Trace(err)
	}

	c.messageToWrite <- msg
	return nil
}

func (c *Conn) WriteResponse(resp *Response) error {
	if err := validate.Struct(resp); err != nil {
		return Trace(err)
	}

	msg, err := json.Marshal(resp)
	if err != nil {
		return Trace(err)
	}

	c.messageToWrite <- msg
	return nil
}

func (c *Conn) Run() error {

	go c.startHandleRead()
	go c.startHandleWrite()
	go c.startDispatch()

	<-c.Ctx.Done()
	close(c.messageToRead)
	close(c.messageToWrite)
	close(c.requests)
	close(c.responses)
	return Trace(c.lastErr)
}

func (c *Conn) startHandleRead() {
	for {
		select {
		case <-c.Ctx.Done():
			log.Infof("stop read message")
			return
		default:
			msg, err := c.jsonRPCReadWriteCloser.Read(c.Ctx)
			if c.handleJsonRPCReadWriteCloserError(Trace(err)) {
				c.messageToRead <- msg
			}
		}
	}
}

func (c *Conn) startHandleWrite() {
	for {
		select {
		case <-c.Ctx.Done():
			log.With(c.Ctx).Infof("stop write message")
			return
		case msg := <-c.messageToWrite:
			err := c.jsonRPCReadWriteCloser.Write(c.Ctx, msg)
			c.handleJsonRPCReadWriteCloserError(Trace(err))
		}
	}
}

func (c *Conn) startDispatch() {
	for {
		select {
		case <-c.Ctx.Done():
			log.With(c.Ctx).Infof("stop dispatch message")
			return
		case msg := <-c.messageToRead:
			err := c.dispatchMessage(msg)
			if err != nil {
				log.With(c.Ctx).Errorf(
					"failed to dispatch message %v: %v",
					msg, err)
			}
		}
	}
}

func (c *Conn) handleJsonRPCReadWriteCloserError(err error) bool {
	if err == nil {
		return true
	}

	if errors.Is(err, ErrNoMessageToRead) {
		return true
	}

	if !errors.Is(err, context.Canceled) {
		c.lastErrMutex.Lock()
		defer c.lastErrMutex.Unlock()
		log.With(c.Ctx).
			Debugf("set connection error then cancel context")
		c.lastErr = Trace(err)
		c.Cancel()
		log.With(c.Ctx).Debug("done")
	}

	return false

}

func (c *Conn) dispatchMessage(msg json.RawMessage) error {
	var req Request
	var resp Response

	if messageAs(msg, &req) {
		c.requests <- &req
		return nil
	}

	if messageAs(msg, &resp) {
		c.responses <- &resp
		return nil
	}

	return Trace(ErrInvalidMessage)
}
