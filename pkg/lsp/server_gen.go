// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
	"context"
	"github.com/black-desk/notels/pkg/jsonrpc2"
)

// jsonrpc2gen: LspServer adaptor
// jsonrpc2gen: LspServer proxy
type LspServer interface {

	// A request to resolve the implementation locations of a symbol at a
	// given text document position. The request's parameter is of type
	// [TextDocumentPositionParams] (#TextDocumentPositionParams) the
	// response is of type {@link Definition} or a Thenable that resolves to
	// such.
	TextDocumentImplementation(
		ctx context.Context,
		params *ImplementationParams,
	) (
		result *TextDocumentImplementation_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/implementation"

	// A request to resolve the type definition locations of a symbol at a
	// given text document position. The request's parameter is of type
	// [TextDocumentPositionParams] (#TextDocumentPositionParams) the
	// response is of type {@link Definition} or a Thenable that resolves to
	// such.
	TextDocumentTypeDefinition(
		ctx context.Context,
		params *TypeDefinitionParams,
	) (
		result *TextDocumentTypeDefinition_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/typeDefinition"

	// A request to list all color symbols found in a given text document.
	// The request's parameter is of type {@link DocumentColorParams} the
	// response is of type {@link ColorInformation ColorInformation[]} or a
	// Thenable that resolves to such.
	TextDocumentDocumentColor(
		ctx context.Context,
		params *DocumentColorParams,
	) (
		result []ColorInformation,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/documentColor"

	// A request to list all presentation for a color. The request's
	// parameter is of type {@link ColorPresentationParams} the response is
	// of type {@link ColorInformation ColorInformation[]} or a Thenable
	// that resolves to such.
	TextDocumentColorPresentation(
		ctx context.Context,
		params *ColorPresentationParams,
	) (
		result []ColorPresentation,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/colorPresentation"

	// A request to provide folding ranges in a document. The request's
	// parameter is of type {@link FoldingRangeParams}, the response is of
	// type {@link FoldingRangeList} or a Thenable that resolves to such.
	TextDocumentFoldingRange(
		ctx context.Context,
		params *FoldingRangeParams,
	) (
		result *TextDocumentFoldingRange_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/foldingRange"

	// A request to resolve the type definition locations of a symbol at a
	// given text document position. The request's parameter is of type
	// [TextDocumentPositionParams] (#TextDocumentPositionParams) the
	// response is of type {@link Declaration} or a typed array of {@link
	// DeclarationLink} or a Thenable that resolves to such.
	TextDocumentDeclaration(
		ctx context.Context,
		params *DeclarationParams,
	) (
		result *TextDocumentDeclaration_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/declaration"

	// A request to provide selection ranges in a document. The request's
	// parameter is of type {@link SelectionRangeParams}, the response is of
	// type {@link SelectionRange SelectionRange[]} or a Thenable that
	// resolves to such.
	TextDocumentSelectionRange(
		ctx context.Context,
		params *SelectionRangeParams,
	) (
		result *TextDocumentSelectionRange_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/selectionRange"

	// A request to result a `CallHierarchyItem` in a document at a given
	// position. Can be used as an input to an incoming or outgoing call
	// hierarchy.  @since 3.16.0
	TextDocumentPrepareCallHierarchy(
		ctx context.Context,
		params *CallHierarchyPrepareParams,
	) (
		result *TextDocumentPrepareCallHierarchy_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/prepareCallHierarchy"

	// A request to resolve the incoming calls for a given
	// `CallHierarchyItem`.  @since 3.16.0
	CallHierarchyIncomingCalls(
		ctx context.Context,
		params *CallHierarchyIncomingCallsParams,
	) (
		result *CallHierarchyIncomingCalls_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"callHierarchy/incomingCalls"

	// A request to resolve the outgoing calls for a given
	// `CallHierarchyItem`.  @since 3.16.0
	CallHierarchyOutgoingCalls(
		ctx context.Context,
		params *CallHierarchyOutgoingCallsParams,
	) (
		result *CallHierarchyOutgoingCalls_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"callHierarchy/outgoingCalls"

	// @since 3.16.0
	// registration method: textDocument/semanticTokens
	TextDocumentSemanticTokensFull(
		ctx context.Context,
		params *SemanticTokensParams,
	) (
		result *TextDocumentSemanticTokensFull_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/semanticTokens/full"

	// @since 3.16.0
	// registration method: textDocument/semanticTokens
	TextDocumentSemanticTokensFullDelta(
		ctx context.Context,
		params *SemanticTokensDeltaParams,
	) (
		result *TextDocumentSemanticTokensFullDelta_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/semanticTokens/full/delta"

	// @since 3.16.0
	// registration method: textDocument/semanticTokens
	TextDocumentSemanticTokensRange(
		ctx context.Context,
		params *SemanticTokensRangeParams,
	) (
		result *TextDocumentSemanticTokensRange_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/semanticTokens/range"

	// A request to provide ranges that can be edited together.  @since
	// 3.16.0
	TextDocumentLinkedEditingRange(
		ctx context.Context,
		params *LinkedEditingRangeParams,
	) (
		result *TextDocumentLinkedEditingRange_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/linkedEditingRange"

	// The will create files request is sent from the client to the server
	// before files are actually created as long as the creation is
	// triggered from within the client.  @since 3.16.0
	WorkspaceWillCreateFiles(
		ctx context.Context,
		params *CreateFilesParams,
	) (
		result *WorkspaceWillCreateFiles_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/willCreateFiles"

	// The will rename files request is sent from the client to the server
	// before files are actually renamed as long as the rename is triggered
	// from within the client.  @since 3.16.0
	WorkspaceWillRenameFiles(
		ctx context.Context,
		params *RenameFilesParams,
	) (
		result *WorkspaceWillRenameFiles_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/willRenameFiles"

	// The did delete files notification is sent from the client to the
	// server when files were deleted from within the client.  @since 3.16.0
	WorkspaceWillDeleteFiles(
		ctx context.Context,
		params *DeleteFilesParams,
	) (
		result *WorkspaceWillDeleteFiles_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/willDeleteFiles"

	// A request to get the moniker of a symbol at a given text document
	// position. The request parameter is of type {@link
	// TextDocumentPositionParams}. The response is of type {@link Moniker
	// Moniker[]} or `null`.
	TextDocumentMoniker(
		ctx context.Context,
		params *MonikerParams,
	) (
		result *TextDocumentMoniker_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/moniker"

	// A request to result a `TypeHierarchyItem` in a document at a given
	// position. Can be used as an input to a subtypes or supertypes type
	// hierarchy.  @since 3.17.0
	TextDocumentPrepareTypeHierarchy(
		ctx context.Context,
		params *TypeHierarchyPrepareParams,
	) (
		result *TextDocumentPrepareTypeHierarchy_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/prepareTypeHierarchy"

	// A request to resolve the supertypes for a given `TypeHierarchyItem`.
	// @since 3.17.0
	TypeHierarchySupertypes(
		ctx context.Context,
		params *TypeHierarchySupertypesParams,
	) (
		result *TypeHierarchySupertypes_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"typeHierarchy/supertypes"

	// A request to resolve the subtypes for a given `TypeHierarchyItem`.
	// @since 3.17.0
	TypeHierarchySubtypes(
		ctx context.Context,
		params *TypeHierarchySubtypesParams,
	) (
		result *TypeHierarchySubtypes_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"typeHierarchy/subtypes"

	// A request to provide inline values in a document. The request's
	// parameter is of type {@link InlineValueParams}, the response is of
	// type {@link InlineValue InlineValue[]} or a Thenable that resolves to
	// such.  @since 3.17.0
	TextDocumentInlineValue(
		ctx context.Context,
		params *InlineValueParams,
	) (
		result *TextDocumentInlineValue_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/inlineValue"

	// A request to provide inlay hints in a document. The request's
	// parameter is of type {@link InlayHintsParams}, the response is of
	// type {@link InlayHint InlayHint[]} or a Thenable that resolves to
	// such.  @since 3.17.0
	TextDocumentInlayHint(
		ctx context.Context,
		params *InlayHintParams,
	) (
		result *TextDocumentInlayHint_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/inlayHint"

	// A request to resolve additional properties for an inlay hint. The
	// request's parameter is of type {@link InlayHint}, the response is of
	// type {@link InlayHint} or a Thenable that resolves to such.  @since
	// 3.17.0
	InlayHintResolve(
		ctx context.Context,
		params *InlayHint,
	) (
		result *InlayHint,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"inlayHint/resolve"

	// The document diagnostic request definition.  @since 3.17.0
	TextDocumentDiagnostic(
		ctx context.Context,
		params *DocumentDiagnosticParams,
	) (
		result *DocumentDiagnosticReport,
		errorData *DiagnosticServerCancellationData,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/diagnostic"

	// The workspace diagnostic request definition.  @since 3.17.0
	WorkspaceDiagnostic(
		ctx context.Context,
		params *WorkspaceDiagnosticParams,
	) (
		result *WorkspaceDiagnosticReport,
		errorData *DiagnosticServerCancellationData,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/diagnostic"

	// The initialize request is sent from the client to the server. It is
	// sent once as the request after starting up the server. The requests
	// parameter is of type {@link InitializeParams} the response if of type
	// {@link InitializeResult} of a Thenable that resolves to such.
	Initialize(
		ctx context.Context,
		params *InitializeParams,
	) (
		result *InitializeResult,
		errorData *InitializeError,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"initialize"

	// A shutdown request is sent from the client to the server. It is sent
	// once when the client decides to shutdown the server. The only
	// notification that is sent after a shutdown request is the exit event.
	Shutdown(
		ctx context.Context,
	) (
		result *Null,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"shutdown"

	// A document will save request is sent from the client to the server
	// before the document is actually saved. The request can return an
	// array of TextEdits which will be applied to the text document before
	// it is saved. Please note that clients might drop results if computing
	// the text edits took too long or if a server constantly fails on this
	// request. This is done to keep the save fast and reliable.
	TextDocumentWillSaveWaitUntil(
		ctx context.Context,
		params *WillSaveTextDocumentParams,
	) (
		result *TextDocumentWillSaveWaitUntil_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/willSaveWaitUntil"

	// Request to request completion at a given text document position. The
	// request's parameter is of type {@link TextDocumentPosition} the
	// response is of type {@link CompletionItem CompletionItem[]} or {@link
	// CompletionList} or a Thenable that resolves to such.  The request can
	// delay the computation of the {@link CompletionItem.detail `detail`}
	// and {@link CompletionItem.documentation `documentation`} properties
	// to the `completionItem/resolve` request. However, properties that are
	// needed for the initial sorting and filtering, like `sortText`,
	// `filterText`, `insertText`, and `textEdit`, must not be changed
	// during resolve.
	TextDocumentCompletion(
		ctx context.Context,
		params *CompletionParams,
	) (
		result *TextDocumentCompletion_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/completion"

	// Request to resolve additional information for a given completion
	// item.The request's parameter is of type {@link CompletionItem} the
	// response is of type {@link CompletionItem} or a Thenable that
	// resolves to such.
	CompletionItemResolve(
		ctx context.Context,
		params *CompletionItem,
	) (
		result *CompletionItem,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"completionItem/resolve"

	// Request to request hover information at a given text document
	// position. The request's parameter is of type {@link
	// TextDocumentPosition} the response is of type {@link Hover} or a
	// Thenable that resolves to such.
	TextDocumentHover(
		ctx context.Context,
		params *HoverParams,
	) (
		result *TextDocumentHover_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/hover"

	TextDocumentSignatureHelp(
		ctx context.Context,
		params *SignatureHelpParams,
	) (
		result *TextDocumentSignatureHelp_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/signatureHelp"

	// A request to resolve the definition location of a symbol at a given
	// text document position. The request's parameter is of type
	// [TextDocumentPosition] (#TextDocumentPosition) the response is of
	// either type {@link Definition} or a typed array of {@link
	// DefinitionLink} or a Thenable that resolves to such.
	TextDocumentDefinition(
		ctx context.Context,
		params *DefinitionParams,
	) (
		result *TextDocumentDefinition_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/definition"

	// A request to resolve project-wide references for the symbol denoted
	// by the given text document position. The request's parameter is of
	// type {@link ReferenceParams} the response is of type {@link Location
	// Location[]} or a Thenable that resolves to such.
	TextDocumentReferences(
		ctx context.Context,
		params *ReferenceParams,
	) (
		result *TextDocumentReferences_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/references"

	// Request to resolve a {@link DocumentHighlight} for a given text
	// document position. The request's parameter is of type
	// [TextDocumentPosition] (#TextDocumentPosition) the request response
	// is of type [DocumentHighlight[]] (#DocumentHighlight) or a Thenable
	// that resolves to such.
	TextDocumentDocumentHighlight(
		ctx context.Context,
		params *DocumentHighlightParams,
	) (
		result *TextDocumentDocumentHighlight_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/documentHighlight"

	// A request to list all symbols found in a given text document. The
	// request's parameter is of type {@link TextDocumentIdentifier} the
	// response is of type {@link SymbolInformation SymbolInformation[]} or
	// a Thenable that resolves to such.
	TextDocumentDocumentSymbol(
		ctx context.Context,
		params *DocumentSymbolParams,
	) (
		result *TextDocumentDocumentSymbol_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/documentSymbol"

	// A request to provide commands for the given text document and range.
	TextDocumentCodeAction(
		ctx context.Context,
		params *CodeActionParams,
	) (
		result *TextDocumentCodeAction_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/codeAction"

	// Request to resolve additional information for a given code action.The
	// request's parameter is of type {@link CodeAction} the response is of
	// type {@link CodeAction} or a Thenable that resolves to such.
	CodeActionResolve(
		ctx context.Context,
		params *CodeAction,
	) (
		result *CodeAction,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"codeAction/resolve"

	// A request to list project-wide symbols matching the query string
	// given by the {@link WorkspaceSymbolParams}. The response is of type
	// {@link SymbolInformation SymbolInformation[]} or a Thenable that
	// resolves to such.  @since 3.17.0 - support for WorkspaceSymbol in the
	// returned data. Clients  need to advertise support for
	// WorkspaceSymbols via the client capability
	// `workspace.symbol.resolveSupport`.
	WorkspaceSymbol(
		ctx context.Context,
		params *WorkspaceSymbolParams,
	) (
		result *WorkspaceSymbol_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/symbol"

	// A request to resolve the range inside the workspace symbol's
	// location.  @since 3.17.0
	WorkspaceSymbolResolve(
		ctx context.Context,
		params *WorkspaceSymbol,
	) (
		result *WorkspaceSymbol,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspaceSymbol/resolve"

	// A request to provide code lens for the given text document.
	TextDocumentCodeLens(
		ctx context.Context,
		params *CodeLensParams,
	) (
		result *TextDocumentCodeLens_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/codeLens"

	// A request to resolve a command for a given code lens.
	CodeLensResolve(
		ctx context.Context,
		params *CodeLens,
	) (
		result *CodeLens,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"codeLens/resolve"

	// A request to provide document links
	TextDocumentDocumentLink(
		ctx context.Context,
		params *DocumentLinkParams,
	) (
		result *TextDocumentDocumentLink_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/documentLink"

	// Request to resolve additional information for a given document link.
	// The request's parameter is of type {@link DocumentLink} the response
	// is of type {@link DocumentLink} or a Thenable that resolves to such.
	DocumentLinkResolve(
		ctx context.Context,
		params *DocumentLink,
	) (
		result *DocumentLink,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"documentLink/resolve"

	// A request to to format a whole document.
	TextDocumentFormatting(
		ctx context.Context,
		params *DocumentFormattingParams,
	) (
		result *TextDocumentFormatting_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/formatting"

	// A request to to format a range in a document.
	TextDocumentRangeFormatting(
		ctx context.Context,
		params *DocumentRangeFormattingParams,
	) (
		result *TextDocumentRangeFormatting_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/rangeFormatting"

	// A request to format a document on type.
	TextDocumentOnTypeFormatting(
		ctx context.Context,
		params *DocumentOnTypeFormattingParams,
	) (
		result *TextDocumentOnTypeFormatting_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/onTypeFormatting"

	// A request to rename a symbol.
	TextDocumentRename(
		ctx context.Context,
		params *RenameParams,
	) (
		result *TextDocumentRename_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/rename"

	// A request to test and perform the setup necessary for a rename.
	// @since 3.16 - support for default behavior
	TextDocumentPrepareRename(
		ctx context.Context,
		params *PrepareRenameParams,
	) (
		result *TextDocumentPrepareRename_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"textDocument/prepareRename"

	// A request send from the client to the server to execute a command.
	// The request might return a workspace edit which the client will apply
	// to the workspace.
	WorkspaceExecuteCommand(
		ctx context.Context,
		params *ExecuteCommandParams,
	) (
		result *WorkspaceExecuteCommand_Result_Or,
		code jsonrpc2.Code,
		err error,
	) // jsonrpc2gen:request"workspace/executeCommand"

	// The `workspace/didChangeWorkspaceFolders` notification is sent from
	// the client to the server when the workspace folder configuration
	// changes.
	WorkspaceDidChangeWorkspaceFolders(
		ctx context.Context,
		params *DidChangeWorkspaceFoldersParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didChangeWorkspaceFolders"

	// The `window/workDoneProgress/cancel` notification is sent from  the
	// client to the server to cancel a progress initiated on the server
	// side.
	WindowWorkDoneProgressCancel(
		ctx context.Context,
		params *WorkDoneProgressCancelParams,
	) (
		err error,
	) // jsonrpc2gen:notification"window/workDoneProgress/cancel"

	// The did create files notification is sent from the client to the
	// server when files were created from within the client.  @since 3.16.0
	WorkspaceDidCreateFiles(
		ctx context.Context,
		params *CreateFilesParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didCreateFiles"

	// The did rename files notification is sent from the client to the
	// server when files were renamed from within the client.  @since 3.16.0
	WorkspaceDidRenameFiles(
		ctx context.Context,
		params *RenameFilesParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didRenameFiles"

	// The will delete files request is sent from the client to the server
	// before files are actually deleted as long as the deletion is
	// triggered from within the client.  @since 3.16.0
	WorkspaceDidDeleteFiles(
		ctx context.Context,
		params *DeleteFilesParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didDeleteFiles"

	// A notification sent when a notebook opens.  @since 3.17.0
	// registration method: notebookDocument/sync
	NotebookDocumentDidOpen(
		ctx context.Context,
		params *DidOpenNotebookDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"notebookDocument/didOpen"

	// registration method: notebookDocument/sync
	NotebookDocumentDidChange(
		ctx context.Context,
		params *DidChangeNotebookDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"notebookDocument/didChange"

	// A notification sent when a notebook document is saved.  @since 3.17.0
	// registration method: notebookDocument/sync
	NotebookDocumentDidSave(
		ctx context.Context,
		params *DidSaveNotebookDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"notebookDocument/didSave"

	// A notification sent when a notebook closes.  @since 3.17.0
	// registration method: notebookDocument/sync
	NotebookDocumentDidClose(
		ctx context.Context,
		params *DidCloseNotebookDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"notebookDocument/didClose"

	// The initialized notification is sent from the client to the server
	// after the client is fully initialized and the server is allowed to
	// send requests from the server to the client.
	Initialized(
		ctx context.Context,
		params *InitializedParams,
	) (
		err error,
	) // jsonrpc2gen:notification"initialized"

	// The exit event is sent from the client to the server to ask the
	// server to exit its process.
	Exit(
		ctx context.Context,
	) (
		err error,
	) // jsonrpc2gen:notification"exit"

	// The configuration change notification is sent from the client to the
	// server when the client's configuration has changed. The notification
	// contains the changed configuration as defined by the language client.
	WorkspaceDidChangeConfiguration(
		ctx context.Context,
		params *DidChangeConfigurationParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didChangeConfiguration"

	// The document open notification is sent from the client to the server
	// to signal newly opened text documents. The document's truth is now
	// managed by the client and the server must not try to read the
	// document's truth using the document's uri. Open in this sense means
	// it is managed by the client. It doesn't necessarily mean that its
	// content is presented in an editor. An open notification must not be
	// sent more than once without a corresponding close notification send
	// before. This means open and close notification must be balanced and
	// the max open count is one.
	TextDocumentDidOpen(
		ctx context.Context,
		params *DidOpenTextDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/didOpen"

	// The document change notification is sent from the client to the
	// server to signal changes to a text document.
	TextDocumentDidChange(
		ctx context.Context,
		params *DidChangeTextDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/didChange"

	// The document close notification is sent from the client to the server
	// when the document got closed in the client. The document's truth now
	// exists where the document's uri points to (e.g. if the document's uri
	// is a file uri the truth now exists on disk). As with the open
	// notification the close notification is about managing the document's
	// content. Receiving a close notification doesn't mean that the
	// document was open in an editor before. A close notification requires
	// a previous open notification to be sent.
	TextDocumentDidClose(
		ctx context.Context,
		params *DidCloseTextDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/didClose"

	// The document save notification is sent from the client to the server
	// when the document got saved in the client.
	TextDocumentDidSave(
		ctx context.Context,
		params *DidSaveTextDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/didSave"

	// A document will save notification is sent from the client to the
	// server before the document is actually saved.
	TextDocumentWillSave(
		ctx context.Context,
		params *WillSaveTextDocumentParams,
	) (
		err error,
	) // jsonrpc2gen:notification"textDocument/willSave"

	// The watched files notification is sent from the client to the server
	// when the client detects changes to file watched by the language
	// client.
	WorkspaceDidChangeWatchedFiles(
		ctx context.Context,
		params *DidChangeWatchedFilesParams,
	) (
		err error,
	) // jsonrpc2gen:notification"workspace/didChangeWatchedFiles"

	LspSetTrace(
		ctx context.Context,
		params *SetTraceParams,
	) (
		err error,
	) // jsonrpc2gen:notification"$/setTrace"

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
