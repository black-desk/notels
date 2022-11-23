// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
	"encoding/json"
	"fmt"
)

var StructureValidateFailed = func(name string) error {
	return fmt.Errorf("structure \"%s\"validate failed", name)
}

type ImplementationParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// Represents a location inside a resource, such as a line inside a text file.
type Location struct {
	Uri *DocumentUri `json:"uri"`

	Range *Range `json:"range"`
}

func (this *Location) UnmarshalJSON(data []byte) error {
	type LocationUnmarshal Location
	var tmpUnmarshal LocationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("Location")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("Location")
	}

	*this = Location(tmpUnmarshal)
	return nil
}

func (this *Location) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("Location")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("Location")
	}

	type LocationMarshal Location
	tmpMarshal := LocationMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Uri *URI `json:"uri"`

	// The name of the workspace folder. Used to refer to this workspace
	// folder in the user interface.
	Name *String `json:"name"`
}

func (this *WorkspaceFolder) UnmarshalJSON(data []byte) error {
	type WorkspaceFolderUnmarshal WorkspaceFolder
	var tmpUnmarshal WorkspaceFolderUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("WorkspaceFolder")
	}

	if tmpUnmarshal.Name == nil {
		return StructureValidateFailed("WorkspaceFolder")
	}

	*this = WorkspaceFolder(tmpUnmarshal)
	return nil
}

func (this *WorkspaceFolder) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("WorkspaceFolder")
	}

	if this.Name == nil {
		return nil, StructureValidateFailed("WorkspaceFolder")
	}

	type WorkspaceFolderMarshal WorkspaceFolder
	tmpMarshal := WorkspaceFolderMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {

	// The actual workspace folder change event.
	Event *WorkspaceFoldersChangeEvent `json:"event"`
}

func (this *DidChangeWorkspaceFoldersParams) UnmarshalJSON(data []byte) error {
	type DidChangeWorkspaceFoldersParamsUnmarshal DidChangeWorkspaceFoldersParams
	var tmpUnmarshal DidChangeWorkspaceFoldersParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Event == nil {
		return StructureValidateFailed(
			"DidChangeWorkspaceFoldersParams",
		)
	}

	*this = DidChangeWorkspaceFoldersParams(tmpUnmarshal)
	return nil
}

func (this *DidChangeWorkspaceFoldersParams) MarshalJSON() ([]byte, error) {

	if this.Event == nil {
		return nil, StructureValidateFailed(
			"DidChangeWorkspaceFoldersParams",
		)
	}

	type DidChangeWorkspaceFoldersParamsMarshal DidChangeWorkspaceFoldersParams
	tmpMarshal := DidChangeWorkspaceFoldersParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters of a configuration request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}

func (this *ConfigurationParams) UnmarshalJSON(data []byte) error {
	type ConfigurationParamsUnmarshal ConfigurationParams
	var tmpUnmarshal ConfigurationParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Items == nil {
		return StructureValidateFailed("ConfigurationParams")
	}

	*this = ConfigurationParams(tmpUnmarshal)
	return nil
}

func (this *ConfigurationParams) MarshalJSON() ([]byte, error) {

	if this.Items == nil {
		return nil, StructureValidateFailed("ConfigurationParams")
	}

	type ConfigurationParamsMarshal ConfigurationParams
	tmpMarshal := ConfigurationParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Parameters for a {@link DocumentColorRequest}.
type DocumentColorParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *DocumentColorParams) UnmarshalJSON(data []byte) error {
	type DocumentColorParamsUnmarshal DocumentColorParams
	var tmpUnmarshal DocumentColorParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentColorParams")
	}

	*this = DocumentColorParams(tmpUnmarshal)
	return nil
}

func (this *DocumentColorParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DocumentColorParams")
	}

	type DocumentColorParamsMarshal DocumentColorParams
	tmpMarshal := DocumentColorParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a color range from a document.
type ColorInformation struct {

	// The range in the document where this color appears.
	Range *Range `json:"range"`

	// The actual color value for this color range.
	Color *Color `json:"color"`
}

func (this *ColorInformation) UnmarshalJSON(data []byte) error {
	type ColorInformationUnmarshal ColorInformation
	var tmpUnmarshal ColorInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("ColorInformation")
	}

	if tmpUnmarshal.Color == nil {
		return StructureValidateFailed("ColorInformation")
	}

	*this = ColorInformation(tmpUnmarshal)
	return nil
}

func (this *ColorInformation) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("ColorInformation")
	}

	if this.Color == nil {
		return nil, StructureValidateFailed("ColorInformation")
	}

	type ColorInformationMarshal ColorInformation
	tmpMarshal := ColorInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type DocumentColorRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentColorOptions

	// mixins

	StaticRegistrationOptions
}

// Parameters for a {@link ColorPresentationRequest}.
type ColorPresentationParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The color to request presentations for.
	Color *Color `json:"color"`

	// The range where the color would be inserted. Serves as a context.
	Range *Range `json:"range"`
}

func (this *ColorPresentationParams) UnmarshalJSON(data []byte) error {
	type ColorPresentationParamsUnmarshal ColorPresentationParams
	var tmpUnmarshal ColorPresentationParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("ColorPresentationParams")
	}

	if tmpUnmarshal.Color == nil {
		return StructureValidateFailed("ColorPresentationParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("ColorPresentationParams")
	}

	*this = ColorPresentationParams(tmpUnmarshal)
	return nil
}

func (this *ColorPresentationParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("ColorPresentationParams")
	}

	if this.Color == nil {
		return nil, StructureValidateFailed("ColorPresentationParams")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("ColorPresentationParams")
	}

	type ColorPresentationParamsMarshal ColorPresentationParams
	tmpMarshal := ColorPresentationParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type ColorPresentation struct {

	// The label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when
	// selecting this color presentation.
	Label *String `json:"label"`

	// An {@link TextEdit edit} which is applied to a document when
	// selecting this presentation for the color.  When `falsy` the {@link
	// ColorPresentation.label label} is used.
	TextEdit *TextEdit `json:"textEdit"`

	// An optional array of additional {@link TextEdit text edits} that are
	// applied when selecting this color presentation. Edits must not
	// overlap with the main {@link ColorPresentation.textEdit edit} nor
	// with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits"`
}

func (this *ColorPresentation) UnmarshalJSON(data []byte) error {
	type ColorPresentationUnmarshal ColorPresentation
	var tmpUnmarshal ColorPresentationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("ColorPresentation")
	}

	*this = ColorPresentation(tmpUnmarshal)
	return nil
}

func (this *ColorPresentation) MarshalJSON() ([]byte, error) {

	if this.Label == nil {
		return nil, StructureValidateFailed("ColorPresentation")
	}

	type ColorPresentationMarshal ColorPresentation
	tmpMarshal := ColorPresentationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *Boolean `json:"workDoneProgress"`
}

// General text document registration options.
type TextDocumentRegistrationOptions struct {

	// A document selector to identify the scope of the registration. If set
	// to null the document selector provided on the client side will be
	// used.
	DocumentSelector *TextDocumentRegistrationOptions_DocumentSelector_Or `json:"documentSelector"`
}

func (this *TextDocumentRegistrationOptions) UnmarshalJSON(data []byte) error {
	type TextDocumentRegistrationOptionsUnmarshal TextDocumentRegistrationOptions
	var tmpUnmarshal TextDocumentRegistrationOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.DocumentSelector == nil {
		return StructureValidateFailed(
			"TextDocumentRegistrationOptions",
		)
	}

	*this = TextDocumentRegistrationOptions(tmpUnmarshal)
	return nil
}

func (this *TextDocumentRegistrationOptions) MarshalJSON() ([]byte, error) {

	if this.DocumentSelector == nil {
		return nil, StructureValidateFailed(
			"TextDocumentRegistrationOptions",
		)
	}

	type TextDocumentRegistrationOptionsMarshal TextDocumentRegistrationOptions
	tmpMarshal := TextDocumentRegistrationOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Parameters for a {@link FoldingRangeRequest}.
type FoldingRangeParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *FoldingRangeParams) UnmarshalJSON(data []byte) error {
	type FoldingRangeParamsUnmarshal FoldingRangeParams
	var tmpUnmarshal FoldingRangeParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("FoldingRangeParams")
	}

	*this = FoldingRangeParams(tmpUnmarshal)
	return nil
}

func (this *FoldingRangeParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("FoldingRangeParams")
	}

	type FoldingRangeParamsMarshal FoldingRangeParams
	tmpMarshal := FoldingRangeParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a folding range. To be valid, start and end line must be bigger
// than zero and smaller than the number of lines in the document. Clients are
// free to ignore invalid ranges.
type FoldingRange struct {

	// The zero-based start line of the range to fold. The folded area
	// starts after the line's last character. To be valid, the end must be
	// zero or larger and smaller than the number of lines in the document.
	StartLine *Uinteger `json:"startLine"`

	// The zero-based character offset from where the folded range starts.
	// If not defined, defaults to the length of the start line.
	StartCharacter *Uinteger `json:"startCharacter"`

	// The zero-based end line of the range to fold. The folded area ends
	// with the line's last character. To be valid, the end must be zero or
	// larger and smaller than the number of lines in the document.
	EndLine *Uinteger `json:"endLine"`

	// The zero-based character offset before the folded range ends. If not
	// defined, defaults to the length of the end line.
	EndCharacter *Uinteger `json:"endCharacter"`

	// Describes the kind of the folding range such as `comment' or
	// 'region'. The kind is used to categorize folding ranges and used by
	// commands like 'Fold all comments'. See {@link FoldingRangeKind} for
	// an enumeration of standardized kinds.
	Kind *FoldingRangeKind `json:"kind"`

	// The text that the client should show when the specified range is
	// collapsed. If not defined or not supported by the client, a default
	// will be chosen by the client.  @since 3.17.0
	CollapsedText *String `json:"collapsedText"`
}

func (this *FoldingRange) UnmarshalJSON(data []byte) error {
	type FoldingRangeUnmarshal FoldingRange
	var tmpUnmarshal FoldingRangeUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.StartLine == nil {
		return StructureValidateFailed("FoldingRange")
	}

	if tmpUnmarshal.EndLine == nil {
		return StructureValidateFailed("FoldingRange")
	}

	*this = FoldingRange(tmpUnmarshal)
	return nil
}

func (this *FoldingRange) MarshalJSON() ([]byte, error) {

	if this.StartLine == nil {
		return nil, StructureValidateFailed("FoldingRange")
	}

	if this.EndLine == nil {
		return nil, StructureValidateFailed("FoldingRange")
	}

	type FoldingRangeMarshal FoldingRange
	tmpMarshal := FoldingRangeMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The positions inside the text document.
	Positions []Position `json:"positions"`
}

func (this *SelectionRangeParams) UnmarshalJSON(data []byte) error {
	type SelectionRangeParamsUnmarshal SelectionRangeParams
	var tmpUnmarshal SelectionRangeParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("SelectionRangeParams")
	}

	if tmpUnmarshal.Positions == nil {
		return StructureValidateFailed("SelectionRangeParams")
	}

	*this = SelectionRangeParams(tmpUnmarshal)
	return nil
}

func (this *SelectionRangeParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("SelectionRangeParams")
	}

	if this.Positions == nil {
		return nil, StructureValidateFailed("SelectionRangeParams")
	}

	type SelectionRangeParamsMarshal SelectionRangeParams
	tmpMarshal := SelectionRangeParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A selection range represents a part of a selection hierarchy. A selection
// range may have a parent selection range that contains it.
type SelectionRange struct {

	// The {@link Range range} of this selection range.
	Range *Range `json:"range"`

	// The parent selection range containing this range. Therefore
	// `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent"`
}

func (this *SelectionRange) UnmarshalJSON(data []byte) error {
	type SelectionRangeUnmarshal SelectionRange
	var tmpUnmarshal SelectionRangeUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("SelectionRange")
	}

	*this = SelectionRange(tmpUnmarshal)
	return nil
}

func (this *SelectionRange) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("SelectionRange")
	}

	type SelectionRangeMarshal SelectionRange
	tmpMarshal := SelectionRangeMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Token *ProgressToken `json:"token"`
}

func (this *WorkDoneProgressCreateParams) UnmarshalJSON(data []byte) error {
	type WorkDoneProgressCreateParamsUnmarshal WorkDoneProgressCreateParams
	var tmpUnmarshal WorkDoneProgressCreateParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Token == nil {
		return StructureValidateFailed("WorkDoneProgressCreateParams")
	}

	*this = WorkDoneProgressCreateParams(tmpUnmarshal)
	return nil
}

func (this *WorkDoneProgressCreateParams) MarshalJSON() ([]byte, error) {

	if this.Token == nil {
		return nil, StructureValidateFailed(
			"WorkDoneProgressCreateParams",
		)
	}

	type WorkDoneProgressCreateParamsMarshal WorkDoneProgressCreateParams
	tmpMarshal := WorkDoneProgressCreateParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressCancelParams struct {

	// The token to be used to report progress.
	Token *ProgressToken `json:"token"`
}

func (this *WorkDoneProgressCancelParams) UnmarshalJSON(data []byte) error {
	type WorkDoneProgressCancelParamsUnmarshal WorkDoneProgressCancelParams
	var tmpUnmarshal WorkDoneProgressCancelParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Token == nil {
		return StructureValidateFailed("WorkDoneProgressCancelParams")
	}

	*this = WorkDoneProgressCancelParams(tmpUnmarshal)
	return nil
}

func (this *WorkDoneProgressCancelParams) MarshalJSON() ([]byte, error) {

	if this.Token == nil {
		return nil, StructureValidateFailed(
			"WorkDoneProgressCancelParams",
		)
	}

	type WorkDoneProgressCancelParamsMarshal WorkDoneProgressCancelParams
	tmpMarshal := WorkDoneProgressCancelParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Name *String `json:"name"`

	// The kind of this item.
	Kind *SymbolKind `json:"kind"`

	// Tags for this item.
	Tags []SymbolTag `json:"tags"`

	// More detail for this item, e.g. the signature of a function.
	Detail *String `json:"detail"`

	// The resource identifier of this item.
	Uri *DocumentUri `json:"uri"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else, e.g. comments and code.
	Range *Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g. the name of a function. Must be contained by the
	// {@link CallHierarchyItem.range `range`}.
	SelectionRange *Range `json:"selectionRange"`

	// A data entry field that is preserved between a call hierarchy prepare
	// and incoming calls or outgoing calls requests.
	Data *LSPAny `json:"data"`
}

func (this *CallHierarchyItem) UnmarshalJSON(data []byte) error {
	type CallHierarchyItemUnmarshal CallHierarchyItem
	var tmpUnmarshal CallHierarchyItemUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureValidateFailed("CallHierarchyItem")
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("CallHierarchyItem")
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("CallHierarchyItem")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("CallHierarchyItem")
	}

	if tmpUnmarshal.SelectionRange == nil {
		return StructureValidateFailed("CallHierarchyItem")
	}

	*this = CallHierarchyItem(tmpUnmarshal)
	return nil
}

func (this *CallHierarchyItem) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureValidateFailed("CallHierarchyItem")
	}

	if this.Kind == nil {
		return nil, StructureValidateFailed("CallHierarchyItem")
	}

	if this.Uri == nil {
		return nil, StructureValidateFailed("CallHierarchyItem")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("CallHierarchyItem")
	}

	if this.SelectionRange == nil {
		return nil, StructureValidateFailed("CallHierarchyItem")
	}

	type CallHierarchyItemMarshal CallHierarchyItem
	tmpMarshal := CallHierarchyItemMarshal(*this)
	return json.Marshal(&tmpMarshal)
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

	Item *CallHierarchyItem `json:"item"`
}

func (this *CallHierarchyIncomingCallsParams) UnmarshalJSON(data []byte) error {
	type CallHierarchyIncomingCallsParamsUnmarshal CallHierarchyIncomingCallsParams
	var tmpUnmarshal CallHierarchyIncomingCallsParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Item == nil {
		return StructureValidateFailed(
			"CallHierarchyIncomingCallsParams",
		)
	}

	*this = CallHierarchyIncomingCallsParams(tmpUnmarshal)
	return nil
}

func (this *CallHierarchyIncomingCallsParams) MarshalJSON() ([]byte, error) {

	if this.Item == nil {
		return nil, StructureValidateFailed(
			"CallHierarchyIncomingCallsParams",
		)
	}

	type CallHierarchyIncomingCallsParamsMarshal CallHierarchyIncomingCallsParams
	tmpMarshal := CallHierarchyIncomingCallsParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents an incoming call, e.g. a caller of a method or constructor.
// @since 3.16.0
type CallHierarchyIncomingCall struct {

	// The item that makes the call.
	From *CallHierarchyItem `json:"from"`

	// The ranges at which the calls appear. This is relative to the caller
	// denoted by {@link CallHierarchyIncomingCall.from `this.from`}.
	FromRanges []Range `json:"fromRanges"`
}

func (this *CallHierarchyIncomingCall) UnmarshalJSON(data []byte) error {
	type CallHierarchyIncomingCallUnmarshal CallHierarchyIncomingCall
	var tmpUnmarshal CallHierarchyIncomingCallUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.From == nil {
		return StructureValidateFailed("CallHierarchyIncomingCall")
	}

	if tmpUnmarshal.FromRanges == nil {
		return StructureValidateFailed("CallHierarchyIncomingCall")
	}

	*this = CallHierarchyIncomingCall(tmpUnmarshal)
	return nil
}

func (this *CallHierarchyIncomingCall) MarshalJSON() ([]byte, error) {

	if this.From == nil {
		return nil, StructureValidateFailed("CallHierarchyIncomingCall")
	}

	if this.FromRanges == nil {
		return nil, StructureValidateFailed("CallHierarchyIncomingCall")
	}

	type CallHierarchyIncomingCallMarshal CallHierarchyIncomingCall
	tmpMarshal := CallHierarchyIncomingCallMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameter of a `callHierarchy/outgoingCalls` request.  @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item *CallHierarchyItem `json:"item"`
}

func (this *CallHierarchyOutgoingCallsParams) UnmarshalJSON(data []byte) error {
	type CallHierarchyOutgoingCallsParamsUnmarshal CallHierarchyOutgoingCallsParams
	var tmpUnmarshal CallHierarchyOutgoingCallsParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Item == nil {
		return StructureValidateFailed(
			"CallHierarchyOutgoingCallsParams",
		)
	}

	*this = CallHierarchyOutgoingCallsParams(tmpUnmarshal)
	return nil
}

func (this *CallHierarchyOutgoingCallsParams) MarshalJSON() ([]byte, error) {

	if this.Item == nil {
		return nil, StructureValidateFailed(
			"CallHierarchyOutgoingCallsParams",
		)
	}

	type CallHierarchyOutgoingCallsParamsMarshal CallHierarchyOutgoingCallsParams
	tmpMarshal := CallHierarchyOutgoingCallsParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents an outgoing call, e.g. calling a getter from a method or a method
// from a constructor etc.  @since 3.16.0
type CallHierarchyOutgoingCall struct {

	// The item that is called.
	To *CallHierarchyItem `json:"to"`

	// The range at which this item is called. This is the range relative to
	// the caller, e.g the item passed to {@link
	// CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls
	// `provideCallHierarchyOutgoingCalls`} and not {@link
	// CallHierarchyOutgoingCall.to `this.to`}.
	FromRanges []Range `json:"fromRanges"`
}

func (this *CallHierarchyOutgoingCall) UnmarshalJSON(data []byte) error {
	type CallHierarchyOutgoingCallUnmarshal CallHierarchyOutgoingCall
	var tmpUnmarshal CallHierarchyOutgoingCallUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.To == nil {
		return StructureValidateFailed("CallHierarchyOutgoingCall")
	}

	if tmpUnmarshal.FromRanges == nil {
		return StructureValidateFailed("CallHierarchyOutgoingCall")
	}

	*this = CallHierarchyOutgoingCall(tmpUnmarshal)
	return nil
}

func (this *CallHierarchyOutgoingCall) MarshalJSON() ([]byte, error) {

	if this.To == nil {
		return nil, StructureValidateFailed("CallHierarchyOutgoingCall")
	}

	if this.FromRanges == nil {
		return nil, StructureValidateFailed("CallHierarchyOutgoingCall")
	}

	type CallHierarchyOutgoingCallMarshal CallHierarchyOutgoingCall
	tmpMarshal := CallHierarchyOutgoingCallMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *SemanticTokensParams) UnmarshalJSON(data []byte) error {
	type SemanticTokensParamsUnmarshal SemanticTokensParams
	var tmpUnmarshal SemanticTokensParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("SemanticTokensParams")
	}

	*this = SemanticTokensParams(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("SemanticTokensParams")
	}

	type SemanticTokensParamsMarshal SemanticTokensParams
	tmpMarshal := SemanticTokensParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokens struct {

	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token
	// request. A server can then instead of computing all semantic tokens
	// again simply send a delta.
	ResultId *String `json:"resultId"`

	// The actual tokens.
	Data []Uinteger `json:"data"`
}

func (this *SemanticTokens) UnmarshalJSON(data []byte) error {
	type SemanticTokensUnmarshal SemanticTokens
	var tmpUnmarshal SemanticTokensUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Data == nil {
		return StructureValidateFailed("SemanticTokens")
	}

	*this = SemanticTokens(tmpUnmarshal)
	return nil
}

func (this *SemanticTokens) MarshalJSON() ([]byte, error) {

	if this.Data == nil {
		return nil, StructureValidateFailed("SemanticTokens")
	}

	type SemanticTokensMarshal SemanticTokens
	tmpMarshal := SemanticTokensMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensPartialResult struct {
	Data []Uinteger `json:"data"`
}

func (this *SemanticTokensPartialResult) UnmarshalJSON(data []byte) error {
	type SemanticTokensPartialResultUnmarshal SemanticTokensPartialResult
	var tmpUnmarshal SemanticTokensPartialResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Data == nil {
		return StructureValidateFailed("SemanticTokensPartialResult")
	}

	*this = SemanticTokensPartialResult(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensPartialResult) MarshalJSON() ([]byte, error) {

	if this.Data == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensPartialResult",
		)
	}

	type SemanticTokensPartialResultMarshal SemanticTokensPartialResult
	tmpMarshal := SemanticTokensPartialResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The result id of a previous response. The result Id can either point
	// to a full response or a delta response depending on what was received
	// last.
	PreviousResultId *String `json:"previousResultId"`
}

func (this *SemanticTokensDeltaParams) UnmarshalJSON(data []byte) error {
	type SemanticTokensDeltaParamsUnmarshal SemanticTokensDeltaParams
	var tmpUnmarshal SemanticTokensDeltaParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("SemanticTokensDeltaParams")
	}

	if tmpUnmarshal.PreviousResultId == nil {
		return StructureValidateFailed("SemanticTokensDeltaParams")
	}

	*this = SemanticTokensDeltaParams(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensDeltaParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("SemanticTokensDeltaParams")
	}

	if this.PreviousResultId == nil {
		return nil, StructureValidateFailed("SemanticTokensDeltaParams")
	}

	type SemanticTokensDeltaParamsMarshal SemanticTokensDeltaParams
	tmpMarshal := SemanticTokensDeltaParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensDelta struct {
	ResultId *String `json:"resultId"`

	// The semantic token edits to transform a previous result into a new
	// result.
	Edits []SemanticTokensEdit `json:"edits"`
}

func (this *SemanticTokensDelta) UnmarshalJSON(data []byte) error {
	type SemanticTokensDeltaUnmarshal SemanticTokensDelta
	var tmpUnmarshal SemanticTokensDeltaUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Edits == nil {
		return StructureValidateFailed("SemanticTokensDelta")
	}

	*this = SemanticTokensDelta(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensDelta) MarshalJSON() ([]byte, error) {

	if this.Edits == nil {
		return nil, StructureValidateFailed("SemanticTokensDelta")
	}

	type SemanticTokensDeltaMarshal SemanticTokensDelta
	tmpMarshal := SemanticTokensDeltaMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit `json:"edits"`
}

func (this *SemanticTokensDeltaPartialResult) UnmarshalJSON(data []byte) error {
	type SemanticTokensDeltaPartialResultUnmarshal SemanticTokensDeltaPartialResult
	var tmpUnmarshal SemanticTokensDeltaPartialResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Edits == nil {
		return StructureValidateFailed(
			"SemanticTokensDeltaPartialResult",
		)
	}

	*this = SemanticTokensDeltaPartialResult(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensDeltaPartialResult) MarshalJSON() ([]byte, error) {

	if this.Edits == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensDeltaPartialResult",
		)
	}

	type SemanticTokensDeltaPartialResultMarshal SemanticTokensDeltaPartialResult
	tmpMarshal := SemanticTokensDeltaPartialResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensRangeParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The range the semantic tokens are requested for.
	Range *Range `json:"range"`
}

func (this *SemanticTokensRangeParams) UnmarshalJSON(data []byte) error {
	type SemanticTokensRangeParamsUnmarshal SemanticTokensRangeParams
	var tmpUnmarshal SemanticTokensRangeParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("SemanticTokensRangeParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("SemanticTokensRangeParams")
	}

	*this = SemanticTokensRangeParams(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensRangeParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("SemanticTokensRangeParams")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("SemanticTokensRangeParams")
	}

	type SemanticTokensRangeParamsMarshal SemanticTokensRangeParams
	tmpMarshal := SemanticTokensRangeParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Params to show a document.  @since 3.16.0
type ShowDocumentParams struct {

	// The document uri to show.
	Uri *URI `json:"uri"`

	// Indicates to show the resource in an external program. To show for
	// example `https://code.visualstudio.com/` in the default WEB browser
	// set `external` to `true`.
	External *Boolean `json:"external"`

	// An optional property to indicate whether the editor showing the
	// document should take focus or not. Clients might ignore this property
	// if an external program is started.
	TakeFocus *Boolean `json:"takeFocus"`

	// An optional selection range if the document is a text document.
	// Clients might ignore the property if an external program is started
	// or the file is not a text file.
	Selection *Range `json:"selection"`
}

func (this *ShowDocumentParams) UnmarshalJSON(data []byte) error {
	type ShowDocumentParamsUnmarshal ShowDocumentParams
	var tmpUnmarshal ShowDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("ShowDocumentParams")
	}

	*this = ShowDocumentParams(tmpUnmarshal)
	return nil
}

