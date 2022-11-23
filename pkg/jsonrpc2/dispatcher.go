package jsonrpc2

import (
	"context"

	"github.com/black-desk/notels/internal/utils/logger"
)

var log = logger.Get("jsonrpc2")

type RequestHandler struct {
	adaptor Adaptor
}

func NewDispatcher(adaptor Adaptor) *RequestHandler {
	return &RequestHandler{
		adaptor: adaptor,
	}
}

func (d *RequestHandler) Listen(conn *Conn) {
	requests := conn.RequestsToRead()
	for request := range requests {
		go d.HandleRequest(request, conn)
	}
	return
}

func (d *RequestHandler) HandleRequest(
	request *RequestMessage,
	conn *Conn,
) {
	ctx := context.WithValue(
		context.Background(),
		"jsonrpc2",
		&jsonrpcContext{
			ID:   request.ID,
			Conn: conn,
		},
	)

	if request.ID != nil {
		result, err := d.adaptor.Call(
			ctx,
			request.Method,
			request.Params,
		)

		var error *ResponseMessageError = nil

		if err != nil {
			log.Errorw(
				"handle request error",
				"id (nil for Notification)", request.ID,
				"error", err.Error(),
				"code", err.Code(),
			)

			if request.ID == nil {
				return
			}

			error = &ResponseMessageError{
				Code:    err.Code(),
				Message: err.Error(),
				Data:    err.Data(),
			}

			if result != nil {
				log.Warnw(
					"jsonrpc 2.0 method should not return both result and error",
					"ctx",
					ctx,
					"method",
					request.Method,
					"error",
					error,
				)
				result = nil
			}
		}

		response := ResponseMessage{
			Message: message,
			ID:      request.ID,
			Result:  result,
			Error:   error,
		}

		conn.ResponsesToWrite() <- &response
	} else {
		err := d.adaptor.Notify(
			ctx,
			request.Method,
			request.Params,
		)
		if err != nil {
			log.Errorw("handle notification error",
				"conn", conn,
				"error", err,
			)
		}
	}
}
