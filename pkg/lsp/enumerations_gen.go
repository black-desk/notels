// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"errors"
)

var EnumerationValidateFailed error = errors.New(
	"enumeration validate failed",
)

// A set of predefined token types. This set is not fixed an clients can specify
// additional token types via the corresponding client capabilities.  @since
// 3.16.0
type SemanticTokenTypes string

const (
	SemanticTokenTypesNamespace = "namespace"

	// Represents a generic type. Acts as a fallback for types which can't
	// be mapped to a specific type like class or enum.
	SemanticTokenTypesType = "type"

	SemanticTokenTypesClass = "class"

	SemanticTokenTypesEnum = "enum"

	SemanticTokenTypesInterface = "interface"

	SemanticTokenTypesStruct = "struct"

	SemanticTokenTypesTypeParameter = "typeParameter"

	SemanticTokenTypesParameter = "parameter"

	SemanticTokenTypesVariable = "variable"

	SemanticTokenTypesProperty = "property"

	SemanticTokenTypesEnumMember = "enumMember"

	SemanticTokenTypesEvent = "event"

	SemanticTokenTypesFunction = "function"

	SemanticTokenTypesMethod = "method"

	SemanticTokenTypesMacro = "macro"

	SemanticTokenTypesKeyword = "keyword"

	SemanticTokenTypesModifier = "modifier"

	SemanticTokenTypesComment = "comment"

	SemanticTokenTypesString = "string"

	SemanticTokenTypesNumber = "number"

	SemanticTokenTypesRegexp = "regexp"

	SemanticTokenTypesOperator = "operator"

	// @since 3.17.0
	SemanticTokenTypesDecorator = "decorator"
)

var _SemanticTokenTypes = []SemanticTokenTypes{

	SemanticTokenTypesNamespace,

	SemanticTokenTypesType,

	SemanticTokenTypesClass,

	SemanticTokenTypesEnum,

	SemanticTokenTypesInterface,

	SemanticTokenTypesStruct,

	SemanticTokenTypesTypeParameter,

	SemanticTokenTypesParameter,

	SemanticTokenTypesVariable,

	SemanticTokenTypesProperty,

	SemanticTokenTypesEnumMember,

	SemanticTokenTypesEvent,

	SemanticTokenTypesFunction,

	SemanticTokenTypesMethod,

	SemanticTokenTypesMacro,

	SemanticTokenTypesKeyword,

	SemanticTokenTypesModifier,

	SemanticTokenTypesComment,

	SemanticTokenTypesString,

	SemanticTokenTypesNumber,

	SemanticTokenTypesRegexp,

	SemanticTokenTypesOperator,

	SemanticTokenTypesDecorator,
}