func (this *ShowDocumentParams) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("ShowDocumentParams")
	}

	type ShowDocumentParamsMarshal ShowDocumentParams
	tmpMarshal := ShowDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The result of a showDocument request.  @since 3.16.0
type ShowDocumentResult struct {

	// A boolean indicating if the show was successful.
	Success *Boolean `json:"success"`
}

func (this *ShowDocumentResult) UnmarshalJSON(data []byte) error {
	type ShowDocumentResultUnmarshal ShowDocumentResult
	var tmpUnmarshal ShowDocumentResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Success == nil {
		return StructureValidateFailed("ShowDocumentResult")
	}

	*this = ShowDocumentResult(tmpUnmarshal)
	return nil
}

func (this *ShowDocumentResult) MarshalJSON() ([]byte, error) {

	if this.Success == nil {
		return nil, StructureValidateFailed("ShowDocumentResult")
	}

	type ShowDocumentResultMarshal ShowDocumentResult
	tmpMarshal := ShowDocumentResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	WordPattern *String `json:"wordPattern"`
}

func (this *LinkedEditingRanges) UnmarshalJSON(data []byte) error {
	type LinkedEditingRangesUnmarshal LinkedEditingRanges
	var tmpUnmarshal LinkedEditingRangesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Ranges == nil {
		return StructureValidateFailed("LinkedEditingRanges")
	}

	*this = LinkedEditingRanges(tmpUnmarshal)
	return nil
}

