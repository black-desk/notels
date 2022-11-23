package jsonrpc2

import "encoding/json"

type Error interface {
	Error() string
	Code() Code
	Data() json.RawMessage
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
        return nil
}
