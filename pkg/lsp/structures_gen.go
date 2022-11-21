// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"encoding/json"
	"errors"
)

type ImplementationParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// Represents a location inside a resource, such as a line inside a text file.
type Location struct {
	Uri DocumentUri `json:"uri"`

	Range Range `json:"range"`
}

type ImplementationRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	ImplementationOptions

	// mixins

	StaticRegistrationOptions
}

type TypeDefinitionParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

type TypeDefinitionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	TypeDefinitionOptions

	// mixins

	StaticRegistrationOptions
}

// A workspace folder inside a client.
type WorkspaceFolder struct {

	// The associated URI for this workspace folder.
	Uri URI `json:"uri"`

	// The name of the workspace folder. Used to refer to this workspace
	// folder in the user interface.
	Name string `json:"name"`
}

// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {

	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

// The parameters of a configuration request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

type PartialResultParams struct {

	// An optional token that a server can use to report partial results
	// (e.g. streaming) to the client.
	PartialResultToken *ProgressToken `json:"partialResultToken"`
}

// Parameters for a [DocumentColorRequest](#DocumentColorRequest).
type DocumentColorParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// Represents a color range from a document.
type ColorInformation struct {

	// The range in the document where this color appears.
	Range Range `json:"range"`

	// The actual color value for this color range.
	Color Color `json:"color"`
}

type DocumentColorRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentColorOptions

	// mixins

	StaticRegistrationOptions
}

// Parameters for a [ColorPresentationRequest](#ColorPresentationRequest).
type ColorPresentationParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The color to request presentations for.
	Color Color `json:"color"`

	// The range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
}

type ColorPresentation struct {

	// The label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when
	// selecting this color presentation.
	Label string `json:"label"`

	// An [edit](#TextEdit) which is applied to a document when selecting
	// this presentation for the color.  When `falsy` the
	// [label](#ColorPresentation.label) is used.
	TextEdit *TextEdit `json:"textEdit"`

	// An optional array of additional [text edits](#TextEdit) that are
	// applied when selecting this color presentation. Edits must not
	// overlap with the main [edit](#ColorPresentation.textEdit) nor with
	// themselves.
	AdditionalTextEdits *[]TextEdit `json:"additionalTextEdits"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress"`
}

// General text document registration options.
type TextDocumentRegistrationOptions struct {

	// A document selector to identify the scope of the registration. If set
	// to null the document selector provided on the client side will be
	// used.
	DocumentSelector TextDocumentRegistrationOptions_DocumentSelector__Or `json:"documentSelector"`
}

// Parameters for a [FoldingRangeRequest](#FoldingRangeRequest).
type FoldingRangeParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// Represents a folding range. To be valid, start and end line must be bigger
// than zero and smaller than the number of lines in the document. Clients are
// free to ignore invalid ranges.
type FoldingRange struct {

	// The zero-based start line of the range to fold. The folded area
	// starts after the line's last character. To be valid, the end must be
	// zero or larger and smaller than the number of lines in the document.
	StartLine uint64 `json:"startLine"`

	// The zero-based character offset from where the folded range starts.
	// If not defined, defaults to the length of the start line.
	StartCharacter *uint64 `json:"startCharacter"`

	// The zero-based end line of the range to fold. The folded area ends
	// with the line's last character. To be valid, the end must be zero or
	// larger and smaller than the number of lines in the document.
	EndLine uint64 `json:"endLine"`

	// The zero-based character offset before the folded range ends. If not
	// defined, defaults to the length of the end line.
	EndCharacter *uint64 `json:"endCharacter"`

	// Describes the kind of the folding range such as `comment' or
	// 'region'. The kind is used to categorize folding ranges and used by
	// commands like 'Fold all comments'. See
	// [FoldingRangeKind](#FoldingRangeKind) for an enumeration of
	// standardized kinds.
	Kind *FoldingRangeKind `json:"kind"`

	// The text that the client should show when the specified range is
	// collapsed. If not defined or not supported by the client, a default
	// will be chosen by the client.  @since 3.17.0
	CollapsedText *string `json:"collapsedText"`
}

type FoldingRangeRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	FoldingRangeOptions

	// mixins

	StaticRegistrationOptions
}

type DeclarationParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

type DeclarationRegistrationOptions struct {

	// extends

	DeclarationOptions

	TextDocumentRegistrationOptions

	// mixins

	StaticRegistrationOptions
}

// A parameter literal used in selection range requests.
type SelectionRangeParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The positions inside the text document.
	Positions []Position `json:"positions"`
}

// A selection range represents a part of a selection hierarchy. A selection
// range may have a parent selection range that contains it.
type SelectionRange struct {

	// The [range](#Range) of this selection range.
	Range Range `json:"range"`

	// The parent selection range containing this range. Therefore
	// `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent"`
}

type SelectionRangeRegistrationOptions struct {

	// extends

	SelectionRangeOptions

	TextDocumentRegistrationOptions

	// mixins

	StaticRegistrationOptions
}

type WorkDoneProgressCreateParams struct {

	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}

type WorkDoneProgressCancelParams struct {

	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}

// The parameter of a `textDocument/prepareCallHierarchy` request.  @since
// 3.16.0
type CallHierarchyPrepareParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// Represents programming constructs like functions or constructors in the
// context of call hierarchy.  @since 3.16.0
type CallHierarchyItem struct {

	// The name of this item.
	Name string `json:"name"`

	// The kind of this item.
	Kind SymbolKind `json:"kind"`

	// Tags for this item.
	Tags *[]SymbolTag `json:"tags"`

	// More detail for this item, e.g. the signature of a function.
	Detail *string `json:"detail"`

	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else, e.g. comments and code.
	Range Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g. the name of a function. Must be contained by the
	// [`range`](#CallHierarchyItem.range).
	SelectionRange Range `json:"selectionRange"`

	// A data entry field that is preserved between a call hierarchy prepare
	// and incoming calls or outgoing calls requests.
	Data *LSPAny `json:"data"`
}

// Call hierarchy options used during static or dynamic registration.  @since
// 3.16.0
type CallHierarchyRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CallHierarchyOptions

	// mixins

	StaticRegistrationOptions
}

// The parameter of a `callHierarchy/incomingCalls` request.  @since 3.16.0
type CallHierarchyIncomingCallsParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

// Represents an incoming call, e.g. a caller of a method or constructor.
// @since 3.16.0
type CallHierarchyIncomingCall struct {

	// The item that makes the call.
	From CallHierarchyItem `json:"from"`

	// The ranges at which the calls appear. This is relative to the caller
	// denoted by [`this.from`](#CallHierarchyIncomingCall.from).
	FromRanges []Range `json:"fromRanges"`
}

// The parameter of a `callHierarchy/outgoingCalls` request.  @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item CallHierarchyItem `json:"item"`
}

// Represents an outgoing call, e.g. calling a getter from a method or a method
// from a constructor etc.  @since 3.16.0
type CallHierarchyOutgoingCall struct {

	// The item that is called.
	To CallHierarchyItem `json:"to"`

	// The range at which this item is called. This is the range relative to
	// the caller, e.g the item passed to
	// [`provideCallHierarchyOutgoingCalls`](#CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls)
	// and not [`this.to`](#CallHierarchyOutgoingCall.to).
	FromRanges []Range `json:"fromRanges"`
}

// @since 3.16.0
type SemanticTokensParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// @since 3.16.0
type SemanticTokens struct {

	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token
	// request. A server can then instead of computing all semantic tokens
	// again simply send a delta.
	ResultId *string `json:"resultId"`

	// The actual tokens.
	Data []uint64 `json:"data"`
}

// @since 3.16.0
type SemanticTokensPartialResult struct {
	Data []uint64 `json:"data"`
}

// @since 3.16.0
type SemanticTokensRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	SemanticTokensOptions

	// mixins

	StaticRegistrationOptions
}

// @since 3.16.0
type SemanticTokensDeltaParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The result id of a previous response. The result Id can either point
	// to a full response or a delta response depending on what was received
	// last.
	PreviousResultId string `json:"previousResultId"`
}

// @since 3.16.0
type SemanticTokensDelta struct {
	ResultId *string `json:"resultId"`

	// The semantic token edits to transform a previous result into a new
	// result.
	Edits []SemanticTokensEdit `json:"edits"`
}

// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit `json:"edits"`
}

// @since 3.16.0
type SemanticTokensRangeParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range the semantic tokens are requested for.
	Range Range `json:"range"`
}

// Params to show a document.  @since 3.16.0
type ShowDocumentParams struct {

	// The document uri to show.
	Uri URI `json:"uri"`

	// Indicates to show the resource in an external program. To show for
	// example `https://code.visualstudio.com/` in the default WEB browser
	// set `external` to `true`.
	External *bool `json:"external"`

	// An optional property to indicate whether the editor showing the
	// document should take focus or not. Clients might ignore this property
	// if an external program is started.
	TakeFocus *bool `json:"takeFocus"`

	// An optional selection range if the document is a text document.
	// Clients might ignore the property if an external program is started
	// or the file is not a text file.
	Selection *Range `json:"selection"`
}

// The result of a showDocument request.  @since 3.16.0
type ShowDocumentResult struct {

	// A boolean indicating if the show was successful.
	Success bool `json:"success"`
}

type LinkedEditingRangeParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// The result of a linked editing range request.  @since 3.16.0
type LinkedEditingRanges struct {

	// A list of ranges that can be edited together. The ranges must have
	// identical length and contain identical text content. The ranges
	// cannot overlap.
	Ranges []Range `json:"ranges"`

	// An optional word pattern (regular expression) that describes valid
	// contents for the given ranges. If no pattern is provided, the client
	// configuration's word pattern will be used.
	WordPattern *string `json:"wordPattern"`
}

type LinkedEditingRangeRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	LinkedEditingRangeOptions

	// mixins

	StaticRegistrationOptions
}

// The parameters sent in notifications/requests for user-initiated creation of
// files.  @since 3.16.0
type CreateFilesParams struct {

	// An array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
}

// A workspace edit represents changes to many resources managed in the
// workspace. The edit should either provide `changes` or `documentChanges`. If
// documentChanges are present they are preferred over `changes` if the client
// can handle versioned document edits.  Since version 3.13.0 a workspace edit
// can contain resource operations as well. If resource operations are present
// clients need to execute the operations in the order in which they are
// provided. So a workspace edit for example can consist of the following two
// changes: (1) a create file a.txt and (2) a text document edit which insert
// text into file a.txt.  An invalid sequence (e.g. (1) delete file a.txt and
// (2) insert text into file a.txt) will cause failure of the operation. How the
// client recovers from the failure is described by the client capability:
// `workspace.workspaceEdit.failureHandling`
type WorkspaceEdit struct {

	// Holds changes to existing resources.
	Changes *map[DocumentUri]interface{} `json:"changes"`

	// Depending on the client capability
	// `workspace.workspaceEdit.resourceOperations` document changes are
	// either an array of `TextDocumentEdit`s to express changes to n
	// different text documents where each text document edit addresses a
	// specific version of a text document. Or it can contain above
	// `TextDocumentEdit`s mixed with create, rename and delete file /
	// folder operations.  Whether a client supports versioned document
	// edits is expressed via `workspace.workspaceEdit.documentChanges`
	// client capability.  If a client neither supports `documentChanges`
	// nor `workspace.workspaceEdit.resourceOperations` then only plain
	// `TextEdit`s using the `changes` property are supported.
	DocumentChanges *[]WorkspaceEdit_DocumentChanges_Element__Or `json:"documentChanges"`

	// A map of change annotations that can be referenced in
	// `AnnotatedTextEdit`s or create, rename and delete file / folder
	// operations.  Whether clients honor this property depends on the
	// client capability `workspace.changeAnnotationSupport`.  @since 3.16.0
	ChangeAnnotations *map[ChangeAnnotationIdentifier]interface{} `json:"changeAnnotations"`
}

