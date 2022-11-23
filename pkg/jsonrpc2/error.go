package jsonrpc2

import "encoding/json"

type Error interface {
	Error() string
	Code() int64
	Data() json.RawMessage
}

func AnalysisError(
	err error,
	errorData json.RawMessage,
	methodErr error,
	paramsUnmarshal bool,
	call bool,
	errorDataUnmarshal bool,
) Error {
	return nil
}
