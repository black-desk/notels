// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"encoding/json"
	"fmt"
)

var StructureLiteralValidateFailed = func(name string) error {
	return fmt.Errorf("literal structure \"%s\"validate failed", name)
}

type CodeActionClientCapabilities_CodeActionLiteralSupport struct {

	// The code action kind is support with the following value set.
	CodeActionKind *CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind `json:"codeActionKind"`
}

func (this *CodeActionClientCapabilities_CodeActionLiteralSupport) UnmarshalJSON(
	data []byte,
) error {
	type CodeActionClientCapabilities_CodeActionLiteralSupportUnmarshal CodeActionClientCapabilities_CodeActionLiteralSupport
	var tmpUnmarshal CodeActionClientCapabilities_CodeActionLiteralSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.CodeActionKind == nil {
		return StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_CodeActionLiteralSupport",
		)
	}

	*this = CodeActionClientCapabilities_CodeActionLiteralSupport(
		tmpUnmarshal,
	)
	return nil
}

func (this *CodeActionClientCapabilities_CodeActionLiteralSupport) MarshalJSON() ([]byte, error) {

	if this.CodeActionKind == nil {
		return nil, StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_CodeActionLiteralSupport",
		)
	}

	type CodeActionClientCapabilities_CodeActionLiteralSupportMarshal CodeActionClientCapabilities_CodeActionLiteralSupport
	tmpMarshal := CodeActionClientCapabilities_CodeActionLiteralSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind struct {

	// The code action kind values the client supports. When this property
	// exists the client also guarantees that it will handle values outside
	// its set gracefully and falls back to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet"`
}

func (this *CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind) UnmarshalJSON(
	data []byte,
) error {
	type CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKindUnmarshal CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind
	var tmpUnmarshal CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKindUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind",
		)
	}

	*this = CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind(
		tmpUnmarshal,
	)
	return nil
}

func (this *CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind",
		)
	}

	type CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKindMarshal CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKind
	tmpMarshal := CodeActionClientCapabilities_CodeActionLiteralSupport_CodeActionKindMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type CodeActionClientCapabilities_ResolveSupport struct {

	// The properties that a client can resolve lazily.
	Properties []String `json:"properties"`
}

func (this *CodeActionClientCapabilities_ResolveSupport) UnmarshalJSON(
	data []byte,
) error {
	type CodeActionClientCapabilities_ResolveSupportUnmarshal CodeActionClientCapabilities_ResolveSupport
	var tmpUnmarshal CodeActionClientCapabilities_ResolveSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Properties == nil {
		return StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_ResolveSupport",
		)
	}

	*this = CodeActionClientCapabilities_ResolveSupport(tmpUnmarshal)
	return nil
}

func (this *CodeActionClientCapabilities_ResolveSupport) MarshalJSON() ([]byte, error) {

	if this.Properties == nil {
		return nil, StructureLiteralValidateFailed(
			"CodeActionClientCapabilities_ResolveSupport",
		)
	}

	type CodeActionClientCapabilities_ResolveSupportMarshal CodeActionClientCapabilities_ResolveSupport
	tmpMarshal := CodeActionClientCapabilities_ResolveSupportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type CodeAction_Disabled struct {

	// Human readable description of why the code action is currently
	// disabled.  This is displayed in the code actions UI.
	Reason *String `json:"reason"`
}

func (this *CodeAction_Disabled) UnmarshalJSON(data []byte) error {
	type CodeAction_DisabledUnmarshal CodeAction_Disabled
	var tmpUnmarshal CodeAction_DisabledUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Reason == nil {
		return StructureLiteralValidateFailed("CodeAction_Disabled")
	}

	*this = CodeAction_Disabled(tmpUnmarshal)
	return nil
}

func (this *CodeAction_Disabled) MarshalJSON() ([]byte, error) {

	if this.Reason == nil {
		return nil, StructureLiteralValidateFailed(
			"CodeAction_Disabled",
		)
	}

	type CodeAction_DisabledMarshal CodeAction_Disabled
	tmpMarshal := CodeAction_DisabledMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type CompletionClientCapabilities_CompletionItem struct {

	// Client supports snippets as insert text.  A snippet can define tab
	// stops and placeholders with `$1`, `$2` and `${3:foo}`. `$0` defines
	// the final tab stop, it defaults to the end of the snippet.
	// Placeholders with equal identifiers are linked, that is typing in one
	// will update others too.
	SnippetSupport *Boolean `json:"snippetSupport"`

	// Client supports commit characters on a completion item.
	CommitCharactersSupport *Boolean `json:"commitCharactersSupport"`

	// Client supports the following content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat"`

	// Client supports the deprecated property on a completion item.
	DeprecatedSupport *Boolean `json:"deprecatedSupport"`

	// Client supports the preselect property on a completion item.
	PreselectSupport *Boolean `json:"preselectSupport"`

	// Client supports the tag property on a completion item. Clients
	// supporting tags have to handle unknown tags gracefully. Clients
	// especially need to preserve unknown tags when sending a completion
	// item back to the server in a resolve call.  @since 3.15.0
	TagSupport *CompletionClientCapabilities_CompletionItem_TagSupport `json:"tagSupport"`

	// Client support insert replace edit to control different behavior if a
	// completion item is inserted in the text or should replace text.
	// @since 3.16.0
	InsertReplaceSupport *Boolean `json:"insertReplaceSupport"`

	// Indicates which properties a client can resolve lazily on a
	// completion item. Before version 3.16.0 only the predefined properties
	// `documentation` and `details` could be resolved lazily.  @since
	// 3.16.0
	ResolveSupport *CompletionClientCapabilities_CompletionItem_ResolveSupport `json:"resolveSupport"`

	// The client supports the `insertTextMode` property on a completion
	// item to override the whitespace handling mode as defined by the
	// client (see `insertTextMode`).  @since 3.16.0
	InsertTextModeSupport *CompletionClientCapabilities_CompletionItem_InsertTextModeSupport `json:"insertTextModeSupport"`

	// The client has support for completion item label details (see also
	// `CompletionItemLabelDetails`).  @since 3.17.0
	LabelDetailsSupport *Boolean `json:"labelDetailsSupport"`
}

type CompletionClientCapabilities_CompletionItemKind struct {

	// The completion item kind values the client supports. When this
	// property exists the client also guarantees that it will handle values
	// outside its set gracefully and falls back to a default value when
	// unknown.  If this property is not present the client only supports
	// the completion items kinds from `Text` to `Reference` as defined in
	// the initial version of the protocol.
	ValueSet []CompletionItemKind `json:"valueSet"`
}

type CompletionClientCapabilities_CompletionItem_InsertTextModeSupport struct {
	ValueSet []InsertTextMode `json:"valueSet"`
}

func (this *CompletionClientCapabilities_CompletionItem_InsertTextModeSupport) UnmarshalJSON(
	data []byte,
) error {
	type CompletionClientCapabilities_CompletionItem_InsertTextModeSupportUnmarshal CompletionClientCapabilities_CompletionItem_InsertTextModeSupport
	var tmpUnmarshal CompletionClientCapabilities_CompletionItem_InsertTextModeSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_InsertTextModeSupport",
		)
	}

	*this = CompletionClientCapabilities_CompletionItem_InsertTextModeSupport(
		tmpUnmarshal,
	)
	return nil
}