func (this *LinkedEditingRanges) MarshalJSON() ([]byte, error) {

	if this.Ranges == nil {
		return nil, StructureValidateFailed("LinkedEditingRanges")
	}

	type LinkedEditingRangesMarshal LinkedEditingRanges
	tmpMarshal := LinkedEditingRangesMarshal(*this)
	return json.Marshal(&tmpMarshal)
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

func (this *CreateFilesParams) UnmarshalJSON(data []byte) error {
	type CreateFilesParamsUnmarshal CreateFilesParams
	var tmpUnmarshal CreateFilesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Files == nil {
		return StructureValidateFailed("CreateFilesParams")
	}

	*this = CreateFilesParams(tmpUnmarshal)
	return nil
}

func (this *CreateFilesParams) MarshalJSON() ([]byte, error) {

	if this.Files == nil {
		return nil, StructureValidateFailed("CreateFilesParams")
	}

	type CreateFilesParamsMarshal CreateFilesParams
	tmpMarshal := CreateFilesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Changes map[DocumentUri][]TextEdit `json:"changes"`

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
	DocumentChanges []WorkspaceEdit_DocumentChanges_Element_Or `json:"documentChanges"`

	// A map of change annotations that can be referenced in
	// `AnnotatedTextEdit`s or create, rename and delete file / folder
	// operations.  Whether clients honor this property depends on the
	// client capability `workspace.changeAnnotationSupport`.  @since 3.16.0
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations"`
}

// The options to register for file operations.  @since 3.16.0
type FileOperationRegistrationOptions struct {

	// The actual filters.
	Filters []FileOperationFilter `json:"filters"`
}

func (this *FileOperationRegistrationOptions) UnmarshalJSON(data []byte) error {
	type FileOperationRegistrationOptionsUnmarshal FileOperationRegistrationOptions
	var tmpUnmarshal FileOperationRegistrationOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Filters == nil {
		return StructureValidateFailed(
			"FileOperationRegistrationOptions",
		)
	}

	*this = FileOperationRegistrationOptions(tmpUnmarshal)
	return nil
}

func (this *FileOperationRegistrationOptions) MarshalJSON() ([]byte, error) {

	if this.Filters == nil {
		return nil, StructureValidateFailed(
			"FileOperationRegistrationOptions",
		)
	}

	type FileOperationRegistrationOptionsMarshal FileOperationRegistrationOptions
	tmpMarshal := FileOperationRegistrationOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters sent in notifications/requests for user-initiated renames of
// files.  @since 3.16.0
type RenameFilesParams struct {

	// An array of all files/folders renamed in this operation. When a
	// folder is renamed, only the folder will be included, and not its
	// children.
	Files []FileRename `json:"files"`
}

func (this *RenameFilesParams) UnmarshalJSON(data []byte) error {
	type RenameFilesParamsUnmarshal RenameFilesParams
	var tmpUnmarshal RenameFilesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Files == nil {
		return StructureValidateFailed("RenameFilesParams")
	}

	*this = RenameFilesParams(tmpUnmarshal)
	return nil
}

func (this *RenameFilesParams) MarshalJSON() ([]byte, error) {

	if this.Files == nil {
		return nil, StructureValidateFailed("RenameFilesParams")
	}

	type RenameFilesParamsMarshal RenameFilesParams
	tmpMarshal := RenameFilesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters sent in notifications/requests for user-initiated deletes of
// files.  @since 3.16.0
type DeleteFilesParams struct {

	// An array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}

func (this *DeleteFilesParams) UnmarshalJSON(data []byte) error {
	type DeleteFilesParamsUnmarshal DeleteFilesParams
	var tmpUnmarshal DeleteFilesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Files == nil {
		return StructureValidateFailed("DeleteFilesParams")
	}

	*this = DeleteFilesParams(tmpUnmarshal)
	return nil
}

func (this *DeleteFilesParams) MarshalJSON() ([]byte, error) {

	if this.Files == nil {
		return nil, StructureValidateFailed("DeleteFilesParams")
	}

	type DeleteFilesParamsMarshal DeleteFilesParams
	tmpMarshal := DeleteFilesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Scheme *String `json:"scheme"`

	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier *String `json:"identifier"`

	// The scope in which the moniker is unique
	Unique *UniquenessLevel `json:"unique"`

	// The moniker kind if known.
	Kind *MonikerKind `json:"kind"`
}

func (this *Moniker) UnmarshalJSON(data []byte) error {
	type MonikerUnmarshal Moniker
	var tmpUnmarshal MonikerUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Scheme == nil {
		return StructureValidateFailed("Moniker")
	}

	if tmpUnmarshal.Identifier == nil {
		return StructureValidateFailed("Moniker")
	}

	if tmpUnmarshal.Unique == nil {
		return StructureValidateFailed("Moniker")
	}

	*this = Moniker(tmpUnmarshal)
	return nil
}

func (this *Moniker) MarshalJSON() ([]byte, error) {

	if this.Scheme == nil {
		return nil, StructureValidateFailed("Moniker")
	}

	if this.Identifier == nil {
		return nil, StructureValidateFailed("Moniker")
	}

	if this.Unique == nil {
		return nil, StructureValidateFailed("Moniker")
	}

	type MonikerMarshal Moniker
	tmpMarshal := MonikerMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Name *String `json:"name"`

	// The kind of this item.
	Kind *SymbolKind `json:"kind"`

	// Tags for this item.
	Tags []SymbolTag `json:"tags"`

	// More detail for this item, e.g. the signature of a function.
	Detail *String `json:"detail"`

	// The resource identifier of this item.
	Uri *DocumentUri `json:"uri"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else, e.g. comments and code.
	Range *Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g. the name of a function. Must be contained by the
	// {@link TypeHierarchyItem.range `range`}.
	SelectionRange *Range `json:"selectionRange"`

	// A data entry field that is preserved between a type hierarchy prepare
	// and supertypes or subtypes requests. It could also be used to
	// identify the type hierarchy in the server, helping improve the
	// performance on resolving supertypes and subtypes.
	Data *LSPAny `json:"data"`
}

func (this *TypeHierarchyItem) UnmarshalJSON(data []byte) error {
	type TypeHierarchyItemUnmarshal TypeHierarchyItem
	var tmpUnmarshal TypeHierarchyItemUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureValidateFailed("TypeHierarchyItem")
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("TypeHierarchyItem")
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("TypeHierarchyItem")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("TypeHierarchyItem")
	}

	if tmpUnmarshal.SelectionRange == nil {
		return StructureValidateFailed("TypeHierarchyItem")
	}

	*this = TypeHierarchyItem(tmpUnmarshal)
	return nil
}

func (this *TypeHierarchyItem) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureValidateFailed("TypeHierarchyItem")
	}

	if this.Kind == nil {
		return nil, StructureValidateFailed("TypeHierarchyItem")
	}

	if this.Uri == nil {
		return nil, StructureValidateFailed("TypeHierarchyItem")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("TypeHierarchyItem")
	}

	if this.SelectionRange == nil {
		return nil, StructureValidateFailed("TypeHierarchyItem")
	}

	type TypeHierarchyItemMarshal TypeHierarchyItem
	tmpMarshal := TypeHierarchyItemMarshal(*this)
	return json.Marshal(&tmpMarshal)
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

	Item *TypeHierarchyItem `json:"item"`
}

func (this *TypeHierarchySupertypesParams) UnmarshalJSON(data []byte) error {
	type TypeHierarchySupertypesParamsUnmarshal TypeHierarchySupertypesParams
	var tmpUnmarshal TypeHierarchySupertypesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Item == nil {
		return StructureValidateFailed("TypeHierarchySupertypesParams")
	}

	*this = TypeHierarchySupertypesParams(tmpUnmarshal)
	return nil
}

func (this *TypeHierarchySupertypesParams) MarshalJSON() ([]byte, error) {

	if this.Item == nil {
		return nil, StructureValidateFailed(
			"TypeHierarchySupertypesParams",
		)
	}

	type TypeHierarchySupertypesParamsMarshal TypeHierarchySupertypesParams
	tmpMarshal := TypeHierarchySupertypesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameter of a `typeHierarchy/subtypes` request.  @since 3.17.0
type TypeHierarchySubtypesParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Item *TypeHierarchyItem `json:"item"`
}

func (this *TypeHierarchySubtypesParams) UnmarshalJSON(data []byte) error {
	type TypeHierarchySubtypesParamsUnmarshal TypeHierarchySubtypesParams
	var tmpUnmarshal TypeHierarchySubtypesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Item == nil {
		return StructureValidateFailed("TypeHierarchySubtypesParams")
	}

	*this = TypeHierarchySubtypesParams(tmpUnmarshal)
	return nil
}

func (this *TypeHierarchySubtypesParams) MarshalJSON() ([]byte, error) {

	if this.Item == nil {
		return nil, StructureValidateFailed(
			"TypeHierarchySubtypesParams",
		)
	}

	type TypeHierarchySubtypesParamsMarshal TypeHierarchySubtypesParams
	tmpMarshal := TypeHierarchySubtypesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A parameter literal used in inline value requests.  @since 3.17.0
type InlineValueParams struct {

	// mixins

	WorkDoneProgressParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The document range for which inline values should be computed.
	Range *Range `json:"range"`

	// Additional information about the context in which inline values were
	// requested.
	Context *InlineValueContext `json:"context"`
}

func (this *InlineValueParams) UnmarshalJSON(data []byte) error {
	type InlineValueParamsUnmarshal InlineValueParams
	var tmpUnmarshal InlineValueParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("InlineValueParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("InlineValueParams")
	}

	if tmpUnmarshal.Context == nil {
		return StructureValidateFailed("InlineValueParams")
	}

	*this = InlineValueParams(tmpUnmarshal)
	return nil
}

func (this *InlineValueParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("InlineValueParams")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("InlineValueParams")
	}

	if this.Context == nil {
		return nil, StructureValidateFailed("InlineValueParams")
	}

	type InlineValueParamsMarshal InlineValueParams
	tmpMarshal := InlineValueParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The document range for which inlay hints should be computed.
	Range *Range `json:"range"`
}

func (this *InlayHintParams) UnmarshalJSON(data []byte) error {
	type InlayHintParamsUnmarshal InlayHintParams
	var tmpUnmarshal InlayHintParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("InlayHintParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("InlayHintParams")
	}

	*this = InlayHintParams(tmpUnmarshal)
	return nil
}

func (this *InlayHintParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("InlayHintParams")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("InlayHintParams")
	}

	type InlayHintParamsMarshal InlayHintParams
	tmpMarshal := InlayHintParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Inlay hint information.  @since 3.17.0
type InlayHint struct {

	// The position of this hint.
	Position *Position `json:"position"`

	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.  *Note* that neither the string nor
	// the label part can be empty.
	Label *InlayHint_Label_Or `json:"label"`

	// The kind of this hint. Can be omitted in which case the client should
	// fall back to a reasonable default.
	Kind *InlayHintKind `json:"kind"`

	// Optional text edits that are performed when accepting this inlay
	// hint.  *Note* that edits are expected to change the document so that
	// the inlay hint (or its nearest variant) is now part of the document
	// and the inlay hint itself is now obsolete.
	TextEdits []TextEdit `json:"textEdits"`

	// The tooltip text when you hover over this item.
	Tooltip *InlayHint_Tooltip_Or `json:"tooltip"`

	// Render padding before the hint.  Note: Padding should use the
	// editor's background color, not the background color of the hint
	// itself. That means padding can be used to visually align/separate an
	// inlay hint.
	PaddingLeft *Boolean `json:"paddingLeft"`

	// Render padding after the hint.  Note: Padding should use the editor's
	// background color, not the background color of the hint itself. That
	// means padding can be used to visually align/separate an inlay hint.
	PaddingRight *Boolean `json:"paddingRight"`

	// A data entry field that is preserved on an inlay hint between a
	// `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data *LSPAny `json:"data"`
}

func (this *InlayHint) UnmarshalJSON(data []byte) error {
	type InlayHintUnmarshal InlayHint
	var tmpUnmarshal InlayHintUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Position == nil {
		return StructureValidateFailed("InlayHint")
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("InlayHint")
	}

	*this = InlayHint(tmpUnmarshal)
	return nil
}

func (this *InlayHint) MarshalJSON() ([]byte, error) {

	if this.Position == nil {
		return nil, StructureValidateFailed("InlayHint")
	}

	if this.Label == nil {
		return nil, StructureValidateFailed("InlayHint")
	}

	type InlayHintMarshal InlayHint
	tmpMarshal := InlayHintMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The additional identifier  provided during registration.
	Identifier *String `json:"identifier"`

	// The result id of a previous response if provided.
	PreviousResultId *String `json:"previousResultId"`
}

func (this *DocumentDiagnosticParams) UnmarshalJSON(data []byte) error {
	type DocumentDiagnosticParamsUnmarshal DocumentDiagnosticParams
	var tmpUnmarshal DocumentDiagnosticParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentDiagnosticParams")
	}

	*this = DocumentDiagnosticParams(tmpUnmarshal)
	return nil
}

func (this *DocumentDiagnosticParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DocumentDiagnosticParams")
	}

	type DocumentDiagnosticParamsMarshal DocumentDiagnosticParams
	tmpMarshal := DocumentDiagnosticParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A partial result for a document diagnostic report.  @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {
	RelatedDocuments map[DocumentUri]DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or `json:"relatedDocuments"`
}

func (this *DocumentDiagnosticReportPartialResult) UnmarshalJSON(
	data []byte,
) error {
	type DocumentDiagnosticReportPartialResultUnmarshal DocumentDiagnosticReportPartialResult
	var tmpUnmarshal DocumentDiagnosticReportPartialResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.RelatedDocuments == nil {
		return StructureValidateFailed(
			"DocumentDiagnosticReportPartialResult",
		)
	}

	*this = DocumentDiagnosticReportPartialResult(tmpUnmarshal)
	return nil
}

func (this *DocumentDiagnosticReportPartialResult) MarshalJSON() ([]byte, error) {

	if this.RelatedDocuments == nil {
		return nil, StructureValidateFailed(
			"DocumentDiagnosticReportPartialResult",
		)
	}

	type DocumentDiagnosticReportPartialResultMarshal DocumentDiagnosticReportPartialResult
	tmpMarshal := DocumentDiagnosticReportPartialResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Cancellation data returned from a diagnostic request.  @since 3.17.0
type DiagnosticServerCancellationData struct {
	RetriggerRequest *Boolean `json:"retriggerRequest"`
}

func (this *DiagnosticServerCancellationData) UnmarshalJSON(data []byte) error {
	type DiagnosticServerCancellationDataUnmarshal DiagnosticServerCancellationData
	var tmpUnmarshal DiagnosticServerCancellationDataUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.RetriggerRequest == nil {
		return StructureValidateFailed(
			"DiagnosticServerCancellationData",
		)
	}

	*this = DiagnosticServerCancellationData(tmpUnmarshal)
	return nil
}

func (this *DiagnosticServerCancellationData) MarshalJSON() ([]byte, error) {

	if this.RetriggerRequest == nil {
		return nil, StructureValidateFailed(
			"DiagnosticServerCancellationData",
		)
	}

	type DiagnosticServerCancellationDataMarshal DiagnosticServerCancellationData
	tmpMarshal := DiagnosticServerCancellationDataMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Identifier *String `json:"identifier"`

	// The currently known diagnostic reports with their previous result
	// ids.
	PreviousResultIds []PreviousResultId `json:"previousResultIds"`
}

func (this *WorkspaceDiagnosticParams) UnmarshalJSON(data []byte) error {
	type WorkspaceDiagnosticParamsUnmarshal WorkspaceDiagnosticParams
	var tmpUnmarshal WorkspaceDiagnosticParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.PreviousResultIds == nil {
		return StructureValidateFailed("WorkspaceDiagnosticParams")
	}

	*this = WorkspaceDiagnosticParams(tmpUnmarshal)
	return nil
}

func (this *WorkspaceDiagnosticParams) MarshalJSON() ([]byte, error) {

	if this.PreviousResultIds == nil {
		return nil, StructureValidateFailed("WorkspaceDiagnosticParams")
	}

	type WorkspaceDiagnosticParamsMarshal WorkspaceDiagnosticParams
	tmpMarshal := WorkspaceDiagnosticParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A workspace diagnostic report.  @since 3.17.0
type WorkspaceDiagnosticReport struct {
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

func (this *WorkspaceDiagnosticReport) UnmarshalJSON(data []byte) error {
	type WorkspaceDiagnosticReportUnmarshal WorkspaceDiagnosticReport
	var tmpUnmarshal WorkspaceDiagnosticReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Items == nil {
		return StructureValidateFailed("WorkspaceDiagnosticReport")
	}

	*this = WorkspaceDiagnosticReport(tmpUnmarshal)
	return nil
}

func (this *WorkspaceDiagnosticReport) MarshalJSON() ([]byte, error) {

	if this.Items == nil {
		return nil, StructureValidateFailed("WorkspaceDiagnosticReport")
	}

	type WorkspaceDiagnosticReportMarshal WorkspaceDiagnosticReport
	tmpMarshal := WorkspaceDiagnosticReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A partial result for a workspace diagnostic report.  @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {
	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}

func (this *WorkspaceDiagnosticReportPartialResult) UnmarshalJSON(
	data []byte,
) error {
	type WorkspaceDiagnosticReportPartialResultUnmarshal WorkspaceDiagnosticReportPartialResult
	var tmpUnmarshal WorkspaceDiagnosticReportPartialResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Items == nil {
		return StructureValidateFailed(
			"WorkspaceDiagnosticReportPartialResult",
		)
	}

	*this = WorkspaceDiagnosticReportPartialResult(tmpUnmarshal)
	return nil
}

func (this *WorkspaceDiagnosticReportPartialResult) MarshalJSON() ([]byte, error) {

	if this.Items == nil {
		return nil, StructureValidateFailed(
			"WorkspaceDiagnosticReportPartialResult",
		)
	}

	type WorkspaceDiagnosticReportPartialResultMarshal WorkspaceDiagnosticReportPartialResult
	tmpMarshal := WorkspaceDiagnosticReportPartialResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The params sent in an open notebook document notification.  @since 3.17.0
type DidOpenNotebookDocumentParams struct {

	// The notebook document that got opened.
	NotebookDocument *NotebookDocument `json:"notebookDocument"`

	// The text documents that represent the content of a notebook cell.
	CellTextDocuments []TextDocumentItem `json:"cellTextDocuments"`
}

func (this *DidOpenNotebookDocumentParams) UnmarshalJSON(data []byte) error {
	type DidOpenNotebookDocumentParamsUnmarshal DidOpenNotebookDocumentParams
	var tmpUnmarshal DidOpenNotebookDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookDocument == nil {
		return StructureValidateFailed("DidOpenNotebookDocumentParams")
	}

	if tmpUnmarshal.CellTextDocuments == nil {
		return StructureValidateFailed("DidOpenNotebookDocumentParams")
	}

	*this = DidOpenNotebookDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidOpenNotebookDocumentParams) MarshalJSON() ([]byte, error) {

	if this.NotebookDocument == nil {
		return nil, StructureValidateFailed(
			"DidOpenNotebookDocumentParams",
		)
	}

	if this.CellTextDocuments == nil {
		return nil, StructureValidateFailed(
			"DidOpenNotebookDocumentParams",
		)
	}

	type DidOpenNotebookDocumentParamsMarshal DidOpenNotebookDocumentParams
	tmpMarshal := DidOpenNotebookDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The params sent in a change notebook document notification.  @since 3.17.0
type DidChangeNotebookDocumentParams struct {

	// The notebook document that did change. The version number points to
	// the version after all provided changes have been applied. If only the
	// text document content of a cell changes the notebook version doesn't
	// necessarily have to change.
	NotebookDocument *VersionedNotebookDocumentIdentifier `json:"notebookDocument"`

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
	Change *NotebookDocumentChangeEvent `json:"change"`
}

func (this *DidChangeNotebookDocumentParams) UnmarshalJSON(data []byte) error {
	type DidChangeNotebookDocumentParamsUnmarshal DidChangeNotebookDocumentParams
	var tmpUnmarshal DidChangeNotebookDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookDocument == nil {
		return StructureValidateFailed(
			"DidChangeNotebookDocumentParams",
		)
	}

	if tmpUnmarshal.Change == nil {
		return StructureValidateFailed(
			"DidChangeNotebookDocumentParams",
		)
	}

	*this = DidChangeNotebookDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidChangeNotebookDocumentParams) MarshalJSON() ([]byte, error) {

	if this.NotebookDocument == nil {
		return nil, StructureValidateFailed(
			"DidChangeNotebookDocumentParams",
		)
	}

	if this.Change == nil {
		return nil, StructureValidateFailed(
			"DidChangeNotebookDocumentParams",
		)
	}

	type DidChangeNotebookDocumentParamsMarshal DidChangeNotebookDocumentParams
	tmpMarshal := DidChangeNotebookDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The params sent in a save notebook document notification.  @since 3.17.0
type DidSaveNotebookDocumentParams struct {

	// The notebook document that got saved.
	NotebookDocument *NotebookDocumentIdentifier `json:"notebookDocument"`
}

func (this *DidSaveNotebookDocumentParams) UnmarshalJSON(data []byte) error {
	type DidSaveNotebookDocumentParamsUnmarshal DidSaveNotebookDocumentParams
	var tmpUnmarshal DidSaveNotebookDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookDocument == nil {
		return StructureValidateFailed("DidSaveNotebookDocumentParams")
	}

	*this = DidSaveNotebookDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidSaveNotebookDocumentParams) MarshalJSON() ([]byte, error) {

	if this.NotebookDocument == nil {
		return nil, StructureValidateFailed(
			"DidSaveNotebookDocumentParams",
		)
	}

	type DidSaveNotebookDocumentParamsMarshal DidSaveNotebookDocumentParams
	tmpMarshal := DidSaveNotebookDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The params sent in a close notebook document notification.  @since 3.17.0
type DidCloseNotebookDocumentParams struct {

	// The notebook document that got closed.
	NotebookDocument *NotebookDocumentIdentifier `json:"notebookDocument"`

	// The text documents that represent the content of a notebook cell that
	// got closed.
	CellTextDocuments []TextDocumentIdentifier `json:"cellTextDocuments"`
}

func (this *DidCloseNotebookDocumentParams) UnmarshalJSON(data []byte) error {
	type DidCloseNotebookDocumentParamsUnmarshal DidCloseNotebookDocumentParams
	var tmpUnmarshal DidCloseNotebookDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookDocument == nil {
		return StructureValidateFailed("DidCloseNotebookDocumentParams")
	}

	if tmpUnmarshal.CellTextDocuments == nil {
		return StructureValidateFailed("DidCloseNotebookDocumentParams")
	}

	*this = DidCloseNotebookDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidCloseNotebookDocumentParams) MarshalJSON() ([]byte, error) {

	if this.NotebookDocument == nil {
		return nil, StructureValidateFailed(
			"DidCloseNotebookDocumentParams",
		)
	}

	if this.CellTextDocuments == nil {
		return nil, StructureValidateFailed(
			"DidCloseNotebookDocumentParams",
		)
	}

	type DidCloseNotebookDocumentParamsMarshal DidCloseNotebookDocumentParams
	tmpMarshal := DidCloseNotebookDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

func (this *RegistrationParams) UnmarshalJSON(data []byte) error {
	type RegistrationParamsUnmarshal RegistrationParams
	var tmpUnmarshal RegistrationParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Registrations == nil {
		return StructureValidateFailed("RegistrationParams")
	}

	*this = RegistrationParams(tmpUnmarshal)
	return nil
}

func (this *RegistrationParams) MarshalJSON() ([]byte, error) {

	if this.Registrations == nil {
		return nil, StructureValidateFailed("RegistrationParams")
	}

	type RegistrationParamsMarshal RegistrationParams
	tmpMarshal := RegistrationParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type UnregistrationParams struct {
	Unregisterations []Unregistration `json:"unregisterations"`
}

func (this *UnregistrationParams) UnmarshalJSON(data []byte) error {
	type UnregistrationParamsUnmarshal UnregistrationParams
	var tmpUnmarshal UnregistrationParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Unregisterations == nil {
		return StructureValidateFailed("UnregistrationParams")
	}

	*this = UnregistrationParams(tmpUnmarshal)
	return nil
}

func (this *UnregistrationParams) MarshalJSON() ([]byte, error) {

	if this.Unregisterations == nil {
		return nil, StructureValidateFailed("UnregistrationParams")
	}

	type UnregistrationParamsMarshal UnregistrationParams
	tmpMarshal := UnregistrationParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type InitializeParams struct {

	// extends

	XInitializeParams

	WorkspaceFoldersInitializeParams
}

// The result returned from an initialize request.
type InitializeResult struct {

	// The capabilities the language server provides.
	Capabilities *ServerCapabilities `json:"capabilities"`

	// Information about the server.  @since 3.15.0
	ServerInfo *InitializeResult_ServerInfo `json:"serverInfo"`
}

func (this *InitializeResult) UnmarshalJSON(data []byte) error {
	type InitializeResultUnmarshal InitializeResult
	var tmpUnmarshal InitializeResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Capabilities == nil {
		return StructureValidateFailed("InitializeResult")
	}

	*this = InitializeResult(tmpUnmarshal)
	return nil
}

func (this *InitializeResult) MarshalJSON() ([]byte, error) {

	if this.Capabilities == nil {
		return nil, StructureValidateFailed("InitializeResult")
	}

	type InitializeResultMarshal InitializeResult
	tmpMarshal := InitializeResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The data type of the ResponseError if the initialize request fails.
type InitializeError struct {

	// Indicates whether the client execute the following retry logic: (1)
	// show the message provided by the ResponseError to the user (2) user
	// selects retry or cancel (3) if user selected retry the initialize
	// method is sent again.
	Retry *Boolean `json:"retry"`
}

func (this *InitializeError) UnmarshalJSON(data []byte) error {
	type InitializeErrorUnmarshal InitializeError
	var tmpUnmarshal InitializeErrorUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Retry == nil {
		return StructureValidateFailed("InitializeError")
	}

	*this = InitializeError(tmpUnmarshal)
	return nil
}

func (this *InitializeError) MarshalJSON() ([]byte, error) {

	if this.Retry == nil {
		return nil, StructureValidateFailed("InitializeError")
	}

	type InitializeErrorMarshal InitializeError
	tmpMarshal := InitializeErrorMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type InitializedParams struct {
}

// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {

	// The actual changed settings
	Settings *LSPAny `json:"settings"`
}

func (this *DidChangeConfigurationParams) UnmarshalJSON(data []byte) error {
	type DidChangeConfigurationParamsUnmarshal DidChangeConfigurationParams
	var tmpUnmarshal DidChangeConfigurationParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Settings == nil {
		return StructureValidateFailed("DidChangeConfigurationParams")
	}

	*this = DidChangeConfigurationParams(tmpUnmarshal)
	return nil
}

func (this *DidChangeConfigurationParams) MarshalJSON() ([]byte, error) {

	if this.Settings == nil {
		return nil, StructureValidateFailed(
			"DidChangeConfigurationParams",
		)
	}

	type DidChangeConfigurationParamsMarshal DidChangeConfigurationParams
	tmpMarshal := DidChangeConfigurationParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type DidChangeConfigurationRegistrationOptions struct {
	Section *DidChangeConfigurationRegistrationOptions_Section_Or `json:"section"`
}

// The parameters of a notification message.
type ShowMessageParams struct {

	// The message type. See {@link MessageType}
	Type *MessageType `json:"type"`

	// The actual message.
	Message *String `json:"message"`
}

func (this *ShowMessageParams) UnmarshalJSON(data []byte) error {
	type ShowMessageParamsUnmarshal ShowMessageParams
	var tmpUnmarshal ShowMessageParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Type == nil {
		return StructureValidateFailed("ShowMessageParams")
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("ShowMessageParams")
	}

	*this = ShowMessageParams(tmpUnmarshal)
	return nil
}

func (this *ShowMessageParams) MarshalJSON() ([]byte, error) {

	if this.Type == nil {
		return nil, StructureValidateFailed("ShowMessageParams")
	}

	if this.Message == nil {
		return nil, StructureValidateFailed("ShowMessageParams")
	}

	type ShowMessageParamsMarshal ShowMessageParams
	tmpMarshal := ShowMessageParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type ShowMessageRequestParams struct {

	// The message type. See {@link MessageType}
	Type *MessageType `json:"type"`

	// The actual message.
	Message *String `json:"message"`

	// The message action items to present.
	Actions []MessageActionItem `json:"actions"`
}

func (this *ShowMessageRequestParams) UnmarshalJSON(data []byte) error {
	type ShowMessageRequestParamsUnmarshal ShowMessageRequestParams
	var tmpUnmarshal ShowMessageRequestParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Type == nil {
		return StructureValidateFailed("ShowMessageRequestParams")
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("ShowMessageRequestParams")
	}

	*this = ShowMessageRequestParams(tmpUnmarshal)
	return nil
}

func (this *ShowMessageRequestParams) MarshalJSON() ([]byte, error) {

	if this.Type == nil {
		return nil, StructureValidateFailed("ShowMessageRequestParams")
	}

	if this.Message == nil {
		return nil, StructureValidateFailed("ShowMessageRequestParams")
	}

	type ShowMessageRequestParamsMarshal ShowMessageRequestParams
	tmpMarshal := ShowMessageRequestParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type MessageActionItem struct {

	// A short title like 'Retry', 'Open Log' etc.
	Title *String `json:"title"`
}

func (this *MessageActionItem) UnmarshalJSON(data []byte) error {
	type MessageActionItemUnmarshal MessageActionItem
	var tmpUnmarshal MessageActionItemUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Title == nil {
		return StructureValidateFailed("MessageActionItem")
	}

	*this = MessageActionItem(tmpUnmarshal)
	return nil
}

func (this *MessageActionItem) MarshalJSON() ([]byte, error) {

	if this.Title == nil {
		return nil, StructureValidateFailed("MessageActionItem")
	}

	type MessageActionItemMarshal MessageActionItem
	tmpMarshal := MessageActionItemMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The log message parameters.
type LogMessageParams struct {

	// The message type. See {@link MessageType}
	Type *MessageType `json:"type"`

	// The actual message.
	Message *String `json:"message"`
}

func (this *LogMessageParams) UnmarshalJSON(data []byte) error {
	type LogMessageParamsUnmarshal LogMessageParams
	var tmpUnmarshal LogMessageParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Type == nil {
		return StructureValidateFailed("LogMessageParams")
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("LogMessageParams")
	}

	*this = LogMessageParams(tmpUnmarshal)
	return nil
}

func (this *LogMessageParams) MarshalJSON() ([]byte, error) {

	if this.Type == nil {
		return nil, StructureValidateFailed("LogMessageParams")
	}

	if this.Message == nil {
		return nil, StructureValidateFailed("LogMessageParams")
	}

	type LogMessageParamsMarshal LogMessageParams
	tmpMarshal := LogMessageParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {

	// The document that was opened.
	TextDocument *TextDocumentItem `json:"textDocument"`
}

func (this *DidOpenTextDocumentParams) UnmarshalJSON(data []byte) error {
	type DidOpenTextDocumentParamsUnmarshal DidOpenTextDocumentParams
	var tmpUnmarshal DidOpenTextDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DidOpenTextDocumentParams")
	}

	*this = DidOpenTextDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidOpenTextDocumentParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DidOpenTextDocumentParams")
	}

	type DidOpenTextDocumentParamsMarshal DidOpenTextDocumentParams
	tmpMarshal := DidOpenTextDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The change text document notification's parameters.
type DidChangeTextDocumentParams struct {

	// The document that did change. The version number points to the
	// version after all provided content changes have been applied.
	TextDocument *VersionedTextDocumentIdentifier `json:"textDocument"`

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

func (this *DidChangeTextDocumentParams) UnmarshalJSON(data []byte) error {
	type DidChangeTextDocumentParamsUnmarshal DidChangeTextDocumentParams
	var tmpUnmarshal DidChangeTextDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DidChangeTextDocumentParams")
	}

	if tmpUnmarshal.ContentChanges == nil {
		return StructureValidateFailed("DidChangeTextDocumentParams")
	}

	*this = DidChangeTextDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidChangeTextDocumentParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"DidChangeTextDocumentParams",
		)
	}

	if this.ContentChanges == nil {
		return nil, StructureValidateFailed(
			"DidChangeTextDocumentParams",
		)
	}

	type DidChangeTextDocumentParamsMarshal DidChangeTextDocumentParams
	tmpMarshal := DidChangeTextDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	// How documents are synced to the server.
	SyncKind *TextDocumentSyncKind `json:"syncKind"`
}

func (this *TextDocumentChangeRegistrationOptions) UnmarshalJSON(
	data []byte,
) error {
	type TextDocumentChangeRegistrationOptionsUnmarshal TextDocumentChangeRegistrationOptions
	var tmpUnmarshal TextDocumentChangeRegistrationOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.SyncKind == nil {
		return StructureValidateFailed(
			"TextDocumentChangeRegistrationOptions",
		)
	}

	*this = TextDocumentChangeRegistrationOptions(tmpUnmarshal)
	return nil
}

func (this *TextDocumentChangeRegistrationOptions) MarshalJSON() ([]byte, error) {

	if this.SyncKind == nil {
		return nil, StructureValidateFailed(
			"TextDocumentChangeRegistrationOptions",
		)
	}

	type TextDocumentChangeRegistrationOptionsMarshal TextDocumentChangeRegistrationOptions
	tmpMarshal := TextDocumentChangeRegistrationOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {

	// The document that was closed.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *DidCloseTextDocumentParams) UnmarshalJSON(data []byte) error {
	type DidCloseTextDocumentParamsUnmarshal DidCloseTextDocumentParams
	var tmpUnmarshal DidCloseTextDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DidCloseTextDocumentParams")
	}

	*this = DidCloseTextDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidCloseTextDocumentParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"DidCloseTextDocumentParams",
		)
	}

	type DidCloseTextDocumentParamsMarshal DidCloseTextDocumentParams
	tmpMarshal := DidCloseTextDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {

	// The document that was saved.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text *String `json:"text"`
}

func (this *DidSaveTextDocumentParams) UnmarshalJSON(data []byte) error {
	type DidSaveTextDocumentParamsUnmarshal DidSaveTextDocumentParams
	var tmpUnmarshal DidSaveTextDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DidSaveTextDocumentParams")
	}

	*this = DidSaveTextDocumentParams(tmpUnmarshal)
	return nil
}

func (this *DidSaveTextDocumentParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DidSaveTextDocumentParams")
	}

	type DidSaveTextDocumentParamsMarshal DidSaveTextDocumentParams
	tmpMarshal := DidSaveTextDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The 'TextDocumentSaveReason'.
	Reason *TextDocumentSaveReason `json:"reason"`
}

func (this *WillSaveTextDocumentParams) UnmarshalJSON(data []byte) error {
	type WillSaveTextDocumentParamsUnmarshal WillSaveTextDocumentParams
	var tmpUnmarshal WillSaveTextDocumentParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("WillSaveTextDocumentParams")
	}

	if tmpUnmarshal.Reason == nil {
		return StructureValidateFailed("WillSaveTextDocumentParams")
	}

	*this = WillSaveTextDocumentParams(tmpUnmarshal)
	return nil
}

func (this *WillSaveTextDocumentParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"WillSaveTextDocumentParams",
		)
	}

	if this.Reason == nil {
		return nil, StructureValidateFailed(
			"WillSaveTextDocumentParams",
		)
	}

	type WillSaveTextDocumentParamsMarshal WillSaveTextDocumentParams
	tmpMarshal := WillSaveTextDocumentParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A text edit applicable to a text document.
type TextEdit struct {

	// The range of the text document to be manipulated. To insert text into
	// a document create a range where start === end.
	Range *Range `json:"range"`

	// The string to be inserted. For delete operations use an empty string.
	NewText *String `json:"newText"`
}

func (this *TextEdit) UnmarshalJSON(data []byte) error {
	type TextEditUnmarshal TextEdit
	var tmpUnmarshal TextEditUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("TextEdit")
	}

	if tmpUnmarshal.NewText == nil {
		return StructureValidateFailed("TextEdit")
	}

	*this = TextEdit(tmpUnmarshal)
	return nil
}

func (this *TextEdit) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("TextEdit")
	}

	if this.NewText == nil {
		return nil, StructureValidateFailed("TextEdit")
	}

	type TextEditMarshal TextEdit
	tmpMarshal := TextEditMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {

	// The actual file events.
	Changes []FileEvent `json:"changes"`
}

func (this *DidChangeWatchedFilesParams) UnmarshalJSON(data []byte) error {
	type DidChangeWatchedFilesParamsUnmarshal DidChangeWatchedFilesParams
	var tmpUnmarshal DidChangeWatchedFilesParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Changes == nil {
		return StructureValidateFailed("DidChangeWatchedFilesParams")
	}

	*this = DidChangeWatchedFilesParams(tmpUnmarshal)
	return nil
}

func (this *DidChangeWatchedFilesParams) MarshalJSON() ([]byte, error) {

	if this.Changes == nil {
		return nil, StructureValidateFailed(
			"DidChangeWatchedFilesParams",
		)
	}

	type DidChangeWatchedFilesParamsMarshal DidChangeWatchedFilesParams
	tmpMarshal := DidChangeWatchedFilesParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {

	// The watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}

func (this *DidChangeWatchedFilesRegistrationOptions) UnmarshalJSON(
	data []byte,
) error {
	type DidChangeWatchedFilesRegistrationOptionsUnmarshal DidChangeWatchedFilesRegistrationOptions
	var tmpUnmarshal DidChangeWatchedFilesRegistrationOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Watchers == nil {
		return StructureValidateFailed(
			"DidChangeWatchedFilesRegistrationOptions",
		)
	}

	*this = DidChangeWatchedFilesRegistrationOptions(tmpUnmarshal)
	return nil
}

func (this *DidChangeWatchedFilesRegistrationOptions) MarshalJSON() ([]byte, error) {

	if this.Watchers == nil {
		return nil, StructureValidateFailed(
			"DidChangeWatchedFilesRegistrationOptions",
		)
	}

	type DidChangeWatchedFilesRegistrationOptionsMarshal DidChangeWatchedFilesRegistrationOptions
	tmpMarshal := DidChangeWatchedFilesRegistrationOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {

	// The URI for which diagnostic information is reported.
	Uri *DocumentUri `json:"uri"`

	// Optional the version number of the document the diagnostics are
	// published for.  @since 3.15.0
	Version *Integer `json:"version"`

	// An array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

func (this *PublishDiagnosticsParams) UnmarshalJSON(data []byte) error {
	type PublishDiagnosticsParamsUnmarshal PublishDiagnosticsParams
	var tmpUnmarshal PublishDiagnosticsParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("PublishDiagnosticsParams")
	}

	if tmpUnmarshal.Diagnostics == nil {
		return StructureValidateFailed("PublishDiagnosticsParams")
	}

	*this = PublishDiagnosticsParams(tmpUnmarshal)
	return nil
}

func (this *PublishDiagnosticsParams) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("PublishDiagnosticsParams")
	}

	if this.Diagnostics == nil {
		return nil, StructureValidateFailed("PublishDiagnosticsParams")
	}

	type PublishDiagnosticsParamsMarshal PublishDiagnosticsParams
	tmpMarshal := PublishDiagnosticsParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Label *String `json:"label"`

	// Additional details for the label  @since 3.17.0
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails"`

	// The kind of this completion item. Based of the kind an icon is chosen
	// by the editor.
	Kind *CompletionItemKind `json:"kind"`

	// Tags for this completion item.  @since 3.15.0
	Tags []CompletionItemTag `json:"tags"`

	// A human-readable string with additional information about this item,
	// like type or symbol information.
	Detail *String `json:"detail"`

	// A human-readable string that represents a doc-comment.
	Documentation *CompletionItem_Documentation_Or `json:"documentation"`

	// Indicates if this item is deprecated. @deprecated Use `tags` instead.
	Deprecated *Boolean `json:"deprecated"`

	// Select this item when showing.  *Note* that only one completion item
	// can be selected and that the tool / client decides which item that
	// is. The rule is that the *first* item of those that match best is
	// selected.
	Preselect *Boolean `json:"preselect"`

	// A string that should be used when comparing this item with other
	// items. When `falsy` the {@link CompletionItem.label label} is used.
	SortText *String `json:"sortText"`

	// A string that should be used when filtering a set of completion
	// items. When `falsy` the {@link CompletionItem.label label} is used.
	FilterText *String `json:"filterText"`

	// A string that should be inserted into a document when selecting this
	// completion. When `falsy` the {@link CompletionItem.label label} is
	// used.  The `insertText` is subject to interpretation by the client
	// side. Some tools might not take the string literally. For example VS
	// Code when code complete is requested in this example `con<cursor
	// position>` and a completion item with an `insertText` of `console` is
	// provided it will only insert `sole`. Therefore it is recommended to
	// use `textEdit` instead since it avoids additional client side
	// interpretation.
	InsertText *String `json:"insertText"`

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

	// An {@link TextEdit edit} which is applied to a document when
	// selecting this completion. When an edit is provided the value of
	// {@link CompletionItem.insertText insertText} is ignored.  Most
	// editors support two different operations when accepting a completion
	// item. One is to insert a completion text and the other is to replace
	// an existing text with a completion text. Since this can usually not
	// be predetermined by a server it can report both ranges. Clients need
	// to signal support for `InsertReplaceEdits` via the
	// `textDocument.completion.insertReplaceSupport` client capability
	// property.  *Note 1:* The text edit's range as well as both ranges
	// from an insert replace edit must be a [single line] and they must
	// contain the position at which completion has been requested. *Note
	// 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.  @since 3.16.0
	// additional type `InsertReplaceEdit`
	TextEdit *CompletionItem_TextEdit_Or `json:"textEdit"`

	// The edit text used if the completion item is part of a CompletionList
	// and CompletionList defines an item default for the text edit range.
	// Clients will only honor this property if they opt into completion
	// list item defaults using the capability
	// `completionList.itemDefaults`.  If not provided and a list's default
	// range is provided the label property is used as a text.  @since
	// 3.17.0
	TextEditText *String `json:"textEditText"`

	// An optional array of additional {@link TextEdit text edits} that are
	// applied when selecting this completion. Edits must not overlap
	// (including the same insert position) with the main {@link
	// CompletionItem.textEdit edit} nor with themselves.  Additional text
	// edits should be used to change text unrelated to the current cursor
	// position (for example adding an import statement at the top of the
	// file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits"`

	// An optional set of characters that when pressed while this completion
	// is active will accept it first and then type that character. *Note*
	// that all commit characters should have `length=1` and that
	// superfluous characters will be ignored.
	CommitCharacters []String `json:"commitCharacters"`

	// An optional {@link Command command} that is executed *after*
	// inserting this completion. *Note* that additional modifications to
	// the current document should be described with the {@link
	// CompletionItem.additionalTextEdits additionalTextEdits}-property.
	Command *Command `json:"command"`

	// A data entry field that is preserved on a completion item between a
	// {@link CompletionRequest} and a {@link CompletionResolveRequest}.
	Data *LSPAny `json:"data"`
}

func (this *CompletionItem) UnmarshalJSON(data []byte) error {
	type CompletionItemUnmarshal CompletionItem
	var tmpUnmarshal CompletionItemUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("CompletionItem")
	}

	*this = CompletionItem(tmpUnmarshal)
	return nil
}

func (this *CompletionItem) MarshalJSON() ([]byte, error) {

	if this.Label == nil {
		return nil, StructureValidateFailed("CompletionItem")
	}

	type CompletionItemMarshal CompletionItem
	tmpMarshal := CompletionItemMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a collection of {@link CompletionItem completion items} to be
// presented in the editor.
type CompletionList struct {

	// This list it not complete. Further typing results in recomputing this
	// list.  Recomputed lists have all their items replaced (not appended)
	// in the incomplete completion sessions.
	IsIncomplete *Boolean `json:"isIncomplete"`

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

func (this *CompletionList) UnmarshalJSON(data []byte) error {
	type CompletionListUnmarshal CompletionList
	var tmpUnmarshal CompletionListUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.IsIncomplete == nil {
		return StructureValidateFailed("CompletionList")
	}

	if tmpUnmarshal.Items == nil {
		return StructureValidateFailed("CompletionList")
	}

	*this = CompletionList(tmpUnmarshal)
	return nil
}

func (this *CompletionList) MarshalJSON() ([]byte, error) {

	if this.IsIncomplete == nil {
		return nil, StructureValidateFailed("CompletionList")
	}

	if this.Items == nil {
		return nil, StructureValidateFailed("CompletionList")
	}

	type CompletionListMarshal CompletionList
	tmpMarshal := CompletionListMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link CompletionRequest}.
type CompletionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CompletionOptions
}

// Parameters for a {@link HoverRequest}.
type HoverParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams
}

// The result of a hover request.
type Hover struct {

	// The hover's content
	Contents *Hover_Contents_Or `json:"contents"`

	// An optional range inside the text document that is used to visualize
	// the hover, e.g. by changing the background color.
	Range *Range `json:"range"`
}

func (this *Hover) UnmarshalJSON(data []byte) error {
	type HoverUnmarshal Hover
	var tmpUnmarshal HoverUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Contents == nil {
		return StructureValidateFailed("Hover")
	}

	*this = Hover(tmpUnmarshal)
	return nil
}

func (this *Hover) MarshalJSON() ([]byte, error) {

	if this.Contents == nil {
		return nil, StructureValidateFailed("Hover")
	}

	type HoverMarshal Hover
	tmpMarshal := HoverMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link HoverRequest}.
type HoverRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	HoverOptions
}

// Parameters for a {@link SignatureHelpRequest}.
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
	ActiveSignature *Uinteger `json:"activeSignature"`

	// The active parameter of the active signature. If omitted or the value
	// lies outside the range of `signatures[activeSignature].parameters`
	// defaults to 0 if the active signature has parameters. If the active
	// signature has no parameters it is ignored. In future version of the
	// protocol this property might become mandatory to better express the
	// active parameter if the active signature does have any.
	ActiveParameter *Uinteger `json:"activeParameter"`
}

func (this *SignatureHelp) UnmarshalJSON(data []byte) error {
	type SignatureHelpUnmarshal SignatureHelp
	var tmpUnmarshal SignatureHelpUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Signatures == nil {
		return StructureValidateFailed("SignatureHelp")
	}

	*this = SignatureHelp(tmpUnmarshal)
	return nil
}

func (this *SignatureHelp) MarshalJSON() ([]byte, error) {

	if this.Signatures == nil {
		return nil, StructureValidateFailed("SignatureHelp")
	}

	type SignatureHelpMarshal SignatureHelp
	tmpMarshal := SignatureHelpMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link SignatureHelpRequest}.
type SignatureHelpRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	SignatureHelpOptions
}

// Parameters for a {@link DefinitionRequest}.
type DefinitionParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams
}

// Registration options for a {@link DefinitionRequest}.
type DefinitionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DefinitionOptions
}

// Parameters for a {@link ReferencesRequest}.
type ReferenceParams struct {

	// extends

	TextDocumentPositionParams

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	Context *ReferenceContext `json:"context"`
}

func (this *ReferenceParams) UnmarshalJSON(data []byte) error {
	type ReferenceParamsUnmarshal ReferenceParams
	var tmpUnmarshal ReferenceParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Context == nil {
		return StructureValidateFailed("ReferenceParams")
	}

	*this = ReferenceParams(tmpUnmarshal)
	return nil
}

func (this *ReferenceParams) MarshalJSON() ([]byte, error) {

	if this.Context == nil {
		return nil, StructureValidateFailed("ReferenceParams")
	}

	type ReferenceParamsMarshal ReferenceParams
	tmpMarshal := ReferenceParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link ReferencesRequest}.
type ReferenceRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	ReferenceOptions
}

// Parameters for a {@link DocumentHighlightRequest}.
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
	Range *Range `json:"range"`

	// The highlight kind, default is {@link DocumentHighlightKind.Text
	// text}.
	Kind *DocumentHighlightKind `json:"kind"`
}

func (this *DocumentHighlight) UnmarshalJSON(data []byte) error {
	type DocumentHighlightUnmarshal DocumentHighlight
	var tmpUnmarshal DocumentHighlightUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("DocumentHighlight")
	}

	*this = DocumentHighlight(tmpUnmarshal)
	return nil
}

func (this *DocumentHighlight) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("DocumentHighlight")
	}

	type DocumentHighlightMarshal DocumentHighlight
	tmpMarshal := DocumentHighlightMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentHighlightRequest}.
type DocumentHighlightRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentHighlightOptions
}

