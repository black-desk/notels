package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	Error() string
	Code() Code
	Data() json.RawMessage
}

type ErrorImpl struct {
	msg  string
	code Code
	data json.RawMessage
}

func AnalysisError(
	err error,
	errorData json.RawMessage,
	code Code,
	methodErr error,
	paramsUnmarshal bool,
	call bool,
	errorDataUnmarshal bool,
) Error {
	var Err ErrorImpl
	if err != nil {
		if !paramsUnmarshal {
			if syntax, ok := err.(*json.SyntaxError); ok {
				Err.code = ParseErrorCode
			} else if err.(*json.Error) {

			}
			Err.msg = fmt.Sprintf("%v", err)
		}
	}
	return nil
}