func (this *CompletionClientCapabilities_CompletionItem_InsertTextModeSupport) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_InsertTextModeSupport",
		)
	}

	type CompletionClientCapabilities_CompletionItem_InsertTextModeSupportMarshal CompletionClientCapabilities_CompletionItem_InsertTextModeSupport
	tmpMarshal := CompletionClientCapabilities_CompletionItem_InsertTextModeSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type CompletionClientCapabilities_CompletionItem_ResolveSupport struct {

	// The properties that a client can resolve lazily.
	Properties []String `json:"properties"`
}

func (this *CompletionClientCapabilities_CompletionItem_ResolveSupport) UnmarshalJSON(
	data []byte,
) error {
	type CompletionClientCapabilities_CompletionItem_ResolveSupportUnmarshal CompletionClientCapabilities_CompletionItem_ResolveSupport
	var tmpUnmarshal CompletionClientCapabilities_CompletionItem_ResolveSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Properties == nil {
		return StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_ResolveSupport",
		)
	}

	*this = CompletionClientCapabilities_CompletionItem_ResolveSupport(
		tmpUnmarshal,
	)
	return nil
}

func (this *CompletionClientCapabilities_CompletionItem_ResolveSupport) MarshalJSON() ([]byte, error) {

	if this.Properties == nil {
		return nil, StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_ResolveSupport",
		)
	}

	type CompletionClientCapabilities_CompletionItem_ResolveSupportMarshal CompletionClientCapabilities_CompletionItem_ResolveSupport
	tmpMarshal := CompletionClientCapabilities_CompletionItem_ResolveSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type CompletionClientCapabilities_CompletionItem_TagSupport struct {

	// The tags supported by the client.
	ValueSet []CompletionItemTag `json:"valueSet"`
}

func (this *CompletionClientCapabilities_CompletionItem_TagSupport) UnmarshalJSON(
	data []byte,
) error {
	type CompletionClientCapabilities_CompletionItem_TagSupportUnmarshal CompletionClientCapabilities_CompletionItem_TagSupport
	var tmpUnmarshal CompletionClientCapabilities_CompletionItem_TagSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_TagSupport",
		)
	}

	*this = CompletionClientCapabilities_CompletionItem_TagSupport(
		tmpUnmarshal,
	)
	return nil
}

func (this *CompletionClientCapabilities_CompletionItem_TagSupport) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"CompletionClientCapabilities_CompletionItem_TagSupport",
		)
	}

	type CompletionClientCapabilities_CompletionItem_TagSupportMarshal CompletionClientCapabilities_CompletionItem_TagSupport
	tmpMarshal := CompletionClientCapabilities_CompletionItem_TagSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type CompletionClientCapabilities_CompletionList struct {

	// The client supports the following itemDefaults on a completion list.
	// The value lists the supported property names of the
	// `CompletionList.itemDefaults` object. If omitted no properties are
	// supported.  @since 3.17.0
	ItemDefaults []String `json:"itemDefaults"`
}

type CompletionList_ItemDefaults struct {

	// A default commit character set.  @since 3.17.0
	CommitCharacters []String `json:"commitCharacters"`

	// A default edit range.  @since 3.17.0
	EditRange *CompletionList_ItemDefaults_EditRange_Or `json:"editRange"`

	// A default insert text format.  @since 3.17.0
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat"`

	// A default insert text mode.  @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode"`

	// A default data value.  @since 3.17.0
	Data *LSPAny `json:"data"`
}