// Parameters for a {@link DocumentSymbolRequest}.
type DocumentSymbolParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *DocumentSymbolParams) UnmarshalJSON(data []byte) error {
	type DocumentSymbolParamsUnmarshal DocumentSymbolParams
	var tmpUnmarshal DocumentSymbolParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentSymbolParams")
	}

	*this = DocumentSymbolParams(tmpUnmarshal)
	return nil
}

func (this *DocumentSymbolParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DocumentSymbolParams")
	}

	type DocumentSymbolParamsMarshal DocumentSymbolParams
	tmpMarshal := DocumentSymbolParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {

	// extends

	BaseSymbolInformation

	// Indicates if this symbol is deprecated.  @deprecated Use tags instead
	Deprecated *Boolean `json:"deprecated"`

	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in
	// the tool the range's start information is used to position the
	// cursor. So the range usually spans more than the actual symbol's name
	// and does normally include things like visibility modifiers.  The
	// range doesn't have to denote a node range in the sense of an abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy
	// of the symbols.
	Location *Location `json:"location"`
}

func (this *SymbolInformation) UnmarshalJSON(data []byte) error {
	type SymbolInformationUnmarshal SymbolInformation
	var tmpUnmarshal SymbolInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Location == nil {
		return StructureValidateFailed("SymbolInformation")
	}

	*this = SymbolInformation(tmpUnmarshal)
	return nil
}

func (this *SymbolInformation) MarshalJSON() ([]byte, error) {

	if this.Location == nil {
		return nil, StructureValidateFailed("SymbolInformation")
	}

	type SymbolInformationMarshal SymbolInformation
	tmpMarshal := SymbolInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they have
// two ranges: one that encloses its definition and one that points to its most
// interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {

	// The name of this symbol. Will be displayed in the user interface and
	// therefore must not be an empty string or a string only consisting of
	// white spaces.
	Name *String `json:"name"`

	// More detail for this symbol, e.g the signature of a function.
	Detail *String `json:"detail"`

	// The kind of this symbol.
	Kind *SymbolKind `json:"kind"`

	// Tags for this document symbol.  @since 3.16.0
	Tags []SymbolTag `json:"tags"`

	// Indicates if this symbol is deprecated.  @deprecated Use tags instead
	Deprecated *Boolean `json:"deprecated"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else like comments. This information is
	// typically used to determine if the clients cursor is inside the
	// symbol to reveal in the symbol in the UI.
	Range *Range `json:"range"`

	// The range that should be selected and revealed when this symbol is
	// being picked, e.g the name of a function. Must be contained by the
	// `range`.
	SelectionRange *Range `json:"selectionRange"`

	// Children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children"`
}

func (this *DocumentSymbol) UnmarshalJSON(data []byte) error {
	type DocumentSymbolUnmarshal DocumentSymbol
	var tmpUnmarshal DocumentSymbolUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureValidateFailed("DocumentSymbol")
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("DocumentSymbol")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("DocumentSymbol")
	}

	if tmpUnmarshal.SelectionRange == nil {
		return StructureValidateFailed("DocumentSymbol")
	}

	*this = DocumentSymbol(tmpUnmarshal)
	return nil
}

func (this *DocumentSymbol) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureValidateFailed("DocumentSymbol")
	}

	if this.Kind == nil {
		return nil, StructureValidateFailed("DocumentSymbol")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("DocumentSymbol")
	}

	if this.SelectionRange == nil {
		return nil, StructureValidateFailed("DocumentSymbol")
	}

	type DocumentSymbolMarshal DocumentSymbol
	tmpMarshal := DocumentSymbolMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentSymbolRequest}.
type DocumentSymbolRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentSymbolOptions
}

// The parameters of a {@link CodeActionRequest}.
type CodeActionParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document in which the command was invoked.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The range for which the command was invoked.
	Range *Range `json:"range"`

	// Context carrying additional information.
	Context *CodeActionContext `json:"context"`
}

func (this *CodeActionParams) UnmarshalJSON(data []byte) error {
	type CodeActionParamsUnmarshal CodeActionParams
	var tmpUnmarshal CodeActionParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("CodeActionParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("CodeActionParams")
	}

	if tmpUnmarshal.Context == nil {
		return StructureValidateFailed("CodeActionParams")
	}

	*this = CodeActionParams(tmpUnmarshal)
	return nil
}

func (this *CodeActionParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("CodeActionParams")
	}

	if this.Range == nil {
		return nil, StructureValidateFailed("CodeActionParams")
	}

	if this.Context == nil {
		return nil, StructureValidateFailed("CodeActionParams")
	}

	type CodeActionParamsMarshal CodeActionParams
	tmpMarshal := CodeActionParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a reference to a command. Provides a title which will be used to
// represent a command in the UI and, optionally, an array of arguments which
// will be passed to the command handler function when invoked.
type Command struct {

	// Title of the command, like `save`.
	Title *String `json:"title"`

	// The identifier of the actual command handler.
	Command *String `json:"command"`

	// Arguments that the command handler should be invoked with.
	Arguments []LSPAny `json:"arguments"`
}

func (this *Command) UnmarshalJSON(data []byte) error {
	type CommandUnmarshal Command
	var tmpUnmarshal CommandUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Title == nil {
		return StructureValidateFailed("Command")
	}

	if tmpUnmarshal.Command == nil {
		return StructureValidateFailed("Command")
	}

	*this = Command(tmpUnmarshal)
	return nil
}

func (this *Command) MarshalJSON() ([]byte, error) {

	if this.Title == nil {
		return nil, StructureValidateFailed("Command")
	}

	if this.Command == nil {
		return nil, StructureValidateFailed("Command")
	}

	type CommandMarshal Command
	tmpMarshal := CommandMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A code action represents a change that can be performed in code, e.g. to fix
// a problem or to refactor code.  A CodeAction must set either `edit` and/or a
// `command`. If both are supplied, the `edit` is applied first, then the
// `command` is executed.
type CodeAction struct {

	// A short, human-readable, title for this code action.
	Title *String `json:"title"`

	// The kind of the code action.  Used to filter code actions.
	Kind *CodeActionKind `json:"kind"`

	// The diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Marks this as a preferred action. Preferred actions are used by the
	// `auto fix` command and can be targeted by keybindings.  A quick fix
	// should be marked preferred if it properly addresses the underlying
	// error. A refactoring should be marked preferred if it is the most
	// reasonable choice of actions to take.  @since 3.15.0
	IsPreferred *Boolean `json:"isPreferred"`

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

func (this *CodeAction) UnmarshalJSON(data []byte) error {
	type CodeActionUnmarshal CodeAction
	var tmpUnmarshal CodeActionUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Title == nil {
		return StructureValidateFailed("CodeAction")
	}

	*this = CodeAction(tmpUnmarshal)
	return nil
}

func (this *CodeAction) MarshalJSON() ([]byte, error) {

	if this.Title == nil {
		return nil, StructureValidateFailed("CodeAction")
	}

	type CodeActionMarshal CodeAction
	tmpMarshal := CodeActionMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link CodeActionRequest}.
type CodeActionRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CodeActionOptions
}

// The parameters of a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// A query string to filter symbols by. Clients may send an empty string
	// here to request all symbols.
	Query *String `json:"query"`
}

func (this *WorkspaceSymbolParams) UnmarshalJSON(data []byte) error {
	type WorkspaceSymbolParamsUnmarshal WorkspaceSymbolParams
	var tmpUnmarshal WorkspaceSymbolParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Query == nil {
		return StructureValidateFailed("WorkspaceSymbolParams")
	}

	*this = WorkspaceSymbolParams(tmpUnmarshal)
	return nil
}

func (this *WorkspaceSymbolParams) MarshalJSON() ([]byte, error) {

	if this.Query == nil {
		return nil, StructureValidateFailed("WorkspaceSymbolParams")
	}

	type WorkspaceSymbolParamsMarshal WorkspaceSymbolParams
	tmpMarshal := WorkspaceSymbolParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Location *WorkspaceSymbol_Location_Or `json:"location"`

	// A data entry field that is preserved on a workspace symbol between a
	// workspace symbol request and a workspace symbol resolve request.
	Data *LSPAny `json:"data"`
}

func (this *WorkspaceSymbol) UnmarshalJSON(data []byte) error {
	type WorkspaceSymbolUnmarshal WorkspaceSymbol
	var tmpUnmarshal WorkspaceSymbolUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Location == nil {
		return StructureValidateFailed("WorkspaceSymbol")
	}

	*this = WorkspaceSymbol(tmpUnmarshal)
	return nil
}

func (this *WorkspaceSymbol) MarshalJSON() ([]byte, error) {

	if this.Location == nil {
		return nil, StructureValidateFailed("WorkspaceSymbol")
	}

	type WorkspaceSymbolMarshal WorkspaceSymbol
	tmpMarshal := WorkspaceSymbolMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolRegistrationOptions struct {

	// extends

	WorkspaceSymbolOptions
}

// The parameters of a {@link CodeLensRequest}.
type CodeLensParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document to request code lens for.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *CodeLensParams) UnmarshalJSON(data []byte) error {
	type CodeLensParamsUnmarshal CodeLensParams
	var tmpUnmarshal CodeLensParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("CodeLensParams")
	}

	*this = CodeLensParams(tmpUnmarshal)
	return nil
}

func (this *CodeLensParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("CodeLensParams")
	}

	type CodeLensParamsMarshal CodeLensParams
	tmpMarshal := CodeLensParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A code lens represents a {@link Command command} that should be shown along
// with source text, like the number of references, a way to run tests, etc.  A
// code lens is _unresolved_ when no command is associated to it. For
// performance reasons the creation of a code lens and resolving should be done
// in two stages.
type CodeLens struct {

	// The range in which this code lens is valid. Should only span a single
	// line.
	Range *Range `json:"range"`

	// The command this code lens represents.
	Command *Command `json:"command"`

	// A data entry field that is preserved on a code lens item between a
	// {@link CodeLensRequest} and a [CodeLensResolveRequest]
	// (#CodeLensResolveRequest)
	Data *LSPAny `json:"data"`
}

func (this *CodeLens) UnmarshalJSON(data []byte) error {
	type CodeLensUnmarshal CodeLens
	var tmpUnmarshal CodeLensUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("CodeLens")
	}

	*this = CodeLens(tmpUnmarshal)
	return nil
}

func (this *CodeLens) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("CodeLens")
	}

	type CodeLensMarshal CodeLens
	tmpMarshal := CodeLensMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link CodeLensRequest}.
type CodeLensRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	CodeLensOptions
}

// The parameters of a {@link DocumentLinkRequest}.
type DocumentLinkParams struct {

	// mixins

	WorkDoneProgressParams

	PartialResultParams

	// The document to provide document links for.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

func (this *DocumentLinkParams) UnmarshalJSON(data []byte) error {
	type DocumentLinkParamsUnmarshal DocumentLinkParams
	var tmpUnmarshal DocumentLinkParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentLinkParams")
	}

	*this = DocumentLinkParams(tmpUnmarshal)
	return nil
}

func (this *DocumentLinkParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DocumentLinkParams")
	}

	type DocumentLinkParamsMarshal DocumentLinkParams
	tmpMarshal := DocumentLinkParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A document link is a range in a text document that links to an internal or
// external resource, like another text document or a web site.
type DocumentLink struct {

	// The range this link applies to.
	Range *Range `json:"range"`

	// The uri this link points to. If missing a resolve request is sent
	// later.
	Target *String `json:"target"`

	// The tooltip text when you hover over this link.  If a tooltip is
	// provided, is will be displayed in a string that includes instructions
	// on how to trigger the link, such as `{0} (ctrl + click)`. The
	// specific instructions vary depending on OS, user settings, and
	// localization.  @since 3.15.0
	Tooltip *String `json:"tooltip"`

	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data *LSPAny `json:"data"`
}

func (this *DocumentLink) UnmarshalJSON(data []byte) error {
	type DocumentLinkUnmarshal DocumentLink
	var tmpUnmarshal DocumentLinkUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("DocumentLink")
	}

	*this = DocumentLink(tmpUnmarshal)
	return nil
}

func (this *DocumentLink) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("DocumentLink")
	}

	type DocumentLinkMarshal DocumentLink
	tmpMarshal := DocumentLinkMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentLinkRequest}.
type DocumentLinkRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentLinkOptions
}

// The parameters of a {@link DocumentFormattingRequest}.
type DocumentFormattingParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to format.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The format options.
	Options *FormattingOptions `json:"options"`
}

func (this *DocumentFormattingParams) UnmarshalJSON(data []byte) error {
	type DocumentFormattingParamsUnmarshal DocumentFormattingParams
	var tmpUnmarshal DocumentFormattingParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentFormattingParams")
	}

	if tmpUnmarshal.Options == nil {
		return StructureValidateFailed("DocumentFormattingParams")
	}

	*this = DocumentFormattingParams(tmpUnmarshal)
	return nil
}

func (this *DocumentFormattingParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("DocumentFormattingParams")
	}

	if this.Options == nil {
		return nil, StructureValidateFailed("DocumentFormattingParams")
	}

	type DocumentFormattingParamsMarshal DocumentFormattingParams
	tmpMarshal := DocumentFormattingParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentFormattingRequest}.
type DocumentFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentFormattingOptions
}

// The parameters of a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to format.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The range to format
	Range *Range `json:"range"`

	// The format options
	Options *FormattingOptions `json:"options"`
}

func (this *DocumentRangeFormattingParams) UnmarshalJSON(data []byte) error {
	type DocumentRangeFormattingParamsUnmarshal DocumentRangeFormattingParams
	var tmpUnmarshal DocumentRangeFormattingParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentRangeFormattingParams")
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("DocumentRangeFormattingParams")
	}

	if tmpUnmarshal.Options == nil {
		return StructureValidateFailed("DocumentRangeFormattingParams")
	}

	*this = DocumentRangeFormattingParams(tmpUnmarshal)
	return nil
}

func (this *DocumentRangeFormattingParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"DocumentRangeFormattingParams",
		)
	}

	if this.Range == nil {
		return nil, StructureValidateFailed(
			"DocumentRangeFormattingParams",
		)
	}

	if this.Options == nil {
		return nil, StructureValidateFailed(
			"DocumentRangeFormattingParams",
		)
	}

	type DocumentRangeFormattingParamsMarshal DocumentRangeFormattingParams
	tmpMarshal := DocumentRangeFormattingParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentRangeFormattingOptions
}

// The parameters of a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingParams struct {

	// The document to format.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The position around which the on type formatting should happen. This
	// is not necessarily the exact position where the character denoted by
	// the property `ch` got typed.
	Position *Position `json:"position"`

	// The character that has been typed that triggered the formatting on
	// type request. That is not necessarily the last character that got
	// inserted into the document since the client could auto insert
	// characters as well (e.g. like automatic brace completion).
	Ch *String `json:"ch"`

	// The formatting options.
	Options *FormattingOptions `json:"options"`
}

func (this *DocumentOnTypeFormattingParams) UnmarshalJSON(data []byte) error {
	type DocumentOnTypeFormattingParamsUnmarshal DocumentOnTypeFormattingParams
	var tmpUnmarshal DocumentOnTypeFormattingParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("DocumentOnTypeFormattingParams")
	}

	if tmpUnmarshal.Position == nil {
		return StructureValidateFailed("DocumentOnTypeFormattingParams")
	}

	if tmpUnmarshal.Ch == nil {
		return StructureValidateFailed("DocumentOnTypeFormattingParams")
	}

	if tmpUnmarshal.Options == nil {
		return StructureValidateFailed("DocumentOnTypeFormattingParams")
	}

	*this = DocumentOnTypeFormattingParams(tmpUnmarshal)
	return nil
}

func (this *DocumentOnTypeFormattingParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"DocumentOnTypeFormattingParams",
		)
	}

	if this.Position == nil {
		return nil, StructureValidateFailed(
			"DocumentOnTypeFormattingParams",
		)
	}

	if this.Ch == nil {
		return nil, StructureValidateFailed(
			"DocumentOnTypeFormattingParams",
		)
	}

	if this.Options == nil {
		return nil, StructureValidateFailed(
			"DocumentOnTypeFormattingParams",
		)
	}

	type DocumentOnTypeFormattingParamsMarshal DocumentOnTypeFormattingParams
	tmpMarshal := DocumentOnTypeFormattingParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingRegistrationOptions struct {

	// extends

	TextDocumentRegistrationOptions

	DocumentOnTypeFormattingOptions
}

// The parameters of a {@link RenameRequest}.
type RenameParams struct {

	// mixins

	WorkDoneProgressParams

	// The document to rename.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The position at which this request was sent.
	Position *Position `json:"position"`

	// The new name of the symbol. If the given name is not valid the
	// request must return a {@link ResponseError} with an appropriate
	// message set.
	NewName *String `json:"newName"`
}

func (this *RenameParams) UnmarshalJSON(data []byte) error {
	type RenameParamsUnmarshal RenameParams
	var tmpUnmarshal RenameParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("RenameParams")
	}

	if tmpUnmarshal.Position == nil {
		return StructureValidateFailed("RenameParams")
	}

	if tmpUnmarshal.NewName == nil {
		return StructureValidateFailed("RenameParams")
	}

	*this = RenameParams(tmpUnmarshal)
	return nil
}

func (this *RenameParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("RenameParams")
	}

	if this.Position == nil {
		return nil, StructureValidateFailed("RenameParams")
	}

	if this.NewName == nil {
		return nil, StructureValidateFailed("RenameParams")
	}

	type RenameParamsMarshal RenameParams
	tmpMarshal := RenameParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link RenameRequest}.
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

// The parameters of a {@link ExecuteCommandRequest}.
type ExecuteCommandParams struct {

	// mixins

	WorkDoneProgressParams

	// The identifier of the actual command handler.
	Command *String `json:"command"`

	// Arguments that the command should be invoked with.
	Arguments []LSPAny `json:"arguments"`
}

func (this *ExecuteCommandParams) UnmarshalJSON(data []byte) error {
	type ExecuteCommandParamsUnmarshal ExecuteCommandParams
	var tmpUnmarshal ExecuteCommandParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Command == nil {
		return StructureValidateFailed("ExecuteCommandParams")
	}

	*this = ExecuteCommandParams(tmpUnmarshal)
	return nil
}

func (this *ExecuteCommandParams) MarshalJSON() ([]byte, error) {

	if this.Command == nil {
		return nil, StructureValidateFailed("ExecuteCommandParams")
	}

	type ExecuteCommandParamsMarshal ExecuteCommandParams
	tmpMarshal := ExecuteCommandParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Registration options for a {@link ExecuteCommandRequest}.
type ExecuteCommandRegistrationOptions struct {

	// extends

	ExecuteCommandOptions
}

// The parameters passed via a apply workspace edit request.
type ApplyWorkspaceEditParams struct {

	// An optional label of the workspace edit. This label is presented in
	// the user interface for example on an undo stack to undo the workspace
	// edit.
	Label *String `json:"label"`

	// The edits to apply.
	Edit *WorkspaceEdit `json:"edit"`
}

func (this *ApplyWorkspaceEditParams) UnmarshalJSON(data []byte) error {
	type ApplyWorkspaceEditParamsUnmarshal ApplyWorkspaceEditParams
	var tmpUnmarshal ApplyWorkspaceEditParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Edit == nil {
		return StructureValidateFailed("ApplyWorkspaceEditParams")
	}

	*this = ApplyWorkspaceEditParams(tmpUnmarshal)
	return nil
}

func (this *ApplyWorkspaceEditParams) MarshalJSON() ([]byte, error) {

	if this.Edit == nil {
		return nil, StructureValidateFailed("ApplyWorkspaceEditParams")
	}

	type ApplyWorkspaceEditParamsMarshal ApplyWorkspaceEditParams
	tmpMarshal := ApplyWorkspaceEditParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The result returned from the apply workspace edit request.  @since 3.17
// renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {

	// Indicates whether the edit was applied or not.
	Applied *Boolean `json:"applied"`

	// An optional textual description for why the edit was not applied.
	// This may be used by the server for diagnostic logging or to provide a
	// suitable error for a request that triggered the edit.
	FailureReason *String `json:"failureReason"`

	// Depending on the client's failure handling strategy `failedChange`
	// might contain the index of the change that failed. This property is
	// only available if the client signals a `failureHandlingStrategy` in
	// its client capabilities.
	FailedChange *Uinteger `json:"failedChange"`
}

func (this *ApplyWorkspaceEditResult) UnmarshalJSON(data []byte) error {
	type ApplyWorkspaceEditResultUnmarshal ApplyWorkspaceEditResult
	var tmpUnmarshal ApplyWorkspaceEditResultUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Applied == nil {
		return StructureValidateFailed("ApplyWorkspaceEditResult")
	}

	*this = ApplyWorkspaceEditResult(tmpUnmarshal)
	return nil
}

func (this *ApplyWorkspaceEditResult) MarshalJSON() ([]byte, error) {

	if this.Applied == nil {
		return nil, StructureValidateFailed("ApplyWorkspaceEditResult")
	}

	type ApplyWorkspaceEditResultMarshal ApplyWorkspaceEditResult
	tmpMarshal := ApplyWorkspaceEditResultMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressBegin struct {
	Kind *string `json:"kind"`

	// Mandatory title of the progress operation. Used to briefly inform
	// about the kind of operation being performed.  Examples: "Indexing" or
	// "Linking dependencies".
	Title *String `json:"title"`

	// Controls if a cancel button should show to allow the user to cancel
	// the long running operation. Clients that don't support cancellation
	// are allowed to ignore the setting.
	Cancellable *Boolean `json:"cancellable"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.  Examples: "3/25 files",
	// "project/src/module2", "node_modules/some_dep". If unset, the
	// previous progress message (if any) is still valid.
	Message *String `json:"message"`

	// Optional progress percentage to display (value 100 is considered
	// 100%). If not provided infinite progress is assumed and clients are
	// allowed to ignore the `percentage` value in subsequent in report
	// notifications.  The value should be steadily rising. Clients are free
	// to ignore values that are not following this rule. The value range is
	// [0, 100].
	Percentage *Uinteger `json:"percentage"`
}

func (this *WorkDoneProgressBegin) UnmarshalJSON(data []byte) error {
	type WorkDoneProgressBeginUnmarshal WorkDoneProgressBegin
	var tmpUnmarshal WorkDoneProgressBeginUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("WorkDoneProgressBegin")
	}

	if tmpUnmarshal.Title == nil {
		return StructureValidateFailed("WorkDoneProgressBegin")
	}

	*this = WorkDoneProgressBegin(tmpUnmarshal)
	return nil
}

func (this *WorkDoneProgressBegin) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("WorkDoneProgressBegin")
	}

	if this.Title == nil {
		return nil, StructureValidateFailed("WorkDoneProgressBegin")
	}

	type WorkDoneProgressBeginMarshal WorkDoneProgressBegin
	tmpMarshal := WorkDoneProgressBeginMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressReport struct {
	Kind *string `json:"kind"`

	// Controls enablement state of a cancel button.  Clients that don't
	// support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable *Boolean `json:"cancellable"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.  Examples: "3/25 files",
	// "project/src/module2", "node_modules/some_dep". If unset, the
	// previous progress message (if any) is still valid.
	Message *String `json:"message"`

	// Optional progress percentage to display (value 100 is considered
	// 100%). If not provided infinite progress is assumed and clients are
	// allowed to ignore the `percentage` value in subsequent in report
	// notifications.  The value should be steadily rising. Clients are free
	// to ignore values that are not following this rule. The value range is
	// [0, 100]
	Percentage *Uinteger `json:"percentage"`
}

func (this *WorkDoneProgressReport) UnmarshalJSON(data []byte) error {
	type WorkDoneProgressReportUnmarshal WorkDoneProgressReport
	var tmpUnmarshal WorkDoneProgressReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("WorkDoneProgressReport")
	}

	*this = WorkDoneProgressReport(tmpUnmarshal)
	return nil
}

func (this *WorkDoneProgressReport) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("WorkDoneProgressReport")
	}

	type WorkDoneProgressReportMarshal WorkDoneProgressReport
	tmpMarshal := WorkDoneProgressReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressEnd struct {
	Kind *string `json:"kind"`

	// Optional, a final message indicating to for example indicate the
	// outcome of the operation.
	Message *String `json:"message"`
}

func (this *WorkDoneProgressEnd) UnmarshalJSON(data []byte) error {
	type WorkDoneProgressEndUnmarshal WorkDoneProgressEnd
	var tmpUnmarshal WorkDoneProgressEndUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("WorkDoneProgressEnd")
	}

	*this = WorkDoneProgressEnd(tmpUnmarshal)
	return nil
}

func (this *WorkDoneProgressEnd) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("WorkDoneProgressEnd")
	}

	type WorkDoneProgressEndMarshal WorkDoneProgressEnd
	tmpMarshal := WorkDoneProgressEndMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type SetTraceParams struct {
	Value *TraceValues `json:"value"`
}

func (this *SetTraceParams) UnmarshalJSON(data []byte) error {
	type SetTraceParamsUnmarshal SetTraceParams
	var tmpUnmarshal SetTraceParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Value == nil {
		return StructureValidateFailed("SetTraceParams")
	}

	*this = SetTraceParams(tmpUnmarshal)
	return nil
}

func (this *SetTraceParams) MarshalJSON() ([]byte, error) {

	if this.Value == nil {
		return nil, StructureValidateFailed("SetTraceParams")
	}

	type SetTraceParamsMarshal SetTraceParams
	tmpMarshal := SetTraceParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type LogTraceParams struct {
	Message *String `json:"message"`

	Verbose *String `json:"verbose"`
}

func (this *LogTraceParams) UnmarshalJSON(data []byte) error {
	type LogTraceParamsUnmarshal LogTraceParams
	var tmpUnmarshal LogTraceParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("LogTraceParams")
	}

	*this = LogTraceParams(tmpUnmarshal)
	return nil
}