// The options to register for file operations.  @since 3.16.0
type FileOperationRegistrationOptions struct {

	// The actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

// The parameters sent in notifications/requests for user-initiated renames of
// files.  @since 3.16.0
type RenameFilesParams struct {

	// An array of all files/folders renamed in this operation. When a
	// folder is renamed, only the folder will be included, and not its
	// children.
	Files []FileRename `json:"files"`
}

// The parameters sent in notifications/requests for user-initiated deletes of
// files.  @since 3.16.0
type DeleteFilesParams struct {

	// An array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}

type MonikerParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// Moniker definition to match LSIF 0.5 moniker definition.  @since 3.16.0
type Moniker struct {

	// The scheme of the moniker. For example tsc or .Net
	Scheme string `json:"scheme"`

	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`

	// The scope in which the moniker is unique
	Unique UniquenessLevel `json:"unique"`

	// The moniker kind if known.
	Kind *MonikerKind `json:"kind"`
}

type MonikerRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	MonikerOptions
}

// The parameter of a `textDocument/prepareTypeHierarchy` request.  @since
// 3.17.0
type TypeHierarchyPrepareParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// @since 3.17.0
type TypeHierarchyItem struct {

	// The name of this item.
	Name string `json:"name"`

	// The kind of this item.
	Kind SymbolKind `json:"kind"`

	// Tags for this item.
	Tags *[]SymbolTag `json:"tags"`

	// More detail for this item, e.g. the signature of a function.
	Detail *string `json:"detail"`

	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else, e.g. comments and code.
	Range Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g. the name of a function. Must be contained by the
	// [`range`](#TypeHierarchyItem.range).
	SelectionRange Range `json:"selectionRange"`

	// A data entry field that is preserved between a type hierarchy prepare
	// and supertypes or subtypes requests. It could also be used to
	// identify the type hierarchy in the server, helping improve the
	// performance on resolving supertypes and subtypes.
	Data *LSPAny `json:"data"`
}

// Type hierarchy options used during static or dynamic registration.  @since
// 3.17.0
type TypeHierarchyRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	TypeHierarchyOptions

	// mixins

	StaticRegistrationOptions
}

// The parameter of a `typeHierarchy/supertypes` request.  @since 3.17.0
type TypeHierarchySupertypesParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

// The parameter of a `typeHierarchy/subtypes` request.  @since 3.17.0
type TypeHierarchySubtypesParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

// A parameter literal used in inline value requests.  @since 3.17.0
type InlineValueParams struct {

	// mixins

	WorkDoneProgressParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The document range for which inline values should be computed.
	Range Range `json:"range"`

	// Additional information about the context in which inline values were
	// requested.
	Context InlineValueContext `json:"context"`
}

// Inline value options used during static or dynamic registration.  @since
// 3.17.0
type InlineValueRegistrationOptions struct {

	// extends

	InlineValueOptions

	TextDocumentRegistrationOptions

	// mixins

	StaticRegistrationOptions
}

// A parameter literal used in inlay hint requests.  @since 3.17.0
type InlayHintParams struct {

	// mixins

	WorkDoneProgressParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The document range for which inlay hints should be computed.
	Range Range `json:"range"`
}

// Inlay hint information.  @since 3.17.0
type InlayHint struct {

	// The position of this hint.
	Position Position `json:"position"`

	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.  *Note* that neither the string nor
	// the label part can be empty.
	Label InlayHint_Label__Or `json:"label"`

	// The kind of this hint. Can be omitted in which case the client should
	// fall back to a reasonable default.
	Kind *InlayHintKind `json:"kind"`

	// Optional text edits that are performed when accepting this inlay
	// hint.  *Note* that edits are expected to change the document so that
	// the inlay hint (or its nearest variant) is now part of the document
	// and the inlay hint itself is now obsolete.
	TextEdits *[]TextEdit `json:"textEdits"`

	// The tooltip text when you hover over this item.
	Tooltip *InlayHint_Tooltip__Or `json:"tooltip"`

	// Render padding before the hint.  Note: Padding should use the
	// editor's background color, not the background color of the hint
	// itself. That means padding can be used to visually align/separate an
	// inlay hint.
	PaddingLeft *bool `json:"paddingLeft"`

	// Render padding after the hint.  Note: Padding should use the editor's
	// background color, not the background color of the hint itself. That
	// means padding can be used to visually align/separate an inlay hint.
	PaddingRight *bool `json:"paddingRight"`

	// A data entry field that is preserved on an inlay hint between a
	// `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data *LSPAny `json:"data"`
}

// Inlay hint options used during static or dynamic registration.  @since 3.17.0
type InlayHintRegistrationOptions struct {

	// extends

	InlayHintOptions

	TextDocumentRegistrationOptions

	// mixins

	StaticRegistrationOptions
}

// Parameters of the document diagnostic request.  @since 3.17.0
type DocumentDiagnosticParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The additional identifier  provided during registration.
	Identifier *string `json:"identifier"`

	// The result id of a previous response if provided.
	PreviousResultId *string `json:"previousResultId"`
}

// A partial result for a document diagnostic report.  @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {
	RelatedDocuments map[DocumentUri]interface{} `json:"relatedDocuments"`
}

// Cancellation data returned from a diagnostic request.  @since 3.17.0
type DiagnosticServerCancellationData struct {
	RetriggerRequest bool `json:"retriggerRequest"`
}

// Diagnostic registration options.  @since 3.17.0
type DiagnosticRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DiagnosticOptions

	// mixins

	StaticRegistrationOptions
}

// Parameters of the workspace diagnostic request.  @since 3.17.0
type WorkspaceDiagnosticParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The additional identifier provided during registration.
	Identifier *string `json:"identifier"`

	// The currently known diagnostic reports with their previous result
	// ids.
	PreviousResultIds []PreviousResultId `json:"previousResultIds"`
}

// A workspace diagnostic report.  @since 3.17.0
type WorkspaceDiagnosticReport struct {
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

// A partial result for a workspace diagnostic report.  @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

// The params sent in an open notebook document notification.  @since 3.17.0
type DidOpenNotebookDocumentParams struct {

	// The notebook document that got opened.
	NotebookDocument NotebookDocument `json:"notebookDocument"`

	// The text documents that represent the content of a notebook cell.
	CellTextDocuments []TextDocumentItem `json:"cellTextDocuments"`
}

// The params sent in a change notebook document notification.  @since 3.17.0
type DidChangeNotebookDocumentParams struct {

	// The notebook document that did change. The version number points to
	// the version after all provided changes have been applied. If only the
	// text document content of a cell changes the notebook version doesn't
	// necessarily have to change.
	NotebookDocument VersionedNotebookDocumentIdentifier `json:"notebookDocument"`

	// The actual changes to the notebook document.  The changes describe
	// single state changes to the notebook document. So if there are two
	// changes c1 (at array index 0) and c2 (at array index 1) for a
	// notebook in state S then c1 moves the notebook from S to S' and c2
	// from S' to S''. So c1 is computed on the state S and c2 is computed
	// on the state S'.  To mirror the content of a notebook using change
	// events use the following approach: - start with the same initial
	// content - apply the 'notebookDocument/didChange' notifications in the
	// order you receive them. - apply the `NotebookChangeEvent`s in a
	// single notification in the order   you receive them.
	Change NotebookDocumentChangeEvent `json:"change"`
}

// The params sent in a save notebook document notification.  @since 3.17.0
type DidSaveNotebookDocumentParams struct {

	// The notebook document that got saved.
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`
}

// The params sent in a close notebook document notification.  @since 3.17.0
type DidCloseNotebookDocumentParams struct {

	// The notebook document that got closed.
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`

	// The text documents that represent the content of a notebook cell that
	// got closed.
	CellTextDocuments []TextDocumentIdentifier `json:"cellTextDocuments"`
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

type UnregistrationParams struct {
	Unregisterations []Unregistration `json:"unregisterations"`
}

type InitializeParams struct {

	// extends

	_InitializeParams

	WorkspaceFoldersInitializeParams
}

// The result returned from an initialize request.
type InitializeResult struct {

	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// Information about the server.  @since 3.15.0
	ServerInfo *InitializeResult_ServerInfo `json:"serverInfo"`
}

// The data type of the ResponseError if the initialize request fails.
type InitializeError struct {

	// Indicates whether the client execute the following retry logic: (1)
	// show the message provided by the ResponseError to the user (2) user
	// selects retry or cancel (3) if user selected retry the initialize
	// method is sent again.
	Retry bool `json:"retry"`
}

type InitializedParams struct {
}

// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {

	// The actual changed settings
	Settings LSPAny `json:"settings"`
}

type DidChangeConfigurationRegistrationOptions struct {
	Section *DidChangeConfigurationRegistrationOptions_Section__Or `json:"section"`
}

// The parameters of a notification message.
type ShowMessageParams struct {

	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}

type ShowMessageRequestParams struct {

	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`

	// The message action items to present.
	Actions *[]MessageActionItem `json:"actions"`
}

type MessageActionItem struct {

	// A short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

// The log message parameters.
type LogMessageParams struct {

	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}

// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {

	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// The change text document notification's parameters.
type DidChangeTextDocumentParams struct {

	// The document that did change. The version number points to the
	// version after all provided content changes have been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// The actual content changes. The content changes describe single state
	// changes to the document. So if there are two content changes c1 (at
	// array index 0) and c2 (at array index 1) for a document in state S
	// then c1 moves the document from S to S' and c2 from S' to S''. So c1
	// is computed on the state S and c2 is computed on the state S'.  To
	// mirror the content of a document using change events use the
	// following approach: - start with the same initial content - apply the
	// 'textDocument/didChange' notifications in the order you receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single
	// notification in the order   you receive them.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {

	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {

	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text *string `json:"text"`
}

// Save registration options.
type TextDocumentSaveRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	SaveOptions
}

// The parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {

	// The document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
}

// A text edit applicable to a text document.
type TextEdit struct {

	// The range of the text document to be manipulated. To insert text into
	// a document create a range where start === end.
	Range Range `json:"range"`

	// The string to be inserted. For delete operations use an empty string.
	NewText string `json:"newText"`
}

// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {

	// The actual file events.
	Changes []FileEvent `json:"changes"`
}

// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {

	// The watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}

// The publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {

	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`

	// Optional the version number of the document the diagnostics are
	// published for.  @since 3.15.0
	Version *int64 `json:"version"`

	// An array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// Completion parameters
type CompletionParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The completion context. This is only available it the client
	// specifies to send this using the client capability
	// `textDocument.completion.contextSupport === true`
	Context *CompletionContext `json:"context"`
}

// A completion item represents a text snippet that is proposed to complete text
// that is being typed.
type CompletionItem struct {

	// The label of this completion item.  The label property is also by
	// default the text that is inserted when selecting this completion.  If
	// label details are provided the label itself should be an unqualified
	// name of the completion item.
	Label string `json:"label"`

	// Additional details for the label  @since 3.17.0
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails"`

	// The kind of this completion item. Based of the kind an icon is chosen
	// by the editor.
	Kind *CompletionItemKind `json:"kind"`

	// Tags for this completion item.  @since 3.15.0
	Tags *[]CompletionItemTag `json:"tags"`

	// A human-readable string with additional information about this item,
	// like type or symbol information.
	Detail *string `json:"detail"`

	// A human-readable string that represents a doc-comment.
	Documentation *CompletionItem_Documentation__Or `json:"documentation"`

	// Indicates if this item is deprecated. @deprecated Use `tags` instead.
	Deprecated *bool `json:"deprecated"`

	// Select this item when showing.  *Note* that only one completion item
	// can be selected and that the tool / client decides which item that
	// is. The rule is that the *first* item of those that match best is
	// selected.
	Preselect *bool `json:"preselect"`

	// A string that should be used when comparing this item with other
	// items. When `falsy` the [label](#CompletionItem.label) is used.
	SortText *string `json:"sortText"`

	// A string that should be used when filtering a set of completion
	// items. When `falsy` the [label](#CompletionItem.label) is used.
	FilterText *string `json:"filterText"`

	// A string that should be inserted into a document when selecting this
	// completion. When `falsy` the [label](#CompletionItem.label) is used.
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example VS Code
	// when code complete is requested in this example `con<cursor
	// position>` and a completion item with an `insertText` of `console` is
	// provided it will only insert `sole`. Therefore it is recommended to
	// use `textEdit` instead since it avoids additional client side
	// interpretation.
	InsertText *string `json:"insertText"`

	// The format of the insert text. The format applies to both the
	// `insertText` property and the `newText` property of a provided
	// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	// Please note that the insertTextFormat doesn't apply to
	// `additionalTextEdits`.
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat"`

	// How whitespace and indentation is handled during completion item
	// insertion. If not provided the clients default value depends on the
	// `textDocument.completion.insertTextMode` client capability.  @since
	// 3.16.0
	InsertTextMode *InsertTextMode `json:"insertTextMode"`

	// An [edit](#TextEdit) which is applied to a document when selecting
	// this completion. When an edit is provided the value of
	// [insertText](#CompletionItem.insertText) is ignored.  Most editors
	// support two different operations when accepting a completion item.
	// One is to insert a completion text and the other is to replace an
	// existing text with a completion text. Since this can usually not be
	// predetermined by a server it can report both ranges. Clients need to
	// signal support for `InsertReplaceEdits` via the
	// `textDocument.completion.insertReplaceSupport` client capability
	// property.  *Note 1:* The text edit's range as well as both ranges
	// from an insert replace edit must be a [single line] and they must
	// contain the position at which completion has been requested. *Note
	// 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.  @since 3.16.0
	// additional type `InsertReplaceEdit`
	TextEdit *CompletionItem_TextEdit__Or `json:"textEdit"`

	// The edit text used if the completion item is part of a CompletionList
	// and CompletionList defines an item default for the text edit range.
	// Clients will only honor this property if they opt into completion
	// list item defaults using the capability
	// `completionList.itemDefaults`.  If not provided and a list's default
	// range is provided the label property is used as a text.  @since
	// 3.17.0
	TextEditText *string `json:"textEditText"`

	// An optional array of additional [text edits](#TextEdit) that are
	// applied when selecting this completion. Edits must not overlap
	// (including the same insert position) with the main
	// [edit](#CompletionItem.textEdit) nor with themselves.  Additional
	// text edits should be used to change text unrelated to the current
	// cursor position (for example adding an import statement at the top of
	// the file if the completion item will insert an unqualified type).
	AdditionalTextEdits *[]TextEdit `json:"additionalTextEdits"`

	// An optional set of characters that when pressed while this completion
	// is active will accept it first and then type that character. *Note*
	// that all commit characters should have `length=1` and that
	// superfluous characters will be ignored.
	CommitCharacters *[]string `json:"commitCharacters"`

	// An optional [command](#Command) that is executed *after* inserting
	// this completion. *Note* that additional modifications to the current
	// document should be described with the
	// [additionalTextEdits](#CompletionItem.additionalTextEdits)-property.
	Command *Command `json:"command"`

	// A data entry field that is preserved on a completion item between a
	// [CompletionRequest](#CompletionRequest) and a
	// [CompletionResolveRequest](#CompletionResolveRequest).
	Data *LSPAny `json:"data"`
}

// Represents a collection of [completion items](#CompletionItem) to be
// presented in the editor.
type CompletionList struct {

	// This list it not complete. Further typing results in recomputing this
	// list.  Recomputed lists have all their items replaced (not appended)
	// in the incomplete completion sessions.
	IsIncomplete bool `json:"isIncomplete"`

	// In many cases the items of an actual completion result share the same
	// value for properties like `commitCharacters` or the range of a text
	// edit. A completion list can therefore define item defaults which will
	// be used if a completion item itself doesn't specify the value.  If a
	// completion list specifies a default value and a completion item also
	// specifies a corresponding value the one from the item is used.
	// Servers are only allowed to return default values if the client
	// signals support for this via the `completionList.itemDefaults`
	// capability.  @since 3.17.0
	ItemDefaults *CompletionList_ItemDefaults `json:"itemDefaults"`

	// The completion items.
	Items []CompletionItem `json:"items"`
}

// Registration options for a [CompletionRequest](#CompletionRequest).
type CompletionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CompletionOptions
}