type CompletionList_ItemDefaults_EditRange_Or_1 struct {
	Insert *Range `json:"insert"`

	Replace *Range `json:"replace"`
}

func (this *CompletionList_ItemDefaults_EditRange_Or_1) UnmarshalJSON(
	data []byte,
) error {
	type CompletionList_ItemDefaults_EditRange_Or_1Unmarshal CompletionList_ItemDefaults_EditRange_Or_1
	var tmpUnmarshal CompletionList_ItemDefaults_EditRange_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Insert == nil {
		return StructureLiteralValidateFailed(
			"CompletionList_ItemDefaults_EditRange_Or_1",
		)
	}

	if tmpUnmarshal.Replace == nil {
		return StructureLiteralValidateFailed(
			"CompletionList_ItemDefaults_EditRange_Or_1",
		)
	}

	*this = CompletionList_ItemDefaults_EditRange_Or_1(tmpUnmarshal)
	return nil
}

func (this *CompletionList_ItemDefaults_EditRange_Or_1) MarshalJSON() ([]byte, error) {

	if this.Insert == nil {
		return nil, StructureLiteralValidateFailed(
			"CompletionList_ItemDefaults_EditRange_Or_1",
		)
	}

	if this.Replace == nil {
		return nil, StructureLiteralValidateFailed(
			"CompletionList_ItemDefaults_EditRange_Or_1",
		)
	}

	type CompletionList_ItemDefaults_EditRange_Or_1Marshal CompletionList_ItemDefaults_EditRange_Or_1
	tmpMarshal := CompletionList_ItemDefaults_EditRange_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type CompletionOptions_CompletionItem struct {

	// The server has support for completion item label details (see also
	// `CompletionItemLabelDetails`) when receiving a completion item in a
	// resolve call.  @since 3.17.0
	LabelDetailsSupport *Boolean `json:"labelDetailsSupport"`
}

type DocumentSymbolClientCapabilities_SymbolKind struct {

	// The symbol kind values the client supports. When this property exists
	// the client also guarantees that it will handle values outside its set
	// gracefully and falls back to a default value when unknown.  If this
	// property is not present the client only supports the symbol kinds
	// from `File` to `Array` as defined in the initial version of the
	// protocol.
	ValueSet []SymbolKind `json:"valueSet"`
}

type DocumentSymbolClientCapabilities_TagSupport struct {

	// The tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}

func (this *DocumentSymbolClientCapabilities_TagSupport) UnmarshalJSON(
	data []byte,
) error {
	type DocumentSymbolClientCapabilities_TagSupportUnmarshal DocumentSymbolClientCapabilities_TagSupport
	var tmpUnmarshal DocumentSymbolClientCapabilities_TagSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"DocumentSymbolClientCapabilities_TagSupport",
		)
	}

	*this = DocumentSymbolClientCapabilities_TagSupport(tmpUnmarshal)
	return nil
}

func (this *DocumentSymbolClientCapabilities_TagSupport) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"DocumentSymbolClientCapabilities_TagSupport",
		)
	}

	type DocumentSymbolClientCapabilities_TagSupportMarshal DocumentSymbolClientCapabilities_TagSupport
	tmpMarshal := DocumentSymbolClientCapabilities_TagSupportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type FoldingRangeClientCapabilities_FoldingRange struct {

	// If set, the client signals that it supports setting collapsedText on
	// folding ranges to display custom labels instead of the default text.
	// @since 3.17.0
	CollapsedText *Boolean `json:"collapsedText"`
}

type FoldingRangeClientCapabilities_FoldingRangeKind struct {

	// The folding range kind values the client supports. When this property
	// exists the client also guarantees that it will handle values outside
	// its set gracefully and falls back to a default value when unknown.
	ValueSet []FoldingRangeKind `json:"valueSet"`
}

type GeneralClientCapabilities_StaleRequestSupport struct {

	// The client will actively cancel the request.
	Cancel *Boolean `json:"cancel"`

	// The list of requests for which the client will retry the request if
	// it receives a response with error code `ContentModified`
	RetryOnContentModified []String `json:"retryOnContentModified"`
}

func (this *GeneralClientCapabilities_StaleRequestSupport) UnmarshalJSON(
	data []byte,
) error {
	type GeneralClientCapabilities_StaleRequestSupportUnmarshal GeneralClientCapabilities_StaleRequestSupport
	var tmpUnmarshal GeneralClientCapabilities_StaleRequestSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Cancel == nil {
		return StructureLiteralValidateFailed(
			"GeneralClientCapabilities_StaleRequestSupport",
		)
	}

	if tmpUnmarshal.RetryOnContentModified == nil {
		return StructureLiteralValidateFailed(
			"GeneralClientCapabilities_StaleRequestSupport",
		)
	}

	*this = GeneralClientCapabilities_StaleRequestSupport(tmpUnmarshal)
	return nil
}