func (this *SemanticTokenTypes) Validate() error {
	for _, x := range _SemanticTokenTypes {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A set of predefined token modifiers. This set is not fixed an clients can
// specify additional token types via the corresponding client capabilities.
// @since 3.16.0
type SemanticTokenModifiers string

const (
	SemanticTokenModifiersDeclaration = "declaration"

	SemanticTokenModifiersDefinition = "definition"

	SemanticTokenModifiersReadonly = "readonly"

	SemanticTokenModifiersStatic = "static"

	SemanticTokenModifiersDeprecated = "deprecated"

	SemanticTokenModifiersAbstract = "abstract"

	SemanticTokenModifiersAsync = "async"

	SemanticTokenModifiersModification = "modification"

	SemanticTokenModifiersDocumentation = "documentation"

	SemanticTokenModifiersDefaultLibrary = "defaultLibrary"
)

var _SemanticTokenModifiers = []SemanticTokenModifiers{

	SemanticTokenModifiersDeclaration,

	SemanticTokenModifiersDefinition,

	SemanticTokenModifiersReadonly,

	SemanticTokenModifiersStatic,

	SemanticTokenModifiersDeprecated,

	SemanticTokenModifiersAbstract,

	SemanticTokenModifiersAsync,

	SemanticTokenModifiersModification,

	SemanticTokenModifiersDocumentation,

	SemanticTokenModifiersDefaultLibrary,
}

func (this *SemanticTokenModifiers) Validate() error {
	for _, x := range _SemanticTokenModifiers {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The document diagnostic report kinds.  @since 3.17.0
type DocumentDiagnosticReportKind string

const (

	// A diagnostic report with a full set of problems.
	DocumentDiagnosticReportKindFull = "full"

	// A report indicating that the last returned report is still accurate.
	DocumentDiagnosticReportKindUnchanged = "unchanged"
)

var _DocumentDiagnosticReportKind = []DocumentDiagnosticReportKind{

	DocumentDiagnosticReportKindFull,

	DocumentDiagnosticReportKindUnchanged,
}

func (this *DocumentDiagnosticReportKind) Validate() error {
	for _, x := range _DocumentDiagnosticReportKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Predefined error codes.
type ErrorCodes int

const (
	ErrorCodesParseError = -32700

	ErrorCodesInvalidRequest = -32600

	ErrorCodesMethodNotFound = -32601

	ErrorCodesInvalidParams = -32602

	ErrorCodesInternalError = -32603

	// Error code indicating that a server received a notification or
	// request before the server has received the `initialize` request.
	ErrorCodesServerNotInitialized = -32002

	ErrorCodesUnknownErrorCode = -32001
)

var _ErrorCodes = []ErrorCodes{

	ErrorCodesParseError,

	ErrorCodesInvalidRequest,

	ErrorCodesMethodNotFound,

	ErrorCodesInvalidParams,

	ErrorCodesInternalError,

	ErrorCodesServerNotInitialized,

	ErrorCodesUnknownErrorCode,
}

func (this *ErrorCodes) Validate() error {
	for _, x := range _ErrorCodes {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type LSPErrorCodes int

const (

	// A request failed but it was syntactically correct, e.g the method
	// name was known and the parameters were valid. The error message
	// should contain human readable information about why the request
	// failed.  @since 3.17.0
	LSPErrorCodesRequestFailed = -32803

	// The server cancelled the request. This error code should only be used
	// for requests that explicitly support being server cancellable.
	// @since 3.17.0
	LSPErrorCodesServerCancelled = -32802

	// The server detected that the content of a document got modified
	// outside normal conditions. A server should NOT send this error code
	// if it detects a content change in it unprocessed messages. The result
	// even computed on an older state might still be useful for the client.
	//  If a client decides that a result is not of any use anymore the
	// client should cancel the request.
	LSPErrorCodesContentModified = -32801

	// The client has canceled a request and a server as detected the
	// cancel.
	LSPErrorCodesRequestCancelled = -32800
)

var _LSPErrorCodes = []LSPErrorCodes{

	LSPErrorCodesRequestFailed,

	LSPErrorCodesServerCancelled,

	LSPErrorCodesContentModified,

	LSPErrorCodesRequestCancelled,
}

func (this *LSPErrorCodes) Validate() error {
	for _, x := range _LSPErrorCodes {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A set of predefined range kinds.
type FoldingRangeKind string

const (

	// Folding range for a comment
	FoldingRangeKindComment = "comment"

	// Folding range for an import or include
	FoldingRangeKindImports = "imports"

	// Folding range for a region (e.g. `#region`)
	FoldingRangeKindRegion = "region"
)

var _FoldingRangeKind = []FoldingRangeKind{

	FoldingRangeKindComment,

	FoldingRangeKindImports,

	FoldingRangeKindRegion,
}

func (this *FoldingRangeKind) Validate() error {
	for _, x := range _FoldingRangeKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A symbol kind.
type SymbolKind uint

const (
	SymbolKindFile = 1

	SymbolKindModule = 2

	SymbolKindNamespace = 3

	SymbolKindPackage = 4

	SymbolKindClass = 5

	SymbolKindMethod = 6

	SymbolKindProperty = 7

	SymbolKindField = 8

	SymbolKindConstructor = 9

	SymbolKindEnum = 10

	SymbolKindInterface = 11

	SymbolKindFunction = 12

	SymbolKindVariable = 13

	SymbolKindConstant = 14

	SymbolKindString = 15

	SymbolKindNumber = 16

	SymbolKindBoolean = 17

	SymbolKindArray = 18

	SymbolKindObject = 19

	SymbolKindKey = 20

	SymbolKindNull = 21

	SymbolKindEnumMember = 22

	SymbolKindStruct = 23

	SymbolKindEvent = 24

	SymbolKindOperator = 25

	SymbolKindTypeParameter = 26
)

var _SymbolKind = []SymbolKind{

	SymbolKindFile,

	SymbolKindModule,

	SymbolKindNamespace,

	SymbolKindPackage,

	SymbolKindClass,

	SymbolKindMethod,

	SymbolKindProperty,

	SymbolKindField,

	SymbolKindConstructor,

	SymbolKindEnum,

	SymbolKindInterface,

	SymbolKindFunction,

	SymbolKindVariable,

	SymbolKindConstant,

	SymbolKindString,

	SymbolKindNumber,

	SymbolKindBoolean,

	SymbolKindArray,

	SymbolKindObject,

	SymbolKindKey,

	SymbolKindNull,

	SymbolKindEnumMember,

	SymbolKindStruct,

	SymbolKindEvent,

	SymbolKindOperator,

	SymbolKindTypeParameter,
}

func (this *SymbolKind) Validate() error {
	for _, x := range _SymbolKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Symbol tags are extra annotations that tweak the rendering of a symbol.
// @since 3.16
type SymbolTag uint

const (

	// Render a symbol as obsolete, usually using a strike-out.
	SymbolTagDeprecated = 1
)

var _SymbolTag = []SymbolTag{

	SymbolTagDeprecated,
}

func (this *SymbolTag) Validate() error {
	for _, x := range _SymbolTag {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Moniker uniqueness level to define scope of the moniker.  @since 3.16.0
type UniquenessLevel string

const (

	// The moniker is only unique inside a document
	UniquenessLevelDocument = "document"

	// The moniker is unique inside a project for which a dump got created
	UniquenessLevelProject = "project"

	// The moniker is unique inside the group to which a project belongs
	UniquenessLevelGroup = "group"

	// The moniker is unique inside the moniker scheme.
	UniquenessLevelScheme = "scheme"

	// The moniker is globally unique
	UniquenessLevelGlobal = "global"
)

var _UniquenessLevel = []UniquenessLevel{

	UniquenessLevelDocument,

	UniquenessLevelProject,

	UniquenessLevelGroup,

	UniquenessLevelScheme,

	UniquenessLevelGlobal,
}

func (this *UniquenessLevel) Validate() error {
	for _, x := range _UniquenessLevel {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The moniker kind.  @since 3.16.0
type MonikerKind string

const (

	// The moniker represent a symbol that is imported into a project
	MonikerKindImport = "import"

	// The moniker represents a symbol that is exported from a project
	MonikerKindExport = "export"

	// The moniker represents a symbol that is local to a project (e.g. a
	// local variable of a function, a class not visible outside the
	// project, ...)
	MonikerKindLocal = "local"
)

var _MonikerKind = []MonikerKind{

	MonikerKindImport,

	MonikerKindExport,

	MonikerKindLocal,
}

func (this *MonikerKind) Validate() error {
	for _, x := range _MonikerKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Inlay hint kinds.  @since 3.17.0
type InlayHintKind uint

const (

	// An inlay hint that for a type annotation.
	InlayHintKindType = 1

	// An inlay hint that is for a parameter.
	InlayHintKindParameter = 2
)

var _InlayHintKind = []InlayHintKind{

	InlayHintKindType,

	InlayHintKindParameter,
}

func (this *InlayHintKind) Validate() error {
	for _, x := range _InlayHintKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The message type
type MessageType uint

const (

	// An error message.
	MessageTypeError = 1

	// A warning message.
	MessageTypeWarning = 2

	// An information message.
	MessageTypeInfo = 3

	// A log message.
	MessageTypeLog = 4
)

var _MessageType = []MessageType{

	MessageTypeError,

	MessageTypeWarning,

	MessageTypeInfo,

	MessageTypeLog,
}

func (this *MessageType) Validate() error {
	for _, x := range _MessageType {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Defines how the host (editor) should sync document changes to the language
// server.
type TextDocumentSyncKind uint

const (

	// Documents should not be synced at all.
	TextDocumentSyncKindNone = 0

	// Documents are synced by always sending the full content of the
	// document.
	TextDocumentSyncKindFull = 1

	// Documents are synced by sending the full content on open. After that
	// only incremental updates to the document are send.
	TextDocumentSyncKindIncremental = 2
)

var _TextDocumentSyncKind = []TextDocumentSyncKind{

	TextDocumentSyncKindNone,

	TextDocumentSyncKindFull,

	TextDocumentSyncKindIncremental,
}

func (this *TextDocumentSyncKind) Validate() error {
	for _, x := range _TextDocumentSyncKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Represents reasons why a text document is saved.
type TextDocumentSaveReason uint

const (

	// Manually triggered, e.g. by the user pressing save, by starting
	// debugging, or by an API call.
	TextDocumentSaveReasonManual = 1

	// Automatic after a delay.
	TextDocumentSaveReasonAfterDelay = 2

	// When the editor lost focus.
	TextDocumentSaveReasonFocusOut = 3
)

var _TextDocumentSaveReason = []TextDocumentSaveReason{

	TextDocumentSaveReasonManual,

	TextDocumentSaveReasonAfterDelay,

	TextDocumentSaveReasonFocusOut,
}

func (this *TextDocumentSaveReason) Validate() error {
	for _, x := range _TextDocumentSaveReason {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The kind of a completion entry.
type CompletionItemKind uint

const (
	CompletionItemKindText = 1

	CompletionItemKindMethod = 2

	CompletionItemKindFunction = 3

	CompletionItemKindConstructor = 4

	CompletionItemKindField = 5

	CompletionItemKindVariable = 6

	CompletionItemKindClass = 7

	CompletionItemKindInterface = 8

	CompletionItemKindModule = 9

	CompletionItemKindProperty = 10

	CompletionItemKindUnit = 11

	CompletionItemKindValue = 12

	CompletionItemKindEnum = 13

	CompletionItemKindKeyword = 14

	CompletionItemKindSnippet = 15

	CompletionItemKindColor = 16

	CompletionItemKindFile = 17

	CompletionItemKindReference = 18

	CompletionItemKindFolder = 19

	CompletionItemKindEnumMember = 20

	CompletionItemKindConstant = 21

	CompletionItemKindStruct = 22

	CompletionItemKindEvent = 23

	CompletionItemKindOperator = 24

	CompletionItemKindTypeParameter = 25
)

var _CompletionItemKind = []CompletionItemKind{

	CompletionItemKindText,

	CompletionItemKindMethod,

	CompletionItemKindFunction,

	CompletionItemKindConstructor,

	CompletionItemKindField,

	CompletionItemKindVariable,

	CompletionItemKindClass,

	CompletionItemKindInterface,

	CompletionItemKindModule,

	CompletionItemKindProperty,

	CompletionItemKindUnit,

	CompletionItemKindValue,

	CompletionItemKindEnum,

	CompletionItemKindKeyword,

	CompletionItemKindSnippet,

	CompletionItemKindColor,

	CompletionItemKindFile,

	CompletionItemKindReference,

	CompletionItemKindFolder,

	CompletionItemKindEnumMember,

	CompletionItemKindConstant,

	CompletionItemKindStruct,

	CompletionItemKindEvent,

	CompletionItemKindOperator,

	CompletionItemKindTypeParameter,
}

func (this *CompletionItemKind) Validate() error {
	for _, x := range _CompletionItemKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Completion item tags are extra annotations that tweak the rendering of a
// completion item.  @since 3.15.0
type CompletionItemTag uint

const (

	// Render a completion as obsolete, usually using a strike-out.
	CompletionItemTagDeprecated = 1
)

var _CompletionItemTag = []CompletionItemTag{

	CompletionItemTagDeprecated,
}

func (this *CompletionItemTag) Validate() error {
	for _, x := range _CompletionItemTag {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
type InsertTextFormat uint

const (

	// The primary text to be inserted is treated as a plain string.
	InsertTextFormatPlainText = 1

	// The primary text to be inserted is treated as a snippet.  A snippet
	// can define tab stops and placeholders with `$1`, `$2` and `${3:foo}`.
	// `$0` defines the final tab stop, it defaults to the end of the
	// snippet. Placeholders with equal identifiers are linked, that is
	// typing in one will update others too.  See also:
	// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax
	InsertTextFormatSnippet = 2
)

var _InsertTextFormat = []InsertTextFormat{

	InsertTextFormatPlainText,

	InsertTextFormatSnippet,
}

func (this *InsertTextFormat) Validate() error {
	for _, x := range _InsertTextFormat {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// How whitespace and indentation is handled during completion item insertion.
// @since 3.16.0
type InsertTextMode uint

const (

	// The insertion or replace strings is taken as it is. If the value is
	// multi line the lines below the cursor will be inserted using the
	// indentation defined in the string value. The client will not apply
	// any kind of adjustments to the string.
	InsertTextModeAsIs = 1

	// The editor adjusts leading whitespace of new lines so that they match
	// the indentation up to the cursor of the line for which the item is
	// accepted.  Consider a line like this: <2tabs><cursor><3tabs>foo.
	// Accepting a multi line completion item is indented using 2 tabs and
	// all following lines inserted will be indented using 2 tabs as well.
	InsertTextModeAdjustIndentation = 2
)

var _InsertTextMode = []InsertTextMode{

	InsertTextModeAsIs,

	InsertTextModeAdjustIndentation,
}

func (this *InsertTextMode) Validate() error {
	for _, x := range _InsertTextMode {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A document highlight kind.
type DocumentHighlightKind uint

const (

	// A textual occurrence.
	DocumentHighlightKindText = 1

	// Read-access of a symbol, like reading a variable.
	DocumentHighlightKindRead = 2

	// Write-access of a symbol, like writing to a variable.
	DocumentHighlightKindWrite = 3
)

var _DocumentHighlightKind = []DocumentHighlightKind{

	DocumentHighlightKindText,

	DocumentHighlightKindRead,

	DocumentHighlightKindWrite,
}

func (this *DocumentHighlightKind) Validate() error {
	for _, x := range _DocumentHighlightKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A set of predefined code action kinds
type CodeActionKind string

const (

	// Empty kind.
	CodeActionKindEmpty = ""

	// Base kind for quickfix actions: 'quickfix'
	CodeActionKindQuickFix = "quickfix"

	// Base kind for refactoring actions: 'refactor'
	CodeActionKindRefactor = "refactor"

	// Base kind for refactoring extraction actions: 'refactor.extract'
	// Example extract actions:  - Extract method - Extract function -
	// Extract variable - Extract interface from class - ...
	CodeActionKindRefactorExtract = "refactor.extract"

	// Base kind for refactoring inline actions: 'refactor.inline'  Example
	// inline actions:  - Inline function - Inline variable - Inline
	// constant - ...
	CodeActionKindRefactorInline = "refactor.inline"

	// Base kind for refactoring rewrite actions: 'refactor.rewrite'
	// Example rewrite actions:  - Convert JavaScript function to class -
	// Add or remove parameter - Encapsulate field - Make method static -
	// Move method to base class - ...
	CodeActionKindRefactorRewrite = "refactor.rewrite"

	// Base kind for source actions: `source`  Source code actions apply to
	// the entire file.
	CodeActionKindSource = "source"

	// Base kind for an organize imports source action:
	// `source.organizeImports`
	CodeActionKindSourceOrganizeImports = "source.organizeImports"

	// Base kind for auto-fix source actions: `source.fixAll`.  Fix all
	// actions automatically fix errors that have a clear fix that do not
	// require user input. They should not suppress errors or perform unsafe
	// fixes such as generating new types or classes.  @since 3.15.0
	CodeActionKindSourceFixAll = "source.fixAll"
)

var _CodeActionKind = []CodeActionKind{

	CodeActionKindEmpty,

	CodeActionKindQuickFix,

	CodeActionKindRefactor,

	CodeActionKindRefactorExtract,

	CodeActionKindRefactorInline,

	CodeActionKindRefactorRewrite,

	CodeActionKindSource,

	CodeActionKindSourceOrganizeImports,

	CodeActionKindSourceFixAll,
}

func (this *CodeActionKind) Validate() error {
	for _, x := range _CodeActionKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type TraceValues string

const (

	// Turn tracing off.
	TraceValuesOff = "off"

	// Trace messages only.
	TraceValuesMessages = "messages"

	// Verbose message tracing.
	TraceValuesVerbose = "verbose"
)

var _TraceValues = []TraceValues{

	TraceValuesOff,

	TraceValuesMessages,

	TraceValuesVerbose,
}

func (this *TraceValues) Validate() error {
	for _, x := range _TraceValues {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// Describes the content type that a client supports in various result literals
// like `Hover`, `ParameterInfo` or `CompletionItem`.  Please note that
// `MarkupKinds` must not start with a `$`. This kinds are reserved for internal
// usage.
type MarkupKind string

const (

	// Plain text is supported as a content format
	MarkupKindPlainText = "plaintext"

	// Markdown is supported as a content format
	MarkupKindMarkdown = "markdown"
)

var _MarkupKind = []MarkupKind{

	MarkupKindPlainText,

	MarkupKindMarkdown,
}

func (this *MarkupKind) Validate() error {
	for _, x := range _MarkupKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A set of predefined position encoding kinds.  @since 3.17.0
type PositionEncodingKind string

const (

	// Character offsets count UTF-8 code units.
	PositionEncodingKindUTF8 = "utf-8"

	// Character offsets count UTF-16 code units.  This is the default and
	// must always be supported by servers
	PositionEncodingKindUTF16 = "utf-16"

	// Character offsets count UTF-32 code units.  Implementation note:
	// these are the same as Unicode code points, so this
	// `PositionEncodingKind` may also be used for an encoding-agnostic
	// representation of character offsets.
	PositionEncodingKindUTF32 = "utf-32"
)

var _PositionEncodingKind = []PositionEncodingKind{

	PositionEncodingKindUTF8,

	PositionEncodingKindUTF16,

	PositionEncodingKindUTF32,
}

func (this *PositionEncodingKind) Validate() error {
	for _, x := range _PositionEncodingKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The file event type
type FileChangeType uint

const (

	// The file got created.
	FileChangeTypeCreated = 1

	// The file got changed.
	FileChangeTypeChanged = 2

	// The file got deleted.
	FileChangeTypeDeleted = 3
)

var _FileChangeType = []FileChangeType{

	FileChangeTypeCreated,

	FileChangeTypeChanged,

	FileChangeTypeDeleted,
}

func (this *FileChangeType) Validate() error {
	for _, x := range _FileChangeType {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type WatchKind uint

const (

	// Interested in create events.
	WatchKindCreate = 1

	// Interested in change events
	WatchKindChange = 2

	// Interested in delete events
	WatchKindDelete = 4
)

var _WatchKind = []WatchKind{

	WatchKindCreate,

	WatchKindChange,

	WatchKindDelete,
}

func (this *WatchKind) Validate() error {
	for _, x := range _WatchKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The diagnostic's severity.
type DiagnosticSeverity uint

const (

	// Reports an error.
	DiagnosticSeverityError = 1

	// Reports a warning.
	DiagnosticSeverityWarning = 2

	// Reports an information.
	DiagnosticSeverityInformation = 3

	// Reports a hint.
	DiagnosticSeverityHint = 4
)

var _DiagnosticSeverity = []DiagnosticSeverity{

	DiagnosticSeverityError,

	DiagnosticSeverityWarning,

	DiagnosticSeverityInformation,

	DiagnosticSeverityHint,
}

func (this *DiagnosticSeverity) Validate() error {
	for _, x := range _DiagnosticSeverity {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The diagnostic tags.  @since 3.15.0
type DiagnosticTag uint

const (

	// Unused or unnecessary code.  Clients are allowed to render
	// diagnostics with this tag faded out instead of having an error
	// squiggle.
	DiagnosticTagUnnecessary = 1

	// Deprecated or obsolete code.  Clients are allowed to rendered
	// diagnostics with this tag strike through.
	DiagnosticTagDeprecated = 2
)

var _DiagnosticTag = []DiagnosticTag{

	DiagnosticTagUnnecessary,

	DiagnosticTagDeprecated,
}

func (this *DiagnosticTag) Validate() error {
	for _, x := range _DiagnosticTag {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// How a completion was triggered
type CompletionTriggerKind uint

const (

	// Completion was triggered by typing an identifier (24x7 code
	// complete), manual invocation (e.g Ctrl+Space) or via API.
	CompletionTriggerKindInvoked = 1

	// Completion was triggered by a trigger character specified by the
	// `triggerCharacters` properties of the
	// `CompletionRegistrationOptions`.
	CompletionTriggerKindTriggerCharacter = 2

	// Completion was re-triggered as current completion list is incomplete
	CompletionTriggerKindTriggerForIncompleteCompletions = 3
)

var _CompletionTriggerKind = []CompletionTriggerKind{

	CompletionTriggerKindInvoked,

	CompletionTriggerKindTriggerCharacter,

	CompletionTriggerKindTriggerForIncompleteCompletions,
}

func (this *CompletionTriggerKind) Validate() error {
	for _, x := range _CompletionTriggerKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// How a signature help was triggered.  @since 3.15.0
type SignatureHelpTriggerKind uint

const (

	// Signature help was invoked manually by the user or by a command.
	SignatureHelpTriggerKindInvoked = 1

	// Signature help was triggered by a trigger character.
	SignatureHelpTriggerKindTriggerCharacter = 2

	// Signature help was triggered by the cursor moving or by the document
	// content changing.
	SignatureHelpTriggerKindContentChange = 3
)

var _SignatureHelpTriggerKind = []SignatureHelpTriggerKind{

	SignatureHelpTriggerKindInvoked,

	SignatureHelpTriggerKindTriggerCharacter,

	SignatureHelpTriggerKindContentChange,
}

func (this *SignatureHelpTriggerKind) Validate() error {
	for _, x := range _SignatureHelpTriggerKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// The reason why code actions were requested.  @since 3.17.0
type CodeActionTriggerKind uint

const (

	// Code actions were explicitly requested by the user or by an
	// extension.
	CodeActionTriggerKindInvoked = 1

	// Code actions were requested automatically.  This typically happens
	// when current selection in a file changes, but can also be triggered
	// when file content changes.
	CodeActionTriggerKindAutomatic = 2
)

var _CodeActionTriggerKind = []CodeActionTriggerKind{

	CodeActionTriggerKindInvoked,

	CodeActionTriggerKindAutomatic,
}

func (this *CodeActionTriggerKind) Validate() error {
	for _, x := range _CodeActionTriggerKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A pattern kind describing if a glob pattern matches a file a folder or both.
// @since 3.16.0
type FileOperationPatternKind string

const (

	// The pattern matches a file only.
	FileOperationPatternKindFile = "file"

	// The pattern matches a folder only.
	FileOperationPatternKindFolder = "folder"
)

var _FileOperationPatternKind = []FileOperationPatternKind{

	FileOperationPatternKindFile,

	FileOperationPatternKindFolder,
}

func (this *FileOperationPatternKind) Validate() error {
	for _, x := range _FileOperationPatternKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

// A notebook cell kind.  @since 3.17.0
type NotebookCellKind uint

const (

	// A markup-cell is formatted source that is used for display.
	NotebookCellKindMarkup = 1

	// A code-cell is source code.
	NotebookCellKindCode = 2
)

var _NotebookCellKind = []NotebookCellKind{

	NotebookCellKindMarkup,

	NotebookCellKindCode,
}

func (this *NotebookCellKind) Validate() error {
	for _, x := range _NotebookCellKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type ResourceOperationKind string

const (

	// Supports creating new files and folders.
	ResourceOperationKindCreate = "create"

	// Supports renaming existing files and folders.
	ResourceOperationKindRename = "rename"

	// Supports deleting existing files and folders.
	ResourceOperationKindDelete = "delete"
)

var _ResourceOperationKind = []ResourceOperationKind{

	ResourceOperationKindCreate,

	ResourceOperationKindRename,

	ResourceOperationKindDelete,
}

func (this *ResourceOperationKind) Validate() error {
	for _, x := range _ResourceOperationKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type FailureHandlingKind string

const (

	// Applying the workspace change is simply aborted if one of the changes
	// provided fails. All operations executed before the failing operation
	// stay executed.
	FailureHandlingKindAbort = "abort"

	// All operations are executed transactional. That means they either all
	// succeed or no changes at all are applied to the workspace.
	FailureHandlingKindTransactional = "transactional"

	// If the workspace edit contains only textual file changes they are
	// executed transactional. If resource changes (create, rename or delete
	// file) are part of the change the failure handling strategy is abort.
	FailureHandlingKindTextOnlyTransactional = "textOnlyTransactional"

	// The client tries to undo the operations already executed. But there
	// is no guarantee that this is succeeding.
	FailureHandlingKindUndo = "undo"
)

var _FailureHandlingKind = []FailureHandlingKind{

	FailureHandlingKindAbort,

	FailureHandlingKindTransactional,

	FailureHandlingKindTextOnlyTransactional,

	FailureHandlingKindUndo,
}

func (this *FailureHandlingKind) Validate() error {
	for _, x := range _FailureHandlingKind {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type PrepareSupportDefaultBehavior uint

const (

	// The client's default behavior is to select the identifier according
	// the to language's syntax rule.
	PrepareSupportDefaultBehaviorIdentifier = 1
)

var _PrepareSupportDefaultBehavior = []PrepareSupportDefaultBehavior{

	PrepareSupportDefaultBehaviorIdentifier,
}

func (this *PrepareSupportDefaultBehavior) Validate() error {
	for _, x := range _PrepareSupportDefaultBehavior {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}

type TokenFormat string

const (
	TokenFormatRelative = "relative"
)

var _TokenFormat = []TokenFormat{

	TokenFormatRelative,
}

func (this *TokenFormat) Validate() error {
	for _, x := range _TokenFormat {
		if *this == x {
			return nil
		}
	}
	return EnumerationValidateFailed
}