// Parameters for a [HoverRequest](#HoverRequest).
type HoverParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// The result of a hover request.
type Hover struct {

	// The hover's content
	Contents Hover_Contents__Or `json:"contents"`

	// An optional range inside the text document that is used to visualize
	// the hover, e.g. by changing the background color.
	Range *Range `json:"range"`
}

// Registration options for a [HoverRequest](#HoverRequest).
type HoverRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	HoverOptions
}

// Parameters for a [SignatureHelpRequest](#SignatureHelpRequest).
type SignatureHelpParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	// The signature help context. This is only available if the client
	// specifies to send this using the client capability
	// `textDocument.signatureHelp.contextSupport === true`  @since 3.15.0
	Context *SignatureHelpContext `json:"context"`
}

// Signature help represents the signature of something callable. There can be
// multiple signature but only one active and only one active parameter.
type SignatureHelp struct {

	// One or more signatures.
	Signatures []SignatureInformation `json:"signatures"`

	// The active signature. If omitted or the value lies outside the range
	// of `signatures` the value defaults to zero or is ignored if the
	// `SignatureHelp` has no signatures.  Whenever possible implementors
	// should make an active decision about the active signature and
	// shouldn't rely on a default value.  In future version of the protocol
	// this property might become mandatory to better express this.
	ActiveSignature *uint64 `json:"activeSignature"`

	// The active parameter of the active signature. If omitted or the value
	// lies outside the range of `signatures[activeSignature].parameters`
	// defaults to 0 if the active signature has parameters. If the active
	// signature has no parameters it is ignored. In future version of the
	// protocol this property might become mandatory to better express the
	// active parameter if the active signature does have any.
	ActiveParameter *uint64 `json:"activeParameter"`
}

// Registration options for a [SignatureHelpRequest](#SignatureHelpRequest).
type SignatureHelpRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	SignatureHelpOptions
}

// Parameters for a [DefinitionRequest](#DefinitionRequest).
type DefinitionParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// Registration options for a [DefinitionRequest](#DefinitionRequest).
type DefinitionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DefinitionOptions
}

// Parameters for a [ReferencesRequest](#ReferencesRequest).
type ReferenceParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Context ReferenceContext `json:"context"`
}

// Registration options for a [ReferencesRequest](#ReferencesRequest).
type ReferenceRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	ReferenceOptions
}

// Parameters for a [DocumentHighlightRequest](#DocumentHighlightRequest).
type DocumentHighlightParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// A document highlight is a range inside a text document which deserves special
// attention. Usually a document highlight is visualized by changing the
// background color of its range.
type DocumentHighlight struct {

	// The range this highlight applies to.
	Range Range `json:"range"`

	// The highlight kind, default is [text](#DocumentHighlightKind.Text).
	Kind *DocumentHighlightKind `json:"kind"`
}

// Registration options for a
// [DocumentHighlightRequest](#DocumentHighlightRequest).
type DocumentHighlightRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentHighlightOptions
}

// Parameters for a [DocumentSymbolRequest](#DocumentSymbolRequest).
type DocumentSymbolParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// Represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {

	// extends

	BaseSymbolInformation

	// Indicates if this symbol is deprecated.  @deprecated Use tags instead
	Deprecated *bool `json:"deprecated"`

	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in
	// the tool the range's start information is used to position the
	// cursor. So the range usually spans more than the actual symbol's name
	// and does normally include things like visibility modifiers.  The
	// range doesn't have to denote a node range in the sense of an abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy
	// of the symbols.
	Location Location `json:"location"`
}

// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they have
// two ranges: one that encloses its definition and one that points to its most
// interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {

	// The name of this symbol. Will be displayed in the user interface and
	// therefore must not be an empty string or a string only consisting of
	// white spaces.
	Name string `json:"name"`

	// More detail for this symbol, e.g the signature of a function.
	Detail *string `json:"detail"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this document symbol.  @since 3.16.0
	Tags *[]SymbolTag `json:"tags"`

	// Indicates if this symbol is deprecated.  @deprecated Use tags instead
	Deprecated *bool `json:"deprecated"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else like comments. This information is
	// typically used to determine if the clients cursor is inside the
	// symbol to reveal in the symbol in the UI.
	Range Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g the name of a function. Must be contained by the
	// `range`.
	SelectionRange Range `json:"selectionRange"`

	// Children of this symbol, e.g. properties of a class.
	Children *[]DocumentSymbol `json:"children"`
}

// Registration options for a [DocumentSymbolRequest](#DocumentSymbolRequest).
type DocumentSymbolRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentSymbolOptions
}

// The parameters of a [CodeActionRequest](#CodeActionRequest).
type CodeActionParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range for which the command was invoked.
	Range Range `json:"range"`

	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
}

// Represents a reference to a command. Provides a title which will be used to
// represent a command in the UI and, optionally, an array of arguments which
// will be passed to the command handler function when invoked.
type Command struct {

	// Title of the command, like `save`.
	Title string `json:"title"`

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command handler should be invoked with.
	Arguments *[]LSPAny `json:"arguments"`
}

// A code action represents a change that can be performed in code, e.g. to fix
// a problem or to refactor code.  A CodeAction must set either `edit` and/or a
// `command`. If both are supplied, the `edit` is applied first, then the
// `command` is executed.
type CodeAction struct {

	// A short, human-readable, title for this code action.
	Title string `json:"title"`

	// The kind of the code action.  Used to filter code actions.
	Kind *CodeActionKind `json:"kind"`

	// The diagnostics that this code action resolves.
	Diagnostics *[]Diagnostic `json:"diagnostics"`

	// Marks this as a preferred action. Preferred actions are used by the
	// `auto fix` command and can be targeted by keybindings.  A quick fix
	// should be marked preferred if it properly addresses the underlying
	// error. A refactoring should be marked preferred if it is the most
	// reasonable choice of actions to take.  @since 3.15.0
	IsPreferred *bool `json:"isPreferred"`

	// Marks that the code action cannot currently be applied.  Clients
	// should follow the following guidelines regarding disabled code
	// actions:    - Disabled code actions are not shown in automatic
	// [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action)
	//     code action menus.    - Disabled actions are shown as faded out
	// in the code action menu when the user requests a more specific type
	//   of code action, such as refactorings.    - If the user has a
	// [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions)
	//     that auto applies a code action and only disabled code actions
	// are returned, the client should show the user an     error message
	// with `reason` in the editor.  @since 3.16.0
	Disabled *CodeAction_Disabled `json:"disabled"`

	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit"`

	// A command this code action executes. If a code action provides an
	// edit and a command, first the edit is executed and then the command.
	Command *Command `json:"command"`

	// A data entry field that is preserved on a code action between a
	// `textDocument/codeAction` and a `codeAction/resolve` request.  @since
	// 3.16.0
	Data *LSPAny `json:"data"`
}

// Registration options for a [CodeActionRequest](#CodeActionRequest).
type CodeActionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CodeActionOptions
}

// The parameters of a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
type WorkspaceSymbolParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// A query string to filter symbols by. Clients may send an empty string
	// here to request all symbols.
	Query string `json:"query"`
}

// A special workspace symbol that supports locations without a range.  See also
// SymbolInformation.  @since 3.17.0
type WorkspaceSymbol struct {

	// extends

	BaseSymbolInformation

	// The location of the symbol. Whether a server is allowed to return a
	// location without a range depends on the client capability
	// `workspace.symbol.resolveSupport`.  See SymbolInformation#location
	// for more details.
	Location WorkspaceSymbol_Location__Or `json:"location"`

	// A data entry field that is preserved on a workspace symbol between a
	// workspace symbol request and a workspace symbol resolve request.
	Data *LSPAny `json:"data"`
}

// Registration options for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
type WorkspaceSymbolRegistrationOptions struct {

	// extends

	WorkspaceSymbolOptions
}

// The parameters of a [CodeLensRequest](#CodeLensRequest).
type CodeLensParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// A code lens represents a [command](#Command) that should be shown along with
// source text, like the number of references, a way to run tests, etc.  A code
// lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two
// stages.
type CodeLens struct {

	// The range in which this code lens is valid. Should only span a single
	// line.
	Range Range `json:"range"`

	// The command this code lens represents.
	Command *Command `json:"command"`

	// A data entry field that is preserved on a code lens item between a
	// [CodeLensRequest](#CodeLensRequest) and a [CodeLensResolveRequest]
	// (#CodeLensResolveRequest)
	Data *LSPAny `json:"data"`
}

// Registration options for a [CodeLensRequest](#CodeLensRequest).
type CodeLensRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CodeLensOptions
}

// The parameters of a [DocumentLinkRequest](#DocumentLinkRequest).
type DocumentLinkParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// A document link is a range in a text document that links to an internal or
// external resource, like another text document or a web site.
type DocumentLink struct {

	// The range this link applies to.
	Range Range `json:"range"`

	// The uri this link points to. If missing a resolve request is sent
	// later.
	Target *string `json:"target"`

	// The tooltip text when you hover over this link.  If a tooltip is
	// provided, is will be displayed in a string that includes instructions
	// on how to trigger the link, such as `{0} (ctrl + click)`. The
	// specific instructions vary depending on OS, user settings, and
	// localization.  @since 3.15.0
	Tooltip *string `json:"tooltip"`

	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data *LSPAny `json:"data"`
}

// Registration options for a [DocumentLinkRequest](#DocumentLinkRequest).
type DocumentLinkRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentLinkOptions
}

// The parameters of a [DocumentFormattingRequest](#DocumentFormattingRequest).
type DocumentFormattingParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The format options.
	Options FormattingOptions `json:"options"`
}

// Registration options for a
// [DocumentFormattingRequest](#DocumentFormattingRequest).
type DocumentFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentFormattingOptions
}

// The parameters of a
// [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
type DocumentRangeFormattingParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range to format
	Range Range `json:"range"`

	// The format options
	Options FormattingOptions `json:"options"`
}

// Registration options for a
// [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
type DocumentRangeFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentRangeFormattingOptions
}

// The parameters of a
// [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
type DocumentOnTypeFormattingParams struct {

	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The position around which the on type formatting should happen. This
	// is not necessarily the exact position where the character denoted by
	// the property `ch` got typed.
	Position Position `json:"position"`

	// The character that has been typed that triggered the formatting on
	// type request. That is not necessarily the last character that got
	// inserted into the document since the client could auto insert
	// characters as well (e.g. like automatic brace completion).
	Ch string `json:"ch"`

	// The formatting options.
	Options FormattingOptions `json:"options"`
}

// Registration options for a
// [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
type DocumentOnTypeFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentOnTypeFormattingOptions
}

// The parameters of a [RenameRequest](#RenameRequest).
type RenameParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to rename.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The position at which this request was sent.
	Position Position `json:"position"`

	// The new name of the symbol. If the given name is not valid the
	// request must return a [ResponseError](#ResponseError) with an
	// appropriate message set.
	NewName string `json:"newName"`
}

// Registration options for a [RenameRequest](#RenameRequest).
type RenameRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	RenameOptions
}

type PrepareRenameParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// The parameters of a [ExecuteCommandRequest](#ExecuteCommandRequest).
type ExecuteCommandParams struct {

	// mixins

	WorkDoneProgressParams

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command should be invoked with.
	Arguments *[]LSPAny `json:"arguments"`
}

// Registration options for a [ExecuteCommandRequest](#ExecuteCommandRequest).
type ExecuteCommandRegistrationOptions struct {

	// extends

	ExecuteCommandOptions
}

// The parameters passed via a apply workspace edit request.
type ApplyWorkspaceEditParams struct {

	// An optional label of the workspace edit. This label is presented in
	// the user interface for example on an undo stack to undo the workspace
	// edit.
	Label *string `json:"label"`

	// The edits to apply.
	Edit WorkspaceEdit `json:"edit"`
}

// The result returned from the apply workspace edit request.  @since 3.17
// renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {

	// Indicates whether the edit was applied or not.
	Applied bool `json:"applied"`

	// An optional textual description for why the edit was not applied.
	// This may be used by the server for diagnostic logging or to provide a
	// suitable error for a request that triggered the edit.
	FailureReason *string `json:"failureReason"`

	// Depending on the client's failure handling strategy `failedChange`
	// might contain the index of the change that failed. This property is
	// only available if the client signals a `failureHandlingStrategy` in
	// its client capabilities.
	FailedChange *uint64 `json:"failedChange"`
}