func (this *GeneralClientCapabilities_StaleRequestSupport) MarshalJSON() ([]byte, error) {

	if this.Cancel == nil {
		return nil, StructureLiteralValidateFailed(
			"GeneralClientCapabilities_StaleRequestSupport",
		)
	}

	if this.RetryOnContentModified == nil {
		return nil, StructureLiteralValidateFailed(
			"GeneralClientCapabilities_StaleRequestSupport",
		)
	}

	type GeneralClientCapabilities_StaleRequestSupportMarshal GeneralClientCapabilities_StaleRequestSupport
	tmpMarshal := GeneralClientCapabilities_StaleRequestSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type InitializeResult_ServerInfo struct {

	// The name of the server as defined by the server.
	Name *String `json:"name"`

	// The server's version as defined by the server.
	Version *String `json:"version"`
}

func (this *InitializeResult_ServerInfo) UnmarshalJSON(data []byte) error {
	type InitializeResult_ServerInfoUnmarshal InitializeResult_ServerInfo
	var tmpUnmarshal InitializeResult_ServerInfoUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureLiteralValidateFailed(
			"InitializeResult_ServerInfo",
		)
	}

	*this = InitializeResult_ServerInfo(tmpUnmarshal)
	return nil
}

func (this *InitializeResult_ServerInfo) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureLiteralValidateFailed(
			"InitializeResult_ServerInfo",
		)
	}

	type InitializeResult_ServerInfoMarshal InitializeResult_ServerInfo
	tmpMarshal := InitializeResult_ServerInfoMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type InlayHintClientCapabilities_ResolveSupport struct {

	// The properties that a client can resolve lazily.
	Properties []String `json:"properties"`
}

func (this *InlayHintClientCapabilities_ResolveSupport) UnmarshalJSON(
	data []byte,
) error {
	type InlayHintClientCapabilities_ResolveSupportUnmarshal InlayHintClientCapabilities_ResolveSupport
	var tmpUnmarshal InlayHintClientCapabilities_ResolveSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Properties == nil {
		return StructureLiteralValidateFailed(
			"InlayHintClientCapabilities_ResolveSupport",
		)
	}

	*this = InlayHintClientCapabilities_ResolveSupport(tmpUnmarshal)
	return nil
}

func (this *InlayHintClientCapabilities_ResolveSupport) MarshalJSON() ([]byte, error) {

	if this.Properties == nil {
		return nil, StructureLiteralValidateFailed(
			"InlayHintClientCapabilities_ResolveSupport",
		)
	}

	type InlayHintClientCapabilities_ResolveSupportMarshal InlayHintClientCapabilities_ResolveSupport
	tmpMarshal := InlayHintClientCapabilities_ResolveSupportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type MarkedString_Or_1 struct {
	Language *String `json:"language"`

	Value *String `json:"value"`
}

func (this *MarkedString_Or_1) UnmarshalJSON(data []byte) error {
	type MarkedString_Or_1Unmarshal MarkedString_Or_1
	var tmpUnmarshal MarkedString_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Language == nil {
		return StructureLiteralValidateFailed("MarkedString_Or_1")
	}

	if tmpUnmarshal.Value == nil {
		return StructureLiteralValidateFailed("MarkedString_Or_1")
	}

	*this = MarkedString_Or_1(tmpUnmarshal)
	return nil
}

func (this *MarkedString_Or_1) MarshalJSON() ([]byte, error) {

	if this.Language == nil {
		return nil, StructureLiteralValidateFailed("MarkedString_Or_1")
	}

	if this.Value == nil {
		return nil, StructureLiteralValidateFailed("MarkedString_Or_1")
	}

	type MarkedString_Or_1Marshal MarkedString_Or_1
	tmpMarshal := MarkedString_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentChangeEvent_Cells struct {

	// Changes to the cell structure to add or remove cells.
	Structure *NotebookDocumentChangeEvent_Cells_Structure `json:"structure"`

	// Changes to notebook cells properties like its kind, execution summary
	// or metadata.
	Data []NotebookCell `json:"data"`

	// Changes to the text content of notebook cells.
	TextContent []NotebookDocumentChangeEvent_Cells_TextContent_Element `json:"textContent"`
}

type NotebookDocumentChangeEvent_Cells_Structure struct {

	// The change to the cell array.
	Array *NotebookCellArrayChange `json:"array"`

	// Additional opened cell text documents.
	DidOpen []TextDocumentItem `json:"didOpen"`

	// Additional closed cell text documents.
	DidClose []TextDocumentIdentifier `json:"didClose"`
}

func (this *NotebookDocumentChangeEvent_Cells_Structure) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentChangeEvent_Cells_StructureUnmarshal NotebookDocumentChangeEvent_Cells_Structure
	var tmpUnmarshal NotebookDocumentChangeEvent_Cells_StructureUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Array == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_Structure",
		)
	}

	*this = NotebookDocumentChangeEvent_Cells_Structure(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentChangeEvent_Cells_Structure) MarshalJSON() ([]byte, error) {

	if this.Array == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_Structure",
		)
	}

	type NotebookDocumentChangeEvent_Cells_StructureMarshal NotebookDocumentChangeEvent_Cells_Structure
	tmpMarshal := NotebookDocumentChangeEvent_Cells_StructureMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentChangeEvent_Cells_TextContent_Element struct {
	Document *VersionedTextDocumentIdentifier `json:"document"`

	Changes []TextDocumentContentChangeEvent `json:"changes"`
}

func (this *NotebookDocumentChangeEvent_Cells_TextContent_Element) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentChangeEvent_Cells_TextContent_ElementUnmarshal NotebookDocumentChangeEvent_Cells_TextContent_Element
	var tmpUnmarshal NotebookDocumentChangeEvent_Cells_TextContent_ElementUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Document == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_TextContent_Element",
		)
	}

	if tmpUnmarshal.Changes == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_TextContent_Element",
		)
	}

	*this = NotebookDocumentChangeEvent_Cells_TextContent_Element(
		tmpUnmarshal,
	)
	return nil
}

