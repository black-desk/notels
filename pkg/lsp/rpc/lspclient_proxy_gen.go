// Code generated by "jsonrpc2gen". DO NOT EDIT

package rpc

import (
	"context"
	"encoding/json"
	"github.com/black-desk/notels/pkg/jsonrpc2"
	"github.com/black-desk/notels/pkg/lsp"
	"github.com/google/uuid"
)

type LspClientJSONRPCProxy struct {
	conn *jsonrpc2.Conn
}

func NewLspClientJSONRPCProxy(conn *jsonrpc2.Conn) *LspClientJSONRPCProxy {
	return &LspClientJSONRPCProxy{
		conn: conn,
	}
}

func (p *LspClientJSONRPCProxy) WorkspaceWorkspaceFolders(
	ctx context.Context,

) (
	result *lsp.WorkspaceWorkspaceFolders_Result_Or,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/workspaceFolders"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceConfiguration(
	ctx context.Context,
	params *lsp.ConfigurationParams,
) (
	result []lsp.LSPAny,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/configuration"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WindowWorkDoneProgressCreate(
	ctx context.Context,
	params *lsp.WorkDoneProgressCreateParams,
) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "window/workDoneProgress/create"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceSemanticTokensRefresh(
	ctx context.Context,

) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/semanticTokens/refresh"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WindowShowDocument(
	ctx context.Context,
	params *lsp.ShowDocumentParams,
) (
	result *lsp.ShowDocumentResult,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "window/showDocument"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceInlineValueRefresh(
	ctx context.Context,

) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/inlineValue/refresh"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceInlayHintRefresh(
	ctx context.Context,

) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/inlayHint/refresh"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceDiagnosticRefresh(
	ctx context.Context,

) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/diagnostic/refresh"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) ClientRegisterCapability(
	ctx context.Context,
	params *lsp.RegistrationParams,
) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "client/registerCapability"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) ClientUnregisterCapability(
	ctx context.Context,
	params *lsp.UnregistrationParams,
) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "client/unregisterCapability"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WindowShowMessageRequest(
	ctx context.Context,
	params *lsp.ShowMessageRequestParams,
) (
	result *lsp.WindowShowMessageRequest_Result_Or,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "window/showMessageRequest"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceCodeLensRefresh(
	ctx context.Context,

) (
	result *lsp.Null,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/codeLens/refresh"

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WorkspaceApplyEdit(
	ctx context.Context,
	params *lsp.ApplyWorkspaceEditParams,
) (
	result *lsp.ApplyWorkspaceEditResult,

	err error,
) {
	var requestMsg jsonrpc2.RequestMessage
	requestMsg.Version = "2.0"
	var uuidStr jsonrpc2.String = jsonrpc2.String(uuid.New().String())
	requestMsg.ID = &uuidStr
	requestMsg.Method = "workspace/applyEdit"

	requestMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	var responseMessage *jsonrpc2.ResponseMessage
	responseMessage, err = p.conn.Call(&requestMsg)

	if responseMessage.Result != nil {
		err = json.Unmarshal(responseMessage.Result, &result)
		if err != nil {
			return
		}
	}

	return
}

func (p *LspClientJSONRPCProxy) WindowShowMessage(
	ctx context.Context,
	params *lsp.ShowMessageParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "window/showMessage"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) WindowLogMessage(
	ctx context.Context,
	params *lsp.LogMessageParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "window/logMessage"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) TelemetryEvent(
	ctx context.Context,
	params *lsp.LSPAny,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "telemetry/event"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) TextDocumentPublishDiagnostics(
	ctx context.Context,
	params *lsp.PublishDiagnosticsParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "textDocument/publishDiagnostics"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) LspLogTrace(
	ctx context.Context,
	params *lsp.LogTraceParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "$/logTrace"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) LspCancelRequest(
	ctx context.Context,
	params *lsp.CancelParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "$/cancelRequest"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

func (p *LspClientJSONRPCProxy) LspProgress(
	ctx context.Context,
	params *lsp.ProgressParams,
) (
	err error,
) {
	var notificationMsg jsonrpc2.RequestMessage
	notificationMsg.Version = "2.0"
	notificationMsg.Method = "$/progress"

	notificationMsg.Params, err = json.Marshal(params)
	if err != nil {
		return
	}

	p.conn.RequestsWrite() <- &notificationMsg
	return
}

var _ lsp.LspClient = &LspClientJSONRPCProxy{}