type WorkDoneProgressBegin struct {
	Kind string `json:"kind"`

	// Mandatory title of the progress operation. Used to briefly inform
	// about the kind of operation being performed.  Examples: "Indexing" or
	// "Linking dependencies".
	Title string `json:"title"`

	// Controls if a cancel button should show to allow the user to cancel
	// the long running operation. Clients that don't support cancellation
	// are allowed to ignore the setting.
	Cancellable *bool `json:"cancellable"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.  Examples: "3/25 files",
	// "project/src/module2", "node_modules/some_dep". If unset, the
	// previous progress message (if any) is still valid.
	Message *string `json:"message"`

	// Optional progress percentage to display (value 100 is considered
	// 100%). If not provided infinite progress is assumed and clients are
	// allowed to ignore the `percentage` value in subsequent in report
	// notifications.  The value should be steadily rising. Clients are free
	// to ignore values that are not following this rule. The value range is
	// [0, 100].
	Percentage *uint64 `json:"percentage"`
}

type WorkDoneProgressReport struct {
	Kind string `json:"kind"`

	// Controls enablement state of a cancel button.  Clients that don't
	// support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable *bool `json:"cancellable"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.  Examples: "3/25 files",
	// "project/src/module2", "node_modules/some_dep". If unset, the
	// previous progress message (if any) is still valid.
	Message *string `json:"message"`

	// Optional progress percentage to display (value 100 is considered
	// 100%). If not provided infinite progress is assumed and clients are
	// allowed to ignore the `percentage` value in subsequent in report
	// notifications.  The value should be steadily rising. Clients are free
	// to ignore values that are not following this rule. The value range is
	// [0, 100]
	Percentage *uint64 `json:"percentage"`
}

type WorkDoneProgressEnd struct {
	Kind string `json:"kind"`

	// Optional, a final message indicating to for example indicate the
	// outcome of the operation.
	Message *string `json:"message"`
}

type SetTraceParams struct {
	Value TraceValues `json:"value"`
}

type LogTraceParams struct {
	Message string `json:"message"`

	Verbose *string `json:"verbose"`
}

type CancelParams struct {

	// The request id to cancel.
	Id CancelParams_Id__Or `json:"id"`
}

type ProgressParams struct {

	// The progress token provided by the client or server.
	Token ProgressToken `json:"token"`

	// The progress data.
	Value LSPAny `json:"value"`
}

// A parameter literal used in requests to pass a text document and a position
// inside that document.
type TextDocumentPositionParams struct {

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The position inside the text document.
	Position Position `json:"position"`
}

type WorkDoneProgressParams struct {

	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken"`
}

// Represents the connection of two locations. Provides additional metadata over
// normal [locations](#Location), including an origin range.
type LocationLink struct {

	// Span of the origin of this link.  Used as the underlined span for
	// mouse interaction. Defaults to the word range at the definition
	// position.
	OriginSelectionRange *Range `json:"originSelectionRange"`

	// The target resource identifier of this link.
	TargetUri DocumentUri `json:"targetUri"`

	// The full target range of this link. If the target for example is a
	// symbol then target range is the range enclosing this symbol not
	// including leading/trailing whitespace but everything else like
	// comments. This information is typically used to highlight the range
	// in the editor.
	TargetRange Range `json:"targetRange"`

	// The range that should be selected and revealed when this link is
	// being followed, e.g the name of a function. Must be contained by the
	// `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

// A range in a text document expressed as (zero-based) start and end positions.
//
//	If you want to specify a range that contains a line including the line
//
// ending character(s) then use an end position denoting the start of the next
// line. For example: ```ts {     start: { line: 5, character: 23 }     end : {
// line 6, character : 0 } } ```
type Range struct {

	// The range's start position.
	Start Position `json:"start"`

	// The range's end position.
	End Position `json:"end"`
}

type ImplementationOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Static registration options to be returned in the initialize request.
type StaticRegistrationOptions struct {

	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id *string `json:"id"`
}

type TypeDefinitionOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// The workspace folder change event.
type WorkspaceFoldersChangeEvent struct {

	// The array of added workspace folders
	Added []WorkspaceFolder `json:"added"`

	// The array of the removed workspace folders
	Removed []WorkspaceFolder `json:"removed"`
}

type ConfigurationItem struct {

	// The scope to get the configuration section for.
	ScopeUri *string `json:"scopeUri"`

	// The configuration section asked for.
	Section *string `json:"section"`
}

// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {

	// The text document's uri.
	Uri DocumentUri `json:"uri"`
}

// Represents a color in RGBA space.
type Color struct {

	// The red component of this color in the range [0-1].
	Red float64 `json:"red"`

	// The green component of this color in the range [0-1].
	Green float64 `json:"green"`

	// The blue component of this color in the range [0-1].
	Blue float64 `json:"blue"`

	// The alpha component of this color in the range [0-1].
	Alpha float64 `json:"alpha"`
}

type DocumentColorOptions struct {

	// mixins

	WorkDoneProgressOptions
}

type FoldingRangeOptions struct {

	// mixins

	WorkDoneProgressOptions
}

type DeclarationOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Position in a text document expressed as zero-based line and character
// offset. Prior to 3.17 the offsets were always based on a UTF-16 string
// representation. So a string of the form `a𐐀b` the character offset of the
// character `a` is 0, the character offset of `𐐀` is 1 and the character
// offset of b is 3 since `𐐀` is represented using two code units in UTF-16.
// Since 3.17 clients and servers can agree on a different string encoding
// representation (e.g. UTF-8). The client announces it's supported encoding via
// the client capability [`general.positionEncodings`](#clientCapabilities). The
// value is an array of position encodings the client supports, with decreasing
// preference (e.g. the encoding at index `0` is the most preferred one). To
// stay backwards compatible the only mandatory encoding is UTF-16 represented
// via the string `utf-16`. The server can pick one of the encodings offered by
// the client and signals that encoding back to the client via the initialize
// result's property [`capabilities.positionEncoding`](#serverCapabilities). If
// the string value `utf-16` is missing from the client's capability
// `general.positionEncodings` servers can safely assume that the client
// supports UTF-16. If the server omits the position encoding in its initialize
// result the encoding defaults to the string value `utf-16`. Implementation
// considerations: since the conversion from one encoding into another requires
// the content of the file / line the conversion is best done where the file is
// read which is usually on the server side.  Positions are line end character
// agnostic. So you can not specify a position that denotes `\r|\n` or `\n|`
// where `|` represents the character offset.  @since 3.17.0 - support for
// negotiated position encoding.
type Position struct {

	// Line position in a document (zero-based).  If a line number is
	// greater than the number of lines in a document, it defaults back to
	// the number of lines in the document. If a line number is negative, it
	// defaults to 0.
	Line uint64 `json:"line"`

	// Character offset on a line in a document (zero-based).  The meaning
	// of this offset is determined by the negotiated
	// `PositionEncodingKind`.  If the character value is greater than the
	// line length it defaults back to the line length.
	Character uint64 `json:"character"`
}

type SelectionRangeOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Call hierarchy options used during static registration.  @since 3.16.0
type CallHierarchyOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// @since 3.16.0
type SemanticTokensOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The legend used by the server
	Legend SemanticTokensLegend `json:"legend"`

	// Server supports providing semantic tokens for a specific range of a
	// document.
	Range *SemanticTokensOptions_Range__Or `json:"range"`

	// Server supports providing semantic tokens for a full document.
	Full *SemanticTokensOptions_Full__Or `json:"full"`
}

// @since 3.16.0
type SemanticTokensEdit struct {

	// The start offset of the edit.
	Start uint64 `json:"start"`

	// The count of elements to remove.
	DeleteCount uint64 `json:"deleteCount"`

	// The elements to insert.
	Data *[]uint64 `json:"data"`
}

type LinkedEditingRangeOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Represents information on a file/folder create.  @since 3.16.0
type FileCreate struct {

	// A file:// URI for the location of the file/folder being created.
	Uri string `json:"uri"`
}

// Describes textual changes on a text document. A TextDocumentEdit describes
// all changes on a document version Si and after they are applied move the
// document to version Si+1. So the creator of a TextDocumentEdit doesn't need
// to sort the array of edits or do any kind of ordering. However the edits must
// be non overlapping.
type TextDocumentEdit struct {

	// The text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`

	// The edits to be applied.  @since 3.16.0 - support for
	// AnnotatedTextEdit. This is guarded using a client capability.
	Edits []TextDocumentEdit_Edits_Element__Or `json:"edits"`
}

// Create file operation.
type CreateFile struct {

	// extends

	ResourceOperation

	// A create
	Kind string `json:"kind"`

	// The resource to create.
	Uri DocumentUri `json:"uri"`

	// Additional options
	Options *CreateFileOptions `json:"options"`
}

// Rename file operation
type RenameFile struct {

	// extends

	ResourceOperation

	// A rename
	Kind string `json:"kind"`

	// The old (existing) location.
	OldUri DocumentUri `json:"oldUri"`

	// The new location.
	NewUri DocumentUri `json:"newUri"`

	// Rename options.
	Options *RenameFileOptions `json:"options"`
}

// Delete file operation
type DeleteFile struct {

	// extends

	ResourceOperation

	// A delete
	Kind string `json:"kind"`

	// The file to delete.
	Uri DocumentUri `json:"uri"`

	// Delete options.
	Options *DeleteFileOptions `json:"options"`
}

// Additional information that describes document changes.  @since 3.16.0
type ChangeAnnotation struct {

	// A human-readable string describing the actual change. The string is
	// rendered prominent in the user interface.
	Label string `json:"label"`

	// A flag which indicates that user confirmation is needed before
	// applying the change.
	NeedsConfirmation *bool `json:"needsConfirmation"`

	// A human-readable string which is rendered less prominent in the user
	// interface.
	Description *string `json:"description"`
}

// A filter to describe in which file operation requests or notifications the
// server is interested in receiving.  @since 3.16.0
type FileOperationFilter struct {

	// A Uri scheme like `file` or `untitled`.
	Scheme *string `json:"scheme"`

	// The actual file operation pattern.
	Pattern FileOperationPattern `json:"pattern"`
}

// Represents information on a file/folder rename.  @since 3.16.0
type FileRename struct {

	// A file:// URI for the original location of the file/folder being
	// renamed.
	OldUri string `json:"oldUri"`

	// A file:// URI for the new location of the file/folder being renamed.
	NewUri string `json:"newUri"`
}

// Represents information on a file/folder delete.  @since 3.16.0
type FileDelete struct {

	// A file:// URI for the location of the file/folder being deleted.
	Uri string `json:"uri"`
}

type MonikerOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Type hierarchy options used during static registration.  @since 3.17.0
type TypeHierarchyOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// @since 3.17.0
type InlineValueContext struct {

	// The stack frame (as a DAP Id) where the execution has stopped.
	FrameId int64 `json:"frameId"`

	// The document range where execution has stopped. Typically the end
	// position of the range denotes the line where the inline values are
	// shown.
	StoppedLocation Range `json:"stoppedLocation"`
}

// Provide inline value as text.  @since 3.17.0
type InlineValueText struct {

	// The document range for which the inline value applies.
	Range Range `json:"range"`

	// The text of the inline value.
	Text string `json:"text"`
}

// Provide inline value through a variable lookup. If only a range is specified,
// the variable name will be extracted from the underlying document. An optional
// variable name can be used to override the extracted name.  @since 3.17.0
type InlineValueVariableLookup struct {

	// The document range for which the inline value applies. The range is
	// used to extract the variable name from the underlying document.
	Range Range `json:"range"`

	// If specified the name of the variable to look up.
	VariableName *string `json:"variableName"`

	// How to perform the lookup.
	CaseSensitiveLookup bool `json:"caseSensitiveLookup"`
}

// Provide an inline value through an expression evaluation. If only a range is
// specified, the expression will be extracted from the underlying document. An
// optional expression can be used to override the extracted expression.  @since
// 3.17.0
type InlineValueEvaluatableExpression struct {

	// The document range for which the inline value applies. The range is
	// used to extract the evaluatable expression from the underlying
	// document.
	Range Range `json:"range"`

	// If specified the expression overrides the extracted expression.
	Expression *string `json:"expression"`
}

// Inline value options used during static registration.  @since 3.17.0
type InlineValueOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// An inlay hint label part allows for interactive and composite labels of inlay
// hints.  @since 3.17.0
type InlayHintLabelPart struct {

	// The value of this label part.
	Value string `json:"value"`

	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might
	// resolve this property late using the resolve request.
	Tooltip *InlayHintLabelPart_Tooltip__Or `json:"tooltip"`

	// An optional source code location that represents this label part.
	// The editor will use this location for the hover and for code
	// navigation features: This part will become a clickable link that
	// resolves to the definition of the symbol at the given location (not
	// necessarily the location itself), it shows the hover that shows at
	// the given location, and it shows a context menu with further code
	// navigation commands.  Depending on the client capability
	// `inlayHint.resolveSupport` clients might resolve this property late
	// using the resolve request.
	Location *Location `json:"location"`

	// An optional command for this label part.  Depending on the client
	// capability `inlayHint.resolveSupport` clients might resolve this
	// property late using the resolve request.
	Command *Command `json:"command"`
}