func (this *NotebookDocumentChangeEvent_Cells_TextContent_Element) MarshalJSON() ([]byte, error) {

	if this.Document == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_TextContent_Element",
		)
	}

	if this.Changes == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentChangeEvent_Cells_TextContent_Element",
		)
	}

	type NotebookDocumentChangeEvent_Cells_TextContent_ElementMarshal NotebookDocumentChangeEvent_Cells_TextContent_Element
	tmpMarshal := NotebookDocumentChangeEvent_Cells_TextContent_ElementMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentFilter_Or_0 struct {

	// The type of the enclosing notebook.
	NotebookType *String `json:"notebookType"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern.
	Pattern *String `json:"pattern"`
}

func (this *NotebookDocumentFilter_Or_0) UnmarshalJSON(data []byte) error {
	type NotebookDocumentFilter_Or_0Unmarshal NotebookDocumentFilter_Or_0
	var tmpUnmarshal NotebookDocumentFilter_Or_0Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.NotebookType == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_0",
		)
	}

	*this = NotebookDocumentFilter_Or_0(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentFilter_Or_0) MarshalJSON() ([]byte, error) {

	if this.NotebookType == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_0",
		)
	}

	type NotebookDocumentFilter_Or_0Marshal NotebookDocumentFilter_Or_0
	tmpMarshal := NotebookDocumentFilter_Or_0Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentFilter_Or_1 struct {

	// The type of the enclosing notebook.
	NotebookType *String `json:"notebookType"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern.
	Pattern *String `json:"pattern"`
}

