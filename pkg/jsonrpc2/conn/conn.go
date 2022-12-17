package conn

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	. "github.com/black-desk/lib/go/errwrap"
	"github.com/black-desk/notels/pkg/jsonrpc2"
	. "github.com/black-desk/notels/pkg/jsonrpc2/internal/log"
	"github.com/black-desk/notels/pkg/jsonrpc2/message"
	"github.com/sourcegraph/conc/pool"
)

type Conn struct {
	ctx      context.Context
	cancel   context.CancelFunc
	transfer jsonrpc2.Transfer
	debug    bool

	freezedMux sync.Mutex
	freezed    bool
	adaptors   []jsonrpc2.Adaptor

	cancels         sync.Map
	pendingRequests sync.Map

	p *pool.ErrorPool
}

func New(opts ...option) (conn *Conn, err error) {
	result := &Conn{
		ctx:      nil,
		cancel:   nil,
		transfer: nil,
	}

	for i := range opts {
		result = opts[i](result)
	}

	if result.ctx == nil && result.cancel == nil {
		ctx, cancel := context.WithCancel(context.Background())
		result.ctx = ctx
		result.cancel = cancel
	}

	verify := struct {
		Ctx      context.Context    `validate:"required"`
		Cancel   context.CancelFunc `validate:"required"`
		Transfer jsonrpc2.Transfer  `validate:"required"`
	}{
		Ctx:      result.ctx,
		Cancel:   result.cancel,
		Transfer: result.transfer,
	}

	err = validator().Struct(verify)
	if err != nil {
		err = Trace(err)
		return
	}

	conn = result

	conn.p = pool.New().WithErrors()
	return
}

func (c *Conn) Register(adaptor jsonrpc2.Adaptor) (err error) {
	c.freezedMux.Lock()
	defer c.freezedMux.Unlock()

	if c.freezed {
		err = Trace(ErrRegisterToRunning)
		return
	}

	c.adaptors = append(c.adaptors, adaptor)
	return
}

func (c *Conn) Loop() (err error) {
	err = Trace(c.freeze())
	if err != nil {
		return err
	}

	defer func() {
		err = Trace(c.p.Wait())
		if err != nil {
			c.cancel()
		}
	}()

	c.p.Go(func() error { return Trace(c.transfer.Run()) })
	c.p.Go(func() error { return Trace(c.loopRead()) })

	return
}

func (c *Conn) freeze() (err error) {
	c.freezedMux.Lock()
	defer c.freezedMux.Unlock()

	if c.freezed {
		return Trace(ErrDoubleFreeze)
	}

	c.freezed = true
	return
}

func (c *Conn) loopRead() (err error) {
	defer c.transfer.Close()
	for {
		select {
		case <-c.ctx.Done():
			err = Trace(ErrCanceled)
			return
		case msg := <-c.transfer.Read():
			c.dispatchMessage(msg)
		}
	}
}

func (c *Conn) dispatchMessage(raw json.RawMessage) (err error) {
	req := &message.Request{}
	err = json.Unmarshal(raw, req)
	if err == nil {
		c.dispatchRequest(req)
		return
	}

	requestUnmarshalErr := err

	resp := &message.Response{}
	err = json.Unmarshal(raw, resp)
	if err == nil {
		err = Trace(c.dispatchResponse(resp))
		return
	}

	responseUnmarshalErr := err
	err = nil

	c.response(
		c.ctx,
		c.getIDFromRawJSONMessage(raw),
		nil,
		fmt.Errorf("%w: [\n%s\n%s\n]", ErrInvalidMessage,
			requestUnmarshalErr.Error(),
			responseUnmarshalErr.Error(),
		),
	)

	return
}

func (c *Conn) dispatchRequest(req *message.Request) {
	for i := range c.adaptors {
		adaptor := c.adaptors[i]
		method := adaptor.Method(req.Method)
		if method == nil {
			continue
		}

		c.callMethod(req, method)
		return
	}

	c.response(
		c.ctx,
		req.Id,
		nil,
		Trace(fmt.Errorf("%w [method=\"%s\"]",
			ErrMethodNotFound, req.Method)),
	)
	return
}

func (c *Conn) callMethod(req *message.Request, method jsonrpc2.Method) {
	ctx, cancel := context.WithCancel(c.ctx)

	if err := c.remberCancel(req.Id, cancel); err != nil {
		Log.With("request", req).Infof("refused: %s", err.Error())
		return
	}

	c.p.Go(func() error {
		response, err := method(ctx, req)
		if err != nil {
			response = nil
		}

		if response == nil && err == nil {
			return nil
		}

		if req.Id != nil {
			c.cancels.Delete(req.Id.HashKey())
		}

		c.response(ctx, req.Id, response, Trace(err))
		return nil
	})

	return
}

func (c *Conn) remberCancel(id *message.ID, cancel context.CancelFunc) (err error) {
	if id == nil {
		return nil
	}

	_, loaded := c.cancels.LoadOrStore(id.HashKey(), cancel)
	if loaded {
		return Trace(ErrDuplicateID)
	}

	return nil
}

func (c *Conn) response(ctx context.Context, id *message.ID, response *message.Response, err error) {
	if response == nil {
		c.write(ctx, c.generateRawResponseFromError(id, err))
		return
	}

	rawResp, err := json.Marshal(response)
	if err != nil {
		rawResp = c.generateRawResponseFromError(id, Trace(err))
	}

	c.write(ctx, rawResp)
}