// A `MarkupContent` literal represents a string value which content is
// interpreted base on its kind flag. Currently the protocol supports
// `plaintext` and `markdown` as markup kinds.  If the kind is `markdown` then
// the value can contain fenced code blocks like in GitHub issues. See
// https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
//
//	Here is an example how such a string can be constructed using JavaScript /
//
// TypeScript: ```ts let markdown: MarkdownContent = {  kind:
// MarkupKind.Markdown,  value: [    '# Header',    'Some text',
// '```typescript',    'someCode();',    '```'  ].join('\n') }; ```  *Please
// Note* that clients might sanitize the return markdown. A client could decide
// to remove HTML from the markdown to avoid script execution.
type MarkupContent struct {

	// The type of the Markup
	Kind MarkupKind `json:"kind"`

	// The content itself
	Value string `json:"value"`
}

// Inlay hint options used during static registration.  @since 3.17.0
type InlayHintOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The server provides support to resolve additional information for an
	// inlay hint item.
	ResolveProvider *bool `json:"resolveProvider"`
}

// A full diagnostic report with a set of related documents.  @since 3.17.0
type RelatedFullDocumentDiagnosticReport struct {

	// extends

	FullDocumentDiagnosticReport

	// Diagnostics of related documents. This information is useful in
	// programming languages where code in a file A can generate diagnostics
	// in a file B which A depends on. An example of such a language is
	// C/C++ where marco definitions in a file a.cpp and result in errors in
	// a header file b.hpp.  @since 3.17.0
	RelatedDocuments *map[DocumentUri]interface{} `json:"relatedDocuments"`
}

// An unchanged diagnostic report with a set of related documents.  @since
// 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {

	// extends

	UnchangedDocumentDiagnosticReport

	// Diagnostics of related documents. This information is useful in
	// programming languages where code in a file A can generate diagnostics
	// in a file B which A depends on. An example of such a language is
	// C/C++ where marco definitions in a file a.cpp and result in errors in
	// a header file b.hpp.  @since 3.17.0
	RelatedDocuments *map[DocumentUri]interface{} `json:"relatedDocuments"`
}

// A diagnostic report with a full set of problems.  @since 3.17.0
type FullDocumentDiagnosticReport struct {

	// A full document diagnostic report.
	Kind string `json:"kind"`

	// An optional result id. If provided it will be sent on the next
	// diagnostic request for the same document.
	ResultId *string `json:"resultId"`

	// The actual items.
	Items []Diagnostic `json:"items"`
}

// A diagnostic report indicating that the last returned report is still
// accurate.  @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {

	// A document diagnostic report indicating no changes to the last
	// result. A server can only return `unchanged` if result ids are
	// provided.
	Kind string `json:"kind"`

	// A result id which will be sent on the next diagnostic request for the
	// same document.
	ResultId string `json:"resultId"`
}

// Diagnostic options.  @since 3.17.0
type DiagnosticOptions struct {

	// mixins

	WorkDoneProgressOptions

	// An optional identifier under which the diagnostics are managed by the
	// client.
	Identifier *string `json:"identifier"`

	// Whether the language has inter file dependencies meaning that editing
	// code in one file can result in a different diagnostic set in another
	// file. Inter file dependencies are common for most programming
	// languages and typically uncommon for linters.
	InterFileDependencies bool `json:"interFileDependencies"`

	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}

// A previous result id in a workspace pull request.  @since 3.17.0
type PreviousResultId struct {

	// The URI for which the client knowns a result id.
	Uri DocumentUri `json:"uri"`

	// The value of the previous result id.
	Value string `json:"value"`
}

// A notebook document.  @since 3.17.0
type NotebookDocument struct {

	// The notebook document's uri.
	Uri URI `json:"uri"`

	// The type of the notebook.
	NotebookType string `json:"notebookType"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int64 `json:"version"`

	// Additional metadata stored with the notebook document.  Note: should
	// always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata"`

	// The cells of a notebook.
	Cells []NotebookCell `json:"cells"`
}

// An item to transfer a text document from the client to the server.
type TextDocumentItem struct {

	// The text document's uri.
	Uri DocumentUri `json:"uri"`

	// The text document's language identifier.
	LanguageId string `json:"languageId"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int64 `json:"version"`

	// The content of the opened text document.
	Text string `json:"text"`
}

// A versioned notebook document identifier.  @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {

	// The version number of this notebook document.
	Version int64 `json:"version"`

	// The notebook document's uri.
	Uri URI `json:"uri"`
}

// A change event for a notebook document.  @since 3.17.0
type NotebookDocumentChangeEvent struct {

	// The changed meta data if any.  Note: should always be an object
	// literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata"`

	// Changes to cells
	Cells *NotebookDocumentChangeEvent_Cells `json:"cells"`
}

// A literal to identify a notebook document in the client.  @since 3.17.0
type NotebookDocumentIdentifier struct {

	// The notebook document's uri.
	Uri URI `json:"uri"`
}

// General parameters to to register for an notification or to register a
// provider.
type Registration struct {

	// The id used to register the request. The id can be used to deregister
	// the request again.
	Id string `json:"id"`

	// The method / capability to register for.
	Method string `json:"method"`

	// Options necessary for the registration.
	RegisterOptions *LSPAny `json:"registerOptions"`
}

// General parameters to unregister a request or notification.
type Unregistration struct {

	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id string `json:"id"`

	// The method to unregister for.
	Method string `json:"method"`
}

// The initialize parameters
type _InitializeParams struct {

	// mixins

	WorkDoneProgressParams

	// The process Id of the parent process that started the server.  Is
	// `null` if the process has not been started by another process. If the
	// parent process is not alive then the server should exit.
	ProcessId _InitializeParams_ProcessId__Or `json:"processId"`

	// Information about the client  @since 3.15.0
	ClientInfo *_InitializeParams_ClientInfo `json:"clientInfo"`

	// The locale the client is currently showing the user interface in.
	// This must not necessarily be the locale of the operating system.
	// Uses IETF language tags as the value's syntax (See
	// https://en.wikipedia.org/wiki/IETF_language_tag)  @since 3.16.0
	Locale *string `json:"locale"`

	// The rootPath of the workspace. Is null if no folder is open.
	// @deprecated in favour of rootUri.
	RootPath *_InitializeParams_RootPath__Or `json:"rootPath"`

	// The rootUri of the workspace. Is null if no folder is open. If both
	// `rootPath` and `rootUri` are set `rootUri` wins.  @deprecated in
	// favour of workspaceFolders.
	RootUri _InitializeParams_RootUri__Or `json:"rootUri"`

	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// User provided initialization options.
	InitializationOptions *LSPAny `json:"initializationOptions"`

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *_InitializeParams_Trace__Or `json:"trace"`
}

type WorkspaceFoldersInitializeParams struct {

	// The workspace folders configured in the client when the server
	// starts.  This property is only available if the client supports
	// workspace folders. It can be `null` if the client supports workspace
	// folders but none are configured.  @since 3.6.0
	WorkspaceFolders *WorkspaceFoldersInitializeParams_WorkspaceFolders__Or `json:"workspaceFolders"`
}

// Defines the capabilities provided by a language server.
type ServerCapabilities struct {

	// The position encoding the server picked from the encodings offered by
	// the client via the client capability `general.positionEncodings`.  If
	// the client didn't provide any position encodings the only valid value
	// that a server can return is 'utf-16'.  If omitted it defaults to
	// 'utf-16'.  @since 3.17.0
	PositionEncoding *PositionEncodingKind `json:"positionEncoding"`

	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number.
	TextDocumentSync *ServerCapabilities_TextDocumentSync__Or `json:"textDocumentSync"`

	// Defines how notebook documents are synced.  @since 3.17.0
	NotebookDocumentSync *ServerCapabilities_NotebookDocumentSync__Or `json:"notebookDocumentSync"`

	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider"`

	// The server provides hover support.
	HoverProvider *ServerCapabilities_HoverProvider__Or `json:"hoverProvider"`

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider"`

	// The server provides Goto Declaration support.
	DeclarationProvider *ServerCapabilities_DeclarationProvider__Or `json:"declarationProvider"`

	// The server provides goto definition support.
	DefinitionProvider *ServerCapabilities_DefinitionProvider__Or `json:"definitionProvider"`

	// The server provides Goto Type Definition support.
	TypeDefinitionProvider *ServerCapabilities_TypeDefinitionProvider__Or `json:"typeDefinitionProvider"`

	// The server provides Goto Implementation support.
	ImplementationProvider *ServerCapabilities_ImplementationProvider__Or `json:"implementationProvider"`

	// The server provides find references support.
	ReferencesProvider *ServerCapabilities_ReferencesProvider__Or `json:"referencesProvider"`

	// The server provides document highlight support.
	DocumentHighlightProvider *ServerCapabilities_DocumentHighlightProvider__Or `json:"documentHighlightProvider"`

	// The server provides document symbol support.
	DocumentSymbolProvider *ServerCapabilities_DocumentSymbolProvider__Or `json:"documentSymbolProvider"`

	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *ServerCapabilities_CodeActionProvider__Or `json:"codeActionProvider"`

	// The server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider"`

	// The server provides color provider support.
	ColorProvider *ServerCapabilities_ColorProvider__Or `json:"colorProvider"`

	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *ServerCapabilities_WorkspaceSymbolProvider__Or `json:"workspaceSymbolProvider"`

	// The server provides document formatting.
	DocumentFormattingProvider *ServerCapabilities_DocumentFormattingProvider__Or `json:"documentFormattingProvider"`

	// The server provides document range formatting.
	DocumentRangeFormattingProvider *ServerCapabilities_DocumentRangeFormattingProvider__Or `json:"documentRangeFormattingProvider"`

	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider"`

	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports `prepareSupport` in
	// its initial `initialize` request.
	RenameProvider *ServerCapabilities_RenameProvider__Or `json:"renameProvider"`

	// The server provides folding provider support.
	FoldingRangeProvider *ServerCapabilities_FoldingRangeProvider__Or `json:"foldingRangeProvider"`

	// The server provides selection range support.
	SelectionRangeProvider *ServerCapabilities_SelectionRangeProvider__Or `json:"selectionRangeProvider"`

	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider"`

	// The server provides call hierarchy support.  @since 3.16.0
	CallHierarchyProvider *ServerCapabilities_CallHierarchyProvider__Or `json:"callHierarchyProvider"`

	// The server provides linked editing range support.  @since 3.16.0
	LinkedEditingRangeProvider *ServerCapabilities_LinkedEditingRangeProvider__Or `json:"linkedEditingRangeProvider"`

	// The server provides semantic tokens support.  @since 3.16.0
	SemanticTokensProvider *ServerCapabilities_SemanticTokensProvider__Or `json:"semanticTokensProvider"`

	// The server provides moniker support.  @since 3.16.0
	MonikerProvider *ServerCapabilities_MonikerProvider__Or `json:"monikerProvider"`

	// The server provides type hierarchy support.  @since 3.17.0
	TypeHierarchyProvider *ServerCapabilities_TypeHierarchyProvider__Or `json:"typeHierarchyProvider"`

	// The server provides inline values.  @since 3.17.0
	InlineValueProvider *ServerCapabilities_InlineValueProvider__Or `json:"inlineValueProvider"`

	// The server provides inlay hints.  @since 3.17.0
	InlayHintProvider *ServerCapabilities_InlayHintProvider__Or `json:"inlayHintProvider"`

	// The server has support for pull model diagnostics.  @since 3.17.0
	DiagnosticProvider *ServerCapabilities_DiagnosticProvider__Or `json:"diagnosticProvider"`

	// Workspace specific server capabilities.
	Workspace *ServerCapabilities_Workspace `json:"workspace"`

	// Experimental server capabilities.
	Experimental *LSPAny `json:"experimental"`
}

// A text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {

	// extends

	TextDocumentIdentifier

	// The version number of this document.
	Version int64 `json:"version"`
}

// Save options.
type SaveOptions struct {

	// The client is supposed to include the content on save.
	IncludeText *bool `json:"includeText"`
}

// An event describing a file change.
type FileEvent struct {

	// The file's uri.
	Uri DocumentUri `json:"uri"`

	// The change type.
	Type FileChangeType `json:"type"`
}

type FileSystemWatcher struct {

	// The glob pattern to watch. See {@link GlobPattern glob pattern} for
	// more detail.  @since 3.17.0 support for relative patterns.
	GlobPattern GlobPattern `json:"globPattern"`

	// The kind of events of interest. If omitted it defaults to
	// WatchKind.Create | WatchKind.Change | WatchKind.Delete which is 7.
	Kind *WatchKind `json:"kind"`
}

// Represents a diagnostic, such as a compiler error or warning. Diagnostic
// objects are only valid in the scope of a resource.
type Diagnostic struct {

	// The range at which the message applies
	Range Range `json:"range"`

	// The diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity *DiagnosticSeverity `json:"severity"`

	// The diagnostic's code, which usually appear in the user interface.
	Code *Diagnostic_Code__Or `json:"code"`

	// An optional property to describe the error code. Requires the code
	// field (above) to be present/not null.  @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription"`

	// A human-readable string describing the source of this diagnostic,
	// e.g. 'typescript' or 'super lint'. It usually appears in the user
	// interface.
	Source *string `json:"source"`

	// The diagnostic's message. It usually appears in the user interface
	Message string `json:"message"`

	// Additional metadata about the diagnostic.  @since 3.15.0
	Tags *[]DiagnosticTag `json:"tags"`

	// An array of related diagnostic information, e.g. when symbol-names
	// within a scope collide all definitions can be marked via this
	// property.
	RelatedInformation *[]DiagnosticRelatedInformation `json:"relatedInformation"`

	// A data entry field that is preserved between a
	// `textDocument/publishDiagnostics` notification and
	// `textDocument/codeAction` request.  @since 3.16.0
	Data *LSPAny `json:"data"`
}