func (this *NotebookDocumentFilter_Or_1) UnmarshalJSON(data []byte) error {
	type NotebookDocumentFilter_Or_1Unmarshal NotebookDocumentFilter_Or_1
	var tmpUnmarshal NotebookDocumentFilter_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Scheme == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_1",
		)
	}

	*this = NotebookDocumentFilter_Or_1(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentFilter_Or_1) MarshalJSON() ([]byte, error) {

	if this.Scheme == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_1",
		)
	}

	type NotebookDocumentFilter_Or_1Marshal NotebookDocumentFilter_Or_1
	tmpMarshal := NotebookDocumentFilter_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentFilter_Or_2 struct {

	// The type of the enclosing notebook.
	NotebookType *String `json:"notebookType"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern.
	Pattern *String `json:"pattern"`
}

func (this *NotebookDocumentFilter_Or_2) UnmarshalJSON(data []byte) error {
	type NotebookDocumentFilter_Or_2Unmarshal NotebookDocumentFilter_Or_2
	var tmpUnmarshal NotebookDocumentFilter_Or_2Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Pattern == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_2",
		)
	}

	*this = NotebookDocumentFilter_Or_2(tmpUnmarshal)
	return nil
}

func (this *NotebookDocumentFilter_Or_2) MarshalJSON() ([]byte, error) {

	if this.Pattern == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentFilter_Or_2",
		)
	}

	type NotebookDocumentFilter_Or_2Marshal NotebookDocumentFilter_Or_2
	tmpMarshal := NotebookDocumentFilter_Or_2Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0 struct {

	// The notebook to be synced If a string value is provided it matches
	// against the notebook type. '*' matches every notebook.
	Notebook *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or `json:"notebook"`

	// The cells of the matching notebook to be synced.
	Cells []NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element `json:"cells"`
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0Unmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0
	var tmpUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Notebook == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0",
		)
	}

	*this = NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0(
		tmpUnmarshal,
	)
	return nil
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0) MarshalJSON() ([]byte, error) {

	if this.Notebook == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0",
		)
	}

	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0Marshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0
	tmpMarshal := NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0Marshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element struct {
	Language *String `json:"language"`
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_ElementUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element
	var tmpUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_ElementUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Language == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element",
		)
	}

	*this = NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element(
		tmpUnmarshal,
	)
	return nil
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element) MarshalJSON() ([]byte, error) {

	if this.Language == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element",
		)
	}

	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_ElementMarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_Element
	tmpMarshal := NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Cells_ElementMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1 struct {

	// The notebook to be synced If a string value is provided it matches
	// against the notebook type. '*' matches every notebook.
	Notebook *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or `json:"notebook"`

	// The cells of the matching notebook to be synced.
	Cells []NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element `json:"cells"`
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1Unmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1
	var tmpUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Cells == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1",
		)
	}

	*this = NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1(
		tmpUnmarshal,
	)
	return nil
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1) MarshalJSON() ([]byte, error) {

	if this.Cells == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1",
		)
	}

	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1Marshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1
	tmpMarshal := NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1Marshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element struct {
	Language *String `json:"language"`
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element) UnmarshalJSON(
	data []byte,
) error {
	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_ElementUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element
	var tmpUnmarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_ElementUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Language == nil {
		return StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element",
		)
	}

	*this = NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element(
		tmpUnmarshal,
	)
	return nil
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element) MarshalJSON() ([]byte, error) {

	if this.Language == nil {
		return nil, StructureLiteralValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element",
		)
	}

	type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_ElementMarshal NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_Element
	tmpMarshal := NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Cells_ElementMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type PrepareRenameResult_Or_1 struct {
	Range *Range `json:"range"`

	Placeholder *String `json:"placeholder"`
}

func (this *PrepareRenameResult_Or_1) UnmarshalJSON(data []byte) error {
	type PrepareRenameResult_Or_1Unmarshal PrepareRenameResult_Or_1
	var tmpUnmarshal PrepareRenameResult_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_1",
		)
	}

	if tmpUnmarshal.Placeholder == nil {
		return StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_1",
		)
	}

	*this = PrepareRenameResult_Or_1(tmpUnmarshal)
	return nil
}

func (this *PrepareRenameResult_Or_1) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_1",
		)
	}

	if this.Placeholder == nil {
		return nil, StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_1",
		)
	}

	type PrepareRenameResult_Or_1Marshal PrepareRenameResult_Or_1
	tmpMarshal := PrepareRenameResult_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type PrepareRenameResult_Or_2 struct {
	DefaultBehavior *Boolean `json:"defaultBehavior"`
}

func (this *PrepareRenameResult_Or_2) UnmarshalJSON(data []byte) error {
	type PrepareRenameResult_Or_2Unmarshal PrepareRenameResult_Or_2
	var tmpUnmarshal PrepareRenameResult_Or_2Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.DefaultBehavior == nil {
		return StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_2",
		)
	}

	*this = PrepareRenameResult_Or_2(tmpUnmarshal)
	return nil
}

func (this *PrepareRenameResult_Or_2) MarshalJSON() ([]byte, error) {

	if this.DefaultBehavior == nil {
		return nil, StructureLiteralValidateFailed(
			"PrepareRenameResult_Or_2",
		)
	}

	type PrepareRenameResult_Or_2Marshal PrepareRenameResult_Or_2
	tmpMarshal := PrepareRenameResult_Or_2Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type PublishDiagnosticsClientCapabilities_TagSupport struct {

	// The tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}

func (this *PublishDiagnosticsClientCapabilities_TagSupport) UnmarshalJSON(
	data []byte,
) error {
	type PublishDiagnosticsClientCapabilities_TagSupportUnmarshal PublishDiagnosticsClientCapabilities_TagSupport
	var tmpUnmarshal PublishDiagnosticsClientCapabilities_TagSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"PublishDiagnosticsClientCapabilities_TagSupport",
		)
	}

	*this = PublishDiagnosticsClientCapabilities_TagSupport(tmpUnmarshal)
	return nil
}

func (this *PublishDiagnosticsClientCapabilities_TagSupport) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"PublishDiagnosticsClientCapabilities_TagSupport",
		)
	}

	type PublishDiagnosticsClientCapabilities_TagSupportMarshal PublishDiagnosticsClientCapabilities_TagSupport
	tmpMarshal := PublishDiagnosticsClientCapabilities_TagSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type SemanticTokensClientCapabilities_Requests struct {

	// The client will send the `textDocument/semanticTokens/range` request
	// if the server provides a corresponding handler.
	Range *SemanticTokensClientCapabilities_Requests_Range_Or `json:"range"`

	// The client will send the `textDocument/semanticTokens/full` request
	// if the server provides a corresponding handler.
	Full *SemanticTokensClientCapabilities_Requests_Full_Or `json:"full"`
}

type SemanticTokensClientCapabilities_Requests_Full_Or_1 struct {

	// The client will send the `textDocument/semanticTokens/full/delta`
	// request if the server provides a corresponding handler.
	Delta *Boolean `json:"delta"`
}

type SemanticTokensClientCapabilities_Requests_Range_Or_1 struct {
}

type SemanticTokensOptions_Full_Or_1 struct {

	// The server supports deltas for full documents.
	Delta *Boolean `json:"delta"`
}

type SemanticTokensOptions_Range_Or_1 struct {
}

type ServerCapabilities_Workspace struct {

	// The server supports workspace folder.  @since 3.6.0
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders"`

	// The server is interested in notifications/requests for operations on
	// files.  @since 3.16.0
	FileOperations *FileOperationOptions `json:"fileOperations"`
}

type ShowMessageRequestClientCapabilities_MessageActionItem struct {

	// Whether the client supports additional attributes which are preserved
	// and send back to the server in the request's response.
	AdditionalPropertiesSupport *Boolean `json:"additionalPropertiesSupport"`
}

type SignatureHelpClientCapabilities_SignatureInformation struct {

	// Client supports the following content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat"`

	// Client capabilities specific to parameter information.
	ParameterInformation *SignatureHelpClientCapabilities_SignatureInformation_ParameterInformation `json:"parameterInformation"`

	// The client supports the `activeParameter` property on
	// `SignatureInformation` literal.  @since 3.16.0
	ActiveParameterSupport *Boolean `json:"activeParameterSupport"`
}

type SignatureHelpClientCapabilities_SignatureInformation_ParameterInformation struct {

	// The client supports processing label offsets instead of a simple
	// label string.  @since 3.14.0
	LabelOffsetSupport *Boolean `json:"labelOffsetSupport"`
}

type TextDocumentContentChangeEvent_Or_0 struct {

	// The range of the document that changed.
	Range *Range `json:"range"`

	// The optional length of the range that got replaced.  @deprecated use
	// range instead.
	RangeLength *Uinteger `json:"rangeLength"`

	// The new text for the provided range.
	Text *String `json:"text"`
}