func (this *LogTraceParams) MarshalJSON() ([]byte, error) {

	if this.Message == nil {
		return nil, StructureValidateFailed("LogTraceParams")
	}

	type LogTraceParamsMarshal LogTraceParams
	tmpMarshal := LogTraceParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type CancelParams struct {

	// The request id to cancel.
	Id *CancelParams_Id_Or `json:"id"`
}

func (this *CancelParams) UnmarshalJSON(data []byte) error {
	type CancelParamsUnmarshal CancelParams
	var tmpUnmarshal CancelParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Id == nil {
		return StructureValidateFailed("CancelParams")
	}

	*this = CancelParams(tmpUnmarshal)
	return nil
}

func (this *CancelParams) MarshalJSON() ([]byte, error) {

	if this.Id == nil {
		return nil, StructureValidateFailed("CancelParams")
	}

	type CancelParamsMarshal CancelParams
	tmpMarshal := CancelParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type ProgressParams struct {

	// The progress token provided by the client or server.
	Token *ProgressToken `json:"token"`

	// The progress data.
	Value *LSPAny `json:"value"`
}

func (this *ProgressParams) UnmarshalJSON(data []byte) error {
	type ProgressParamsUnmarshal ProgressParams
	var tmpUnmarshal ProgressParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Token == nil {
		return StructureValidateFailed("ProgressParams")
	}

	if tmpUnmarshal.Value == nil {
		return StructureValidateFailed("ProgressParams")
	}

	*this = ProgressParams(tmpUnmarshal)
	return nil
}

func (this *ProgressParams) MarshalJSON() ([]byte, error) {

	if this.Token == nil {
		return nil, StructureValidateFailed("ProgressParams")
	}

	if this.Value == nil {
		return nil, StructureValidateFailed("ProgressParams")
	}

	type ProgressParamsMarshal ProgressParams
	tmpMarshal := ProgressParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A parameter literal used in requests to pass a text document and a position
// inside that document.
type TextDocumentPositionParams struct {

	// The text document.
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The position inside the text document.
	Position *Position `json:"position"`
}

func (this *TextDocumentPositionParams) UnmarshalJSON(data []byte) error {
	type TextDocumentPositionParamsUnmarshal TextDocumentPositionParams
	var tmpUnmarshal TextDocumentPositionParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("TextDocumentPositionParams")
	}

	if tmpUnmarshal.Position == nil {
		return StructureValidateFailed("TextDocumentPositionParams")
	}

	*this = TextDocumentPositionParams(tmpUnmarshal)
	return nil
}

func (this *TextDocumentPositionParams) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed(
			"TextDocumentPositionParams",
		)
	}

	if this.Position == nil {
		return nil, StructureValidateFailed(
			"TextDocumentPositionParams",
		)
	}

	type TextDocumentPositionParamsMarshal TextDocumentPositionParams
	tmpMarshal := TextDocumentPositionParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkDoneProgressParams struct {

	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken"`
}

type PartialResultParams struct {

	// An optional token that a server can use to report partial results
	// (e.g. streaming) to the client.
	PartialResultToken *ProgressToken `json:"partialResultToken"`
}

// Represents the connection of two locations. Provides additional metadata over
// normal {@link Location locations}, including an origin range.
type LocationLink struct {

	// Span of the origin of this link.  Used as the underlined span for
	// mouse interaction. Defaults to the word range at the definition
	// position.
	OriginSelectionRange *Range `json:"originSelectionRange"`

	// The target resource identifier of this link.
	TargetUri *DocumentUri `json:"targetUri"`

	// The full target range of this link. If the target for example is a
	// symbol then target range is the range enclosing this symbol not
	// including leading/trailing whitespace but everything else like
	// comments. This information is typically used to highlight the range
	// in the editor.
	TargetRange *Range `json:"targetRange"`

	// The range that should be selected and revealed when this link is
	// being followed, e.g the name of a function. Must be contained by the
	// `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange *Range `json:"targetSelectionRange"`
}

func (this *LocationLink) UnmarshalJSON(data []byte) error {
	type LocationLinkUnmarshal LocationLink
	var tmpUnmarshal LocationLinkUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TargetUri == nil {
		return StructureValidateFailed("LocationLink")
	}

	if tmpUnmarshal.TargetRange == nil {
		return StructureValidateFailed("LocationLink")
	}

	if tmpUnmarshal.TargetSelectionRange == nil {
		return StructureValidateFailed("LocationLink")
	}

	*this = LocationLink(tmpUnmarshal)
	return nil
}

func (this *LocationLink) MarshalJSON() ([]byte, error) {

	if this.TargetUri == nil {
		return nil, StructureValidateFailed("LocationLink")
	}

	if this.TargetRange == nil {
		return nil, StructureValidateFailed("LocationLink")
	}

	if this.TargetSelectionRange == nil {
		return nil, StructureValidateFailed("LocationLink")
	}

	type LocationLinkMarshal LocationLink
	tmpMarshal := LocationLinkMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Start *Position `json:"start"`

	// The range's end position.
	End *Position `json:"end"`
}

func (this *Range) UnmarshalJSON(data []byte) error {
	type RangeUnmarshal Range
	var tmpUnmarshal RangeUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Start == nil {
		return StructureValidateFailed("Range")
	}

	if tmpUnmarshal.End == nil {
		return StructureValidateFailed("Range")
	}

	*this = Range(tmpUnmarshal)
	return nil
}

func (this *Range) MarshalJSON() ([]byte, error) {

	if this.Start == nil {
		return nil, StructureValidateFailed("Range")
	}

	if this.End == nil {
		return nil, StructureValidateFailed("Range")
	}

	type RangeMarshal Range
	tmpMarshal := RangeMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type ImplementationOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Static registration options to be returned in the initialize request.
type StaticRegistrationOptions struct {

	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id *String `json:"id"`
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

func (this *WorkspaceFoldersChangeEvent) UnmarshalJSON(data []byte) error {
	type WorkspaceFoldersChangeEventUnmarshal WorkspaceFoldersChangeEvent
	var tmpUnmarshal WorkspaceFoldersChangeEventUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Added == nil {
		return StructureValidateFailed("WorkspaceFoldersChangeEvent")
	}

	if tmpUnmarshal.Removed == nil {
		return StructureValidateFailed("WorkspaceFoldersChangeEvent")
	}

	*this = WorkspaceFoldersChangeEvent(tmpUnmarshal)
	return nil
}

func (this *WorkspaceFoldersChangeEvent) MarshalJSON() ([]byte, error) {

	if this.Added == nil {
		return nil, StructureValidateFailed(
			"WorkspaceFoldersChangeEvent",
		)
	}

	if this.Removed == nil {
		return nil, StructureValidateFailed(
			"WorkspaceFoldersChangeEvent",
		)
	}

	type WorkspaceFoldersChangeEventMarshal WorkspaceFoldersChangeEvent
	tmpMarshal := WorkspaceFoldersChangeEventMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type ConfigurationItem struct {

	// The scope to get the configuration section for.
	ScopeUri *String `json:"scopeUri"`

	// The configuration section asked for.
	Section *String `json:"section"`
}

// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {

	// The text document's uri.
	Uri *DocumentUri `json:"uri"`
}

func (this *TextDocumentIdentifier) UnmarshalJSON(data []byte) error {
	type TextDocumentIdentifierUnmarshal TextDocumentIdentifier
	var tmpUnmarshal TextDocumentIdentifierUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("TextDocumentIdentifier")
	}

	*this = TextDocumentIdentifier(tmpUnmarshal)
	return nil
}

func (this *TextDocumentIdentifier) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("TextDocumentIdentifier")
	}

	type TextDocumentIdentifierMarshal TextDocumentIdentifier
	tmpMarshal := TextDocumentIdentifierMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a color in RGBA space.
type Color struct {

	// The red component of this color in the range [0-1].
	Red *Decimal `json:"red"`

	// The green component of this color in the range [0-1].
	Green *Decimal `json:"green"`

	// The blue component of this color in the range [0-1].
	Blue *Decimal `json:"blue"`

	// The alpha component of this color in the range [0-1].
	Alpha *Decimal `json:"alpha"`
}

func (this *Color) UnmarshalJSON(data []byte) error {
	type ColorUnmarshal Color
	var tmpUnmarshal ColorUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Red == nil {
		return StructureValidateFailed("Color")
	}

	if tmpUnmarshal.Green == nil {
		return StructureValidateFailed("Color")
	}

	if tmpUnmarshal.Blue == nil {
		return StructureValidateFailed("Color")
	}

	if tmpUnmarshal.Alpha == nil {
		return StructureValidateFailed("Color")
	}

	*this = Color(tmpUnmarshal)
	return nil
}

func (this *Color) MarshalJSON() ([]byte, error) {

	if this.Red == nil {
		return nil, StructureValidateFailed("Color")
	}

	if this.Green == nil {
		return nil, StructureValidateFailed("Color")
	}

	if this.Blue == nil {
		return nil, StructureValidateFailed("Color")
	}

	if this.Alpha == nil {
		return nil, StructureValidateFailed("Color")
	}

	type ColorMarshal Color
	tmpMarshal := ColorMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Line *Uinteger `json:"line"`

	// Character offset on a line in a document (zero-based).  The meaning
	// of this offset is determined by the negotiated
	// `PositionEncodingKind`.  If the character value is greater than the
	// line length it defaults back to the line length.
	Character *Uinteger `json:"character"`
}

func (this *Position) UnmarshalJSON(data []byte) error {
	type PositionUnmarshal Position
	var tmpUnmarshal PositionUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Line == nil {
		return StructureValidateFailed("Position")
	}

	if tmpUnmarshal.Character == nil {
		return StructureValidateFailed("Position")
	}

	*this = Position(tmpUnmarshal)
	return nil
}

func (this *Position) MarshalJSON() ([]byte, error) {

	if this.Line == nil {
		return nil, StructureValidateFailed("Position")
	}

	if this.Character == nil {
		return nil, StructureValidateFailed("Position")
	}

	type PositionMarshal Position
	tmpMarshal := PositionMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Legend *SemanticTokensLegend `json:"legend"`

	// Server supports providing semantic tokens for a specific range of a
	// document.
	Range *SemanticTokensOptions_Range_Or `json:"range"`

	// Server supports providing semantic tokens for a full document.
	Full *SemanticTokensOptions_Full_Or `json:"full"`
}

func (this *SemanticTokensOptions) UnmarshalJSON(data []byte) error {
	type SemanticTokensOptionsUnmarshal SemanticTokensOptions
	var tmpUnmarshal SemanticTokensOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Legend == nil {
		return StructureValidateFailed("SemanticTokensOptions")
	}

	*this = SemanticTokensOptions(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensOptions) MarshalJSON() ([]byte, error) {

	if this.Legend == nil {
		return nil, StructureValidateFailed("SemanticTokensOptions")
	}

	type SemanticTokensOptionsMarshal SemanticTokensOptions
	tmpMarshal := SemanticTokensOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensEdit struct {

	// The start offset of the edit.
	Start *Uinteger `json:"start"`

	// The count of elements to remove.
	DeleteCount *Uinteger `json:"deleteCount"`

	// The elements to insert.
	Data []Uinteger `json:"data"`
}

func (this *SemanticTokensEdit) UnmarshalJSON(data []byte) error {
	type SemanticTokensEditUnmarshal SemanticTokensEdit
	var tmpUnmarshal SemanticTokensEditUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Start == nil {
		return StructureValidateFailed("SemanticTokensEdit")
	}

	if tmpUnmarshal.DeleteCount == nil {
		return StructureValidateFailed("SemanticTokensEdit")
	}

	*this = SemanticTokensEdit(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensEdit) MarshalJSON() ([]byte, error) {

	if this.Start == nil {
		return nil, StructureValidateFailed("SemanticTokensEdit")
	}

	if this.DeleteCount == nil {
		return nil, StructureValidateFailed("SemanticTokensEdit")
	}

	type SemanticTokensEditMarshal SemanticTokensEdit
	tmpMarshal := SemanticTokensEditMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type LinkedEditingRangeOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Represents information on a file/folder create.  @since 3.16.0
type FileCreate struct {

	// A file:// URI for the location of the file/folder being created.
	Uri *String `json:"uri"`
}

func (this *FileCreate) UnmarshalJSON(data []byte) error {
	type FileCreateUnmarshal FileCreate
	var tmpUnmarshal FileCreateUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("FileCreate")
	}

	*this = FileCreate(tmpUnmarshal)
	return nil
}

func (this *FileCreate) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("FileCreate")
	}

	type FileCreateMarshal FileCreate
	tmpMarshal := FileCreateMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Describes textual changes on a text document. A TextDocumentEdit describes
// all changes on a document version Si and after they are applied move the
// document to version Si+1. So the creator of a TextDocumentEdit doesn't need
// to sort the array of edits or do any kind of ordering. However the edits must
// be non overlapping.
type TextDocumentEdit struct {

	// The text document to change.
	TextDocument *OptionalVersionedTextDocumentIdentifier `json:"textDocument"`

	// The edits to be applied.  @since 3.16.0 - support for
	// AnnotatedTextEdit. This is guarded using a client capability.
	Edits []TextDocumentEdit_Edits_Element_Or `json:"edits"`
}

func (this *TextDocumentEdit) UnmarshalJSON(data []byte) error {
	type TextDocumentEditUnmarshal TextDocumentEdit
	var tmpUnmarshal TextDocumentEditUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TextDocument == nil {
		return StructureValidateFailed("TextDocumentEdit")
	}

	if tmpUnmarshal.Edits == nil {
		return StructureValidateFailed("TextDocumentEdit")
	}

	*this = TextDocumentEdit(tmpUnmarshal)
	return nil
}

func (this *TextDocumentEdit) MarshalJSON() ([]byte, error) {

	if this.TextDocument == nil {
		return nil, StructureValidateFailed("TextDocumentEdit")
	}

	if this.Edits == nil {
		return nil, StructureValidateFailed("TextDocumentEdit")
	}

	type TextDocumentEditMarshal TextDocumentEdit
	tmpMarshal := TextDocumentEditMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Create file operation.
type CreateFile struct {

	// extends

	ResourceOperation

	// A create
	Kind *string `json:"kind"`

	// The resource to create.
	Uri *DocumentUri `json:"uri"`

	// Additional options
	Options *CreateFileOptions `json:"options"`
}

func (this *CreateFile) UnmarshalJSON(data []byte) error {
	type CreateFileUnmarshal CreateFile
	var tmpUnmarshal CreateFileUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("CreateFile")
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("CreateFile")
	}

	*this = CreateFile(tmpUnmarshal)
	return nil
}

func (this *CreateFile) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("CreateFile")
	}

	if this.Uri == nil {
		return nil, StructureValidateFailed("CreateFile")
	}

	type CreateFileMarshal CreateFile
	tmpMarshal := CreateFileMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Rename file operation
type RenameFile struct {

	// extends

	ResourceOperation

	// A rename
	Kind *string `json:"kind"`

	// The old (existing) location.
	OldUri *DocumentUri `json:"oldUri"`

	// The new location.
	NewUri *DocumentUri `json:"newUri"`

	// Rename options.
	Options *RenameFileOptions `json:"options"`
}

func (this *RenameFile) UnmarshalJSON(data []byte) error {
	type RenameFileUnmarshal RenameFile
	var tmpUnmarshal RenameFileUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("RenameFile")
	}

	if tmpUnmarshal.OldUri == nil {
		return StructureValidateFailed("RenameFile")
	}

	if tmpUnmarshal.NewUri == nil {
		return StructureValidateFailed("RenameFile")
	}

	*this = RenameFile(tmpUnmarshal)
	return nil
}

func (this *RenameFile) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("RenameFile")
	}

	if this.OldUri == nil {
		return nil, StructureValidateFailed("RenameFile")
	}

	if this.NewUri == nil {
		return nil, StructureValidateFailed("RenameFile")
	}

	type RenameFileMarshal RenameFile
	tmpMarshal := RenameFileMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Delete file operation
type DeleteFile struct {

	// extends

	ResourceOperation

	// A delete
	Kind *string `json:"kind"`

	// The file to delete.
	Uri *DocumentUri `json:"uri"`

	// Delete options.
	Options *DeleteFileOptions `json:"options"`
}

func (this *DeleteFile) UnmarshalJSON(data []byte) error {
	type DeleteFileUnmarshal DeleteFile
	var tmpUnmarshal DeleteFileUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("DeleteFile")
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("DeleteFile")
	}

	*this = DeleteFile(tmpUnmarshal)
	return nil
}

func (this *DeleteFile) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("DeleteFile")
	}

	if this.Uri == nil {
		return nil, StructureValidateFailed("DeleteFile")
	}

	type DeleteFileMarshal DeleteFile
	tmpMarshal := DeleteFileMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Additional information that describes document changes.  @since 3.16.0
type ChangeAnnotation struct {

	// A human-readable string describing the actual change. The string is
	// rendered prominent in the user interface.
	Label *String `json:"label"`

	// A flag which indicates that user confirmation is needed before
	// applying the change.
	NeedsConfirmation *Boolean `json:"needsConfirmation"`

	// A human-readable string which is rendered less prominent in the user
	// interface.
	Description *String `json:"description"`
}

func (this *ChangeAnnotation) UnmarshalJSON(data []byte) error {
	type ChangeAnnotationUnmarshal ChangeAnnotation
	var tmpUnmarshal ChangeAnnotationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("ChangeAnnotation")
	}

	*this = ChangeAnnotation(tmpUnmarshal)
	return nil
}

func (this *ChangeAnnotation) MarshalJSON() ([]byte, error) {

	if this.Label == nil {
		return nil, StructureValidateFailed("ChangeAnnotation")
	}

	type ChangeAnnotationMarshal ChangeAnnotation
	tmpMarshal := ChangeAnnotationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A filter to describe in which file operation requests or notifications the
// server is interested in receiving.  @since 3.16.0
type FileOperationFilter struct {

	// A Uri scheme like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// The actual file operation pattern.
	Pattern *FileOperationPattern `json:"pattern"`
}

func (this *FileOperationFilter) UnmarshalJSON(data []byte) error {
	type FileOperationFilterUnmarshal FileOperationFilter
	var tmpUnmarshal FileOperationFilterUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Pattern == nil {
		return StructureValidateFailed("FileOperationFilter")
	}

	*this = FileOperationFilter(tmpUnmarshal)
	return nil
}

func (this *FileOperationFilter) MarshalJSON() ([]byte, error) {

	if this.Pattern == nil {
		return nil, StructureValidateFailed("FileOperationFilter")
	}

	type FileOperationFilterMarshal FileOperationFilter
	tmpMarshal := FileOperationFilterMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents information on a file/folder rename.  @since 3.16.0
type FileRename struct {

	// A file:// URI for the original location of the file/folder being
	// renamed.
	OldUri *String `json:"oldUri"`

	// A file:// URI for the new location of the file/folder being renamed.
	NewUri *String `json:"newUri"`
}

func (this *FileRename) UnmarshalJSON(data []byte) error {
	type FileRenameUnmarshal FileRename
	var tmpUnmarshal FileRenameUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.OldUri == nil {
		return StructureValidateFailed("FileRename")
	}

	if tmpUnmarshal.NewUri == nil {
		return StructureValidateFailed("FileRename")
	}

	*this = FileRename(tmpUnmarshal)
	return nil
}

func (this *FileRename) MarshalJSON() ([]byte, error) {

	if this.OldUri == nil {
		return nil, StructureValidateFailed("FileRename")
	}

	if this.NewUri == nil {
		return nil, StructureValidateFailed("FileRename")
	}

	type FileRenameMarshal FileRename
	tmpMarshal := FileRenameMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents information on a file/folder delete.  @since 3.16.0
type FileDelete struct {

	// A file:// URI for the location of the file/folder being deleted.
	Uri *String `json:"uri"`
}

func (this *FileDelete) UnmarshalJSON(data []byte) error {
	type FileDeleteUnmarshal FileDelete
	var tmpUnmarshal FileDeleteUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("FileDelete")
	}

	*this = FileDelete(tmpUnmarshal)
	return nil
}

func (this *FileDelete) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("FileDelete")
	}

	type FileDeleteMarshal FileDelete
	tmpMarshal := FileDeleteMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	FrameId *Integer `json:"frameId"`

	// The document range where execution has stopped. Typically the end
	// position of the range denotes the line where the inline values are
	// shown.
	StoppedLocation *Range `json:"stoppedLocation"`
}

func (this *InlineValueContext) UnmarshalJSON(data []byte) error {
	type InlineValueContextUnmarshal InlineValueContext
	var tmpUnmarshal InlineValueContextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.FrameId == nil {
		return StructureValidateFailed("InlineValueContext")
	}

	if tmpUnmarshal.StoppedLocation == nil {
		return StructureValidateFailed("InlineValueContext")
	}

	*this = InlineValueContext(tmpUnmarshal)
	return nil
}

func (this *InlineValueContext) MarshalJSON() ([]byte, error) {

	if this.FrameId == nil {
		return nil, StructureValidateFailed("InlineValueContext")
	}

	if this.StoppedLocation == nil {
		return nil, StructureValidateFailed("InlineValueContext")
	}

	type InlineValueContextMarshal InlineValueContext
	tmpMarshal := InlineValueContextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provide inline value as text.  @since 3.17.0
type InlineValueText struct {

	// The document range for which the inline value applies.
	Range *Range `json:"range"`

	// The text of the inline value.
	Text *String `json:"text"`
}

func (this *InlineValueText) UnmarshalJSON(data []byte) error {
	type InlineValueTextUnmarshal InlineValueText
	var tmpUnmarshal InlineValueTextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("InlineValueText")
	}

	if tmpUnmarshal.Text == nil {
		return StructureValidateFailed("InlineValueText")
	}

	*this = InlineValueText(tmpUnmarshal)
	return nil
}

func (this *InlineValueText) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("InlineValueText")
	}

	if this.Text == nil {
		return nil, StructureValidateFailed("InlineValueText")
	}

	type InlineValueTextMarshal InlineValueText
	tmpMarshal := InlineValueTextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provide inline value through a variable lookup. If only a range is specified,
// the variable name will be extracted from the underlying document. An optional
// variable name can be used to override the extracted name.  @since 3.17.0
type InlineValueVariableLookup struct {

	// The document range for which the inline value applies. The range is
	// used to extract the variable name from the underlying document.
	Range *Range `json:"range"`

	// If specified the name of the variable to look up.
	VariableName *String `json:"variableName"`

	// How to perform the lookup.
	CaseSensitiveLookup *Boolean `json:"caseSensitiveLookup"`
}

func (this *InlineValueVariableLookup) UnmarshalJSON(data []byte) error {
	type InlineValueVariableLookupUnmarshal InlineValueVariableLookup
	var tmpUnmarshal InlineValueVariableLookupUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("InlineValueVariableLookup")
	}

	if tmpUnmarshal.CaseSensitiveLookup == nil {
		return StructureValidateFailed("InlineValueVariableLookup")
	}

	*this = InlineValueVariableLookup(tmpUnmarshal)
	return nil
}

func (this *InlineValueVariableLookup) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("InlineValueVariableLookup")
	}

	if this.CaseSensitiveLookup == nil {
		return nil, StructureValidateFailed("InlineValueVariableLookup")
	}

	type InlineValueVariableLookupMarshal InlineValueVariableLookup
	tmpMarshal := InlineValueVariableLookupMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provide an inline value through an expression evaluation. If only a range is
// specified, the expression will be extracted from the underlying document. An
// optional expression can be used to override the extracted expression.  @since
// 3.17.0
type InlineValueEvaluatableExpression struct {

	// The document range for which the inline value applies. The range is
	// used to extract the evaluatable expression from the underlying
	// document.
	Range *Range `json:"range"`

	// If specified the expression overrides the extracted expression.
	Expression *String `json:"expression"`
}

func (this *InlineValueEvaluatableExpression) UnmarshalJSON(data []byte) error {
	type InlineValueEvaluatableExpressionUnmarshal InlineValueEvaluatableExpression
	var tmpUnmarshal InlineValueEvaluatableExpressionUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed(
			"InlineValueEvaluatableExpression",
		)
	}

	*this = InlineValueEvaluatableExpression(tmpUnmarshal)
	return nil
}

func (this *InlineValueEvaluatableExpression) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed(
			"InlineValueEvaluatableExpression",
		)
	}

	type InlineValueEvaluatableExpressionMarshal InlineValueEvaluatableExpression
	tmpMarshal := InlineValueEvaluatableExpressionMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Value *String `json:"value"`

	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might
	// resolve this property late using the resolve request.
	Tooltip *InlayHintLabelPart_Tooltip_Or `json:"tooltip"`

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

func (this *InlayHintLabelPart) UnmarshalJSON(data []byte) error {
	type InlayHintLabelPartUnmarshal InlayHintLabelPart
	var tmpUnmarshal InlayHintLabelPartUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Value == nil {
		return StructureValidateFailed("InlayHintLabelPart")
	}

	*this = InlayHintLabelPart(tmpUnmarshal)
	return nil
}

func (this *InlayHintLabelPart) MarshalJSON() ([]byte, error) {

	if this.Value == nil {
		return nil, StructureValidateFailed("InlayHintLabelPart")
	}

	type InlayHintLabelPartMarshal InlayHintLabelPart
	tmpMarshal := InlayHintLabelPartMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Kind *MarkupKind `json:"kind"`

	// The content itself
	Value *String `json:"value"`
}