func (c *Conn) write(ctx context.Context, raw json.RawMessage) {
	if raw == nil {
		return
	}

	select {
	case <-ctx.Done():
		Log.With("message", raw).Debugf("method call canceled, message not send")
		return
	case c.transfer.Write() <- &jsonrpc2.MessageToWrite{
		Ctx: ctx,
		Msg: raw,
	}:
	}
}

func (c *Conn) generateRawResponseFromError(id *message.ID, err error) (
	raw json.RawMessage,
) {
	if err == nil {
		return
	}

	if id == nil {
		// NOTE(black_desk): There is an error but we cannot figure out
		// the corresponds request ID (perhaps c.getIDFromRawJSONMessage
		// fail its job or we are handling a notification). So we cannot
		// send this error as a error in response. We just log it here
		// as a warning.
		Log.Warnw("error occur", "error", err)
		return nil
	}

	resp := message.Response{
		JsonRPC: "2.0",
		Result:  nil,
		Error:   nil,
		Id:      id,
	}

	switch {
	case errors.Is(err, ErrInvalidMessage):
		resp.Error = message.NewError(
			message.CodeInvalidRequest,
			message.MessageInvalidRequest,
			c.generateErrData(err),
		)
	case errors.Is(err, ErrMethodNotFound):
		resp.Error = message.NewError(
			message.CodeMethodNotFound,
			message.MessageMethodNotFound,
			c.generateErrData(err),
		)
	default:
		resp.Error = message.NewError(
			message.CodeInternalError,
			message.MessageInternalError,
			c.generateErrData(err),
		)
	}
	// TODO(black_desk): add more

	raw, err = json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return raw
}

func (c *Conn) generateErrData(err error) (raw json.RawMessage) {
	if !c.debug {
		return
	}

	raw, err = json.Marshal(err.Error())
	if err != nil {
		panic(err)
	}
	return
}

func (c *Conn) getIDFromRawJSONMessage(raw json.RawMessage) (id *message.ID) {
	message := &struct {
		Id *message.ID `json:"id"`
	}{}

	err := json.Unmarshal(raw, message)
	if err != nil {
		Log.With("raw", raw).
			Debug("try to get id from raw JSON failed: %s", err.Error())
	}

	if message.Id != nil {
		id = message.Id
	}
	return
}

func (c *Conn) dispatchResponse(resp *message.Response) (err error) {
	id := resp.Id.HashKey()
	v, loaded := c.pendingRequests.LoadAndDelete(id)
	if !loaded {
		Log.With(id).Info(ErrUnknowID)
		return
	}

	if v == nil {
		Log.With(id).Info("cannot found pending requests: value is nil")
		return
	}

	respChan, ok := v.(chan<- *message.Response)
	if !ok {
		Log.With(id).Info("cannot found funpending requests: value is not a response channel")
	}

	select {
	case respChan <- resp:
	case <-c.ctx.Done():
		err = Trace(ErrCanceled)
	}

	return
}

func (c *Conn) Cancel(ids ...string) (err error) {
	if len(ids) == 0 {
		c.cancel()
		return
	}

	for i := range ids {
		id := ids[i]
		v, loaded := c.cancels.Load(id)
		if !loaded {
			return Trace(fmt.Errorf("%w %v", ErrUnknowID, id))
		}

		if v == nil {
			Log.With(id).Error("cannot found cancel func: value is nil")
			continue
		}

		fn, ok := v.(context.CancelFunc)
		if !ok {
			Log.With(id).Error("cannot found cancel func: value is not a cancel func")
			continue
		}

		fn()
	}

	return
}

func (c *Conn) Request(request *message.Request) (resp *message.Response, err error) {
	return c.RequestC(context.Background(), request)
}

func (c *Conn) RequestC(ctx context.Context, request *message.Request) (
	resp *message.Response, err error,
) {
	if request.IsNotification() {
		err = Trace(ErrNotACall)
		return
	}

	raw, err := json.Marshal(request)
	if err != nil {
		return
	}

	respChan := make(chan *message.Response)
	_, loaded := c.pendingRequests.LoadOrStore(
		request.Id.HashKey(),
		(chan<- *message.Response)(respChan))
	if loaded {
		err = Trace(ErrDuplicateID)
		return
	}

	select {
	case <-c.ctx.Done():
		err = Trace(ErrCanceled)
	case <-ctx.Done():
		err = Trace(ErrCanceled)
	case c.transfer.Write() <- &jsonrpc2.MessageToWrite{Ctx: ctx, Msg: raw}:
	}
	if err != nil {
		return
	}

	select {
	case resp = <-respChan:
	case <-c.ctx.Done():
		err = Trace(ErrCanceled)
	case <-ctx.Done():
		err = Trace(ErrCanceled)
	}

	return
}

func (c *Conn) Notify(notification *message.Request) (err error) {
	if !notification.IsNotification() {
		err = ErrNotANotification
		return
	}

	raw, err := json.Marshal(notification)
	if err != nil {
		return
	}

	select {
	case c.transfer.Write() <- &jsonrpc2.MessageToWrite{
		Ctx: context.Background(),
		Msg: raw,
	}:
	case <-c.ctx.Done():
		err = Trace(ErrCanceled)
	}

	return
}
