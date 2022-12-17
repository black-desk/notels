package jsonrpc2

import (
	"context"
	"encoding/json"
	"errors"

	. "github.com/black-desk/lib/go/errwrap"
)

type JsonRPCReadWriteCloser interface {
	Read(context.Context) (json.RawMessage, error)
	Write(context.Context, json.RawMessage) error
	Close() error
}

type Conn struct {
	jsonRPCReadWriteCloser JsonRPCReadWriteCloser

	err error

	ctx    context.Context
	cancel context.CancelFunc

	messageToRead  chan json.RawMessage
	messageToWrite chan json.RawMessage
}

func NewConn(jsonRPCReadWriteCloser JsonRPCReadWriteCloser) *Conn {
	ctx, cancel := context.WithCancel(context.Background())
	return &Conn{
		jsonRPCReadWriteCloser: jsonRPCReadWriteCloser,
		err:                    nil,
		ctx:                    ctx,
		cancel:                 cancel,
		messageToRead:          make(chan json.RawMessage),
		messageToWrite:         make(chan json.RawMessage),
	}
}

func (c *Conn) Run() error {

	go c.startHandleRead()
	go c.startHandleWrite()
	go c.startDispatch()

	<-c.ctx.Done()
	return Trace(c.err)
}

func (c *Conn) startHandleRead() {
	for {
		select {
		case <-c.ctx.Done():
			log.Infof("stop read message")
			return
		default:
			msg, err := c.jsonRPCReadWriteCloser.Read(c.ctx)
			c.handleJsonRPCReadWriteCloserError(err)
			if c.err != nil {
				continue
			}
			c.messageToRead <- msg
		}
	}
}

func (c *Conn) startHandleWrite() {
	for {
		select {
		case <-c.ctx.Done():
			log.Infof("stop write message")
			return
		case msg := <-c.messageToWrite:
			err := c.jsonRPCReadWriteCloser.Write(c.ctx, msg)
			c.handleJsonRPCReadWriteCloserError(err)
		}
	}
}

func (c *Conn) startDispatch() {
	for {
		select {
		case <-c.ctx.Done():
			log.Infof("stop dispatch message")
			return
		case msg := <-c.messageToRead:
			err := c.dispatchMessage(msg)
			if err != nil {
				c.handleDispatchMessageError(Trace(err))
			}
		}
	}
}

func (c *Conn) handleJsonRPCReadWriteCloserError(err error) {
	if err == nil {
		return
	}

	if errors.Is(err, ErrNoMessageToRead) {
		return
	}

	if !errors.Is(err, context.Canceled) {
		if c.err != nil {
			panic(
				"connection error reset detected",
			)
		}
		c.err = Trace(err)
	}

	c.cancel()
}

func (c *Conn) dispatchMessage(msg json.RawMessage) error {
	var req Request
	var resp Response

	if messageAs(msg, &req) {
	}

	if messageAs(msg, &resp) {
	}

	return nil
}

func (c *Conn) handleDispatchMessageError(err error) {}