func (this *MarkupContent) UnmarshalJSON(data []byte) error {
	type MarkupContentUnmarshal MarkupContent
	var tmpUnmarshal MarkupContentUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("MarkupContent")
	}

	if tmpUnmarshal.Value == nil {
		return StructureValidateFailed("MarkupContent")
	}

	*this = MarkupContent(tmpUnmarshal)
	return nil
}

func (this *MarkupContent) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("MarkupContent")
	}

	if this.Value == nil {
		return nil, StructureValidateFailed("MarkupContent")
	}

	type MarkupContentMarshal MarkupContent
	tmpMarshal := MarkupContentMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Inlay hint options used during static registration.  @since 3.17.0
type InlayHintOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The server provides support to resolve additional information for an
	// inlay hint item.
	ResolveProvider *Boolean `json:"resolveProvider"`
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
	RelatedDocuments map[DocumentUri]RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or `json:"relatedDocuments"`
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
	RelatedDocuments map[DocumentUri]RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or `json:"relatedDocuments"`
}

// A diagnostic report with a full set of problems.  @since 3.17.0
type FullDocumentDiagnosticReport struct {

	// A full document diagnostic report.
	Kind *string `json:"kind"`

	// An optional result id. If provided it will be sent on the next
	// diagnostic request for the same document.
	ResultId *String `json:"resultId"`

	// The actual items.
	Items []Diagnostic `json:"items"`
}

func (this *FullDocumentDiagnosticReport) UnmarshalJSON(data []byte) error {
	type FullDocumentDiagnosticReportUnmarshal FullDocumentDiagnosticReport
	var tmpUnmarshal FullDocumentDiagnosticReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("FullDocumentDiagnosticReport")
	}

	if tmpUnmarshal.Items == nil {
		return StructureValidateFailed("FullDocumentDiagnosticReport")
	}

	*this = FullDocumentDiagnosticReport(tmpUnmarshal)
	return nil
}

func (this *FullDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed(
			"FullDocumentDiagnosticReport",
		)
	}

	if this.Items == nil {
		return nil, StructureValidateFailed(
			"FullDocumentDiagnosticReport",
		)
	}

	type FullDocumentDiagnosticReportMarshal FullDocumentDiagnosticReport
	tmpMarshal := FullDocumentDiagnosticReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A diagnostic report indicating that the last returned report is still
// accurate.  @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {

	// A document diagnostic report indicating no changes to the last
	// result. A server can only return `unchanged` if result ids are
	// provided.
	Kind *string `json:"kind"`

	// A result id which will be sent on the next diagnostic request for the
	// same document.
	ResultId *String `json:"resultId"`
}

func (this *UnchangedDocumentDiagnosticReport) UnmarshalJSON(
	data []byte,
) error {
	type UnchangedDocumentDiagnosticReportUnmarshal UnchangedDocumentDiagnosticReport
	var tmpUnmarshal UnchangedDocumentDiagnosticReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed(
			"UnchangedDocumentDiagnosticReport",
		)
	}

	if tmpUnmarshal.ResultId == nil {
		return StructureValidateFailed(
			"UnchangedDocumentDiagnosticReport",
		)
	}

	*this = UnchangedDocumentDiagnosticReport(tmpUnmarshal)
	return nil
}

func (this *UnchangedDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed(
			"UnchangedDocumentDiagnosticReport",
		)
	}

	if this.ResultId == nil {
		return nil, StructureValidateFailed(
			"UnchangedDocumentDiagnosticReport",
		)
	}

	type UnchangedDocumentDiagnosticReportMarshal UnchangedDocumentDiagnosticReport
	tmpMarshal := UnchangedDocumentDiagnosticReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Diagnostic options.  @since 3.17.0
type DiagnosticOptions struct {

	// mixins

	WorkDoneProgressOptions

	// An optional identifier under which the diagnostics are managed by the
	// client.
	Identifier *String `json:"identifier"`

	// Whether the language has inter file dependencies meaning that editing
	// code in one file can result in a different diagnostic set in another
	// file. Inter file dependencies are common for most programming
	// languages and typically uncommon for linters.
	InterFileDependencies *Boolean `json:"interFileDependencies"`

	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics *Boolean `json:"workspaceDiagnostics"`
}

func (this *DiagnosticOptions) UnmarshalJSON(data []byte) error {
	type DiagnosticOptionsUnmarshal DiagnosticOptions
	var tmpUnmarshal DiagnosticOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.InterFileDependencies == nil {
		return StructureValidateFailed("DiagnosticOptions")
	}

	if tmpUnmarshal.WorkspaceDiagnostics == nil {
		return StructureValidateFailed("DiagnosticOptions")
	}

	*this = DiagnosticOptions(tmpUnmarshal)
	return nil
}

func (this *DiagnosticOptions) MarshalJSON() ([]byte, error) {

	if this.InterFileDependencies == nil {
		return nil, StructureValidateFailed("DiagnosticOptions")
	}

	if this.WorkspaceDiagnostics == nil {
		return nil, StructureValidateFailed("DiagnosticOptions")
	}

	type DiagnosticOptionsMarshal DiagnosticOptions
	tmpMarshal := DiagnosticOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A previous result id in a workspace pull request.  @since 3.17.0
type PreviousResultId struct {

	// The URI for which the client knowns a result id.
	Uri *DocumentUri `json:"uri"`

	// The value of the previous result id.
	Value *String `json:"value"`
}

func (this *PreviousResultId) UnmarshalJSON(data []byte) error {
	type PreviousResultIdUnmarshal PreviousResultId
	var tmpUnmarshal PreviousResultIdUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("PreviousResultId")
	}

	if tmpUnmarshal.Value == nil {
		return StructureValidateFailed("PreviousResultId")
	}

	*this = PreviousResultId(tmpUnmarshal)
	return nil
}

func (this *PreviousResultId) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("PreviousResultId")
	}

	if this.Value == nil {
		return nil, StructureValidateFailed("PreviousResultId")
	}

	type PreviousResultIdMarshal PreviousResultId
	tmpMarshal := PreviousResultIdMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A notebook document.  @since 3.17.0
type NotebookDocument struct {

	// The notebook document's uri.
	Uri *URI `json:"uri"`

	// The type of the notebook.
	NotebookType *String `json:"notebookType"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version *Integer `json:"version"`

	// Additional metadata stored with the notebook document.  Note: should
	// always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata"`

	// The cells of a notebook.
	Cells []NotebookCell `json:"cells"`
}

func (this *NotebookDocument) UnmarshalJSON(data []byte) error {
	type NotebookDocumentUnmarshal NotebookDocument
	var tmpUnmarshal NotebookDocumentUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("NotebookDocument")
	}

	if tmpUnmarshal.NotebookType == nil {
		return StructureValidateFailed("NotebookDocument")
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed("NotebookDocument")
	}

	if tmpUnmarshal.Cells == nil {
		return StructureValidateFailed("NotebookDocument")
	}

	*this = NotebookDocument(tmpUnmarshal)
	return nil
}

func (this *NotebookDocument) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("NotebookDocument")
	}

	if this.NotebookType == nil {
		return nil, StructureValidateFailed("NotebookDocument")
	}

	if this.Version == nil {
		return nil, StructureValidateFailed("NotebookDocument")
	}

	if this.Cells == nil {
		return nil, StructureValidateFailed("NotebookDocument")
	}

	type NotebookDocumentMarshal NotebookDocument
	tmpMarshal := NotebookDocumentMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// An item to transfer a text document from the client to the server.
type TextDocumentItem struct {

	// The text document's uri.
	Uri *DocumentUri `json:"uri"`

	// The text document's language identifier.
	LanguageId *String `json:"languageId"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version *Integer `json:"version"`

	// The content of the opened text document.
	Text *String `json:"text"`
}

func (this *TextDocumentItem) UnmarshalJSON(data []byte) error {
	type TextDocumentItemUnmarshal TextDocumentItem
	var tmpUnmarshal TextDocumentItemUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("TextDocumentItem")
	}

	if tmpUnmarshal.LanguageId == nil {
		return StructureValidateFailed("TextDocumentItem")
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed("TextDocumentItem")
	}

	if tmpUnmarshal.Text == nil {
		return StructureValidateFailed("TextDocumentItem")
	}

	*this = TextDocumentItem(tmpUnmarshal)
	return nil
}

func (this *TextDocumentItem) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("TextDocumentItem")
	}

	if this.LanguageId == nil {
		return nil, StructureValidateFailed("TextDocumentItem")
	}

	if this.Version == nil {
		return nil, StructureValidateFailed("TextDocumentItem")
	}

	if this.Text == nil {
		return nil, StructureValidateFailed("TextDocumentItem")
	}

	type TextDocumentItemMarshal TextDocumentItem
	tmpMarshal := TextDocumentItemMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A versioned notebook document identifier.  @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {

	// The version number of this notebook document.
	Version *Integer `json:"version"`

	// The notebook document's uri.
	Uri *URI `json:"uri"`
}

func (this *VersionedNotebookDocumentIdentifier) UnmarshalJSON(
	data []byte,
) error {
	type VersionedNotebookDocumentIdentifierUnmarshal VersionedNotebookDocumentIdentifier
	var tmpUnmarshal VersionedNotebookDocumentIdentifierUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed(
			"VersionedNotebookDocumentIdentifier",
		)
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed(
			"VersionedNotebookDocumentIdentifier",
		)
	}

	*this = VersionedNotebookDocumentIdentifier(tmpUnmarshal)
	return nil
}

func (this *VersionedNotebookDocumentIdentifier) MarshalJSON() ([]byte, error) {

	if this.Version == nil {
		return nil, StructureValidateFailed(
			"VersionedNotebookDocumentIdentifier",
		)
	}

	if this.Uri == nil {
		return nil, StructureValidateFailed(
			"VersionedNotebookDocumentIdentifier",
		)
	}

	type VersionedNotebookDocumentIdentifierMarshal VersionedNotebookDocumentIdentifier
	tmpMarshal := VersionedNotebookDocumentIdentifierMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Uri *URI `json:"uri"`
}

func (this *NotebookDocumentIdentifier) UnmarshalJSON(data []byte) error {
	type NotebookDocumentIdentifierUnmarshal NotebookDocumentIdentifier
	var tmpUnmarshal NotebookDocumentIdentifierUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("NotebookDocumentIdentifier")
	}

	*this = NotebookDocumentIdentifier(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentIdentifier) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed(
			"NotebookDocumentIdentifier",
		)
	}

	type NotebookDocumentIdentifierMarshal NotebookDocumentIdentifier
	tmpMarshal := NotebookDocumentIdentifierMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// General parameters to to register for an notification or to register a
// provider.
type Registration struct {

	// The id used to register the request. The id can be used to deregister
	// the request again.
	Id *String `json:"id"`

	// The method / capability to register for.
	Method *String `json:"method"`

	// Options necessary for the registration.
	RegisterOptions *LSPAny `json:"registerOptions"`
}

func (this *Registration) UnmarshalJSON(data []byte) error {
	type RegistrationUnmarshal Registration
	var tmpUnmarshal RegistrationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Id == nil {
		return StructureValidateFailed("Registration")
	}

	if tmpUnmarshal.Method == nil {
		return StructureValidateFailed("Registration")
	}

	*this = Registration(tmpUnmarshal)
	return nil
}

func (this *Registration) MarshalJSON() ([]byte, error) {

	if this.Id == nil {
		return nil, StructureValidateFailed("Registration")
	}

	if this.Method == nil {
		return nil, StructureValidateFailed("Registration")
	}

	type RegistrationMarshal Registration
	tmpMarshal := RegistrationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// General parameters to unregister a request or notification.
type Unregistration struct {

	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id *String `json:"id"`

	// The method to unregister for.
	Method *String `json:"method"`
}

func (this *Unregistration) UnmarshalJSON(data []byte) error {
	type UnregistrationUnmarshal Unregistration
	var tmpUnmarshal UnregistrationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Id == nil {
		return StructureValidateFailed("Unregistration")
	}

	if tmpUnmarshal.Method == nil {
		return StructureValidateFailed("Unregistration")
	}

	*this = Unregistration(tmpUnmarshal)
	return nil
}

func (this *Unregistration) MarshalJSON() ([]byte, error) {

	if this.Id == nil {
		return nil, StructureValidateFailed("Unregistration")
	}

	if this.Method == nil {
		return nil, StructureValidateFailed("Unregistration")
	}

	type UnregistrationMarshal Unregistration
	tmpMarshal := UnregistrationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// The initialize parameters
type XInitializeParams struct {

	// mixins

	WorkDoneProgressParams

	// The process Id of the parent process that started the server.  Is
	// `null` if the process has not been started by another process. If the
	// parent process is not alive then the server should exit.
	ProcessId *XInitializeParams_ProcessId_Or `json:"processId"`

	// Information about the client  @since 3.15.0
	ClientInfo *XInitializeParams_ClientInfo `json:"clientInfo"`

	// The locale the client is currently showing the user interface in.
	// This must not necessarily be the locale of the operating system.
	// Uses IETF language tags as the value's syntax (See
	// https://en.wikipedia.org/wiki/IETF_language_tag)  @since 3.16.0
	Locale *String `json:"locale"`

	// The rootPath of the workspace. Is null if no folder is open.
	// @deprecated in favour of rootUri.
	RootPath *XInitializeParams_RootPath_Or `json:"rootPath"`

	// The rootUri of the workspace. Is null if no folder is open. If both
	// `rootPath` and `rootUri` are set `rootUri` wins.  @deprecated in
	// favour of workspaceFolders.
	RootUri *XInitializeParams_RootUri_Or `json:"rootUri"`

	// The capabilities provided by the client (editor or tool)
	Capabilities *ClientCapabilities `json:"capabilities"`

	// User provided initialization options.
	InitializationOptions *LSPAny `json:"initializationOptions"`

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValues `json:"trace"`
}

func (this *XInitializeParams) UnmarshalJSON(data []byte) error {
	type XInitializeParamsUnmarshal XInitializeParams
	var tmpUnmarshal XInitializeParamsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ProcessId == nil {
		return StructureValidateFailed("XInitializeParams")
	}

	if tmpUnmarshal.RootUri == nil {
		return StructureValidateFailed("XInitializeParams")
	}

	if tmpUnmarshal.Capabilities == nil {
		return StructureValidateFailed("XInitializeParams")
	}

	*this = XInitializeParams(tmpUnmarshal)
	return nil
}

func (this *XInitializeParams) MarshalJSON() ([]byte, error) {

	if this.ProcessId == nil {
		return nil, StructureValidateFailed("XInitializeParams")
	}

	if this.RootUri == nil {
		return nil, StructureValidateFailed("XInitializeParams")
	}

	if this.Capabilities == nil {
		return nil, StructureValidateFailed("XInitializeParams")
	}

	type XInitializeParamsMarshal XInitializeParams
	tmpMarshal := XInitializeParamsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkspaceFoldersInitializeParams struct {

	// The workspace folders configured in the client when the server
	// starts.  This property is only available if the client supports
	// workspace folders. It can be `null` if the client supports workspace
	// folders but none are configured.  @since 3.6.0
	WorkspaceFolders *WorkspaceFoldersInitializeParams_WorkspaceFolders_Or `json:"workspaceFolders"`
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
	TextDocumentSync *ServerCapabilities_TextDocumentSync_Or `json:"textDocumentSync"`

	// Defines how notebook documents are synced.  @since 3.17.0
	NotebookDocumentSync *ServerCapabilities_NotebookDocumentSync_Or `json:"notebookDocumentSync"`

	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider"`

	// The server provides hover support.
	HoverProvider *ServerCapabilities_HoverProvider_Or `json:"hoverProvider"`

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider"`

	// The server provides Goto Declaration support.
	DeclarationProvider *ServerCapabilities_DeclarationProvider_Or `json:"declarationProvider"`

	// The server provides goto definition support.
	DefinitionProvider *ServerCapabilities_DefinitionProvider_Or `json:"definitionProvider"`

	// The server provides Goto Type Definition support.
	TypeDefinitionProvider *ServerCapabilities_TypeDefinitionProvider_Or `json:"typeDefinitionProvider"`

	// The server provides Goto Implementation support.
	ImplementationProvider *ServerCapabilities_ImplementationProvider_Or `json:"implementationProvider"`

	// The server provides find references support.
	ReferencesProvider *ServerCapabilities_ReferencesProvider_Or `json:"referencesProvider"`

	// The server provides document highlight support.
	DocumentHighlightProvider *ServerCapabilities_DocumentHighlightProvider_Or `json:"documentHighlightProvider"`

	// The server provides document symbol support.
	DocumentSymbolProvider *ServerCapabilities_DocumentSymbolProvider_Or `json:"documentSymbolProvider"`

	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *ServerCapabilities_CodeActionProvider_Or `json:"codeActionProvider"`

	// The server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider"`

	// The server provides color provider support.
	ColorProvider *ServerCapabilities_ColorProvider_Or `json:"colorProvider"`

	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *ServerCapabilities_WorkspaceSymbolProvider_Or `json:"workspaceSymbolProvider"`

	// The server provides document formatting.
	DocumentFormattingProvider *ServerCapabilities_DocumentFormattingProvider_Or `json:"documentFormattingProvider"`

	// The server provides document range formatting.
	DocumentRangeFormattingProvider *ServerCapabilities_DocumentRangeFormattingProvider_Or `json:"documentRangeFormattingProvider"`

	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider"`

	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports `prepareSupport` in
	// its initial `initialize` request.
	RenameProvider *ServerCapabilities_RenameProvider_Or `json:"renameProvider"`

	// The server provides folding provider support.
	FoldingRangeProvider *ServerCapabilities_FoldingRangeProvider_Or `json:"foldingRangeProvider"`

	// The server provides selection range support.
	SelectionRangeProvider *ServerCapabilities_SelectionRangeProvider_Or `json:"selectionRangeProvider"`

	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider"`

	// The server provides call hierarchy support.  @since 3.16.0
	CallHierarchyProvider *ServerCapabilities_CallHierarchyProvider_Or `json:"callHierarchyProvider"`

	// The server provides linked editing range support.  @since 3.16.0
	LinkedEditingRangeProvider *ServerCapabilities_LinkedEditingRangeProvider_Or `json:"linkedEditingRangeProvider"`

	// The server provides semantic tokens support.  @since 3.16.0
	SemanticTokensProvider *ServerCapabilities_SemanticTokensProvider_Or `json:"semanticTokensProvider"`

	// The server provides moniker support.  @since 3.16.0
	MonikerProvider *ServerCapabilities_MonikerProvider_Or `json:"monikerProvider"`

	// The server provides type hierarchy support.  @since 3.17.0
	TypeHierarchyProvider *ServerCapabilities_TypeHierarchyProvider_Or `json:"typeHierarchyProvider"`

	// The server provides inline values.  @since 3.17.0
	InlineValueProvider *ServerCapabilities_InlineValueProvider_Or `json:"inlineValueProvider"`

	// The server provides inlay hints.  @since 3.17.0
	InlayHintProvider *ServerCapabilities_InlayHintProvider_Or `json:"inlayHintProvider"`

	// The server has support for pull model diagnostics.  @since 3.17.0
	DiagnosticProvider *ServerCapabilities_DiagnosticProvider_Or `json:"diagnosticProvider"`

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
	Version *Integer `json:"version"`
}

func (this *VersionedTextDocumentIdentifier) UnmarshalJSON(data []byte) error {
	type VersionedTextDocumentIdentifierUnmarshal VersionedTextDocumentIdentifier
	var tmpUnmarshal VersionedTextDocumentIdentifierUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed(
			"VersionedTextDocumentIdentifier",
		)
	}

	*this = VersionedTextDocumentIdentifier(tmpUnmarshal)
	return nil
}

func (this *VersionedTextDocumentIdentifier) MarshalJSON() ([]byte, error) {

	if this.Version == nil {
		return nil, StructureValidateFailed(
			"VersionedTextDocumentIdentifier",
		)
	}

	type VersionedTextDocumentIdentifierMarshal VersionedTextDocumentIdentifier
	tmpMarshal := VersionedTextDocumentIdentifierMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Save options.
type SaveOptions struct {

	// The client is supposed to include the content on save.
	IncludeText *Boolean `json:"includeText"`
}

// An event describing a file change.
type FileEvent struct {

	// The file's uri.
	Uri *DocumentUri `json:"uri"`

	// The change type.
	Type *FileChangeType `json:"type"`
}

func (this *FileEvent) UnmarshalJSON(data []byte) error {
	type FileEventUnmarshal FileEvent
	var tmpUnmarshal FileEventUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed("FileEvent")
	}

	if tmpUnmarshal.Type == nil {
		return StructureValidateFailed("FileEvent")
	}

	*this = FileEvent(tmpUnmarshal)
	return nil
}

func (this *FileEvent) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed("FileEvent")
	}

	if this.Type == nil {
		return nil, StructureValidateFailed("FileEvent")
	}

	type FileEventMarshal FileEvent
	tmpMarshal := FileEventMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type FileSystemWatcher struct {

	// The glob pattern to watch. See {@link GlobPattern glob pattern} for
	// more detail.  @since 3.17.0 support for relative patterns.
	GlobPattern *GlobPattern `json:"globPattern"`

	// The kind of events of interest. If omitted it defaults to
	// WatchKind.Create | WatchKind.Change | WatchKind.Delete which is 7.
	Kind *WatchKind `json:"kind"`
}

func (this *FileSystemWatcher) UnmarshalJSON(data []byte) error {
	type FileSystemWatcherUnmarshal FileSystemWatcher
	var tmpUnmarshal FileSystemWatcherUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.GlobPattern == nil {
		return StructureValidateFailed("FileSystemWatcher")
	}

	*this = FileSystemWatcher(tmpUnmarshal)
	return nil
}

func (this *FileSystemWatcher) MarshalJSON() ([]byte, error) {

	if this.GlobPattern == nil {
		return nil, StructureValidateFailed("FileSystemWatcher")
	}

	type FileSystemWatcherMarshal FileSystemWatcher
	tmpMarshal := FileSystemWatcherMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a diagnostic, such as a compiler error or warning. Diagnostic
// objects are only valid in the scope of a resource.
type Diagnostic struct {

	// The range at which the message applies
	Range *Range `json:"range"`

	// The diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity *DiagnosticSeverity `json:"severity"`

	// The diagnostic's code, which usually appear in the user interface.
	Code *Diagnostic_Code_Or `json:"code"`

	// An optional property to describe the error code. Requires the code
	// field (above) to be present/not null.  @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription"`

	// A human-readable string describing the source of this diagnostic,
	// e.g. 'typescript' or 'super lint'. It usually appears in the user
	// interface.
	Source *String `json:"source"`

	// The diagnostic's message. It usually appears in the user interface
	Message *String `json:"message"`

	// Additional metadata about the diagnostic.  @since 3.15.0
	Tags []DiagnosticTag `json:"tags"`

	// An array of related diagnostic information, e.g. when symbol-names
	// within a scope collide all definitions can be marked via this
	// property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation"`

	// A data entry field that is preserved between a
	// `textDocument/publishDiagnostics` notification and
	// `textDocument/codeAction` request.  @since 3.16.0
	Data *LSPAny `json:"data"`
}

func (this *Diagnostic) UnmarshalJSON(data []byte) error {
	type DiagnosticUnmarshal Diagnostic
	var tmpUnmarshal DiagnosticUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureValidateFailed("Diagnostic")
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("Diagnostic")
	}

	*this = Diagnostic(tmpUnmarshal)
	return nil
}

func (this *Diagnostic) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureValidateFailed("Diagnostic")
	}

	if this.Message == nil {
		return nil, StructureValidateFailed("Diagnostic")
	}

	type DiagnosticMarshal Diagnostic
	tmpMarshal := DiagnosticMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Contains additional information about the context in which a completion
// request is triggered.
type CompletionContext struct {

	// How the completion was triggered.
	TriggerKind *CompletionTriggerKind `json:"triggerKind"`

	// The trigger character (a single character) that has trigger code
	// complete. Is undefined if `triggerKind !==
	// CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter *String `json:"triggerCharacter"`
}

func (this *CompletionContext) UnmarshalJSON(data []byte) error {
	type CompletionContextUnmarshal CompletionContext
	var tmpUnmarshal CompletionContextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TriggerKind == nil {
		return StructureValidateFailed("CompletionContext")
	}

	*this = CompletionContext(tmpUnmarshal)
	return nil
}