// Contains additional information about the context in which a completion
// request is triggered.
type CompletionContext struct {

	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// The trigger character (a single character) that has trigger code
	// complete. Is undefined if `triggerKind !==
	// CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter *string `json:"triggerCharacter"`
}

// Additional details for a completion item label.  @since 3.17.0
type CompletionItemLabelDetails struct {

	// An optional string which is rendered less prominently directly after
	// {@link CompletionItem.label label}, without any spacing. Should be
	// used for function signatures and type annotations.
	Detail *string `json:"detail"`

	// An optional string which is rendered less prominently after {@link
	// CompletionItem.detail}. Should be used for fully qualified names and
	// file paths.
	Description *string `json:"description"`
}

// A special text edit to provide an insert and a replace operation.  @since
// 3.16.0
type InsertReplaceEdit struct {

	// The string to be inserted.
	NewText string `json:"newText"`

	// The range if the insert is requested
	Insert Range `json:"insert"`

	// The range if the replace is requested.
	Replace Range `json:"replace"`
}

// Completion options.
type CompletionOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Most tools trigger completion request automatically without
	// explicitly requesting it using a keyboard shortcut (e.g. Ctrl+Space).
	// Typically they do so when the user starts to type an identifier. For
	// example if the user types `c` in a JavaScript file code complete will
	// automatically pop up present `console` besides others as a completion
	// item. Characters that make up identifiers don't need to be listed
	// here.  If code complete should automatically be trigger on characters
	// not being valid inside an identifier (for example `.` in JavaScript)
	// list them in `triggerCharacters`.
	TriggerCharacters *[]string `json:"triggerCharacters"`

	// The list of all possible characters that commit a completion. This
	// field can be used if clients don't support individual commit
	// characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	//  If a server provides both `allCommitCharacters` and commit
	// characters on an individual completion item the ones on the
	// completion item win.  @since 3.2.0
	AllCommitCharacters *[]string `json:"allCommitCharacters"`

	// The server provides support to resolve additional information for a
	// completion item.
	ResolveProvider *bool `json:"resolveProvider"`

	// The server supports the following `CompletionItem` specific
	// capabilities.  @since 3.17.0
	CompletionItem *CompletionOptions_CompletionItem `json:"completionItem"`
}

// Hover options.
type HoverOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Additional information about the context in which a signature help request
// was triggered.  @since 3.15.0
type SignatureHelpContext struct {

	// Action that caused signature help to be triggered.
	TriggerKind SignatureHelpTriggerKind `json:"triggerKind"`

	// Character that caused signature help to be triggered.  This is
	// undefined when `triggerKind !==
	// SignatureHelpTriggerKind.TriggerCharacter`
	TriggerCharacter *string `json:"triggerCharacter"`

	// `true` if signature help was already showing when it was triggered.
	// Retriggers occurs when the signature help is already active and can
	// be caused by actions such as typing a trigger character, a cursor
	// move, or document content changes.
	IsRetrigger bool `json:"isRetrigger"`

	// The currently active `SignatureHelp`.  The `activeSignatureHelp` has
	// its `SignatureHelp.activeSignature` field updated based on the user
	// navigating through available signatures.
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp"`
}

// Represents the signature of something callable. A signature can have a label,
// like a function-name, a doc-comment, and a set of parameters.
type SignatureInformation struct {

	// The label of this signature. Will be shown in the UI.
	Label string `json:"label"`

	// The human-readable doc-comment of this signature. Will be shown in
	// the UI but can be omitted.
	Documentation *SignatureInformation_Documentation__Or `json:"documentation"`

	// The parameters of this signature.
	Parameters *[]ParameterInformation `json:"parameters"`

	// The index of the active parameter.  If provided, this is used in
	// place of `SignatureHelp.activeParameter`.  @since 3.16.0
	ActiveParameter *uint64 `json:"activeParameter"`
}

// Server Capabilities for a [SignatureHelpRequest](#SignatureHelpRequest).
type SignatureHelpOptions struct {

	// mixins

	WorkDoneProgressOptions

	// List of characters that trigger signature help automatically.
	TriggerCharacters *[]string `json:"triggerCharacters"`

	// List of characters that re-trigger signature help.  These trigger
	// characters are only active when signature help is already showing.
	// All trigger characters are also counted as re-trigger characters.
	// @since 3.15.0
	RetriggerCharacters *[]string `json:"retriggerCharacters"`
}

// Server Capabilities for a [DefinitionRequest](#DefinitionRequest).
type DefinitionOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Value-object that contains additional information when requesting references.
type ReferenceContext struct {

	// Include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}

// Reference options.
type ReferenceOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a [DocumentHighlightRequest](#DocumentHighlightRequest).
type DocumentHighlightOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// A base for all symbol information.
type BaseSymbolInformation struct {

	// The name of this symbol.
	Name string `json:"name"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this symbol.  @since 3.16.0
	Tags *[]SymbolTag `json:"tags"`

	// The name of the symbol containing this symbol. This information is
	// for user interface purposes (e.g. to render a qualifier in the user
	// interface if necessary). It can't be used to re-infer a hierarchy for
	// the document symbols.
	ContainerName *string `json:"containerName"`
}

// Provider options for a [DocumentSymbolRequest](#DocumentSymbolRequest).
type DocumentSymbolOptions struct {

	// mixins

	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.  @since 3.16.0
	Label *string `json:"label"`
}

// Contains additional diagnostic information about the context in which a [code
// action](#CodeActionProvider.provideCodeActions) is run.
type CodeActionContext struct {

	// An array of diagnostics known on the client side overlapping the
	// range provided to the `textDocument/codeAction` request. They are
	// provided so that the server knows which errors are currently
	// presented to the user for the given range. There is no guarantee that
	// these accurately reflect the error state of the resource. The primary
	// parameter to compute code actions is the provided range.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Requested kind of actions to return.  Actions not of this kind are
	// filtered out by the client before being shown. So servers can omit
	// computing them.
	Only *[]CodeActionKind `json:"only"`

	// The reason why code actions were requested.  @since 3.17.0
	TriggerKind *CodeActionTriggerKind `json:"triggerKind"`
}

// Provider options for a [CodeActionRequest](#CodeActionRequest).
type CodeActionOptions struct {

	// mixins

	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.  The list of kinds may
	// be generic, such as `CodeActionKind.Refactor`, or the server may list
	// out every specific kind they provide.
	CodeActionKinds *[]CodeActionKind `json:"codeActionKinds"`

	// The server provides support to resolve additional information for a
	// code action.  @since 3.16.0
	ResolveProvider *bool `json:"resolveProvider"`
}

// Server capabilities for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
type WorkspaceSymbolOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The server provides support to resolve additional information for a
	// workspace symbol.  @since 3.17.0
	ResolveProvider *bool `json:"resolveProvider"`
}

// Code Lens provider options of a [CodeLensRequest](#CodeLensRequest).
type CodeLensOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider"`
}

// Provider options for a [DocumentLinkRequest](#DocumentLinkRequest).
type DocumentLinkOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Document links have a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider"`
}

// Value-object describing what options formatting should use.
type FormattingOptions struct {

	// Size of a tab in spaces.
	TabSize uint64 `json:"tabSize"`

	// Prefer spaces over tabs.
	InsertSpaces bool `json:"insertSpaces"`

	// Trim trailing whitespace on a line.  @since 3.15.0
	TrimTrailingWhitespace *bool `json:"trimTrailingWhitespace"`

	// Insert a newline character at the end of the file if one does not
	// exist.  @since 3.15.0
	InsertFinalNewline *bool `json:"insertFinalNewline"`

	// Trim all newlines after the final newline at the end of the file.
	// @since 3.15.0
	TrimFinalNewlines *bool `json:"trimFinalNewlines"`
}

// Provider options for a
// [DocumentFormattingRequest](#DocumentFormattingRequest).
type DocumentFormattingOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a
// [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
type DocumentRangeFormattingOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a
// [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
type DocumentOnTypeFormattingOptions struct {

	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// More trigger characters.
	MoreTriggerCharacter *[]string `json:"moreTriggerCharacter"`
}

// Provider options for a [RenameRequest](#RenameRequest).
type RenameOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.  @since
	// version 3.12.0
	PrepareProvider *bool `json:"prepareProvider"`
}

// The server capabilities of a [ExecuteCommandRequest](#ExecuteCommandRequest).
type ExecuteCommandOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The commands to be executed on the server
	Commands []string `json:"commands"`
}

// @since 3.16.0
type SemanticTokensLegend struct {

	// The token types a server uses.
	TokenTypes []string `json:"tokenTypes"`

	// The token modifiers a server uses.
	TokenModifiers []string `json:"tokenModifiers"`
}

// A text document identifier to optionally denote a specific version of a text
// document.
type OptionalVersionedTextDocumentIdentifier struct {

	// extends

	TextDocumentIdentifier

	// The version number of this document. If a versioned text document
	// identifier is sent from the server to the client and the file is not
	// open in the editor (the server has not received an open notification
	// before) the server can send `null` to indicate that the version is
	// unknown and the content on disk is the truth (as specified with
	// document content ownership).
	Version OptionalVersionedTextDocumentIdentifier_Version__Or `json:"version"`
}

// A special text edit with an additional change annotation.  @since 3.16.0.
type AnnotatedTextEdit struct {

	// extends

	TextEdit

	// The actual identifier of the change annotation
	AnnotationId ChangeAnnotationIdentifier `json:"annotationId"`
}

// A generic resource operation.
type ResourceOperation struct {

	// The resource operation kind.
	Kind string `json:"kind"`

	// An optional annotation identifier describing the operation.  @since
	// 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId"`
}

// Options to create a file.
type CreateFileOptions struct {

	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite *bool `json:"overwrite"`

	// Ignore if exists.
	IgnoreIfExists *bool `json:"ignoreIfExists"`
}

// Rename file options
type RenameFileOptions struct {

	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite *bool `json:"overwrite"`

	// Ignores if target exists.
	IgnoreIfExists *bool `json:"ignoreIfExists"`
}

// Delete file options
type DeleteFileOptions struct {

	// Delete the content recursively if a folder is denoted.
	Recursive *bool `json:"recursive"`

	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists *bool `json:"ignoreIfNotExists"`
}

// A pattern to describe in which file operation requests or notifications the
// server is interested in receiving.  @since 3.16.0
type FileOperationPattern struct {

	// The glob pattern to match. Glob patterns can have the following
	// syntax: - `*` to match one or more characters in a path segment - `?`
	// to match on one character in a path segment - `**` to match any
	// number of path segments, including none - `{}` to group sub patterns
	// into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript
	// and JavaScript files) - `[]` to declare a range of characters to
	// match in a path segment (e.g., `example.[0-9]` to match on
	// `example.0`, `example.1`, …) - `[!...]` to negate a range of
	// characters to match in a path segment (e.g., `example.[!0-9]` to
	// match on `example.a`, `example.b`, but not `example.0`)
	Glob string `json:"glob"`

	// Whether to match files or folders with this pattern.  Matches both if
	// undefined.
	Matches *FileOperationPatternKind `json:"matches"`

	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options"`
}

// A full document diagnostic report for a workspace diagnostic result.  @since
// 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {

	// extends

	FullDocumentDiagnosticReport

	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`

	// The version number for which the diagnostics are reported. If the
	// document is not marked as open `null` can be provided.
	Version WorkspaceFullDocumentDiagnosticReport_Version__Or `json:"version"`
}

// An unchanged document diagnostic report for a workspace diagnostic result.
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {

	// extends

	UnchangedDocumentDiagnosticReport

	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`

	// The version number for which the diagnostics are reported. If the
	// document is not marked as open `null` can be provided.
	Version WorkspaceUnchangedDocumentDiagnosticReport_Version__Or `json:"version"`
}

// LSP object definition. @since 3.17.0
type LSPObject struct {
}

// A notebook cell.  A cell's document URI must be unique across ALL notebook
// cells and can therefore be used to uniquely identify a notebook cell or the
// cell's text document.  @since 3.17.0
type NotebookCell struct {

	// The cell's kind
	Kind NotebookCellKind `json:"kind"`

	// The URI of the cell's text document content.
	Document DocumentUri `json:"document"`

	// Additional metadata stored with the cell.  Note: should always be an
	// object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata"`

	// Additional execution summary information if supported by the client.
	ExecutionSummary *ExecutionSummary `json:"executionSummary"`
}

// A change describing how to move a `NotebookCell` array from state S to S'.
// @since 3.17.0
type NotebookCellArrayChange struct {

	// The start oftest of the cell that changed.
	Start uint64 `json:"start"`

	// The deleted cells
	DeleteCount uint64 `json:"deleteCount"`

	// The new cells, if any
	Cells *[]NotebookCell `json:"cells"`
}