func (this *TextDocumentContentChangeEvent_Or_0) UnmarshalJSON(
	data []byte,
) error {
	type TextDocumentContentChangeEvent_Or_0Unmarshal TextDocumentContentChangeEvent_Or_0
	var tmpUnmarshal TextDocumentContentChangeEvent_Or_0Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Range == nil {
		return StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_0",
		)
	}

	if tmpUnmarshal.Text == nil {
		return StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_0",
		)
	}

	*this = TextDocumentContentChangeEvent_Or_0(tmpUnmarshal)
	return nil
}

func (this *TextDocumentContentChangeEvent_Or_0) MarshalJSON() ([]byte, error) {

	if this.Range == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_0",
		)
	}

	if this.Text == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_0",
		)
	}

	type TextDocumentContentChangeEvent_Or_0Marshal TextDocumentContentChangeEvent_Or_0
	tmpMarshal := TextDocumentContentChangeEvent_Or_0Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type TextDocumentContentChangeEvent_Or_1 struct {

	// The new text of the whole document.
	Text *String `json:"text"`
}

func (this *TextDocumentContentChangeEvent_Or_1) UnmarshalJSON(
	data []byte,
) error {
	type TextDocumentContentChangeEvent_Or_1Unmarshal TextDocumentContentChangeEvent_Or_1
	var tmpUnmarshal TextDocumentContentChangeEvent_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Text == nil {
		return StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_1",
		)
	}

	*this = TextDocumentContentChangeEvent_Or_1(tmpUnmarshal)
	return nil
}

func (this *TextDocumentContentChangeEvent_Or_1) MarshalJSON() ([]byte, error) {

	if this.Text == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentContentChangeEvent_Or_1",
		)
	}

	type TextDocumentContentChangeEvent_Or_1Marshal TextDocumentContentChangeEvent_Or_1
	tmpMarshal := TextDocumentContentChangeEvent_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type TextDocumentFilter_Or_0 struct {

	// A language id, like `typescript`.
	Language *String `json:"language"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern, like `*.{ts,js}`.
	Pattern *String `json:"pattern"`
}

func (this *TextDocumentFilter_Or_0) UnmarshalJSON(data []byte) error {
	type TextDocumentFilter_Or_0Unmarshal TextDocumentFilter_Or_0
	var tmpUnmarshal TextDocumentFilter_Or_0Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Language == nil {
		return StructureLiteralValidateFailed("TextDocumentFilter_Or_0")
	}

	*this = TextDocumentFilter_Or_0(tmpUnmarshal)
	return nil
}

func (this *TextDocumentFilter_Or_0) MarshalJSON() ([]byte, error) {

	if this.Language == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentFilter_Or_0",
		)
	}

	type TextDocumentFilter_Or_0Marshal TextDocumentFilter_Or_0
	tmpMarshal := TextDocumentFilter_Or_0Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type TextDocumentFilter_Or_1 struct {

	// A language id, like `typescript`.
	Language *String `json:"language"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern, like `*.{ts,js}`.
	Pattern *String `json:"pattern"`
}

func (this *TextDocumentFilter_Or_1) UnmarshalJSON(data []byte) error {
	type TextDocumentFilter_Or_1Unmarshal TextDocumentFilter_Or_1
	var tmpUnmarshal TextDocumentFilter_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Scheme == nil {
		return StructureLiteralValidateFailed("TextDocumentFilter_Or_1")
	}

	*this = TextDocumentFilter_Or_1(tmpUnmarshal)
	return nil
}