func (this *CompletionContext) MarshalJSON() ([]byte, error) {

	if this.TriggerKind == nil {
		return nil, StructureValidateFailed("CompletionContext")
	}

	type CompletionContextMarshal CompletionContext
	tmpMarshal := CompletionContextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Additional details for a completion item label.  @since 3.17.0
type CompletionItemLabelDetails struct {

	// An optional string which is rendered less prominently directly after
	// {@link CompletionItem.label label}, without any spacing. Should be
	// used for function signatures and type annotations.
	Detail *String `json:"detail"`

	// An optional string which is rendered less prominently after {@link
	// CompletionItem.detail}. Should be used for fully qualified names and
	// file paths.
	Description *String `json:"description"`
}

// A special text edit to provide an insert and a replace operation.  @since
// 3.16.0
type InsertReplaceEdit struct {

	// The string to be inserted.
	NewText *String `json:"newText"`

	// The range if the insert is requested
	Insert *Range `json:"insert"`

	// The range if the replace is requested.
	Replace *Range `json:"replace"`
}

func (this *InsertReplaceEdit) UnmarshalJSON(data []byte) error {
	type InsertReplaceEditUnmarshal InsertReplaceEdit
	var tmpUnmarshal InsertReplaceEditUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NewText == nil {
		return StructureValidateFailed("InsertReplaceEdit")
	}

	if tmpUnmarshal.Insert == nil {
		return StructureValidateFailed("InsertReplaceEdit")
	}

	if tmpUnmarshal.Replace == nil {
		return StructureValidateFailed("InsertReplaceEdit")
	}

	*this = InsertReplaceEdit(tmpUnmarshal)
	return nil
}

func (this *InsertReplaceEdit) MarshalJSON() ([]byte, error) {

	if this.NewText == nil {
		return nil, StructureValidateFailed("InsertReplaceEdit")
	}

	if this.Insert == nil {
		return nil, StructureValidateFailed("InsertReplaceEdit")
	}

	if this.Replace == nil {
		return nil, StructureValidateFailed("InsertReplaceEdit")
	}

	type InsertReplaceEditMarshal InsertReplaceEdit
	tmpMarshal := InsertReplaceEditMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	TriggerCharacters []String `json:"triggerCharacters"`

	// The list of all possible characters that commit a completion. This
	// field can be used if clients don't support individual commit
	// characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	//  If a server provides both `allCommitCharacters` and commit
	// characters on an individual completion item the ones on the
	// completion item win.  @since 3.2.0
	AllCommitCharacters []String `json:"allCommitCharacters"`

	// The server provides support to resolve additional information for a
	// completion item.
	ResolveProvider *Boolean `json:"resolveProvider"`

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
	TriggerKind *SignatureHelpTriggerKind `json:"triggerKind"`

	// Character that caused signature help to be triggered.  This is
	// undefined when `triggerKind !==
	// SignatureHelpTriggerKind.TriggerCharacter`
	TriggerCharacter *String `json:"triggerCharacter"`

	// `true` if signature help was already showing when it was triggered.
	// Retriggers occurs when the signature help is already active and can
	// be caused by actions such as typing a trigger character, a cursor
	// move, or document content changes.
	IsRetrigger *Boolean `json:"isRetrigger"`

	// The currently active `SignatureHelp`.  The `activeSignatureHelp` has
	// its `SignatureHelp.activeSignature` field updated based on the user
	// navigating through available signatures.
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp"`
}

func (this *SignatureHelpContext) UnmarshalJSON(data []byte) error {
	type SignatureHelpContextUnmarshal SignatureHelpContext
	var tmpUnmarshal SignatureHelpContextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TriggerKind == nil {
		return StructureValidateFailed("SignatureHelpContext")
	}

	if tmpUnmarshal.IsRetrigger == nil {
		return StructureValidateFailed("SignatureHelpContext")
	}

	*this = SignatureHelpContext(tmpUnmarshal)
	return nil
}

func (this *SignatureHelpContext) MarshalJSON() ([]byte, error) {

	if this.TriggerKind == nil {
		return nil, StructureValidateFailed("SignatureHelpContext")
	}

	if this.IsRetrigger == nil {
		return nil, StructureValidateFailed("SignatureHelpContext")
	}

	type SignatureHelpContextMarshal SignatureHelpContext
	tmpMarshal := SignatureHelpContextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents the signature of something callable. A signature can have a label,
// like a function-name, a doc-comment, and a set of parameters.
type SignatureInformation struct {

	// The label of this signature. Will be shown in the UI.
	Label *String `json:"label"`

	// The human-readable doc-comment of this signature. Will be shown in
	// the UI but can be omitted.
	Documentation *SignatureInformation_Documentation_Or `json:"documentation"`

	// The parameters of this signature.
	Parameters []ParameterInformation `json:"parameters"`

	// The index of the active parameter.  If provided, this is used in
	// place of `SignatureHelp.activeParameter`.  @since 3.16.0
	ActiveParameter *Uinteger `json:"activeParameter"`
}

func (this *SignatureInformation) UnmarshalJSON(data []byte) error {
	type SignatureInformationUnmarshal SignatureInformation
	var tmpUnmarshal SignatureInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("SignatureInformation")
	}

	*this = SignatureInformation(tmpUnmarshal)
	return nil
}

func (this *SignatureInformation) MarshalJSON() ([]byte, error) {

	if this.Label == nil {
		return nil, StructureValidateFailed("SignatureInformation")
	}

	type SignatureInformationMarshal SignatureInformation
	tmpMarshal := SignatureInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Server Capabilities for a {@link SignatureHelpRequest}.
type SignatureHelpOptions struct {

	// mixins

	WorkDoneProgressOptions

	// List of characters that trigger signature help automatically.
	TriggerCharacters []String `json:"triggerCharacters"`

	// List of characters that re-trigger signature help.  These trigger
	// characters are only active when signature help is already showing.
	// All trigger characters are also counted as re-trigger characters.
	// @since 3.15.0
	RetriggerCharacters []String `json:"retriggerCharacters"`
}

// Server Capabilities for a {@link DefinitionRequest}.
type DefinitionOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Value-object that contains additional information when requesting references.
type ReferenceContext struct {

	// Include the declaration of the current symbol.
	IncludeDeclaration *Boolean `json:"includeDeclaration"`
}

func (this *ReferenceContext) UnmarshalJSON(data []byte) error {
	type ReferenceContextUnmarshal ReferenceContext
	var tmpUnmarshal ReferenceContextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.IncludeDeclaration == nil {
		return StructureValidateFailed("ReferenceContext")
	}

	*this = ReferenceContext(tmpUnmarshal)
	return nil
}

func (this *ReferenceContext) MarshalJSON() ([]byte, error) {

	if this.IncludeDeclaration == nil {
		return nil, StructureValidateFailed("ReferenceContext")
	}

	type ReferenceContextMarshal ReferenceContext
	tmpMarshal := ReferenceContextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Reference options.
type ReferenceOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a {@link DocumentHighlightRequest}.
type DocumentHighlightOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// A base for all symbol information.
type BaseSymbolInformation struct {

	// The name of this symbol.
	Name *String `json:"name"`

	// The kind of this symbol.
	Kind *SymbolKind `json:"kind"`

	// Tags for this symbol.  @since 3.16.0
	Tags []SymbolTag `json:"tags"`

	// The name of the symbol containing this symbol. This information is
	// for user interface purposes (e.g. to render a qualifier in the user
	// interface if necessary). It can't be used to re-infer a hierarchy for
	// the document symbols.
	ContainerName *String `json:"containerName"`
}

func (this *BaseSymbolInformation) UnmarshalJSON(data []byte) error {
	type BaseSymbolInformationUnmarshal BaseSymbolInformation
	var tmpUnmarshal BaseSymbolInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureValidateFailed("BaseSymbolInformation")
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("BaseSymbolInformation")
	}

	*this = BaseSymbolInformation(tmpUnmarshal)
	return nil
}

func (this *BaseSymbolInformation) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureValidateFailed("BaseSymbolInformation")
	}

	if this.Kind == nil {
		return nil, StructureValidateFailed("BaseSymbolInformation")
	}

	type BaseSymbolInformationMarshal BaseSymbolInformation
	tmpMarshal := BaseSymbolInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provider options for a {@link DocumentSymbolRequest}.
type DocumentSymbolOptions struct {

	// mixins

	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.  @since 3.16.0
	Label *String `json:"label"`
}

// Contains additional diagnostic information about the context in which a
// {@link CodeActionProvider.provideCodeActions code action} is run.
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
	Only []CodeActionKind `json:"only"`

	// The reason why code actions were requested.  @since 3.17.0
	TriggerKind *CodeActionTriggerKind `json:"triggerKind"`
}

func (this *CodeActionContext) UnmarshalJSON(data []byte) error {
	type CodeActionContextUnmarshal CodeActionContext
	var tmpUnmarshal CodeActionContextUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Diagnostics == nil {
		return StructureValidateFailed("CodeActionContext")
	}

	*this = CodeActionContext(tmpUnmarshal)
	return nil
}

func (this *CodeActionContext) MarshalJSON() ([]byte, error) {

	if this.Diagnostics == nil {
		return nil, StructureValidateFailed("CodeActionContext")
	}

	type CodeActionContextMarshal CodeActionContext
	tmpMarshal := CodeActionContextMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provider options for a {@link CodeActionRequest}.
type CodeActionOptions struct {

	// mixins

	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.  The list of kinds may
	// be generic, such as `CodeActionKind.Refactor`, or the server may list
	// out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds"`

	// The server provides support to resolve additional information for a
	// code action.  @since 3.16.0
	ResolveProvider *Boolean `json:"resolveProvider"`
}

// Server capabilities for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The server provides support to resolve additional information for a
	// workspace symbol.  @since 3.17.0
	ResolveProvider *Boolean `json:"resolveProvider"`
}

// Code Lens provider options of a {@link CodeLensRequest}.
type CodeLensOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *Boolean `json:"resolveProvider"`
}

// Provider options for a {@link DocumentLinkRequest}.
type DocumentLinkOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Document links have a resolve provider as well.
	ResolveProvider *Boolean `json:"resolveProvider"`
}

// Value-object describing what options formatting should use.
type FormattingOptions struct {

	// Size of a tab in spaces.
	TabSize *Uinteger `json:"tabSize"`

	// Prefer spaces over tabs.
	InsertSpaces *Boolean `json:"insertSpaces"`

	// Trim trailing whitespace on a line.  @since 3.15.0
	TrimTrailingWhitespace *Boolean `json:"trimTrailingWhitespace"`

	// Insert a newline character at the end of the file if one does not
	// exist.  @since 3.15.0
	InsertFinalNewline *Boolean `json:"insertFinalNewline"`

	// Trim all newlines after the final newline at the end of the file.
	// @since 3.15.0
	TrimFinalNewlines *Boolean `json:"trimFinalNewlines"`
}

func (this *FormattingOptions) UnmarshalJSON(data []byte) error {
	type FormattingOptionsUnmarshal FormattingOptions
	var tmpUnmarshal FormattingOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TabSize == nil {
		return StructureValidateFailed("FormattingOptions")
	}

	if tmpUnmarshal.InsertSpaces == nil {
		return StructureValidateFailed("FormattingOptions")
	}

	*this = FormattingOptions(tmpUnmarshal)
	return nil
}

func (this *FormattingOptions) MarshalJSON() ([]byte, error) {

	if this.TabSize == nil {
		return nil, StructureValidateFailed("FormattingOptions")
	}

	if this.InsertSpaces == nil {
		return nil, StructureValidateFailed("FormattingOptions")
	}

	type FormattingOptionsMarshal FormattingOptions
	tmpMarshal := FormattingOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provider options for a {@link DocumentFormattingRequest}.
type DocumentFormattingOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingOptions struct {

	// mixins

	WorkDoneProgressOptions
}

// Provider options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingOptions struct {

	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter *String `json:"firstTriggerCharacter"`

	// More trigger characters.
	MoreTriggerCharacter []String `json:"moreTriggerCharacter"`
}

func (this *DocumentOnTypeFormattingOptions) UnmarshalJSON(data []byte) error {
	type DocumentOnTypeFormattingOptionsUnmarshal DocumentOnTypeFormattingOptions
	var tmpUnmarshal DocumentOnTypeFormattingOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.FirstTriggerCharacter == nil {
		return StructureValidateFailed(
			"DocumentOnTypeFormattingOptions",
		)
	}

	*this = DocumentOnTypeFormattingOptions(tmpUnmarshal)
	return nil
}

func (this *DocumentOnTypeFormattingOptions) MarshalJSON() ([]byte, error) {

	if this.FirstTriggerCharacter == nil {
		return nil, StructureValidateFailed(
			"DocumentOnTypeFormattingOptions",
		)
	}

	type DocumentOnTypeFormattingOptionsMarshal DocumentOnTypeFormattingOptions
	tmpMarshal := DocumentOnTypeFormattingOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Provider options for a {@link RenameRequest}.
type RenameOptions struct {

	// mixins

	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.  @since
	// version 3.12.0
	PrepareProvider *Boolean `json:"prepareProvider"`
}

// The server capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandOptions struct {

	// mixins

	WorkDoneProgressOptions

	// The commands to be executed on the server
	Commands []String `json:"commands"`
}

func (this *ExecuteCommandOptions) UnmarshalJSON(data []byte) error {
	type ExecuteCommandOptionsUnmarshal ExecuteCommandOptions
	var tmpUnmarshal ExecuteCommandOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Commands == nil {
		return StructureValidateFailed("ExecuteCommandOptions")
	}

	*this = ExecuteCommandOptions(tmpUnmarshal)
	return nil
}

func (this *ExecuteCommandOptions) MarshalJSON() ([]byte, error) {

	if this.Commands == nil {
		return nil, StructureValidateFailed("ExecuteCommandOptions")
	}

	type ExecuteCommandOptionsMarshal ExecuteCommandOptions
	tmpMarshal := ExecuteCommandOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// @since 3.16.0
type SemanticTokensLegend struct {

	// The token types a server uses.
	TokenTypes []String `json:"tokenTypes"`

	// The token modifiers a server uses.
	TokenModifiers []String `json:"tokenModifiers"`
}

func (this *SemanticTokensLegend) UnmarshalJSON(data []byte) error {
	type SemanticTokensLegendUnmarshal SemanticTokensLegend
	var tmpUnmarshal SemanticTokensLegendUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.TokenTypes == nil {
		return StructureValidateFailed("SemanticTokensLegend")
	}

	if tmpUnmarshal.TokenModifiers == nil {
		return StructureValidateFailed("SemanticTokensLegend")
	}

	*this = SemanticTokensLegend(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensLegend) MarshalJSON() ([]byte, error) {

	if this.TokenTypes == nil {
		return nil, StructureValidateFailed("SemanticTokensLegend")
	}

	if this.TokenModifiers == nil {
		return nil, StructureValidateFailed("SemanticTokensLegend")
	}

	type SemanticTokensLegendMarshal SemanticTokensLegend
	tmpMarshal := SemanticTokensLegendMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Version *OptionalVersionedTextDocumentIdentifier_Version_Or `json:"version"`
}

func (this *OptionalVersionedTextDocumentIdentifier) UnmarshalJSON(
	data []byte,
) error {
	type OptionalVersionedTextDocumentIdentifierUnmarshal OptionalVersionedTextDocumentIdentifier
	var tmpUnmarshal OptionalVersionedTextDocumentIdentifierUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed(
			"OptionalVersionedTextDocumentIdentifier",
		)
	}

	*this = OptionalVersionedTextDocumentIdentifier(tmpUnmarshal)
	return nil
}

func (this *OptionalVersionedTextDocumentIdentifier) MarshalJSON() ([]byte, error) {

	if this.Version == nil {
		return nil, StructureValidateFailed(
			"OptionalVersionedTextDocumentIdentifier",
		)
	}

	type OptionalVersionedTextDocumentIdentifierMarshal OptionalVersionedTextDocumentIdentifier
	tmpMarshal := OptionalVersionedTextDocumentIdentifierMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A special text edit with an additional change annotation.  @since 3.16.0.
type AnnotatedTextEdit struct {

	// extends

	TextEdit

	// The actual identifier of the change annotation
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId"`
}

func (this *AnnotatedTextEdit) UnmarshalJSON(data []byte) error {
	type AnnotatedTextEditUnmarshal AnnotatedTextEdit
	var tmpUnmarshal AnnotatedTextEditUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.AnnotationId == nil {
		return StructureValidateFailed("AnnotatedTextEdit")
	}

	*this = AnnotatedTextEdit(tmpUnmarshal)
	return nil
}

func (this *AnnotatedTextEdit) MarshalJSON() ([]byte, error) {

	if this.AnnotationId == nil {
		return nil, StructureValidateFailed("AnnotatedTextEdit")
	}

	type AnnotatedTextEditMarshal AnnotatedTextEdit
	tmpMarshal := AnnotatedTextEditMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A generic resource operation.
type ResourceOperation struct {

	// The resource operation kind.
	Kind *String `json:"kind"`

	// An optional annotation identifier describing the operation.  @since
	// 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId"`
}

func (this *ResourceOperation) UnmarshalJSON(data []byte) error {
	type ResourceOperationUnmarshal ResourceOperation
	var tmpUnmarshal ResourceOperationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("ResourceOperation")
	}

	*this = ResourceOperation(tmpUnmarshal)
	return nil
}

func (this *ResourceOperation) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("ResourceOperation")
	}

	type ResourceOperationMarshal ResourceOperation
	tmpMarshal := ResourceOperationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Options to create a file.
type CreateFileOptions struct {

	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite *Boolean `json:"overwrite"`

	// Ignore if exists.
	IgnoreIfExists *Boolean `json:"ignoreIfExists"`
}

// Rename file options
type RenameFileOptions struct {

	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite *Boolean `json:"overwrite"`

	// Ignores if target exists.
	IgnoreIfExists *Boolean `json:"ignoreIfExists"`
}

// Delete file options
type DeleteFileOptions struct {

	// Delete the content recursively if a folder is denoted.
	Recursive *Boolean `json:"recursive"`

	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists *Boolean `json:"ignoreIfNotExists"`
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
	Glob *String `json:"glob"`

	// Whether to match files or folders with this pattern.  Matches both if
	// undefined.
	Matches *FileOperationPatternKind `json:"matches"`

	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options"`
}

func (this *FileOperationPattern) UnmarshalJSON(data []byte) error {
	type FileOperationPatternUnmarshal FileOperationPattern
	var tmpUnmarshal FileOperationPatternUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Glob == nil {
		return StructureValidateFailed("FileOperationPattern")
	}

	*this = FileOperationPattern(tmpUnmarshal)
	return nil
}

func (this *FileOperationPattern) MarshalJSON() ([]byte, error) {

	if this.Glob == nil {
		return nil, StructureValidateFailed("FileOperationPattern")
	}

	type FileOperationPatternMarshal FileOperationPattern
	tmpMarshal := FileOperationPatternMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A full document diagnostic report for a workspace diagnostic result.  @since
// 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {

	// extends

	FullDocumentDiagnosticReport

	// The URI for which diagnostic information is reported.
	Uri *DocumentUri `json:"uri"`

	// The version number for which the diagnostics are reported. If the
	// document is not marked as open `null` can be provided.
	Version *WorkspaceFullDocumentDiagnosticReport_Version_Or `json:"version"`
}

func (this *WorkspaceFullDocumentDiagnosticReport) UnmarshalJSON(
	data []byte,
) error {
	type WorkspaceFullDocumentDiagnosticReportUnmarshal WorkspaceFullDocumentDiagnosticReport
	var tmpUnmarshal WorkspaceFullDocumentDiagnosticReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed(
			"WorkspaceFullDocumentDiagnosticReport",
		)
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed(
			"WorkspaceFullDocumentDiagnosticReport",
		)
	}

	*this = WorkspaceFullDocumentDiagnosticReport(tmpUnmarshal)
	return nil
}

func (this *WorkspaceFullDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed(
			"WorkspaceFullDocumentDiagnosticReport",
		)
	}

	if this.Version == nil {
		return nil, StructureValidateFailed(
			"WorkspaceFullDocumentDiagnosticReport",
		)
	}

	type WorkspaceFullDocumentDiagnosticReportMarshal WorkspaceFullDocumentDiagnosticReport
	tmpMarshal := WorkspaceFullDocumentDiagnosticReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// An unchanged document diagnostic report for a workspace diagnostic result.
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {

	// extends

	UnchangedDocumentDiagnosticReport

	// The URI for which diagnostic information is reported.
	Uri *DocumentUri `json:"uri"`

	// The version number for which the diagnostics are reported. If the
	// document is not marked as open `null` can be provided.
	Version *WorkspaceUnchangedDocumentDiagnosticReport_Version_Or `json:"version"`
}

func (this *WorkspaceUnchangedDocumentDiagnosticReport) UnmarshalJSON(
	data []byte,
) error {
	type WorkspaceUnchangedDocumentDiagnosticReportUnmarshal WorkspaceUnchangedDocumentDiagnosticReport
	var tmpUnmarshal WorkspaceUnchangedDocumentDiagnosticReportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureValidateFailed(
			"WorkspaceUnchangedDocumentDiagnosticReport",
		)
	}

	if tmpUnmarshal.Version == nil {
		return StructureValidateFailed(
			"WorkspaceUnchangedDocumentDiagnosticReport",
		)
	}

	*this = WorkspaceUnchangedDocumentDiagnosticReport(tmpUnmarshal)
	return nil
}

func (this *WorkspaceUnchangedDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureValidateFailed(
			"WorkspaceUnchangedDocumentDiagnosticReport",
		)
	}

	if this.Version == nil {
		return nil, StructureValidateFailed(
			"WorkspaceUnchangedDocumentDiagnosticReport",
		)
	}

	type WorkspaceUnchangedDocumentDiagnosticReportMarshal WorkspaceUnchangedDocumentDiagnosticReport
	tmpMarshal := WorkspaceUnchangedDocumentDiagnosticReportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// LSP object definition. @since 3.17.0
type LSPObject struct {
}

// A notebook cell.  A cell's document URI must be unique across ALL notebook
// cells and can therefore be used to uniquely identify a notebook cell or the
// cell's text document.  @since 3.17.0
type NotebookCell struct {

	// The cell's kind
	Kind *NotebookCellKind `json:"kind"`

	// The URI of the cell's text document content.
	Document *DocumentUri `json:"document"`

	// Additional metadata stored with the cell.  Note: should always be an
	// object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata"`

	// Additional execution summary information if supported by the client.
	ExecutionSummary *ExecutionSummary `json:"executionSummary"`
}

func (this *NotebookCell) UnmarshalJSON(data []byte) error {
	type NotebookCellUnmarshal NotebookCell
	var tmpUnmarshal NotebookCellUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Kind == nil {
		return StructureValidateFailed("NotebookCell")
	}

	if tmpUnmarshal.Document == nil {
		return StructureValidateFailed("NotebookCell")
	}

	*this = NotebookCell(tmpUnmarshal)
	return nil
}