// Defines the capabilities provided by the client.
type ClientCapabilities struct {

	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace"`

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument"`

	// Capabilities specific to the notebook document support.  @since
	// 3.17.0
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument"`

	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window"`

	// General client capabilities.  @since 3.16.0
	General *GeneralClientCapabilities `json:"general"`

	// Experimental client capabilities.
	Experimental *LSPAny `json:"experimental"`
}

type TextDocumentSyncOptions struct {

	// Open and close notifications are sent to the server. If omitted open
	// close notification should not be sent.
	OpenClose *bool `json:"openClose"`

	// Change notifications are sent to the server. See
	// TextDocumentSyncKind.None, TextDocumentSyncKind.Full and
	// TextDocumentSyncKind.Incremental. If omitted it defaults to
	// TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change"`

	// If present will save notifications are sent to the server. If omitted
	// the notification should not be sent.
	WillSave *bool `json:"willSave"`

	// If present will save wait until requests are sent to the server. If
	// omitted the request should not be sent.
	WillSaveWaitUntil *bool `json:"willSaveWaitUntil"`

	// If present save notifications are sent to the server. If omitted the
	// notification should not be sent.
	Save *TextDocumentSyncOptions_Save__Or `json:"save"`
}

// Options specific to a notebook plus its cells to be synced to the server.  If
// a selector provides a notebook document filter but no cell selector all cells
// of a matching notebook document will be synced.  If a selector provides no
// notebook document filter but only a cell selector all notebook document that
// contain at least one matching cell will be synced.  @since 3.17.0
type NotebookDocumentSyncOptions struct {

	// The notebooks to be synced
	NotebookSelector []NotebookDocumentSyncOptions_NotebookSelector_Element__Or `json:"notebookSelector"`

	// Whether save notification should be forwarded to the server. Will
	// only be honored if mode === `notebook`.
	Save *bool `json:"save"`
}

// Registration options specific to a notebook.  @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {

	// extends

	NotebookDocumentSyncOptions

	// mixins

	StaticRegistrationOptions
}

type WorkspaceFoldersServerCapabilities struct {

	// The server has support for workspace folders
	Supported *bool `json:"supported"`

	// Whether the server wants to receive workspace folder change
	// notifications.  If a string is provided the string is treated as an
	// ID under which the notification is registered on the client side. The
	// ID can be used to unregister for these events using the
	// `client/unregisterCapability` request.
	ChangeNotifications *WorkspaceFoldersServerCapabilities_ChangeNotifications__Or `json:"changeNotifications"`
}

// Options for notifications/requests for user operations on files.  @since
// 3.16.0
type FileOperationOptions struct {

	// The server is interested in receiving didCreateFiles notifications.
	DidCreate *FileOperationRegistrationOptions `json:"didCreate"`

	// The server is interested in receiving willCreateFiles requests.
	WillCreate *FileOperationRegistrationOptions `json:"willCreate"`

	// The server is interested in receiving didRenameFiles notifications.
	DidRename *FileOperationRegistrationOptions `json:"didRename"`

	// The server is interested in receiving willRenameFiles requests.
	WillRename *FileOperationRegistrationOptions `json:"willRename"`

	// The server is interested in receiving didDeleteFiles file
	// notifications.
	DidDelete *FileOperationRegistrationOptions `json:"didDelete"`

	// The server is interested in receiving willDeleteFiles file requests.
	WillDelete *FileOperationRegistrationOptions `json:"willDelete"`
}

// Structure to capture a description for an error code.  @since 3.16.0
type CodeDescription struct {

	// An URI to open with more information about the diagnostic error.
	Href URI `json:"href"`
}

// Represents a related message and source code location for a diagnostic. This
// should be used to point to code locations that cause or related to a
// diagnostics, e.g when duplicating a symbol in a scope.
type DiagnosticRelatedInformation struct {

	// The location of this related diagnostic information.
	Location Location `json:"location"`

	// The message of this related diagnostic information.
	Message string `json:"message"`
}

// Represents a parameter of a callable-signature. A parameter can have a label
// and a doc-comment.
type ParameterInformation struct {

	// The label of this parameter information.  Either a string or an
	// inclusive start and exclusive end offsets within its containing
	// signature label. (see SignatureInformation.label). The offsets are
	// based on a UTF-16 string representation as `Position` and `Range`
	// does.  *Note*: a label of type string should be a substring of its
	// containing signature label. Its intended use case is to highlight the
	// parameter label part in the `SignatureInformation.label`.
	Label ParameterInformation_Label__Or `json:"label"`

	// The human-readable doc-comment of this parameter. Will be shown in
	// the UI but can be omitted.
	Documentation *ParameterInformation_Documentation__Or `json:"documentation"`
}

// A notebook cell text document filter denotes a cell text document by
// different properties.  @since 3.17.0
type NotebookCellTextDocumentFilter struct {

	// A filter that matches against the notebook containing the notebook
	// cell. If a string value is provided it matches against the notebook
	// type. '*' matches every notebook.
	Notebook NotebookCellTextDocumentFilter_Notebook__Or `json:"notebook"`

	// A language id like `python`.  Will be matched against the language id
	// of the notebook cell document. '*' matches every language.
	Language *string `json:"language"`
}

// Matching options for the file operation pattern.  @since 3.16.0
type FileOperationPatternOptions struct {

	// The pattern should be matched ignoring casing.
	IgnoreCase *bool `json:"ignoreCase"`
}

type ExecutionSummary struct {

	// A strict monotonically increasing value indicating the execution
	// order of a cell inside a notebook.
	ExecutionOrder uint64 `json:"executionOrder"`

	// Whether the execution was successful or not if known by the client.
	Success *bool `json:"success"`
}

// Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {

	// The client supports applying batch edits to the workspace by
	// supporting the request 'workspace/applyEdit'
	ApplyEdit *bool `json:"applyEdit"`

	// Capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit"`

	// Capabilities specific to the `workspace/didChangeConfiguration`
	// notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration"`

	// Capabilities specific to the `workspace/didChangeWatchedFiles`
	// notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles"`

	// Capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol"`

	// Capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand"`

	// The client has support for workspace folders.  @since 3.6.0
	WorkspaceFolders *bool `json:"workspaceFolders"`

	// The client supports `workspace/configuration` requests.  @since 3.6.0
	Configuration *bool `json:"configuration"`

	// Capabilities specific to the semantic token requests scoped to the
	// workspace.  @since 3.16.0.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens"`

	// Capabilities specific to the code lens requests scoped to the
	// workspace.  @since 3.16.0.
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens"`

	// The client has support for file notifications/requests for user
	// operations on files.  Since 3.16.0
	FileOperations *FileOperationClientCapabilities `json:"fileOperations"`

	// Capabilities specific to the inline values requests scoped to the
	// workspace.  @since 3.17.0.
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue"`

	// Capabilities specific to the inlay hint requests scoped to the
	// workspace.  @since 3.17.0.
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint"`

	// Capabilities specific to the diagnostic requests scoped to the
	// workspace.  @since 3.17.0.
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics"`
}

// Text document specific client capabilities.
type TextDocumentClientCapabilities struct {

	// Defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization"`

	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion"`

	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover"`

	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp"`

	// Capabilities specific to the `textDocument/declaration` request.
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities `json:"declaration"`

	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition"`

	// Capabilities specific to the `textDocument/typeDefinition` request.
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition"`

	// Capabilities specific to the `textDocument/implementation` request.
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities `json:"implementation"`

	// Capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references"`

	// Capabilities specific to the `textDocument/documentHighlight`
	// request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight"`

	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol"`

	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction"`

	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens"`

	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink"`

	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.  @since 3.6.0
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider"`

	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting"`

	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting"`

	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting"`

	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename"`

	// Capabilities specific to the `textDocument/foldingRange` request.
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange"`

	// Capabilities specific to the `textDocument/selectionRange` request.
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange"`

	// Capabilities specific to the `textDocument/publishDiagnostics`
	// notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics"`

	// Capabilities specific to the various call hierarchy requests.  @since
	// 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy"`

	// Capabilities specific to the various semantic token request.  @since
	// 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens"`

	// Capabilities specific to the `textDocument/linkedEditingRange`
	// request.  @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange"`

	// Client capabilities specific to the `textDocument/moniker` request.
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker"`

	// Capabilities specific to the various type hierarchy requests.  @since
	// 3.17.0
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy"`

	// Capabilities specific to the `textDocument/inlineValue` request.
	// @since 3.17.0
	InlineValue *InlineValueClientCapabilities `json:"inlineValue"`

	// Capabilities specific to the `textDocument/inlayHint` request.
	// @since 3.17.0
	InlayHint *InlayHintClientCapabilities `json:"inlayHint"`

	// Capabilities specific to the diagnostic pull model.  @since 3.17.0
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic"`
}

// Capabilities specific to the notebook document support.  @since 3.17.0
type NotebookDocumentClientCapabilities struct {

	// Capabilities specific to notebook document synchronization  @since
	// 3.17.0
	Synchronization NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}

type WindowClientCapabilities struct {

	// It indicates whether the client supports server initiated progress
	// using the `window/workDoneProgress/create` request.  The capability
	// also controls Whether client supports handling of progress
	// notifications. If set servers are allowed to report a
	// `workDoneProgress` property in the request specific server
	// capabilities.  @since 3.15.0
	WorkDoneProgress *bool `json:"workDoneProgress"`

	// Capabilities specific to the showMessage request.  @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage"`

	// Capabilities specific to the showDocument request.  @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument"`
}

// General client capabilities.  @since 3.16.0
type GeneralClientCapabilities struct {

	// Client capability that signals how the client handles stale requests
	// (e.g. a request for which the client will not process the response
	// anymore since the information is outdated).  @since 3.17.0
	StaleRequestSupport *GeneralClientCapabilities_StaleRequestSupport `json:"staleRequestSupport"`

	// Client capabilities specific to regular expressions.  @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions"`

	// Client capabilities specific to the client's markdown parser.  @since
	// 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown"`

	// The position encodings supported by the client. Client and server
	// have to agree on the same position encoding to ensure that offsets
	// (e.g. character position in a line) are interpreted the same on both
	// sides.  To keep the protocol backwards compatible the following
	// applies: if the value 'utf-16' is missing from the array of position
	// encodings servers can assume that the client supports UTF-16. UTF-16
	// is therefore a mandatory encoding.  If omitted it defaults to
	// ['utf-16'].  Implementation considerations: since the conversion from
	// one encoding into another requires the content of the file / line the
	// conversion is best done where the file is read which is usually on
	// the server side.  @since 3.17.0
	PositionEncodings *[]PositionEncodingKind `json:"positionEncodings"`
}

// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.  @since 3.17.0
type RelativePattern struct {

	// A workspace folder or a base URI to which this pattern will be
	// matched against relatively.
	BaseUri RelativePattern_BaseUri__Or `json:"baseUri"`

	// The actual glob pattern;
	Pattern Pattern `json:"pattern"`
}

type WorkspaceEditClientCapabilities struct {

	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges *bool `json:"documentChanges"`

	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.  @since
	// 3.13.0
	ResourceOperations *[]ResourceOperationKind `json:"resourceOperations"`

	// The failure handling strategy of a client if applying the workspace
	// edit fails.  @since 3.13.0
	FailureHandling *FailureHandlingKind `json:"failureHandling"`

	// Whether the client normalizes line endings to the client specific
	// setting. If set to `true` the client will normalize line ending
	// characters in a workspace edit to the client-specified new line
	// character.  @since 3.16.0
	NormalizesLineEndings *bool `json:"normalizesLineEndings"`

	// Whether the client in general supports change annotations on text
	// edits, create file, rename file and delete file changes.  @since
	// 3.16.0
	ChangeAnnotationSupport *WorkspaceEditClientCapabilities_ChangeAnnotationSupport `json:"changeAnnotationSupport"`
}

type DidChangeConfigurationClientCapabilities struct {

	// Did change configuration notification supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DidChangeWatchedFilesClientCapabilities struct {

	// Did change watched files notification supports dynamic registration.
	// Please note that the current protocol doesn't support static
	// configuration for file changes from the server side.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Whether the client has support for {@link  RelativePattern relative
	// pattern} or not.  @since 3.17.0
	RelativePatternSupport *bool `json:"relativePatternSupport"`
}

// Client capabilities for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
type WorkspaceSymbolClientCapabilities struct {

	// Symbol request supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Specific capabilities for the `SymbolKind` in the `workspace/symbol`
	// request.
	SymbolKind *WorkspaceSymbolClientCapabilities_SymbolKind `json:"symbolKind"`

	// The client supports tags on `SymbolInformation`. Clients supporting
	// tags have to handle unknown tags gracefully.  @since 3.16.0
	TagSupport *WorkspaceSymbolClientCapabilities_TagSupport `json:"tagSupport"`

	// The client support partial workspace symbols. The client will send
	// the request `workspaceSymbol/resolve` to the server to resolve
	// additional properties.  @since 3.17.0
	ResolveSupport *WorkspaceSymbolClientCapabilities_ResolveSupport `json:"resolveSupport"`
}

// The client capabilities of a [ExecuteCommandRequest](#ExecuteCommandRequest).
type ExecuteCommandClientCapabilities struct {

	// Execute command supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all semantic tokens currently shown.
	// It should be used with absolute care and is useful for situation
	// where a server for example detects a project wide change that
	// requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport"`
}

// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all code lenses currently shown. It
	// should be used with absolute care and is useful for situation where a
	// server for example detect a project wide change that requires such a
	// calculation.
	RefreshSupport *bool `json:"refreshSupport"`
}

// Capabilities relating to events from file operations by the user in the
// client.  These events do not come from the file system, they come from user
// operations like renaming a file in the UI.  @since 3.16.0
type FileOperationClientCapabilities struct {

	// Whether the client supports dynamic registration for file
	// requests/notifications.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client has support for sending didCreateFiles notifications.
	DidCreate *bool `json:"didCreate"`

	// The client has support for sending willCreateFiles requests.
	WillCreate *bool `json:"willCreate"`

	// The client has support for sending didRenameFiles notifications.
	DidRename *bool `json:"didRename"`

	// The client has support for sending willRenameFiles requests.
	WillRename *bool `json:"willRename"`

	// The client has support for sending didDeleteFiles notifications.
	DidDelete *bool `json:"didDelete"`

	// The client has support for sending willDeleteFiles requests.
	WillDelete *bool `json:"willDelete"`
}

// Client workspace capabilities specific to inline values.  @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all inline values currently shown.
	// It should be used with absolute care and is useful for situation
	// where a server for example detects a project wide change that
	// requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport"`
}

// Client workspace capabilities specific to inlay hints.  @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all inlay hints currently shown. It
	// should be used with absolute care and is useful for situation where a
	// server for example detects a project wide change that requires such a
	// calculation.
	RefreshSupport *bool `json:"refreshSupport"`
}

// Workspace client capabilities specific to diagnostic pull requests.  @since
// 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all pulled diagnostics currently
	// shown. It should be used with absolute care and is useful for
	// situation where a server for example detects a project wide change
	// that requires such a calculation.
	RefreshSupport *bool `json:"refreshSupport"`
}

type TextDocumentSyncClientCapabilities struct {

	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports sending will save notifications.
	WillSave *bool `json:"willSave"`

	// The client supports sending a will save request and waits for a
	// response providing text edits which will be applied to the document
	// before it is saved.
	WillSaveWaitUntil *bool `json:"willSaveWaitUntil"`

	// The client supports did save notifications.
	DidSave *bool `json:"didSave"`
}

// Completion client capabilities
type CompletionClientCapabilities struct {

	// Whether completion supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem *CompletionClientCapabilities_CompletionItem `json:"completionItem"`

	CompletionItemKind *CompletionClientCapabilities_CompletionItemKind `json:"completionItemKind"`

	// Defines how the client handles whitespace and indentation when
	// accepting a completion item that uses multi line text in either
	// `insertText` or `textEdit`.  @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode"`

	// The client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport *bool `json:"contextSupport"`

	// The client supports the following `CompletionList` specific
	// capabilities.  @since 3.17.0
	CompletionList *CompletionClientCapabilities_CompletionList `json:"completionList"`
}

type HoverClientCapabilities struct {

	// Whether hover supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Client supports the following content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat *[]MarkupKind `json:"contentFormat"`
}

// Client Capabilities for a [SignatureHelpRequest](#SignatureHelpRequest).
type SignatureHelpClientCapabilities struct {

	// Whether signature help supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports the following `SignatureInformation` specific
	// properties.
	SignatureInformation *SignatureHelpClientCapabilities_SignatureInformation `json:"signatureInformation"`

	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.  @since 3.15.0
	ContextSupport *bool `json:"contextSupport"`
}

// @since 3.14.0
type DeclarationClientCapabilities struct {

	// Whether declaration supports dynamic registration. If this is set to
	// `true` the client supports the new `DeclarationRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of declaration
	// links.
	LinkSupport *bool `json:"linkSupport"`
}

// Client Capabilities for a [DefinitionRequest](#DefinitionRequest).
type DefinitionClientCapabilities struct {

	// Whether definition supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  @since 3.14.0
	LinkSupport *bool `json:"linkSupport"`
}

// Since 3.6.0
type TypeDefinitionClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `TypeDefinitionRegistrationOptions` return value for the
	// corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  Since 3.14.0
	LinkSupport *bool `json:"linkSupport"`
}

// @since 3.6.0
type ImplementationClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `ImplementationRegistrationOptions` return value for the
	// corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  @since 3.14.0
	LinkSupport *bool `json:"linkSupport"`
}

// Client Capabilities for a [ReferencesRequest](#ReferencesRequest).
type ReferenceClientCapabilities struct {

	// Whether references supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client Capabilities for a
// [DocumentHighlightRequest](#DocumentHighlightRequest).
type DocumentHighlightClientCapabilities struct {

	// Whether document highlight supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client Capabilities for a [DocumentSymbolRequest](#DocumentSymbolRequest).
type DocumentSymbolClientCapabilities struct {

	// Whether document symbol supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind *DocumentSymbolClientCapabilities_SymbolKind `json:"symbolKind"`

	// The client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport *bool `json:"hierarchicalDocumentSymbolSupport"`

	// The client supports tags on `SymbolInformation`. Tags are supported
	// on `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to
	// true. Clients supporting tags have to handle unknown tags gracefully.
	//  @since 3.16.0
	TagSupport *DocumentSymbolClientCapabilities_TagSupport `json:"tagSupport"`

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.  @since 3.16.0
	LabelSupport *bool `json:"labelSupport"`
}

// The Client Capabilities of a [CodeActionRequest](#CodeActionRequest).
type CodeActionClientCapabilities struct {

	// Whether code action supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client support code action literals of type `CodeAction` as a
	// valid response of the `textDocument/codeAction` request. If the
	// property is not set the request can only return `Command` literals.
	// @since 3.8.0
	CodeActionLiteralSupport *CodeActionClientCapabilities_CodeActionLiteralSupport `json:"codeActionLiteralSupport"`

	// Whether code action supports the `isPreferred` property.  @since
	// 3.15.0
	IsPreferredSupport *bool `json:"isPreferredSupport"`

	// Whether code action supports the `disabled` property.  @since 3.16.0
	DisabledSupport *bool `json:"disabledSupport"`

	// Whether code action supports the `data` property which is preserved
	// between a `textDocument/codeAction` and a `codeAction/resolve`
	// request.  @since 3.16.0
	DataSupport *bool `json:"dataSupport"`

	// Whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.  @since
	// 3.16.0
	ResolveSupport *CodeActionClientCapabilities_ResolveSupport `json:"resolveSupport"`

	// Whether the client honors the change annotations in text edits and
	// resource operations returned via the `CodeAction#edit` property by
	// for example presenting the workspace edit in the user interface and
	// asking for confirmation.  @since 3.16.0
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations"`
}

// The client capabilities  of a [CodeLensRequest](#CodeLensRequest).
type CodeLensClientCapabilities struct {

	// Whether code lens supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// The client capabilities of a [DocumentLinkRequest](#DocumentLinkRequest).
type DocumentLinkClientCapabilities struct {

	// Whether document link supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Whether the client supports the `tooltip` property on `DocumentLink`.
	//  @since 3.15.0
	TooltipSupport *bool `json:"tooltipSupport"`
}

type DocumentColorClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `DocumentColorRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client capabilities of a
// [DocumentFormattingRequest](#DocumentFormattingRequest).
type DocumentFormattingClientCapabilities struct {

	// Whether formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client capabilities of a
// [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
type DocumentRangeFormattingClientCapabilities struct {

	// Whether range formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client capabilities of a
// [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
type DocumentOnTypeFormattingClientCapabilities struct {

	// Whether on type formatting supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type RenameClientCapabilities struct {

	// Whether rename supports dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Client supports testing for validity of rename operations before
	// execution.  @since 3.12.0
	PrepareSupport *bool `json:"prepareSupport"`

	// Client supports the default behavior result.  The value indicates the
	// default behavior used by the client.  @since 3.16.0
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior"`

	// Whether the client honors the change annotations in text edits and
	// resource operations returned via the rename request's workspace edit
	// by for example presenting the workspace edit in the user interface
	// and asking for confirmation.  @since 3.16.0
	HonorsChangeAnnotations *bool `json:"honorsChangeAnnotations"`
}

type FoldingRangeClientCapabilities struct {

	// Whether implementation supports dynamic registration for folding
	// range providers. If this is set to `true` the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The maximum number of folding ranges that the client prefers to
	// receive per document. The value serves as a hint, servers are free to
	// follow the limit.
	RangeLimit *uint64 `json:"rangeLimit"`

	// If set, the client signals that it only supports folding complete
	// lines. If set, client will ignore specified `startCharacter` and
	// `endCharacter` properties in a FoldingRange.
	LineFoldingOnly *bool `json:"lineFoldingOnly"`

	// Specific options for the folding range kind.  @since 3.17.0
	FoldingRangeKind *FoldingRangeClientCapabilities_FoldingRangeKind `json:"foldingRangeKind"`

	// Specific options for the folding range.  @since 3.17.0
	FoldingRange *FoldingRangeClientCapabilities_FoldingRange `json:"foldingRange"`
}

type SelectionRangeClientCapabilities struct {

	// Whether implementation supports dynamic registration for selection
	// range providers. If this is set to `true` the client supports the new
	// `SelectionRangeRegistrationOptions` return value for the
	// corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// The publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {

	// Whether the clients accepts diagnostics with related information.
	RelatedInformation *bool `json:"relatedInformation"`

	// Client supports the tag property to provide meta data about a
	// diagnostic. Clients supporting tags have to handle unknown tags
	// gracefully.  @since 3.15.0
	TagSupport *PublishDiagnosticsClientCapabilities_TagSupport `json:"tagSupport"`

	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.  @since
	// 3.15.0
	VersionSupport *bool `json:"versionSupport"`

	// Client supports a codeDescription property  @since 3.16.0
	CodeDescriptionSupport *bool `json:"codeDescriptionSupport"`

	// Whether code action supports the `data` property which is preserved
	// between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.  @since 3.16.0
	DataSupport *bool `json:"dataSupport"`
}

// @since 3.16.0
type CallHierarchyClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// @since 3.16.0
type SemanticTokensClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might
	// not show semantic tokens or degrade some of the user experience if a
	// range or full request is advertised by the client but not provided by
	// the server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or
	// might even decide to not show any semantic tokens at all.
	Requests SemanticTokensClientCapabilities_Requests `json:"requests"`

	// The token types that the client supports.
	TokenTypes []string `json:"tokenTypes"`

	// The token modifiers that the client supports.
	TokenModifiers []string `json:"tokenModifiers"`

	// The token formats the clients supports.
	Formats []TokenFormat `json:"formats"`

	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport *bool `json:"overlappingTokenSupport"`

	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport *bool `json:"multilineTokenSupport"`

	// Whether the client allows the server to actively cancel a semantic
	// token request, e.g. supports returning LSPErrorCodes.ServerCancelled.
	// If a server does the client needs to retrigger the request.  @since
	// 3.17.0
	ServerCancelSupport *bool `json:"serverCancelSupport"`

	// Whether the client uses semantic tokens to augment existing syntax
	// tokens. If set to `true` client side created syntax tokens and
	// semantic tokens are both used for colorization. If set to `false` the
	// client only uses the returned semantic tokens for colorization.  If
	// the value is `undefined` then the client behavior is not specified.
	// @since 3.17.0
	AugmentsSyntaxTokens *bool `json:"augmentsSyntaxTokens"`
}

// Client capabilities for the linked editing range request.  @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client capabilities specific to the moniker request.  @since 3.16.0
type MonikerClientCapabilities struct {

	// Whether moniker supports dynamic registration. If this is set to
	// `true` the client supports the new `MonikerRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// @since 3.17.0
type TypeHierarchyClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Client capabilities specific to inline values.  @since 3.17.0
type InlineValueClientCapabilities struct {

	// Whether implementation supports dynamic registration for inline value
	// providers.
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

// Inlay hint client capabilities.  @since 3.17.0
type InlayHintClientCapabilities struct {

	// Whether inlay hints support dynamic registration.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Indicates which properties a client can resolve lazily on an inlay
	// hint.
	ResolveSupport *InlayHintClientCapabilities_ResolveSupport `json:"resolveSupport"`
}

// Client capabilities specific to diagnostic pull requests.  @since 3.17.0
type DiagnosticClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// Whether the clients supports related documents for document
	// diagnostic pulls.
	RelatedDocumentSupport *bool `json:"relatedDocumentSupport"`
}

// Notebook specific client capabilities.  @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// The client supports sending execution summary data per cell.
	ExecutionSummarySupport *bool `json:"executionSummarySupport"`
}

// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {

	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ShowMessageRequestClientCapabilities_MessageActionItem `json:"messageActionItem"`
}

// Client capabilities for the showDocument request.  @since 3.16.0
type ShowDocumentClientCapabilities struct {

	// The client has support for the showDocument request.
	Support bool `json:"support"`
}

// Client capabilities specific to regular expressions.  @since 3.16.0
type RegularExpressionsClientCapabilities struct {

	// The engine's name.
	Engine string `json:"engine"`

	// The engine's version.
	Version *string `json:"version"`
}

// Client capabilities specific to the used markdown parser.  @since 3.16.0
type MarkdownClientCapabilities struct {

	// The name of the parser.
	Parser string `json:"parser"`

	// The version of the parser.
	Version *string `json:"version"`

	// A list of HTML tags that the client allows / supports in Markdown.
	// @since 3.17.0
	AllowedTags *[]string `json:"allowedTags"`
}
