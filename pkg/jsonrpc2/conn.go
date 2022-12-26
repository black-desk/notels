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

	ctx    context.Context
	Cancel context.CancelFunc

	messageToRead       chan json.RawMessage
	messageToWrite      chan json.RawMessage
	messageToWriteMutex sync.Mutex

	requests chan *Request

	pendingCall sync.Map
	waitGroup   sync.WaitGroup
}

func NewConn(jsonRPCReadWriteCloser JsonRPCReadWriteCloser) *Conn {
	ctx, cancel := context.WithCancel(context.Background())
	return &Conn{
		jsonRPCReadWriteCloser: jsonRPCReadWriteCloser,
		lastErr:                nil,
		lastErrMutex:           sync.Mutex{},
		ctx:                    ctx,
		Cancel:                 cancel,
		messageToRead:          make(chan json.RawMessage),
		messageToWrite:         make(chan json.RawMessage),
		requests:               make(chan *Request),
		pendingCall:            sync.Map{},
	}
}

func (c *Conn) Requests() <-chan *Request {
	return c.requests
}

func (c *Conn) WriteRequest(req *Request) error {
	if err := validate.Struct(req); err != nil {
		return Trace(err)
	}

	if req.IsNotification() {
		panic(
			"should not pass notification message to WriteRequest, use WriteNotification",
		)
	}

	msg, err := json.Marshal(req)
	if err != nil {
		return Trace(err)
	}

	_, loaded := c.pendingCall.LoadOrStore(
		req.Id.toHashKey(), make(chan *Response))

	if loaded {
		panic("JSON-RPC 2.0 ID reused")
	}

	c.messageToWrite <- msg
	return nil
}

func (c *Conn) WriteNotification(req *Request) error {
	if err := validate.Struct(req); err != nil {
		return Trace(err)
	}

	if !req.IsNotification() {
		panic(
			"should not pass method call message to WriteNotification, use WriteRequest",
		)
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

func (c *Conn) write(msg json.RawMessage) {
	c.waitGroup.Add(1)
	go func() {
		defer c.waitGroup.Done()
		c.messageToWrite <- msg
	}()
}

func (c *Conn) Run() error {
	c.waitGroup.Add(3)

	go c.startHandleRead()
	go c.startHandleWrite()
	go c.startDispatch()

	c.waitGroup.Wait()
	return Trace(c.lastErr)
}

func (c *Conn) startHandleRead() {
	defer close(c.messageToRead)
	defer c.waitGroup.Done()
	for {
		select {
		case <-c.ctx.Done():
			log.Infof("stop read message")
			return
		default:
			msg, err := c.jsonRPCReadWriteCloser.Read(c.ctx)
			if c.handleJsonRPCReadWriteCloserError(Trace(err)) {
				c.messageToRead <- msg
			}
		}
	}
}

func (c *Conn) startHandleWrite() {
	defer c.waitGroup.Done()
	for {
		select {
		case <-c.ctx.Done():
			log.With(c.ctx).Infof("stop write message")
			return
		case msg := <-c.messageToWrite:
			err := c.jsonRPCReadWriteCloser.Write(c.ctx, msg)
			c.handleJsonRPCReadWriteCloserError(Trace(err))
		}
	}
}

func (c *Conn) startDispatch() {
	defer close(c.requests)
	defer c.waitGroup.Done()
	for {
		select {
		case <-c.ctx.Done():
			log.With(c.ctx).Infof("stop dispatch message")
			return
		case msg := <-c.messageToRead:
			err := c.dispatchMessage(msg)
			if err != nil {
				log.With(c.ctx).Errorf(
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
		log.With(c.ctx).
			Debugf("set connection error then cancel context")
		c.lastErr = Trace(err)
		c.Cancel()
		log.With(c.ctx).Debug("done")
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
		respChan, loaded := c.pendingCall.LoadAndDelete(
			resp.Id.toHashKey())

		if !loaded {
			log.Warnf("Unexceptional JSON-RPC 2.0 ID %v", resp.Id)
			return nil
		}

		respChan.(chan *Response) <- &resp
		return nil
	}

	return Trace(ErrInvalidMessage)
}
