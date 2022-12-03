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
	/*
	   if method == "textDocument/diagnostic" {
	           var tmpResult *lsp.DocumentDiagnosticReport

	           var tmpParams *lsp.DocumentDiagnosticParams
	           err = json.Unmarshal(params, &tmpParams)
	           if err != nil {
	                   return
	           }

	           paramsUnmarshal = true

	           var tmpErrorData *lsp.DiagnosticServerCancellationData

	           tmpResult, tmpErrorData, code, methodErr = a.service.TextDocumentDiagnostic(
	                   ctx,
	                   tmpParams,
	           )
	           call = true

	           if tmpErrorData != nil {
	                   errorData, err = json.Marshal(tmpErrorData)
	                   if err != nil {
	                           return
	                   }
	           }

	           errorDataUnmarshal = true
	           if tmpResult != nil {
	                   result, err = json.Marshal(tmpResult)
	                   if err != nil {
	                           return
	                   }
	           }
	           return
	   }
	*/

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