func (this *NotebookCell) MarshalJSON() ([]byte, error) {

	if this.Kind == nil {
		return nil, StructureValidateFailed("NotebookCell")
	}

	if this.Document == nil {
		return nil, StructureValidateFailed("NotebookCell")
	}

	type NotebookCellMarshal NotebookCell
	tmpMarshal := NotebookCellMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A change describing how to move a `NotebookCell` array from state S to S'.
// @since 3.17.0
type NotebookCellArrayChange struct {

	// The start oftest of the cell that changed.
	Start *Uinteger `json:"start"`

	// The deleted cells
	DeleteCount *Uinteger `json:"deleteCount"`

	// The new cells, if any
	Cells []NotebookCell `json:"cells"`
}

func (this *NotebookCellArrayChange) UnmarshalJSON(data []byte) error {
	type NotebookCellArrayChangeUnmarshal NotebookCellArrayChange
	var tmpUnmarshal NotebookCellArrayChangeUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Start == nil {
		return StructureValidateFailed("NotebookCellArrayChange")
	}

	if tmpUnmarshal.DeleteCount == nil {
		return StructureValidateFailed("NotebookCellArrayChange")
	}

	*this = NotebookCellArrayChange(tmpUnmarshal)
	return nil
}

func (this *NotebookCellArrayChange) MarshalJSON() ([]byte, error) {

	if this.Start == nil {
		return nil, StructureValidateFailed("NotebookCellArrayChange")
	}

	if this.DeleteCount == nil {
		return nil, StructureValidateFailed("NotebookCellArrayChange")
	}

	type NotebookCellArrayChangeMarshal NotebookCellArrayChange
	tmpMarshal := NotebookCellArrayChangeMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	OpenClose *Boolean `json:"openClose"`

	// Change notifications are sent to the server. See
	// TextDocumentSyncKind.None, TextDocumentSyncKind.Full and
	// TextDocumentSyncKind.Incremental. If omitted it defaults to
	// TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change"`

	// If present will save notifications are sent to the server. If omitted
	// the notification should not be sent.
	WillSave *Boolean `json:"willSave"`

	// If present will save wait until requests are sent to the server. If
	// omitted the request should not be sent.
	WillSaveWaitUntil *Boolean `json:"willSaveWaitUntil"`

	// If present save notifications are sent to the server. If omitted the
	// notification should not be sent.
	Save *TextDocumentSyncOptions_Save_Or `json:"save"`
}

// Options specific to a notebook plus its cells to be synced to the server.  If
// a selector provides a notebook document filter but no cell selector all cells
// of a matching notebook document will be synced.  If a selector provides no
// notebook document filter but only a cell selector all notebook document that
// contain at least one matching cell will be synced.  @since 3.17.0
type NotebookDocumentSyncOptions struct {

	// The notebooks to be synced
	NotebookSelector []NotebookDocumentSyncOptions_NotebookSelector_Element_Or `json:"notebookSelector"`

	// Whether save notification should be forwarded to the server. Will
	// only be honored if mode === `notebook`.
	Save *Boolean `json:"save"`
}

func (this *NotebookDocumentSyncOptions) UnmarshalJSON(data []byte) error {
	type NotebookDocumentSyncOptionsUnmarshal NotebookDocumentSyncOptions
	var tmpUnmarshal NotebookDocumentSyncOptionsUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookSelector == nil {
		return StructureValidateFailed("NotebookDocumentSyncOptions")
	}

	*this = NotebookDocumentSyncOptions(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentSyncOptions) MarshalJSON() ([]byte, error) {

	if this.NotebookSelector == nil {
		return nil, StructureValidateFailed(
			"NotebookDocumentSyncOptions",
		)
	}

	type NotebookDocumentSyncOptionsMarshal NotebookDocumentSyncOptions
	tmpMarshal := NotebookDocumentSyncOptionsMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Supported *Boolean `json:"supported"`

	// Whether the server wants to receive workspace folder change
	// notifications.  If a string is provided the string is treated as an
	// ID under which the notification is registered on the client side. The
	// ID can be used to unregister for these events using the
	// `client/unregisterCapability` request.
	ChangeNotifications *WorkspaceFoldersServerCapabilities_ChangeNotifications_Or `json:"changeNotifications"`
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
	Href *URI `json:"href"`
}

func (this *CodeDescription) UnmarshalJSON(data []byte) error {
	type CodeDescriptionUnmarshal CodeDescription
	var tmpUnmarshal CodeDescriptionUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Href == nil {
		return StructureValidateFailed("CodeDescription")
	}

	*this = CodeDescription(tmpUnmarshal)
	return nil
}

func (this *CodeDescription) MarshalJSON() ([]byte, error) {

	if this.Href == nil {
		return nil, StructureValidateFailed("CodeDescription")
	}

	type CodeDescriptionMarshal CodeDescription
	tmpMarshal := CodeDescriptionMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Represents a related message and source code location for a diagnostic. This
// should be used to point to code locations that cause or related to a
// diagnostics, e.g when duplicating a symbol in a scope.
type DiagnosticRelatedInformation struct {

	// The location of this related diagnostic information.
	Location *Location `json:"location"`

	// The message of this related diagnostic information.
	Message *String `json:"message"`
}

func (this *DiagnosticRelatedInformation) UnmarshalJSON(data []byte) error {
	type DiagnosticRelatedInformationUnmarshal DiagnosticRelatedInformation
	var tmpUnmarshal DiagnosticRelatedInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Location == nil {
		return StructureValidateFailed("DiagnosticRelatedInformation")
	}

	if tmpUnmarshal.Message == nil {
		return StructureValidateFailed("DiagnosticRelatedInformation")
	}

	*this = DiagnosticRelatedInformation(tmpUnmarshal)
	return nil
}

func (this *DiagnosticRelatedInformation) MarshalJSON() ([]byte, error) {

	if this.Location == nil {
		return nil, StructureValidateFailed(
			"DiagnosticRelatedInformation",
		)
	}

	if this.Message == nil {
		return nil, StructureValidateFailed(
			"DiagnosticRelatedInformation",
		)
	}

	type DiagnosticRelatedInformationMarshal DiagnosticRelatedInformation
	tmpMarshal := DiagnosticRelatedInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
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
	Label *ParameterInformation_Label_Or `json:"label"`

	// The human-readable doc-comment of this parameter. Will be shown in
	// the UI but can be omitted.
	Documentation *ParameterInformation_Documentation_Or `json:"documentation"`
}

func (this *ParameterInformation) UnmarshalJSON(data []byte) error {
	type ParameterInformationUnmarshal ParameterInformation
	var tmpUnmarshal ParameterInformationUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Label == nil {
		return StructureValidateFailed("ParameterInformation")
	}

	*this = ParameterInformation(tmpUnmarshal)
	return nil
}

func (this *ParameterInformation) MarshalJSON() ([]byte, error) {

	if this.Label == nil {
		return nil, StructureValidateFailed("ParameterInformation")
	}

	type ParameterInformationMarshal ParameterInformation
	tmpMarshal := ParameterInformationMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// A notebook cell text document filter denotes a cell text document by
// different properties.  @since 3.17.0
type NotebookCellTextDocumentFilter struct {

	// A filter that matches against the notebook containing the notebook
	// cell. If a string value is provided it matches against the notebook
	// type. '*' matches every notebook.
	Notebook *NotebookCellTextDocumentFilter_Notebook_Or `json:"notebook"`

	// A language id like `python`.  Will be matched against the language id
	// of the notebook cell document. '*' matches every language.
	Language *String `json:"language"`
}

func (this *NotebookCellTextDocumentFilter) UnmarshalJSON(data []byte) error {
	type NotebookCellTextDocumentFilterUnmarshal NotebookCellTextDocumentFilter
	var tmpUnmarshal NotebookCellTextDocumentFilterUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Notebook == nil {
		return StructureValidateFailed("NotebookCellTextDocumentFilter")
	}

	*this = NotebookCellTextDocumentFilter(tmpUnmarshal)
	return nil
}

func (this *NotebookCellTextDocumentFilter) MarshalJSON() ([]byte, error) {

	if this.Notebook == nil {
		return nil, StructureValidateFailed(
			"NotebookCellTextDocumentFilter",
		)
	}

	type NotebookCellTextDocumentFilterMarshal NotebookCellTextDocumentFilter
	tmpMarshal := NotebookCellTextDocumentFilterMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Matching options for the file operation pattern.  @since 3.16.0
type FileOperationPatternOptions struct {

	// The pattern should be matched ignoring casing.
	IgnoreCase *Boolean `json:"ignoreCase"`
}

type ExecutionSummary struct {

	// A strict monotonically increasing value indicating the execution
	// order of a cell inside a notebook.
	ExecutionOrder *Uinteger `json:"executionOrder"`

	// Whether the execution was successful or not if known by the client.
	Success *Boolean `json:"success"`
}

func (this *ExecutionSummary) UnmarshalJSON(data []byte) error {
	type ExecutionSummaryUnmarshal ExecutionSummary
	var tmpUnmarshal ExecutionSummaryUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ExecutionOrder == nil {
		return StructureValidateFailed("ExecutionSummary")
	}

	*this = ExecutionSummary(tmpUnmarshal)
	return nil
}

func (this *ExecutionSummary) MarshalJSON() ([]byte, error) {

	if this.ExecutionOrder == nil {
		return nil, StructureValidateFailed("ExecutionSummary")
	}

	type ExecutionSummaryMarshal ExecutionSummary
	tmpMarshal := ExecutionSummaryMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {

	// The client supports applying batch edits to the workspace by
	// supporting the request 'workspace/applyEdit'
	ApplyEdit *Boolean `json:"applyEdit"`

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
	WorkspaceFolders *Boolean `json:"workspaceFolders"`

	// The client supports `workspace/configuration` requests.  @since 3.6.0
	Configuration *Boolean `json:"configuration"`

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
	Synchronization *NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}

func (this *NotebookDocumentClientCapabilities) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentClientCapabilitiesUnmarshal NotebookDocumentClientCapabilities
	var tmpUnmarshal NotebookDocumentClientCapabilitiesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Synchronization == nil {
		return StructureValidateFailed(
			"NotebookDocumentClientCapabilities",
		)
	}

	*this = NotebookDocumentClientCapabilities(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentClientCapabilities) MarshalJSON() ([]byte, error) {

	if this.Synchronization == nil {
		return nil, StructureValidateFailed(
			"NotebookDocumentClientCapabilities",
		)
	}

	type NotebookDocumentClientCapabilitiesMarshal NotebookDocumentClientCapabilities
	tmpMarshal := NotebookDocumentClientCapabilitiesMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WindowClientCapabilities struct {

	// It indicates whether the client supports server initiated progress
	// using the `window/workDoneProgress/create` request.  The capability
	// also controls Whether client supports handling of progress
	// notifications. If set servers are allowed to report a
	// `workDoneProgress` property in the request specific server
	// capabilities.  @since 3.15.0
	WorkDoneProgress *Boolean `json:"workDoneProgress"`

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
	PositionEncodings []PositionEncodingKind `json:"positionEncodings"`
}

// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.  @since 3.17.0
type RelativePattern struct {

	// A workspace folder or a base URI to which this pattern will be
	// matched against relatively.
	BaseUri *RelativePattern_BaseUri_Or `json:"baseUri"`

	// The actual glob pattern;
	Pattern *Pattern `json:"pattern"`
}

func (this *RelativePattern) UnmarshalJSON(data []byte) error {
	type RelativePatternUnmarshal RelativePattern
	var tmpUnmarshal RelativePatternUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.BaseUri == nil {
		return StructureValidateFailed("RelativePattern")
	}

	if tmpUnmarshal.Pattern == nil {
		return StructureValidateFailed("RelativePattern")
	}

	*this = RelativePattern(tmpUnmarshal)
	return nil
}

func (this *RelativePattern) MarshalJSON() ([]byte, error) {

	if this.BaseUri == nil {
		return nil, StructureValidateFailed("RelativePattern")
	}

	if this.Pattern == nil {
		return nil, StructureValidateFailed("RelativePattern")
	}

	type RelativePatternMarshal RelativePattern
	tmpMarshal := RelativePatternMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkspaceEditClientCapabilities struct {

	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges *Boolean `json:"documentChanges"`

	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.  @since
	// 3.13.0
	ResourceOperations []ResourceOperationKind `json:"resourceOperations"`

	// The failure handling strategy of a client if applying the workspace
	// edit fails.  @since 3.13.0
	FailureHandling *FailureHandlingKind `json:"failureHandling"`

	// Whether the client normalizes line endings to the client specific
	// setting. If set to `true` the client will normalize line ending
	// characters in a workspace edit to the client-specified new line
	// character.  @since 3.16.0
	NormalizesLineEndings *Boolean `json:"normalizesLineEndings"`

	// Whether the client in general supports change annotations on text
	// edits, create file, rename file and delete file changes.  @since
	// 3.16.0
	ChangeAnnotationSupport *WorkspaceEditClientCapabilities_ChangeAnnotationSupport `json:"changeAnnotationSupport"`
}

type DidChangeConfigurationClientCapabilities struct {

	// Did change configuration notification supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

type DidChangeWatchedFilesClientCapabilities struct {

	// Did change watched files notification supports dynamic registration.
	// Please note that the current protocol doesn't support static
	// configuration for file changes from the server side.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Whether the client has support for {@link  RelativePattern relative
	// pattern} or not.  @since 3.17.0
	RelativePatternSupport *Boolean `json:"relativePatternSupport"`
}

// Client capabilities for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolClientCapabilities struct {

	// Symbol request supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

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

// The client capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandClientCapabilities struct {

	// Execute command supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all semantic tokens currently shown.
	// It should be used with absolute care and is useful for situation
	// where a server for example detects a project wide change that
	// requires such a calculation.
	RefreshSupport *Boolean `json:"refreshSupport"`
}

// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all code lenses currently shown. It
	// should be used with absolute care and is useful for situation where a
	// server for example detect a project wide change that requires such a
	// calculation.
	RefreshSupport *Boolean `json:"refreshSupport"`
}

// Capabilities relating to events from file operations by the user in the
// client.  These events do not come from the file system, they come from user
// operations like renaming a file in the UI.  @since 3.16.0
type FileOperationClientCapabilities struct {

	// Whether the client supports dynamic registration for file
	// requests/notifications.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client has support for sending didCreateFiles notifications.
	DidCreate *Boolean `json:"didCreate"`

	// The client has support for sending willCreateFiles requests.
	WillCreate *Boolean `json:"willCreate"`

	// The client has support for sending didRenameFiles notifications.
	DidRename *Boolean `json:"didRename"`

	// The client has support for sending willRenameFiles requests.
	WillRename *Boolean `json:"willRename"`

	// The client has support for sending didDeleteFiles notifications.
	DidDelete *Boolean `json:"didDelete"`

	// The client has support for sending willDeleteFiles requests.
	WillDelete *Boolean `json:"willDelete"`
}

// Client workspace capabilities specific to inline values.  @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all inline values currently shown.
	// It should be used with absolute care and is useful for situation
	// where a server for example detects a project wide change that
	// requires such a calculation.
	RefreshSupport *Boolean `json:"refreshSupport"`
}

// Client workspace capabilities specific to inlay hints.  @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent
	// from the server to the client.  Note that this event is global and
	// will force the client to refresh all inlay hints currently shown. It
	// should be used with absolute care and is useful for situation where a
	// server for example detects a project wide change that requires such a
	// calculation.
	RefreshSupport *Boolean `json:"refreshSupport"`
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
	RefreshSupport *Boolean `json:"refreshSupport"`
}

type TextDocumentSyncClientCapabilities struct {

	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports sending will save notifications.
	WillSave *Boolean `json:"willSave"`

	// The client supports sending a will save request and waits for a
	// response providing text edits which will be applied to the document
	// before it is saved.
	WillSaveWaitUntil *Boolean `json:"willSaveWaitUntil"`

	// The client supports did save notifications.
	DidSave *Boolean `json:"didSave"`
}

// Completion client capabilities
type CompletionClientCapabilities struct {

	// Whether completion supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

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
	ContextSupport *Boolean `json:"contextSupport"`

	// The client supports the following `CompletionList` specific
	// capabilities.  @since 3.17.0
	CompletionList *CompletionClientCapabilities_CompletionList `json:"completionList"`
}

type HoverClientCapabilities struct {

	// Whether hover supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Client supports the following content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat"`
}

// Client Capabilities for a {@link SignatureHelpRequest}.
type SignatureHelpClientCapabilities struct {

	// Whether signature help supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports the following `SignatureInformation` specific
	// properties.
	SignatureInformation *SignatureHelpClientCapabilities_SignatureInformation `json:"signatureInformation"`

	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.  @since 3.15.0
	ContextSupport *Boolean `json:"contextSupport"`
}

// @since 3.14.0
type DeclarationClientCapabilities struct {

	// Whether declaration supports dynamic registration. If this is set to
	// `true` the client supports the new `DeclarationRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of declaration
	// links.
	LinkSupport *Boolean `json:"linkSupport"`
}

// Client Capabilities for a {@link DefinitionRequest}.
type DefinitionClientCapabilities struct {

	// Whether definition supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  @since 3.14.0
	LinkSupport *Boolean `json:"linkSupport"`
}

// Since 3.6.0
type TypeDefinitionClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `TypeDefinitionRegistrationOptions` return value for the
	// corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  Since 3.14.0
	LinkSupport *Boolean `json:"linkSupport"`
}

// @since 3.6.0
type ImplementationClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `ImplementationRegistrationOptions` return value for the
	// corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports additional metadata in the form of definition
	// links.  @since 3.14.0
	LinkSupport *Boolean `json:"linkSupport"`
}

// Client Capabilities for a {@link ReferencesRequest}.
type ReferenceClientCapabilities struct {

	// Whether references supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client Capabilities for a {@link DocumentHighlightRequest}.
type DocumentHighlightClientCapabilities struct {

	// Whether document highlight supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client Capabilities for a {@link DocumentSymbolRequest}.
type DocumentSymbolClientCapabilities struct {

	// Whether document symbol supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind *DocumentSymbolClientCapabilities_SymbolKind `json:"symbolKind"`

	// The client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport *Boolean `json:"hierarchicalDocumentSymbolSupport"`

	// The client supports tags on `SymbolInformation`. Tags are supported
	// on `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to
	// true. Clients supporting tags have to handle unknown tags gracefully.
	//  @since 3.16.0
	TagSupport *DocumentSymbolClientCapabilities_TagSupport `json:"tagSupport"`

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.  @since 3.16.0
	LabelSupport *Boolean `json:"labelSupport"`
}

// The Client Capabilities of a {@link CodeActionRequest}.
type CodeActionClientCapabilities struct {

	// Whether code action supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client support code action literals of type `CodeAction` as a
	// valid response of the `textDocument/codeAction` request. If the
	// property is not set the request can only return `Command` literals.
	// @since 3.8.0
	CodeActionLiteralSupport *CodeActionClientCapabilities_CodeActionLiteralSupport `json:"codeActionLiteralSupport"`

	// Whether code action supports the `isPreferred` property.  @since
	// 3.15.0
	IsPreferredSupport *Boolean `json:"isPreferredSupport"`

	// Whether code action supports the `disabled` property.  @since 3.16.0
	DisabledSupport *Boolean `json:"disabledSupport"`

	// Whether code action supports the `data` property which is preserved
	// between a `textDocument/codeAction` and a `codeAction/resolve`
	// request.  @since 3.16.0
	DataSupport *Boolean `json:"dataSupport"`

	// Whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.  @since
	// 3.16.0
	ResolveSupport *CodeActionClientCapabilities_ResolveSupport `json:"resolveSupport"`

	// Whether the client honors the change annotations in text edits and
	// resource operations returned via the `CodeAction#edit` property by
	// for example presenting the workspace edit in the user interface and
	// asking for confirmation.  @since 3.16.0
	HonorsChangeAnnotations *Boolean `json:"honorsChangeAnnotations"`
}

// The client capabilities  of a {@link CodeLensRequest}.
type CodeLensClientCapabilities struct {

	// Whether code lens supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// The client capabilities of a {@link DocumentLinkRequest}.
type DocumentLinkClientCapabilities struct {

	// Whether document link supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Whether the client supports the `tooltip` property on `DocumentLink`.
	//  @since 3.15.0
	TooltipSupport *Boolean `json:"tooltipSupport"`
}

type DocumentColorClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `DocumentColorRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client capabilities of a {@link DocumentFormattingRequest}.
type DocumentFormattingClientCapabilities struct {

	// Whether formatting supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client capabilities of a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingClientCapabilities struct {

	// Whether range formatting supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client capabilities of a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingClientCapabilities struct {

	// Whether on type formatting supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

type RenameClientCapabilities struct {

	// Whether rename supports dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Client supports testing for validity of rename operations before
	// execution.  @since 3.12.0
	PrepareSupport *Boolean `json:"prepareSupport"`

	// Client supports the default behavior result.  The value indicates the
	// default behavior used by the client.  @since 3.16.0
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior"`

	// Whether the client honors the change annotations in text edits and
	// resource operations returned via the rename request's workspace edit
	// by for example presenting the workspace edit in the user interface
	// and asking for confirmation.  @since 3.16.0
	HonorsChangeAnnotations *Boolean `json:"honorsChangeAnnotations"`
}

type FoldingRangeClientCapabilities struct {

	// Whether implementation supports dynamic registration for folding
	// range providers. If this is set to `true` the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The maximum number of folding ranges that the client prefers to
	// receive per document. The value serves as a hint, servers are free to
	// follow the limit.
	RangeLimit *Uinteger `json:"rangeLimit"`

	// If set, the client signals that it only supports folding complete
	// lines. If set, client will ignore specified `startCharacter` and
	// `endCharacter` properties in a FoldingRange.
	LineFoldingOnly *Boolean `json:"lineFoldingOnly"`

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
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// The publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {

	// Whether the clients accepts diagnostics with related information.
	RelatedInformation *Boolean `json:"relatedInformation"`

	// Client supports the tag property to provide meta data about a
	// diagnostic. Clients supporting tags have to handle unknown tags
	// gracefully.  @since 3.15.0
	TagSupport *PublishDiagnosticsClientCapabilities_TagSupport `json:"tagSupport"`

	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.  @since
	// 3.15.0
	VersionSupport *Boolean `json:"versionSupport"`

	// Client supports a codeDescription property  @since 3.16.0
	CodeDescriptionSupport *Boolean `json:"codeDescriptionSupport"`

	// Whether code action supports the `data` property which is preserved
	// between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.  @since 3.16.0
	DataSupport *Boolean `json:"dataSupport"`
}

// @since 3.16.0
type CallHierarchyClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// @since 3.16.0
type SemanticTokensClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might
	// not show semantic tokens or degrade some of the user experience if a
	// range or full request is advertised by the client but not provided by
	// the server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or
	// might even decide to not show any semantic tokens at all.
	Requests *SemanticTokensClientCapabilities_Requests `json:"requests"`

	// The token types that the client supports.
	TokenTypes []String `json:"tokenTypes"`

	// The token modifiers that the client supports.
	TokenModifiers []String `json:"tokenModifiers"`

	// The token formats the clients supports.
	Formats []TokenFormat `json:"formats"`

	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport *Boolean `json:"overlappingTokenSupport"`

	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport *Boolean `json:"multilineTokenSupport"`

	// Whether the client allows the server to actively cancel a semantic
	// token request, e.g. supports returning LSPErrorCodes.ServerCancelled.
	// If a server does the client needs to retrigger the request.  @since
	// 3.17.0
	ServerCancelSupport *Boolean `json:"serverCancelSupport"`

	// Whether the client uses semantic tokens to augment existing syntax
	// tokens. If set to `true` client side created syntax tokens and
	// semantic tokens are both used for colorization. If set to `false` the
	// client only uses the returned semantic tokens for colorization.  If
	// the value is `undefined` then the client behavior is not specified.
	// @since 3.17.0
	AugmentsSyntaxTokens *Boolean `json:"augmentsSyntaxTokens"`
}

func (this *SemanticTokensClientCapabilities) UnmarshalJSON(data []byte) error {
	type SemanticTokensClientCapabilitiesUnmarshal SemanticTokensClientCapabilities
	var tmpUnmarshal SemanticTokensClientCapabilitiesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Requests == nil {
		return StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if tmpUnmarshal.TokenTypes == nil {
		return StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if tmpUnmarshal.TokenModifiers == nil {
		return StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if tmpUnmarshal.Formats == nil {
		return StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	*this = SemanticTokensClientCapabilities(tmpUnmarshal)
	return nil
}

func (this *SemanticTokensClientCapabilities) MarshalJSON() ([]byte, error) {

	if this.Requests == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if this.TokenTypes == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if this.TokenModifiers == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	if this.Formats == nil {
		return nil, StructureValidateFailed(
			"SemanticTokensClientCapabilities",
		)
	}

	type SemanticTokensClientCapabilitiesMarshal SemanticTokensClientCapabilities
	tmpMarshal := SemanticTokensClientCapabilitiesMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Client capabilities for the linked editing range request.  @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client capabilities specific to the moniker request.  @since 3.16.0
type MonikerClientCapabilities struct {

	// Whether moniker supports dynamic registration. If this is set to
	// `true` the client supports the new `MonikerRegistrationOptions`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// @since 3.17.0
type TypeHierarchyClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Client capabilities specific to inline values.  @since 3.17.0
type InlineValueClientCapabilities struct {

	// Whether implementation supports dynamic registration for inline value
	// providers.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`
}

// Inlay hint client capabilities.  @since 3.17.0
type InlayHintClientCapabilities struct {

	// Whether inlay hints support dynamic registration.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

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
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// Whether the clients supports related documents for document
	// diagnostic pulls.
	RelatedDocumentSupport *Boolean `json:"relatedDocumentSupport"`
}

// Notebook specific client capabilities.  @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {

	// Whether implementation supports dynamic registration. If this is set
	// to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration *Boolean `json:"dynamicRegistration"`

	// The client supports sending execution summary data per cell.
	ExecutionSummarySupport *Boolean `json:"executionSummarySupport"`
}

// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {

	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ShowMessageRequestClientCapabilities_MessageActionItem `json:"messageActionItem"`
}

// Client capabilities for the showDocument request.  @since 3.16.0
type ShowDocumentClientCapabilities struct {

	// The client has support for the showDocument request.
	Support *Boolean `json:"support"`
}

func (this *ShowDocumentClientCapabilities) UnmarshalJSON(data []byte) error {
	type ShowDocumentClientCapabilitiesUnmarshal ShowDocumentClientCapabilities
	var tmpUnmarshal ShowDocumentClientCapabilitiesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Support == nil {
		return StructureValidateFailed("ShowDocumentClientCapabilities")
	}

	*this = ShowDocumentClientCapabilities(tmpUnmarshal)
	return nil
}

func (this *ShowDocumentClientCapabilities) MarshalJSON() ([]byte, error) {

	if this.Support == nil {
		return nil, StructureValidateFailed(
			"ShowDocumentClientCapabilities",
		)
	}

	type ShowDocumentClientCapabilitiesMarshal ShowDocumentClientCapabilities
	tmpMarshal := ShowDocumentClientCapabilitiesMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Client capabilities specific to regular expressions.  @since 3.16.0
type RegularExpressionsClientCapabilities struct {

	// The engine's name.
	Engine *String `json:"engine"`

	// The engine's version.
	Version *String `json:"version"`
}

func (this *RegularExpressionsClientCapabilities) UnmarshalJSON(
	data []byte,
) error {
	type RegularExpressionsClientCapabilitiesUnmarshal RegularExpressionsClientCapabilities
	var tmpUnmarshal RegularExpressionsClientCapabilitiesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Engine == nil {
		return StructureValidateFailed(
			"RegularExpressionsClientCapabilities",
		)
	}

	*this = RegularExpressionsClientCapabilities(tmpUnmarshal)
	return nil
}

func (this *RegularExpressionsClientCapabilities) MarshalJSON() ([]byte, error) {

	if this.Engine == nil {
		return nil, StructureValidateFailed(
			"RegularExpressionsClientCapabilities",
		)
	}

	type RegularExpressionsClientCapabilitiesMarshal RegularExpressionsClientCapabilities
	tmpMarshal := RegularExpressionsClientCapabilitiesMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

// Client capabilities specific to the used markdown parser.  @since 3.16.0
type MarkdownClientCapabilities struct {

	// The name of the parser.
	Parser *String `json:"parser"`

	// The version of the parser.
	Version *String `json:"version"`

	// A list of HTML tags that the client allows / supports in Markdown.
	// @since 3.17.0
	AllowedTags []String `json:"allowedTags"`
}

func (this *MarkdownClientCapabilities) UnmarshalJSON(data []byte) error {
	type MarkdownClientCapabilitiesUnmarshal MarkdownClientCapabilities
	var tmpUnmarshal MarkdownClientCapabilitiesUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Parser == nil {
		return StructureValidateFailed("MarkdownClientCapabilities")
	}

	*this = MarkdownClientCapabilities(tmpUnmarshal)
	return nil
}

func (this *MarkdownClientCapabilities) MarshalJSON() ([]byte, error) {

	if this.Parser == nil {
		return nil, StructureValidateFailed(
			"MarkdownClientCapabilities",
		)
	}

	type MarkdownClientCapabilitiesMarshal MarkdownClientCapabilities
	tmpMarshal := MarkdownClientCapabilitiesMarshal(*this)
	return json.Marshal(&tmpMarshal)
}