func (this *TextDocumentFilter_Or_1) MarshalJSON() ([]byte, error) {

	if this.Scheme == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentFilter_Or_1",
		)
	}

	type TextDocumentFilter_Or_1Marshal TextDocumentFilter_Or_1
	tmpMarshal := TextDocumentFilter_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type TextDocumentFilter_Or_2 struct {

	// A language id, like `typescript`.
	Language *String `json:"language"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme *String `json:"scheme"`

	// A glob pattern, like `*.{ts,js}`.
	Pattern *String `json:"pattern"`
}

func (this *TextDocumentFilter_Or_2) UnmarshalJSON(data []byte) error {
	type TextDocumentFilter_Or_2Unmarshal TextDocumentFilter_Or_2
	var tmpUnmarshal TextDocumentFilter_Or_2Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Pattern == nil {
		return StructureLiteralValidateFailed("TextDocumentFilter_Or_2")
	}

	*this = TextDocumentFilter_Or_2(tmpUnmarshal)
	return nil
}

func (this *TextDocumentFilter_Or_2) MarshalJSON() ([]byte, error) {

	if this.Pattern == nil {
		return nil, StructureLiteralValidateFailed(
			"TextDocumentFilter_Or_2",
		)
	}

	type TextDocumentFilter_Or_2Marshal TextDocumentFilter_Or_2
	tmpMarshal := TextDocumentFilter_Or_2Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkspaceEditClientCapabilities_ChangeAnnotationSupport struct {

	// Whether the client groups edits with equal labels into tree nodes,
	// for instance all edits labelled with "Changes in Strings" would be a
	// tree node.
	GroupsOnLabel *Boolean `json:"groupsOnLabel"`
}

type WorkspaceSymbolClientCapabilities_ResolveSupport struct {

	// The properties that a client can resolve lazily. Usually
	// `location.range`
	Properties []String `json:"properties"`
}

func (this *WorkspaceSymbolClientCapabilities_ResolveSupport) UnmarshalJSON(
	data []byte,
) error {
	type WorkspaceSymbolClientCapabilities_ResolveSupportUnmarshal WorkspaceSymbolClientCapabilities_ResolveSupport
	var tmpUnmarshal WorkspaceSymbolClientCapabilities_ResolveSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Properties == nil {
		return StructureLiteralValidateFailed(
			"WorkspaceSymbolClientCapabilities_ResolveSupport",
		)
	}

	*this = WorkspaceSymbolClientCapabilities_ResolveSupport(tmpUnmarshal)
	return nil
}

func (this *WorkspaceSymbolClientCapabilities_ResolveSupport) MarshalJSON() ([]byte, error) {

	if this.Properties == nil {
		return nil, StructureLiteralValidateFailed(
			"WorkspaceSymbolClientCapabilities_ResolveSupport",
		)
	}

	type WorkspaceSymbolClientCapabilities_ResolveSupportMarshal WorkspaceSymbolClientCapabilities_ResolveSupport
	tmpMarshal := WorkspaceSymbolClientCapabilities_ResolveSupportMarshal(
		*this,
	)
	return json.Marshal(&tmpMarshal)
}

type WorkspaceSymbolClientCapabilities_SymbolKind struct {

	// The symbol kind values the client supports. When this property exists
	// the client also guarantees that it will handle values outside its set
	// gracefully and falls back to a default value when unknown.  If this
	// property is not present the client only supports the symbol kinds
	// from `File` to `Array` as defined in the initial version of the
	// protocol.
	ValueSet []SymbolKind `json:"valueSet"`
}

type WorkspaceSymbolClientCapabilities_TagSupport struct {

	// The tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}

func (this *WorkspaceSymbolClientCapabilities_TagSupport) UnmarshalJSON(
	data []byte,
) error {
	type WorkspaceSymbolClientCapabilities_TagSupportUnmarshal WorkspaceSymbolClientCapabilities_TagSupport
	var tmpUnmarshal WorkspaceSymbolClientCapabilities_TagSupportUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.ValueSet == nil {
		return StructureLiteralValidateFailed(
			"WorkspaceSymbolClientCapabilities_TagSupport",
		)
	}

	*this = WorkspaceSymbolClientCapabilities_TagSupport(tmpUnmarshal)
	return nil
}

func (this *WorkspaceSymbolClientCapabilities_TagSupport) MarshalJSON() ([]byte, error) {

	if this.ValueSet == nil {
		return nil, StructureLiteralValidateFailed(
			"WorkspaceSymbolClientCapabilities_TagSupport",
		)
	}

	type WorkspaceSymbolClientCapabilities_TagSupportMarshal WorkspaceSymbolClientCapabilities_TagSupport
	tmpMarshal := WorkspaceSymbolClientCapabilities_TagSupportMarshal(*this)
	return json.Marshal(&tmpMarshal)
}

type WorkspaceSymbol_Location_Or_1 struct {
	Uri *DocumentUri `json:"uri"`
}

func (this *WorkspaceSymbol_Location_Or_1) UnmarshalJSON(data []byte) error {
	type WorkspaceSymbol_Location_Or_1Unmarshal WorkspaceSymbol_Location_Or_1
	var tmpUnmarshal WorkspaceSymbol_Location_Or_1Unmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Uri == nil {
		return StructureLiteralValidateFailed(
			"WorkspaceSymbol_Location_Or_1",
		)
	}

	*this = WorkspaceSymbol_Location_Or_1(tmpUnmarshal)
	return nil
}

func (this *WorkspaceSymbol_Location_Or_1) MarshalJSON() ([]byte, error) {

	if this.Uri == nil {
		return nil, StructureLiteralValidateFailed(
			"WorkspaceSymbol_Location_Or_1",
		)
	}

	type WorkspaceSymbol_Location_Or_1Marshal WorkspaceSymbol_Location_Or_1
	tmpMarshal := WorkspaceSymbol_Location_Or_1Marshal(*this)
	return json.Marshal(&tmpMarshal)
}

type XInitializeParams_ClientInfo struct {

	// The name of the client as defined by the client.
	Name *String `json:"name"`

	// The client's version as defined by the client.
	Version *String `json:"version"`
}

func (this *XInitializeParams_ClientInfo) UnmarshalJSON(data []byte) error {
	type XInitializeParams_ClientInfoUnmarshal XInitializeParams_ClientInfo
	var tmpUnmarshal XInitializeParams_ClientInfoUnmarshal
	err := json.Unmarshal(data, &tmpUnmarshal)
	if err != nil {
		return err
	}

	if tmpUnmarshal.Name == nil {
		return StructureLiteralValidateFailed(
			"XInitializeParams_ClientInfo",
		)
	}

	*this = XInitializeParams_ClientInfo(tmpUnmarshal)
	return nil
}

func (this *XInitializeParams_ClientInfo) MarshalJSON() ([]byte, error) {

	if this.Name == nil {
		return nil, StructureLiteralValidateFailed(
			"XInitializeParams_ClientInfo",
		)
	}

	type XInitializeParams_ClientInfoMarshal XInitializeParams_ClientInfo
	tmpMarshal := XInitializeParams_ClientInfoMarshal(*this)
	return json.Marshal(&tmpMarshal)
}
