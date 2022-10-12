// Code generated from metaModel.json by "lspgen". DO NOT EDIT
package protocol

import "context"

type LspServer interface {

	/*
	   A request to resolve the implementation locations of a symbol at a given text
	   document position. The request's parameter is of type [TextDocumentPositionParams]
	   (#TextDocumentPositionParams) the response is of type [Definition](#Definition) or a
	   Thenable that resolves to such.
	*/
	TextDocumentImplementation(
		ctx context.Context,
		params TextDocumentImplementationParams_,
	) (
		err error,
		result TextDocumentImplementationResult_,
		partialResult TextDocumentImplementationPartialResult_,
	) //jsonrpc2gen:"textDocument/implementation"

	/*
	   A request to resolve the type definition locations of a symbol at a given text
	   document position. The request's parameter is of type [TextDocumentPositionParams]
	   (#TextDocumentPositionParams) the response is of type [Definition](#Definition) or a
	   Thenable that resolves to such.
	*/
	TextDocumentTypeDefinition(
		ctx context.Context,
		params TextDocumentTypeDefinitionParams_,
	) (
		err error,
		result TextDocumentTypeDefinitionResult_,
		partialResult TextDocumentTypeDefinitionPartialResult_,
	) //jsonrpc2gen:"textDocument/typeDefinition"

	/*
	   A request to list all color symbols found in a given text document. The request's
	   parameter is of type [DocumentColorParams](#DocumentColorParams) the
	   response is of type [ColorInformation[]](#ColorInformation) or a Thenable
	   that resolves to such.
	*/
	TextDocumentDocumentColor(
		ctx context.Context,
		params TextDocumentDocumentColorParams_,
	) (
		err error,
		result TextDocumentDocumentColorResult_,
		partialResult TextDocumentDocumentColorPartialResult_,
	) //jsonrpc2gen:"textDocument/documentColor"

	/*
	   A request to list all presentation for a color. The request's
	   parameter is of type [ColorPresentationParams](#ColorPresentationParams) the
	   response is of type [ColorInformation[]](#ColorInformation) or a Thenable
	   that resolves to such.
	*/
	TextDocumentColorPresentation(
		ctx context.Context,
		params TextDocumentColorPresentationParams_,
	) (
		err error,
		result TextDocumentColorPresentationResult_,
		partialResult TextDocumentColorPresentationPartialResult_,
	) //jsonrpc2gen:"textDocument/colorPresentation"

	/*
	   A request to provide folding ranges in a document. The request's
	   parameter is of type [FoldingRangeParams](#FoldingRangeParams), the
	   response is of type [FoldingRangeList](#FoldingRangeList) or a Thenable
	   that resolves to such.
	*/
	TextDocumentFoldingRange(
		ctx context.Context,
		params TextDocumentFoldingRangeParams_,
	) (
		err error,
		result TextDocumentFoldingRangeResult_,
		partialResult TextDocumentFoldingRangePartialResult_,
	) //jsonrpc2gen:"textDocument/foldingRange"

	/*
	   A request to resolve the type definition locations of a symbol at a given text
	   document position. The request's parameter is of type [TextDocumentPositionParams]
	   (#TextDocumentPositionParams) the response is of type [Declaration](#Declaration)
	   or a typed array of [DeclarationLink](#DeclarationLink) or a Thenable that resolves
	   to such.
	*/
	TextDocumentDeclaration(
		ctx context.Context,
		params TextDocumentDeclarationParams_,
	) (
		err error,
		result TextDocumentDeclarationResult_,
		partialResult TextDocumentDeclarationPartialResult_,
	) //jsonrpc2gen:"textDocument/declaration"

	/*
	   A request to provide selection ranges in a document. The request's
	   parameter is of type [SelectionRangeParams](#SelectionRangeParams), the
	   response is of type [SelectionRange[]](#SelectionRange[]) or a Thenable
	   that resolves to such.
	*/
	TextDocumentSelectionRange(
		ctx context.Context,
		params TextDocumentSelectionRangeParams_,
	) (
		err error,
		result TextDocumentSelectionRangeResult_,
		partialResult TextDocumentSelectionRangePartialResult_,
	) //jsonrpc2gen:"textDocument/selectionRange"

	/*
	   A request to result a `CallHierarchyItem` in a document at a given position.
	   Can be used as an input to an incoming or outgoing call hierarchy.

	   @since 3.16.0
	*/
	TextDocumentPrepareCallHierarchy(
		ctx context.Context,
		params TextDocumentPrepareCallHierarchyParams_,
	) (
		err error,
		result TextDocumentPrepareCallHierarchyResult_,
	) //jsonrpc2gen:"textDocument/prepareCallHierarchy"

	/*
	   A request to resolve the incoming calls for a given `CallHierarchyItem`.

	   @since 3.16.0
	*/
	CallHierarchyIncomingCalls(
		ctx context.Context,
		params CallHierarchyIncomingCallsParams_,
	) (
		err error,
		result CallHierarchyIncomingCallsResult_,
		partialResult CallHierarchyIncomingCallsPartialResult_,
	) //jsonrpc2gen:"callHierarchy/incomingCalls"

	/*
	   A request to resolve the outgoing calls for a given `CallHierarchyItem`.

	   @since 3.16.0
	*/
	CallHierarchyOutgoingCalls(
		ctx context.Context,
		params CallHierarchyOutgoingCallsParams_,
	) (
		err error,
		result CallHierarchyOutgoingCallsResult_,
		partialResult CallHierarchyOutgoingCallsPartialResult_,
	) //jsonrpc2gen:"callHierarchy/outgoingCalls"

	/*@since 3.16.0*/
	TextDocumentSemanticTokensFull(
		ctx context.Context,
		params TextDocumentSemanticTokensFullParams_,
	) (
		err error,
		result TextDocumentSemanticTokensFullResult_,
		partialResult TextDocumentSemanticTokensFullPartialResult_,
	) //jsonrpc2gen:"textDocument/semanticTokens/full"
	TextDocumentSemanticTokens(
		ctx context.Context,
		opt TextDocumentSemanticTokensOptions_,
	) (err error) //jsonrpc2gen:"textDocument/semanticTokens"

	/*@since 3.16.0*/
	TextDocumentSemanticTokensFullDelta(
		ctx context.Context,
		params TextDocumentSemanticTokensFullDeltaParams_,
	) (
		err error,
		result TextDocumentSemanticTokensFullDeltaResult_,
		partialResult TextDocumentSemanticTokensFullDeltaPartialResult_,
	) //jsonrpc2gen:"textDocument/semanticTokens/full/delta"

	/*@since 3.16.0*/
	TextDocumentSemanticTokensRange(
		ctx context.Context,
		params TextDocumentSemanticTokensRangeParams_,
	) (
		err error,
		result TextDocumentSemanticTokensRangeResult_,
		partialResult TextDocumentSemanticTokensRangePartialResult_,
	) //jsonrpc2gen:"textDocument/semanticTokens/range"

	/*@since 3.16.0*/
	WorkspaceSemanticTokensRefresh(
		ctx context.Context,
	) (
		err error,
		result WorkspaceSemanticTokensRefreshResult_,
	) //jsonrpc2gen:"workspace/semanticTokens/refresh"

	/*
	   A request to provide ranges that can be edited together.

	   @since 3.16.0
	*/
	TextDocumentLinkedEditingRange(
		ctx context.Context,
		params TextDocumentLinkedEditingRangeParams_,
	) (
		err error,
		result TextDocumentLinkedEditingRangeResult_,
	) //jsonrpc2gen:"textDocument/linkedEditingRange"

	/*
	   The will create files request is sent from the client to the server before files are actually
	   created as long as the creation is triggered from within the client.

	   @since 3.16.0
	*/
	WorkspaceWillCreateFiles(
		ctx context.Context,
		params WorkspaceWillCreateFilesParams_,
	) (
		err error,
		result WorkspaceWillCreateFilesResult_,
	) //jsonrpc2gen:"workspace/willCreateFiles"

	/*
	   The will rename files request is sent from the client to the server before files are actually
	   renamed as long as the rename is triggered from within the client.

	   @since 3.16.0
	*/
	WorkspaceWillRenameFiles(
		ctx context.Context,
		params WorkspaceWillRenameFilesParams_,
	) (
		err error,
		result WorkspaceWillRenameFilesResult_,
	) //jsonrpc2gen:"workspace/willRenameFiles"

	/*
	   The did delete files notification is sent from the client to the server when
	   files were deleted from within the client.

	   @since 3.16.0
	*/
	WorkspaceWillDeleteFiles(
		ctx context.Context,
		params WorkspaceWillDeleteFilesParams_,
	) (
		err error,
		result WorkspaceWillDeleteFilesResult_,
	) //jsonrpc2gen:"workspace/willDeleteFiles"

	/*
	   A request to get the moniker of a symbol at a given text document position.
	   The request parameter is of type [TextDocumentPositionParams](#TextDocumentPositionParams).
	   The response is of type [Moniker[]](#Moniker[]) or `null`.
	*/
	TextDocumentMoniker(
		ctx context.Context,
		params TextDocumentMonikerParams_,
	) (
		err error,
		result TextDocumentMonikerResult_,
		partialResult TextDocumentMonikerPartialResult_,
	) //jsonrpc2gen:"textDocument/moniker"

	/*
	   A request to result a `TypeHierarchyItem` in a document at a given position.
	   Can be used as an input to a subtypes or supertypes type hierarchy.

	   @since 3.17.0
	*/
	TextDocumentPrepareTypeHierarchy(
		ctx context.Context,
		params TextDocumentPrepareTypeHierarchyParams_,
	) (
		err error,
		result TextDocumentPrepareTypeHierarchyResult_,
	) //jsonrpc2gen:"textDocument/prepareTypeHierarchy"

	/*
	   A request to resolve the supertypes for a given `TypeHierarchyItem`.

	   @since 3.17.0
	*/
	TypeHierarchySupertypes(
		ctx context.Context,
		params TypeHierarchySupertypesParams_,
	) (
		err error,
		result TypeHierarchySupertypesResult_,
		partialResult TypeHierarchySupertypesPartialResult_,
	) //jsonrpc2gen:"typeHierarchy/supertypes"

	/*
	   A request to resolve the subtypes for a given `TypeHierarchyItem`.

	   @since 3.17.0
	*/
	TypeHierarchySubtypes(
		ctx context.Context,
		params TypeHierarchySubtypesParams_,
	) (
		err error,
		result TypeHierarchySubtypesResult_,
		partialResult TypeHierarchySubtypesPartialResult_,
	) //jsonrpc2gen:"typeHierarchy/subtypes"

	/*
	   A request to provide inline values in a document. The request's parameter is of
	   type [InlineValueParams](#InlineValueParams), the response is of type
	   [InlineValue[]](#InlineValue[]) or a Thenable that resolves to such.

	   @since 3.17.0
	*/
	TextDocumentInlineValue(
		ctx context.Context,
		params TextDocumentInlineValueParams_,
	) (
		err error,
		result TextDocumentInlineValueResult_,
		partialResult TextDocumentInlineValuePartialResult_,
	) //jsonrpc2gen:"textDocument/inlineValue"

	/*@since 3.17.0*/
	WorkspaceInlineValueRefresh(
		ctx context.Context,
	) (
		err error,
		result WorkspaceInlineValueRefreshResult_,
	) //jsonrpc2gen:"workspace/inlineValue/refresh"

	/*
	   A request to provide inlay hints in a document. The request's parameter is of
	   type [InlayHintsParams](#InlayHintsParams), the response is of type
	   [InlayHint[]](#InlayHint[]) or a Thenable that resolves to such.

	   @since 3.17.0
	*/
	TextDocumentInlayHint(
		ctx context.Context,
		params TextDocumentInlayHintParams_,
	) (
		err error,
		result TextDocumentInlayHintResult_,
		partialResult TextDocumentInlayHintPartialResult_,
	) //jsonrpc2gen:"textDocument/inlayHint"

	/*
	   A request to resolve additional properties for an inlay hint.
	   The request's parameter is of type [InlayHint](#InlayHint), the response is
	   of type [InlayHint](#InlayHint) or a Thenable that resolves to such.

	   @since 3.17.0
	*/
	InlayHintResolve(
		ctx context.Context,
		params InlayHintResolveParams_,
	) (
		err error,
		result InlayHintResolveResult_,
	) //jsonrpc2gen:"inlayHint/resolve"

	/*@since 3.17.0*/
	WorkspaceInlayHintRefresh(
		ctx context.Context,
	) (
		err error,
		result WorkspaceInlayHintRefreshResult_,
	) //jsonrpc2gen:"workspace/inlayHint/refresh"

	/*
	   The document diagnostic request definition.

	   @since 3.17.0
	*/
	TextDocumentDiagnostic(
		ctx context.Context,
		params TextDocumentDiagnosticParams_,
	) (
		err error,
		errorData TextDocumentDiagnosticErrorData_,
		result TextDocumentDiagnosticResult_,
		partialResult TextDocumentDiagnosticPartialResult_,
	) //jsonrpc2gen:"textDocument/diagnostic"

	/*
	   The workspace diagnostic request definition.

	   @since 3.17.0
	*/
	WorkspaceDiagnostic(
		ctx context.Context,
		params WorkspaceDiagnosticParams_,
	) (
		err error,
		errorData WorkspaceDiagnosticErrorData_,
		result WorkspaceDiagnosticResult_,
		partialResult WorkspaceDiagnosticPartialResult_,
	) //jsonrpc2gen:"workspace/diagnostic"

	/*
	   The diagnostic refresh request definition.

	   @since 3.17.0
	*/
	WorkspaceDiagnosticRefresh(
		ctx context.Context,
	) (
		err error,
		result WorkspaceDiagnosticRefreshResult_,
	) //jsonrpc2gen:"workspace/diagnostic/refresh"

	/*
	   The initialize request is sent from the client to the server.
	   It is sent once as the request after starting up the server.
	   The requests parameter is of type [InitializeParams](#InitializeParams)
	   the response if of type [InitializeResult](#InitializeResult) of a Thenable that
	   resolves to such.
	*/
	Initialize(
		ctx context.Context,
		params InitializeParams_,
	) (
		err error,
		errorData InitializeErrorData_,
		result InitializeResult_,
	) //jsonrpc2gen:"initialize"

	/*
	   A shutdown request is sent from the client to the server.
	   It is sent once when the client decides to shutdown the
	   server. The only notification that is sent after a shutdown request
	   is the exit event.
	*/
	Shutdown(
		ctx context.Context,
	) (
		err error,
		result ShutdownResult_,
	) //jsonrpc2gen:"shutdown"

	/*
	   A document will save request is sent from the client to the server before
	   the document is actually saved. The request can return an array of TextEdits
	   which will be applied to the text document before it is saved. Please note that
	   clients might drop results if computing the text edits took too long or if a
	   server constantly fails on this request. This is done to keep the save fast and
	   reliable.
	*/
	TextDocumentWillSaveWaitUntil(
		ctx context.Context,
		params TextDocumentWillSaveWaitUntilParams_,
	) (
		err error,
		result TextDocumentWillSaveWaitUntilResult_,
	) //jsonrpc2gen:"textDocument/willSaveWaitUntil"

	/*
	   Request to request completion at a given text document position. The request's
	   parameter is of type [TextDocumentPosition](#TextDocumentPosition) the response
	   is of type [CompletionItem[]](#CompletionItem) or [CompletionList](#CompletionList)
	   or a Thenable that resolves to such.

	   The request can delay the computation of the [`detail`](#CompletionItem.detail)
	   and [`documentation`](#CompletionItem.documentation) properties to the `completionItem/resolve`
	   request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
	   `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
	*/
	TextDocumentCompletion(
		ctx context.Context,
		params TextDocumentCompletionParams_,
	) (
		err error,
		result TextDocumentCompletionResult_,
		partialResult TextDocumentCompletionPartialResult_,
	) //jsonrpc2gen:"textDocument/completion"

	/*
	   Request to resolve additional information for a given completion item.The request's
	   parameter is of type [CompletionItem](#CompletionItem) the response
	   is of type [CompletionItem](#CompletionItem) or a Thenable that resolves to such.
	*/
	CompletionItemResolve(
		ctx context.Context,
		params CompletionItemResolveParams_,
	) (
		err error,
		result CompletionItemResolveResult_,
	) //jsonrpc2gen:"completionItem/resolve"

	/*
	   Request to request hover information at a given text document position. The request's
	   parameter is of type [TextDocumentPosition](#TextDocumentPosition) the response is of
	   type [Hover](#Hover) or a Thenable that resolves to such.
	*/
	TextDocumentHover(
		ctx context.Context,
		params TextDocumentHoverParams_,
	) (
		err error,
		result TextDocumentHoverResult_,
	) //jsonrpc2gen:"textDocument/hover"

	TextDocumentSignatureHelp(
		ctx context.Context,
		params TextDocumentSignatureHelpParams_,
	) (
		err error,
		result TextDocumentSignatureHelpResult_,
	) //jsonrpc2gen:"textDocument/signatureHelp"

	/*
	   A request to resolve the definition location of a symbol at a given text
	   document position. The request's parameter is of type [TextDocumentPosition]
	   (#TextDocumentPosition) the response is of either type [Definition](#Definition)
	   or a typed array of [DefinitionLink](#DefinitionLink) or a Thenable that resolves
	   to such.
	*/
	TextDocumentDefinition(
		ctx context.Context,
		params TextDocumentDefinitionParams_,
	) (
		err error,
		result TextDocumentDefinitionResult_,
		partialResult TextDocumentDefinitionPartialResult_,
	) //jsonrpc2gen:"textDocument/definition"

	/*
	   A request to resolve project-wide references for the symbol denoted
	   by the given text document position. The request's parameter is of
	   type [ReferenceParams](#ReferenceParams) the response is of type
	   [Location[]](#Location) or a Thenable that resolves to such.
	*/
	TextDocumentReferences(
		ctx context.Context,
		params TextDocumentReferencesParams_,
	) (
		err error,
		result TextDocumentReferencesResult_,
		partialResult TextDocumentReferencesPartialResult_,
	) //jsonrpc2gen:"textDocument/references"

	/*
	   Request to resolve a [DocumentHighlight](#DocumentHighlight) for a given
	   text document position. The request's parameter is of type [TextDocumentPosition]
	   (#TextDocumentPosition) the request response is of type [DocumentHighlight[]]
	   (#DocumentHighlight) or a Thenable that resolves to such.
	*/
	TextDocumentDocumentHighlight(
		ctx context.Context,
		params TextDocumentDocumentHighlightParams_,
	) (
		err error,
		result TextDocumentDocumentHighlightResult_,
		partialResult TextDocumentDocumentHighlightPartialResult_,
	) //jsonrpc2gen:"textDocument/documentHighlight"

	/*
	   A request to list all symbols found in a given text document. The request's
	   parameter is of type [TextDocumentIdentifier](#TextDocumentIdentifier) the
	   response is of type [SymbolInformation[]](#SymbolInformation) or a Thenable
	   that resolves to such.
	*/
	TextDocumentDocumentSymbol(
		ctx context.Context,
		params TextDocumentDocumentSymbolParams_,
	) (
		err error,
		result TextDocumentDocumentSymbolResult_,
		partialResult TextDocumentDocumentSymbolPartialResult_,
	) //jsonrpc2gen:"textDocument/documentSymbol"

	/*A request to provide commands for the given text document and range.*/
	TextDocumentCodeAction(
		ctx context.Context,
		params TextDocumentCodeActionParams_,
	) (
		err error,
		result TextDocumentCodeActionResult_,
		partialResult TextDocumentCodeActionPartialResult_,
	) //jsonrpc2gen:"textDocument/codeAction"

	/*
	   Request to resolve additional information for a given code action.The request's
	   parameter is of type [CodeAction](#CodeAction) the response
	   is of type [CodeAction](#CodeAction) or a Thenable that resolves to such.
	*/
	CodeActionResolve(
		ctx context.Context,
		params CodeActionResolveParams_,
	) (
		err error,
		result CodeActionResolveResult_,
	) //jsonrpc2gen:"codeAction/resolve"

	/*
	   A request to list project-wide symbols matching the query string given
	   by the [WorkspaceSymbolParams](#WorkspaceSymbolParams). The response is
	   of type [SymbolInformation[]](#SymbolInformation) or a Thenable that
	   resolves to such.

	   @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients

	   	need to advertise support for WorkspaceSymbols via the client capability
	   	`workspace.symbol.resolveSupport`.
	*/
	WorkspaceSymbol(
		ctx context.Context,
		params WorkspaceSymbolParams_,
	) (
		err error,
		result WorkspaceSymbolResult_,
		partialResult WorkspaceSymbolPartialResult_,
	) //jsonrpc2gen:"workspace/symbol"

	/*
	   A request to resolve the range inside the workspace
	   symbol's location.

	   @since 3.17.0
	*/
	WorkspaceSymbolResolve(
		ctx context.Context,
		params WorkspaceSymbolResolveParams_,
	) (
		err error,
		result WorkspaceSymbolResolveResult_,
	) //jsonrpc2gen:"workspaceSymbol/resolve"

	/*A request to provide code lens for the given text document.*/
	TextDocumentCodeLens(
		ctx context.Context,
		params TextDocumentCodeLensParams_,
	) (
		err error,
		result TextDocumentCodeLensResult_,
		partialResult TextDocumentCodeLensPartialResult_,
	) //jsonrpc2gen:"textDocument/codeLens"

	/*A request to resolve a command for a given code lens.*/
	CodeLensResolve(
		ctx context.Context,
		params CodeLensResolveParams_,
	) (
		err error,
		result CodeLensResolveResult_,
	) //jsonrpc2gen:"codeLens/resolve"

	/*A request to provide document links*/
	TextDocumentDocumentLink(
		ctx context.Context,
		params TextDocumentDocumentLinkParams_,
	) (
		err error,
		result TextDocumentDocumentLinkResult_,
		partialResult TextDocumentDocumentLinkPartialResult_,
	) //jsonrpc2gen:"textDocument/documentLink"

	/*
	   Request to resolve additional information for a given document link. The request's
	   parameter is of type [DocumentLink](#DocumentLink) the response
	   is of type [DocumentLink](#DocumentLink) or a Thenable that resolves to such.
	*/
	DocumentLinkResolve(
		ctx context.Context,
		params DocumentLinkResolveParams_,
	) (
		err error,
		result DocumentLinkResolveResult_,
	) //jsonrpc2gen:"documentLink/resolve"

	/*A request to to format a whole document.*/
	TextDocumentFormatting(
		ctx context.Context,
		params TextDocumentFormattingParams_,
	) (
		err error,
		result TextDocumentFormattingResult_,
	) //jsonrpc2gen:"textDocument/formatting"

	/*A request to to format a range in a document.*/
	TextDocumentRangeFormatting(
		ctx context.Context,
		params TextDocumentRangeFormattingParams_,
	) (
		err error,
		result TextDocumentRangeFormattingResult_,
	) //jsonrpc2gen:"textDocument/rangeFormatting"

	/*A request to format a document on type.*/
	TextDocumentOnTypeFormatting(
		ctx context.Context,
		params TextDocumentOnTypeFormattingParams_,
	) (
		err error,
		result TextDocumentOnTypeFormattingResult_,
	) //jsonrpc2gen:"textDocument/onTypeFormatting"

	/*A request to rename a symbol.*/
	TextDocumentRename(
		ctx context.Context,
		params TextDocumentRenameParams_,
	) (
		err error,
		result TextDocumentRenameResult_,
	) //jsonrpc2gen:"textDocument/rename"

	/*
	   A request to test and perform the setup necessary for a rename.

	   @since 3.16 - support for default behavior
	*/
	TextDocumentPrepareRename(
		ctx context.Context,
		params TextDocumentPrepareRenameParams_,
	) (
		err error,
		result TextDocumentPrepareRenameResult_,
	) //jsonrpc2gen:"textDocument/prepareRename"

	/*
	   A request send from the client to the server to execute a command. The request might return
	   a workspace edit which the client will apply to the workspace.
	*/
	WorkspaceExecuteCommand(
		ctx context.Context,
		params WorkspaceExecuteCommandParams_,
	) (
		err error,
		result WorkspaceExecuteCommandResult_,
	) //jsonrpc2gen:"workspace/executeCommand"

	/*
	   The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
	   folder configuration changes.
	*/
	WorkspaceDidChangeWorkspaceFolders(
		ctx context.Context,
		params WorkspaceDidChangeWorkspaceFoldersParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didChangeWorkspaceFolders"

	/*
	   The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
	   initiated on the server side.
	*/
	WindowWorkDoneProgressCancel(
		ctx context.Context,
		params WindowWorkDoneProgressCancelParams_,
	) (
		err error,
	) //jsonrpc2gen:"window/workDoneProgress/cancel"

	/*
	   The did create files notification is sent from the client to the server when
	   files were created from within the client.

	   @since 3.16.0
	*/
	WorkspaceDidCreateFiles(
		ctx context.Context,
		params WorkspaceDidCreateFilesParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didCreateFiles"

	/*
	   The did rename files notification is sent from the client to the server when
	   files were renamed from within the client.

	   @since 3.16.0
	*/
	WorkspaceDidRenameFiles(
		ctx context.Context,
		params WorkspaceDidRenameFilesParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didRenameFiles"

	/*
	   The will delete files request is sent from the client to the server before files are actually
	   deleted as long as the deletion is triggered from within the client.

	   @since 3.16.0
	*/
	WorkspaceDidDeleteFiles(
		ctx context.Context,
		params WorkspaceDidDeleteFilesParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didDeleteFiles"

	/*
	   A notification sent when a notebook opens.

	   @since 3.17.0
	*/
	NotebookDocumentDidOpen(
		ctx context.Context,
		params NotebookDocumentDidOpenParams_,
	) (
		err error,
	) //jsonrpc2gen:"notebookDocument/didOpen"
	NotebookDocumentSync(
		ctx context.Context,
	) (err error) //jsonrpc2gen:"notebookDocument/sync"

	NotebookDocumentDidChange(
		ctx context.Context,
		params NotebookDocumentDidChangeParams_,
	) (
		err error,
	) //jsonrpc2gen:"notebookDocument/didChange"

	/*
	   A notification sent when a notebook document is saved.

	   @since 3.17.0
	*/
	NotebookDocumentDidSave(
		ctx context.Context,
		params NotebookDocumentDidSaveParams_,
	) (
		err error,
	) //jsonrpc2gen:"notebookDocument/didSave"

	/*
	   A notification sent when a notebook closes.

	   @since 3.17.0
	*/
	NotebookDocumentDidClose(
		ctx context.Context,
		params NotebookDocumentDidCloseParams_,
	) (
		err error,
	) //jsonrpc2gen:"notebookDocument/didClose"

	/*
	   The initialized notification is sent from the client to the
	   server after the client is fully initialized and the server
	   is allowed to send requests from the server to the client.
	*/
	Initialized(
		ctx context.Context,
		params InitializedParams_,
	) (
		err error,
	) //jsonrpc2gen:"initialized"

	/*
	   The exit event is sent from the client to the server to
	   ask the server to exit its process.
	*/
	Exit(
		ctx context.Context,
	) (
		err error,
	) //jsonrpc2gen:"exit"

	/*
	   The configuration change notification is sent from the client to the server
	   when the client's configuration has changed. The notification contains
	   the changed configuration as defined by the language client.
	*/
	WorkspaceDidChangeConfiguration(
		ctx context.Context,
		params WorkspaceDidChangeConfigurationParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didChangeConfiguration"

	/*
	   The document open notification is sent from the client to the server to signal
	   newly opened text documents. The document's truth is now managed by the client
	   and the server must not try to read the document's truth using the document's
	   uri. Open in this sense means it is managed by the client. It doesn't necessarily
	   mean that its content is presented in an editor. An open notification must not
	   be sent more than once without a corresponding close notification send before.
	   This means open and close notification must be balanced and the max open count
	   is one.
	*/
	TextDocumentDidOpen(
		ctx context.Context,
		params TextDocumentDidOpenParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/didOpen"

	/*
	   The document change notification is sent from the client to the server to signal
	   changes to a text document.
	*/
	TextDocumentDidChange(
		ctx context.Context,
		params TextDocumentDidChangeParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/didChange"

	/*
	   The document close notification is sent from the client to the server when
	   the document got closed in the client. The document's truth now exists where
	   the document's uri points to (e.g. if the document's uri is a file uri the
	   truth now exists on disk). As with the open notification the close notification
	   is about managing the document's content. Receiving a close notification
	   doesn't mean that the document was open in an editor before. A close
	   notification requires a previous open notification to be sent.
	*/
	TextDocumentDidClose(
		ctx context.Context,
		params TextDocumentDidCloseParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/didClose"

	/*
	   The document save notification is sent from the client to the server when
	   the document got saved in the client.
	*/
	TextDocumentDidSave(
		ctx context.Context,
		params TextDocumentDidSaveParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/didSave"

	/*
	   A document will save notification is sent from the client to the server before
	   the document is actually saved.
	*/
	TextDocumentWillSave(
		ctx context.Context,
		params TextDocumentWillSaveParams_,
	) (
		err error,
	) //jsonrpc2gen:"textDocument/willSave"

	/*
	   The watched files notification is sent from the client to the server when
	   the client detects changes to file watched by the language client.
	*/
	WorkspaceDidChangeWatchedFiles(
		ctx context.Context,
		params WorkspaceDidChangeWatchedFilesParams_,
	) (
		err error,
	) //jsonrpc2gen:"workspace/didChangeWatchedFiles"

	LspSetTrace(
		ctx context.Context,
		params LspSetTraceParams_,
	) (
		err error,
	) //jsonrpc2gen:"$/setTrace"

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
