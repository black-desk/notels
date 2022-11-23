// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
	"context"
	"github.com/black-desk/notels/pkg/jsonrpc2"
)

// jsonrpc2gen: LspClient adaptor
// jsonrpc2gen: LspClient proxy
type LspClient interface {

	// The `workspace/workspaceFolders` is sent from the server to the
	// client to fetch the open workspace folders.
	WorkspaceWorkspaceFolders(
		ctx context.Context,
	) (
		result *WorkspaceWorkspaceFolders_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/workspaceFolders"

	// The 'workspace/configuration' request is sent from the server to the
	// client to fetch a certain configuration setting.  This pull model
	// replaces the old push model were the client signaled configuration
	// change via an event. If the server still needs to react to
	// configuration changes (since the server caches the result of
	// `workspace/configuration` requests) the server should register for an
	// empty configuration change event and empty the cache if such an event
	// is received.
	WorkspaceConfiguration(
		ctx context.Context,
		params *ConfigurationParams,
	) (
		result []LSPAny,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/configuration"

	// The `window/workDoneProgress/create` request is sent from the server
	// to the client to initiate progress reporting from the server.
	WindowWorkDoneProgressCreate(
		ctx context.Context,
		params *WorkDoneProgressCreateParams,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"window/workDoneProgress/create"

	// @since 3.16.0
	WorkspaceSemanticTokensRefresh(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/semanticTokens/refresh"

	// A request to show a document. This request might open an external
	// program depending on the value of the URI to open. For example a
	// request to open `https://code.visualstudio.com/` will very likely
	// open the URI in a WEB browser.  @since 3.16.0
	WindowShowDocument(
		ctx context.Context,
		params *ShowDocumentParams,
	) (
		result *ShowDocumentResult,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"window/showDocument"

	// @since 3.17.0
	WorkspaceInlineValueRefresh(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/inlineValue/refresh"

	// @since 3.17.0
	WorkspaceInlayHintRefresh(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/inlayHint/refresh"

	// The diagnostic refresh request definition.  @since 3.17.0
	WorkspaceDiagnosticRefresh(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/diagnostic/refresh"

	// The `client/registerCapability` request is sent from the server to
	// the client to register a new capability handler on the client side.
	ClientRegisterCapability(
		ctx context.Context,
		params *RegistrationParams,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"client/registerCapability"

	// The `client/unregisterCapability` request is sent from the server to
	// the client to unregister a previously registered capability handler
	// on the client side.
	ClientUnregisterCapability(
		ctx context.Context,
		params *UnregistrationParams,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"client/unregisterCapability"

	// The show message request is sent from the server to the client to
	// show a message and a set of options actions to the user.
	WindowShowMessageRequest(
		ctx context.Context,
		params *ShowMessageRequestParams,
	) (
		result *WindowShowMessageRequest_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"window/showMessageRequest"

	// A request to refresh all code actions  @since 3.16.0
	WorkspaceCodeLensRefresh(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/codeLens/refresh"

	// A request sent from the server to the client to modified certain
	// resources.
	WorkspaceApplyEdit(
		ctx context.Context,
		params *ApplyWorkspaceEditParams,
	) (
		result *ApplyWorkspaceEditResult,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/applyEdit"

	// The show message notification is sent from a server to a client to
	// ask the client to display a particular message in the user interface.
	WindowShowMessage(
		ctx context.Context,
		params *ShowMessageParams,
	) (
		err error,
	) // jsonrpc2gen:notification"window/showMessage"

	// The log message notification is sent from the server to the client to
	// ask the client to log a particular message.
	WindowLogMessage(
		ctx context.Context,
		params *LogMessageParams,
	) (
		err error,
	) // jsonrpc2gen:notification"window/logMessage"

	// The telemetry event notification is sent from the server to the
	// client to ask the client to log telemetry data.
	TelemetryEvent(
		ctx context.Context,
		params *LSPAny,
	) (
		err error,
	) // jsonrpc2gen:notification"telemetry/event"

	// Diagnostics notification are sent from the server to the client to
	// signal results of validation runs.
	TextDocumentPublishDiagnostics(
		ctx context.Context,
		params *PublishDiagnosticsParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/publishDiagnostics"

	LspLogTrace(
		ctx context.Context,
		params *LogTraceParams,
	) (
		err error,
	) // jsonrpc2gen:notification"$/logTrace"

	LspCancelRequest(
		ctx context.Context,
		params *CancelParams,
	) (
		err error,
	) // jsonrpc2gen:notification"$/cancelRequest"

	LspProgress(
		ctx context.Context,
		params *ProgressParams,
	) (
		err error,
	) // jsonrpc2gen:notification"$/progress"

}
