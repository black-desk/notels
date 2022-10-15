// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"context"
)

//jsonrpc2gen:proxy
type LspClient interface {

	/*
	   The `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
	*/
	WorkspaceWorkspaceFolders(
		ctx context.Context,
	) (
		err error,
		result WorkspaceWorkspaceFoldersResult_,
	) //jsonrpc2gen:"workspace/workspaceFolders"

	/*
	   The 'workspace/configuration' request is sent from the server to the client to fetch a certain
	   configuration setting.

	   This pull model replaces the old push model were the client signaled configuration change via an
	   event. If the server still needs to react to configuration changes (since the server caches the
	   result of `workspace/configuration` requests) the server should register for an empty configuration
	   change event and empty the cache if such an event is received.
	*/
	WorkspaceConfiguration(
		ctx context.Context,
		params WorkspaceConfigurationParams_,
	) (
		err error,
		result WorkspaceConfigurationResult_,
	) //jsonrpc2gen:"workspace/configuration"

	/*
	   The `window/workDoneProgress/create` request is sent from the server to the client to initiate progress
	   reporting from the server.
	*/
	WindowWorkDoneProgressCreate(
		ctx context.Context,
		params WindowWorkDoneProgressCreateParams_,
	) (
		err error,
		result WindowWorkDoneProgressCreateResult_,
	) //jsonrpc2gen:"window/workDoneProgress/create"

	/*
	   A request to show a document. This request might open an
	   external program depending on the value of the URI to open.
	   For example a request to open `https://code.visualstudio.com/`
	   will very likely open the URI in a WEB browser.

	   @since 3.16.0
	*/
	WindowShowDocument(
		ctx context.Context,
		params WindowShowDocumentParams_,
	) (
		err error,
		result WindowShowDocumentResult_,
	) //jsonrpc2gen:"window/showDocument"

	/*
	   The `client/registerCapability` request is sent from the server to the client to register a new capability
	   handler on the client side.
	*/
	ClientRegisterCapability(
		ctx context.Context,
		params ClientRegisterCapabilityParams_,
	) (
		err error,
		result ClientRegisterCapabilityResult_,
	) //jsonrpc2gen:"client/registerCapability"

	/*
	   The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
	   handler on the client side.
	*/
	ClientUnregisterCapability(
		ctx context.Context,
		params ClientUnregisterCapabilityParams_,
	) (
		err error,
		result ClientUnregisterCapabilityResult_,
	) //jsonrpc2gen:"client/unregisterCapability"

	/*
	   The show message request is sent from the server to the client to show a message
	   and a set of options actions to the user.
	*/
	WindowShowMessageRequest(
		ctx context.Context,
		params WindowShowMessageRequestParams_,
	) (
		err error,
		result WindowShowMessageRequestResult_,
	) //jsonrpc2gen:"window/showMessageRequest"

	/*
	   A request to refresh all code actions

	   @since 3.16.0
	*/
	WorkspaceCodeLensRefresh(
		ctx context.Context,
	) (
		err error,
		result WorkspaceCodeLensRefreshResult_,
	) //jsonrpc2gen:"workspace/codeLens/refresh"

	/*
	   A request sent from the server to the client to modified certain resources.
	*/
	WorkspaceApplyEdit(
		ctx context.Context,
		params WorkspaceApplyEditParams_,
	) (
		err error,
		result WorkspaceApplyEditResult_,
	) //jsonrpc2gen:"workspace/applyEdit"

	/*
	   The show message notification is sent from a server to a client to ask
	   the client to display a particular message in the user interface.
	*/
	WindowShowMessage(
		ctx context.Context,
		params WindowShowMessageParams_,
	) (
		err error,
	) //jsonrpc2gen:"window/showMessage"

	/*
	   The log message notification is sent from the server to the client to ask
	   the client to log a particular message.
	*/
	WindowLogMessage(
		ctx context.Context,
		params WindowLogMessageParams_,
	) (
		err error,
	) //jsonrpc2gen:"window/logMessage"

	/*
	   The telemetry event notification is sent from the server to the client to ask
	   the client to log telemetry data.
	*/
	TelemetryEvent(
		ctx context.Context,
		params TelemetryEventParams_,
	) (
		err error,
	) //jsonrpc2gen:"telemetry/event"

	/*
	   Diagnostics notification are sent from the server to the client to signal
	   results of validation runs.
	*/
	TextDocumentPublishDiagnostics(
		ctx context.Context,
		params TextDocumentPublishDiagnosticsParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/publishDiagnostics"

	LspLogTrace(
		ctx context.Context,
		params LspLogTraceParams_,
	) (
		err error,
	) //jsonrpc2gen:"$/logTrace"

	LspCancelRequest(
		ctx context.Context,
		params LspCancelRequestParams_,
	) (
		err error,
	) //jsonrpc2gen:"$/cancelRequest"

	LspProgress(
		ctx context.Context,
		params LspProgressParams_,
	) (
		err error,
	) //jsonrpc2gen:"$/progress"

}
