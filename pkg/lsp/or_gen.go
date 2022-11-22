// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"encoding/json"
	"fmt"
)

var OrValidateFailed = func(name string) error {
	return fmt.Errorf("or type \"%s\"validate failed", name)
}

type CancelParams_Id_Or struct {
	// Or [ *Integer *String ]
	V interface{}
}

func (this *CancelParams_Id_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Integer
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("CancelParams_Id_Or")
}

func (this *CancelParams_Id_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("CancelParams_Id_Or")
	}
	return json.Marshal(this.V)
}

type CompletionItem_Documentation_Or struct {
	// Or [ *String *MarkupContent ]
	V interface{}
}

func (this *CompletionItem_Documentation_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("CompletionItem_Documentation_Or")
}

func (this *CompletionItem_Documentation_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("CompletionItem_Documentation_Or")
	}
	return json.Marshal(this.V)
}

type CompletionItem_TextEdit_Or struct {
	// Or [ *TextEdit *InsertReplaceEdit ]
	V interface{}
}

func (this *CompletionItem_TextEdit_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *TextEdit
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InsertReplaceEdit
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("CompletionItem_TextEdit_Or")
}

func (this *CompletionItem_TextEdit_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextEdit); ok {
			break
		}

		if _, ok := this.V.(*InsertReplaceEdit); ok {
			break
		}

		return nil, OrValidateFailed("CompletionItem_TextEdit_Or")
	}
	return json.Marshal(this.V)
}

type CompletionList_ItemDefaults_EditRange_Or struct {
	// Or [ *Range *CompletionList_ItemDefaults_EditRange_Or_1 ]
	V interface{}
}

func (this *CompletionList_ItemDefaults_EditRange_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Range
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CompletionList_ItemDefaults_EditRange_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("CompletionList_ItemDefaults_EditRange_Or")
}

func (this *CompletionList_ItemDefaults_EditRange_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Range); ok {
			break
		}

		if _, ok := this.V.(*CompletionList_ItemDefaults_EditRange_Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"CompletionList_ItemDefaults_EditRange_Or",
		)
	}
	return json.Marshal(this.V)
}

type Declaration_Or struct {
	// Or [ *Location []Location ]
	V interface{}
}

func (this *Declaration_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("Declaration_Or")
}

func (this *Declaration_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.([]Location); ok {
			break
		}

		return nil, OrValidateFailed("Declaration_Or")
	}
	return json.Marshal(this.V)
}

type Definition_Or struct {
	// Or [ *Location []Location ]
	V interface{}
}

func (this *Definition_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("Definition_Or")
}

func (this *Definition_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.([]Location); ok {
			break
		}

		return nil, OrValidateFailed("Definition_Or")
	}
	return json.Marshal(this.V)
}

type Diagnostic_Code_Or struct {
	// Or [ *Integer *String ]
	V interface{}
}

func (this *Diagnostic_Code_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Integer
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("Diagnostic_Code_Or")
}

func (this *Diagnostic_Code_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("Diagnostic_Code_Or")
	}
	return json.Marshal(this.V)
}

type DidChangeConfigurationRegistrationOptions_Section_Or struct {
	// Or [ *String []String ]
	V interface{}
}

func (this *DidChangeConfigurationRegistrationOptions_Section_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"DidChangeConfigurationRegistrationOptions_Section_Or",
	)
}

func (this *DidChangeConfigurationRegistrationOptions_Section_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.([]String); ok {
			break
		}

		return nil, OrValidateFailed(
			"DidChangeConfigurationRegistrationOptions_Section_Or",
		)
	}
	return json.Marshal(this.V)
}

type DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or struct {
	// Or [ *FullDocumentDiagnosticReport *UnchangedDocumentDiagnosticReport
	// ]
	V interface{}
}

func (this *DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *FullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *UnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or",
	)
}

func (this *DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"DocumentDiagnosticReportPartialResult_RelatedDocuments_Value_Or",
		)
	}
	return json.Marshal(this.V)
}

type DocumentDiagnosticReport_Or struct {
	// Or [ *RelatedFullDocumentDiagnosticReport
	// *RelatedUnchangedDocumentDiagnosticReport ]
	V interface{}
}

func (this *DocumentDiagnosticReport_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *RelatedFullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *RelatedUnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("DocumentDiagnosticReport_Or")
}

func (this *DocumentDiagnosticReport_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*RelatedFullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*RelatedUnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed("DocumentDiagnosticReport_Or")
	}
	return json.Marshal(this.V)
}

type DocumentFilter_Or struct {
	// Or [ *TextDocumentFilter *NotebookCellTextDocumentFilter ]
	V interface{}
}

func (this *DocumentFilter_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *TextDocumentFilter
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookCellTextDocumentFilter
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("DocumentFilter_Or")
}

func (this *DocumentFilter_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentFilter); ok {
			break
		}

		if _, ok := this.V.(*NotebookCellTextDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed("DocumentFilter_Or")
	}
	return json.Marshal(this.V)
}

type GlobPattern_Or struct {
	// Or [ *Pattern *RelativePattern ]
	V interface{}
}

func (this *GlobPattern_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Pattern
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *RelativePattern
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("GlobPattern_Or")
}

func (this *GlobPattern_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Pattern); ok {
			break
		}

		if _, ok := this.V.(*RelativePattern); ok {
			break
		}

		return nil, OrValidateFailed("GlobPattern_Or")
	}
	return json.Marshal(this.V)
}

type Hover_Contents_Or struct {
	// Or [ *MarkupContent *MarkedString []MarkedString ]
	V interface{}
}

func (this *Hover_Contents_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkedString
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []MarkedString
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("Hover_Contents_Or")
}

func (this *Hover_Contents_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		if _, ok := this.V.(*MarkedString); ok {
			break
		}

		if _, ok := this.V.([]MarkedString); ok {
			break
		}

		return nil, OrValidateFailed("Hover_Contents_Or")
	}
	return json.Marshal(this.V)
}

type InlayHintLabelPart_Tooltip_Or struct {
	// Or [ *String *MarkupContent ]
	V interface{}
}

func (this *InlayHintLabelPart_Tooltip_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("InlayHintLabelPart_Tooltip_Or")
}

func (this *InlayHintLabelPart_Tooltip_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("InlayHintLabelPart_Tooltip_Or")
	}
	return json.Marshal(this.V)
}

type InlayHint_Label_Or struct {
	// Or [ *String []InlayHintLabelPart ]
	V interface{}
}

func (this *InlayHint_Label_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []InlayHintLabelPart
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("InlayHint_Label_Or")
}

func (this *InlayHint_Label_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.([]InlayHintLabelPart); ok {
			break
		}

		return nil, OrValidateFailed("InlayHint_Label_Or")
	}
	return json.Marshal(this.V)
}

type InlayHint_Tooltip_Or struct {
	// Or [ *String *MarkupContent ]
	V interface{}
}

func (this *InlayHint_Tooltip_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("InlayHint_Tooltip_Or")
}

func (this *InlayHint_Tooltip_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("InlayHint_Tooltip_Or")
	}
	return json.Marshal(this.V)
}

type InlineValue_Or struct {
	// Or [ *InlineValueText *InlineValueVariableLookup
	// *InlineValueEvaluatableExpression ]
	V interface{}
}

func (this *InlineValue_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *InlineValueText
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlineValueVariableLookup
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlineValueEvaluatableExpression
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("InlineValue_Or")
}

func (this *InlineValue_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*InlineValueText); ok {
			break
		}

		if _, ok := this.V.(*InlineValueVariableLookup); ok {
			break
		}

		if _, ok := this.V.(*InlineValueEvaluatableExpression); ok {
			break
		}

		return nil, OrValidateFailed("InlineValue_Or")
	}
	return json.Marshal(this.V)
}

type LSPAny_Or struct {
	// Or [ *LSPObject *LSPArray *String *Integer *Uinteger *Decimal
	// *Boolean *Null ]
	V interface{}
}

func (this *LSPAny_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *LSPObject
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *LSPArray
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Integer
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Uinteger
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Decimal
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("LSPAny_Or")
}

func (this *LSPAny_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*LSPObject); ok {
			break
		}

		if _, ok := this.V.(*LSPArray); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*Uinteger); ok {
			break
		}

		if _, ok := this.V.(*Decimal); ok {
			break
		}

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed("LSPAny_Or")
	}
	return json.Marshal(this.V)
}

type MarkedString_Or struct {
	// Or [ *String *MarkedString_Or_1 ]
	V interface{}
}

func (this *MarkedString_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkedString_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("MarkedString_Or")
}

func (this *MarkedString_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkedString_Or_1); ok {
			break
		}

		return nil, OrValidateFailed("MarkedString_Or")
	}
	return json.Marshal(this.V)
}

type NotebookCellTextDocumentFilter_Notebook_Or struct {
	// Or [ *String *NotebookDocumentFilter ]
	V interface{}
}

func (this *NotebookCellTextDocumentFilter_Notebook_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("NotebookCellTextDocumentFilter_Notebook_Or")
}

func (this *NotebookCellTextDocumentFilter_Notebook_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookCellTextDocumentFilter_Notebook_Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentFilter_Or struct {
	// Or [ *NotebookDocumentFilter_Or_0 *NotebookDocumentFilter_Or_1
	// *NotebookDocumentFilter_Or_2 ]
	V interface{}
}

func (this *NotebookDocumentFilter_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *NotebookDocumentFilter_Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter_Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("NotebookDocumentFilter_Or")
}

func (this *NotebookDocumentFilter_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentFilter_Or_0); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter_Or_1); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter_Or_2); ok {
			break
		}

		return nil, OrValidateFailed("NotebookDocumentFilter_Or")
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or struct {
	// Or [ *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0
	// *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1 ]
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element_Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or struct {
	// Or [ *String *NotebookDocumentFilter ]
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_0_Notebook_Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or struct {
	// Or [ *String *NotebookDocumentFilter ]
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element_Or_1_Notebook_Or",
		)
	}
	return json.Marshal(this.V)
}

type ParameterInformation_Documentation_Or struct {
	// Or [ *String *MarkupContent ]
	V interface{}
}

func (this *ParameterInformation_Documentation_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ParameterInformation_Documentation_Or")
}

func (this *ParameterInformation_Documentation_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed(
			"ParameterInformation_Documentation_Or",
		)
	}
	return json.Marshal(this.V)
}

type ParameterInformation_Label_Or struct {
	// Or [ *String *ParameterInformation_Label_Or_1_Tuple ]
	V interface{}
}

func (this *ParameterInformation_Label_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *ParameterInformation_Label_Or_1_Tuple
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ParameterInformation_Label_Or")
}

func (this *ParameterInformation_Label_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*ParameterInformation_Label_Or_1_Tuple); ok {
			break
		}

		return nil, OrValidateFailed("ParameterInformation_Label_Or")
	}
	return json.Marshal(this.V)
}

type PrepareRenameResult_Or struct {
	// Or [ *Range *PrepareRenameResult_Or_1 *PrepareRenameResult_Or_2 ]
	V interface{}
}

func (this *PrepareRenameResult_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Range
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *PrepareRenameResult_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *PrepareRenameResult_Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("PrepareRenameResult_Or")
}

func (this *PrepareRenameResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Range); ok {
			break
		}

		if _, ok := this.V.(*PrepareRenameResult_Or_1); ok {
			break
		}

		if _, ok := this.V.(*PrepareRenameResult_Or_2); ok {
			break
		}

		return nil, OrValidateFailed("PrepareRenameResult_Or")
	}
	return json.Marshal(this.V)
}

type ProgressToken_Or struct {
	// Or [ *Integer *String ]
	V interface{}
}

func (this *ProgressToken_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Integer
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ProgressToken_Or")
}

func (this *ProgressToken_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("ProgressToken_Or")
	}
	return json.Marshal(this.V)
}

type RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or struct {
	// Or [ *FullDocumentDiagnosticReport *UnchangedDocumentDiagnosticReport
	// ]
	V interface{}
}

func (this *RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *FullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *UnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or",
	)
}

func (this *RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value_Or",
		)
	}
	return json.Marshal(this.V)
}

type RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or struct {
	// Or [ *FullDocumentDiagnosticReport *UnchangedDocumentDiagnosticReport
	// ]
	V interface{}
}

func (this *RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *FullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *UnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or",
	)
}

func (this *RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value_Or",
		)
	}
	return json.Marshal(this.V)
}

type RelativePattern_BaseUri_Or struct {
	// Or [ *WorkspaceFolder *URI ]
	V interface{}
}

func (this *RelativePattern_BaseUri_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *WorkspaceFolder
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *URI
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("RelativePattern_BaseUri_Or")
}

func (this *RelativePattern_BaseUri_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*WorkspaceFolder); ok {
			break
		}

		if _, ok := this.V.(*URI); ok {
			break
		}

		return nil, OrValidateFailed("RelativePattern_BaseUri_Or")
	}
	return json.Marshal(this.V)
}

type SemanticTokensClientCapabilities_Requests_Full_Or struct {
	// Or [ *Boolean *SemanticTokensClientCapabilities_Requests_Full_Or_1 ]
	V interface{}
}

func (this *SemanticTokensClientCapabilities_Requests_Full_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensClientCapabilities_Requests_Full_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"SemanticTokensClientCapabilities_Requests_Full_Or",
	)
}

func (this *SemanticTokensClientCapabilities_Requests_Full_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensClientCapabilities_Requests_Full_Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"SemanticTokensClientCapabilities_Requests_Full_Or",
		)
	}
	return json.Marshal(this.V)
}

type SemanticTokensClientCapabilities_Requests_Range_Or struct {
	// Or [ *Boolean *SemanticTokensClientCapabilities_Requests_Range_Or_1 ]
	V interface{}
}

func (this *SemanticTokensClientCapabilities_Requests_Range_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensClientCapabilities_Requests_Range_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"SemanticTokensClientCapabilities_Requests_Range_Or",
	)
}

func (this *SemanticTokensClientCapabilities_Requests_Range_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensClientCapabilities_Requests_Range_Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"SemanticTokensClientCapabilities_Requests_Range_Or",
		)
	}
	return json.Marshal(this.V)
}

type SemanticTokensOptions_Full_Or struct {
	// Or [ *Boolean *SemanticTokensOptions_Full_Or_1 ]
	V interface{}
}

func (this *SemanticTokensOptions_Full_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensOptions_Full_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("SemanticTokensOptions_Full_Or")
}

func (this *SemanticTokensOptions_Full_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensOptions_Full_Or_1); ok {
			break
		}

		return nil, OrValidateFailed("SemanticTokensOptions_Full_Or")
	}
	return json.Marshal(this.V)
}

type SemanticTokensOptions_Range_Or struct {
	// Or [ *Boolean *SemanticTokensOptions_Range_Or_1 ]
	V interface{}
}

func (this *SemanticTokensOptions_Range_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensOptions_Range_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("SemanticTokensOptions_Range_Or")
}

func (this *SemanticTokensOptions_Range_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensOptions_Range_Or_1); ok {
			break
		}

		return nil, OrValidateFailed("SemanticTokensOptions_Range_Or")
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_CallHierarchyProvider_Or struct {
	// Or [ *Boolean *CallHierarchyOptions *CallHierarchyRegistrationOptions
	// ]
	V interface{}
}

func (this *ServerCapabilities_CallHierarchyProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CallHierarchyOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CallHierarchyRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_CallHierarchyProvider_Or")
}

func (this *ServerCapabilities_CallHierarchyProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*CallHierarchyOptions); ok {
			break
		}

		if _, ok := this.V.(*CallHierarchyRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_CallHierarchyProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_CodeActionProvider_Or struct {
	// Or [ *Boolean *CodeActionOptions ]
	V interface{}
}

func (this *ServerCapabilities_CodeActionProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CodeActionOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_CodeActionProvider_Or")
}

func (this *ServerCapabilities_CodeActionProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*CodeActionOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_CodeActionProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ColorProvider_Or struct {
	// Or [ *Boolean *DocumentColorOptions *DocumentColorRegistrationOptions
	// ]
	V interface{}
}

func (this *ServerCapabilities_ColorProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentColorOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentColorRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_ColorProvider_Or")
}

func (this *ServerCapabilities_ColorProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentColorOptions); ok {
			break
		}

		if _, ok := this.V.(*DocumentColorRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_ColorProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DeclarationProvider_Or struct {
	// Or [ *Boolean *DeclarationOptions *DeclarationRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_DeclarationProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DeclarationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DeclarationRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_DeclarationProvider_Or")
}

func (this *ServerCapabilities_DeclarationProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DeclarationOptions); ok {
			break
		}

		if _, ok := this.V.(*DeclarationRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DeclarationProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DefinitionProvider_Or struct {
	// Or [ *Boolean *DefinitionOptions ]
	V interface{}
}

func (this *ServerCapabilities_DefinitionProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DefinitionOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_DefinitionProvider_Or")
}

func (this *ServerCapabilities_DefinitionProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DefinitionOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DefinitionProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DiagnosticProvider_Or struct {
	// Or [ *DiagnosticOptions *DiagnosticRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_DiagnosticProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *DiagnosticOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DiagnosticRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_DiagnosticProvider_Or")
}

func (this *ServerCapabilities_DiagnosticProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*DiagnosticOptions); ok {
			break
		}

		if _, ok := this.V.(*DiagnosticRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DiagnosticProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentFormattingProvider_Or struct {
	// Or [ *Boolean *DocumentFormattingOptions ]
	V interface{}
}

func (this *ServerCapabilities_DocumentFormattingProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentFormattingOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"ServerCapabilities_DocumentFormattingProvider_Or",
	)
}

func (this *ServerCapabilities_DocumentFormattingProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentFormattingOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentFormattingProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentHighlightProvider_Or struct {
	// Or [ *Boolean *DocumentHighlightOptions ]
	V interface{}
}

func (this *ServerCapabilities_DocumentHighlightProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentHighlightOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"ServerCapabilities_DocumentHighlightProvider_Or",
	)
}

func (this *ServerCapabilities_DocumentHighlightProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentHighlightOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentHighlightProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentRangeFormattingProvider_Or struct {
	// Or [ *Boolean *DocumentRangeFormattingOptions ]
	V interface{}
}

func (this *ServerCapabilities_DocumentRangeFormattingProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentRangeFormattingOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"ServerCapabilities_DocumentRangeFormattingProvider_Or",
	)
}

func (this *ServerCapabilities_DocumentRangeFormattingProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentRangeFormattingOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentRangeFormattingProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentSymbolProvider_Or struct {
	// Or [ *Boolean *DocumentSymbolOptions ]
	V interface{}
}

func (this *ServerCapabilities_DocumentSymbolProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DocumentSymbolOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_DocumentSymbolProvider_Or")
}

func (this *ServerCapabilities_DocumentSymbolProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentSymbolOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentSymbolProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_FoldingRangeProvider_Or struct {
	// Or [ *Boolean *FoldingRangeOptions *FoldingRangeRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_FoldingRangeProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *FoldingRangeOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *FoldingRangeRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_FoldingRangeProvider_Or")
}

func (this *ServerCapabilities_FoldingRangeProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*FoldingRangeOptions); ok {
			break
		}

		if _, ok := this.V.(*FoldingRangeRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_FoldingRangeProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_HoverProvider_Or struct {
	// Or [ *Boolean *HoverOptions ]
	V interface{}
}

func (this *ServerCapabilities_HoverProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *HoverOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_HoverProvider_Or")
}

func (this *ServerCapabilities_HoverProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*HoverOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_HoverProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ImplementationProvider_Or struct {
	// Or [ *Boolean *ImplementationOptions
	// *ImplementationRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_ImplementationProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *ImplementationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *ImplementationRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_ImplementationProvider_Or")
}

func (this *ServerCapabilities_ImplementationProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*ImplementationOptions); ok {
			break
		}

		if _, ok := this.V.(*ImplementationRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_ImplementationProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_InlayHintProvider_Or struct {
	// Or [ *Boolean *InlayHintOptions *InlayHintRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_InlayHintProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlayHintOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlayHintRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_InlayHintProvider_Or")
}

func (this *ServerCapabilities_InlayHintProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*InlayHintOptions); ok {
			break
		}

		if _, ok := this.V.(*InlayHintRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_InlayHintProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_InlineValueProvider_Or struct {
	// Or [ *Boolean *InlineValueOptions *InlineValueRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_InlineValueProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlineValueOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *InlineValueRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_InlineValueProvider_Or")
}

func (this *ServerCapabilities_InlineValueProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*InlineValueOptions); ok {
			break
		}

		if _, ok := this.V.(*InlineValueRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_InlineValueProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_LinkedEditingRangeProvider_Or struct {
	// Or [ *Boolean *LinkedEditingRangeOptions
	// *LinkedEditingRangeRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_LinkedEditingRangeProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *LinkedEditingRangeOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *LinkedEditingRangeRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"ServerCapabilities_LinkedEditingRangeProvider_Or",
	)
}

func (this *ServerCapabilities_LinkedEditingRangeProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*LinkedEditingRangeOptions); ok {
			break
		}

		if _, ok := this.V.(*LinkedEditingRangeRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_LinkedEditingRangeProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_MonikerProvider_Or struct {
	// Or [ *Boolean *MonikerOptions *MonikerRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_MonikerProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MonikerOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MonikerRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_MonikerProvider_Or")
}

func (this *ServerCapabilities_MonikerProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*MonikerOptions); ok {
			break
		}

		if _, ok := this.V.(*MonikerRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_MonikerProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_NotebookDocumentSync_Or struct {
	// Or [ *NotebookDocumentSyncOptions
	// *NotebookDocumentSyncRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_NotebookDocumentSync_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *NotebookDocumentSyncOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentSyncRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_NotebookDocumentSync_Or")
}

func (this *ServerCapabilities_NotebookDocumentSync_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentSyncOptions); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentSyncRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_NotebookDocumentSync_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ReferencesProvider_Or struct {
	// Or [ *Boolean *ReferenceOptions ]
	V interface{}
}

func (this *ServerCapabilities_ReferencesProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *ReferenceOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_ReferencesProvider_Or")
}

func (this *ServerCapabilities_ReferencesProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*ReferenceOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_ReferencesProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_RenameProvider_Or struct {
	// Or [ *Boolean *RenameOptions ]
	V interface{}
}

func (this *ServerCapabilities_RenameProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *RenameOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_RenameProvider_Or")
}

func (this *ServerCapabilities_RenameProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*RenameOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_RenameProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_SelectionRangeProvider_Or struct {
	// Or [ *Boolean *SelectionRangeOptions
	// *SelectionRangeRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_SelectionRangeProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SelectionRangeOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SelectionRangeRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_SelectionRangeProvider_Or")
}

func (this *ServerCapabilities_SelectionRangeProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SelectionRangeOptions); ok {
			break
		}

		if _, ok := this.V.(*SelectionRangeRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_SelectionRangeProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_SemanticTokensProvider_Or struct {
	// Or [ *SemanticTokensOptions *SemanticTokensRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_SemanticTokensProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *SemanticTokensOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_SemanticTokensProvider_Or")
}

func (this *ServerCapabilities_SemanticTokensProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*SemanticTokensOptions); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_SemanticTokensProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TextDocumentSync_Or struct {
	// Or [ *TextDocumentSyncOptions *TextDocumentSyncKind ]
	V interface{}
}

func (this *ServerCapabilities_TextDocumentSync_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *TextDocumentSyncOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentSyncKind
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_TextDocumentSync_Or")
}

func (this *ServerCapabilities_TextDocumentSync_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentSyncOptions); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentSyncKind); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_TextDocumentSync_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TypeDefinitionProvider_Or struct {
	// Or [ *Boolean *TypeDefinitionOptions
	// *TypeDefinitionRegistrationOptions ]
	V interface{}
}

func (this *ServerCapabilities_TypeDefinitionProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TypeDefinitionOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TypeDefinitionRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_TypeDefinitionProvider_Or")
}

func (this *ServerCapabilities_TypeDefinitionProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*TypeDefinitionOptions); ok {
			break
		}

		if _, ok := this.V.(*TypeDefinitionRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_TypeDefinitionProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TypeHierarchyProvider_Or struct {
	// Or [ *Boolean *TypeHierarchyOptions *TypeHierarchyRegistrationOptions
	// ]
	V interface{}
}

func (this *ServerCapabilities_TypeHierarchyProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TypeHierarchyOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TypeHierarchyRegistrationOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_TypeHierarchyProvider_Or")
}

func (this *ServerCapabilities_TypeHierarchyProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*TypeHierarchyOptions); ok {
			break
		}

		if _, ok := this.V.(*TypeHierarchyRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_TypeHierarchyProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_WorkspaceSymbolProvider_Or struct {
	// Or [ *Boolean *WorkspaceSymbolOptions ]
	V interface{}
}

func (this *ServerCapabilities_WorkspaceSymbolProvider_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *WorkspaceSymbolOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ServerCapabilities_WorkspaceSymbolProvider_Or")
}

func (this *ServerCapabilities_WorkspaceSymbolProvider_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceSymbolOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_WorkspaceSymbolProvider_Or",
		)
	}
	return json.Marshal(this.V)
}

type SignatureInformation_Documentation_Or struct {
	// Or [ *String *MarkupContent ]
	V interface{}
}

func (this *SignatureInformation_Documentation_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkupContent
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("SignatureInformation_Documentation_Or")
}

func (this *SignatureInformation_Documentation_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed(
			"SignatureInformation_Documentation_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCodeAction_PartialResult_Element_Or struct {
	// Or [ *Command *CodeAction ]
	V interface{}
}

func (this *TextDocumentCodeAction_PartialResult_Element_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Command
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CodeAction
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"TextDocumentCodeAction_PartialResult_Element_Or",
	)
}

func (this *TextDocumentCodeAction_PartialResult_Element_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Command); ok {
			break
		}

		if _, ok := this.V.(*CodeAction); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentCodeAction_PartialResult_Element_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCodeAction_Result_Element_Or struct {
	// Or [ *Command *CodeAction ]
	V interface{}
}

func (this *TextDocumentCodeAction_Result_Element_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Command
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CodeAction
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentCodeAction_Result_Element_Or")
}

func (this *TextDocumentCodeAction_Result_Element_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Command); ok {
			break
		}

		if _, ok := this.V.(*CodeAction); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentCodeAction_Result_Element_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCompletion_Result_Or struct {
	// Or [ []CompletionItem *CompletionList *Null ]
	V interface{}
}

func (this *TextDocumentCompletion_Result_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp []CompletionItem
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CompletionList
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentCompletion_Result_Or")
}

func (this *TextDocumentCompletion_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]CompletionItem); ok {
			break
		}

		if _, ok := this.V.(*CompletionList); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentCompletion_Result_Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentContentChangeEvent_Or struct {
	// Or [ *TextDocumentContentChangeEvent_Or_0
	// *TextDocumentContentChangeEvent_Or_1 ]
	V interface{}
}

func (this *TextDocumentContentChangeEvent_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *TextDocumentContentChangeEvent_Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentContentChangeEvent_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentContentChangeEvent_Or")
}

func (this *TextDocumentContentChangeEvent_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentContentChangeEvent_Or_0); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentContentChangeEvent_Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentContentChangeEvent_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDeclaration_PartialResult_Or struct {
	// Or [ []Location []DeclarationLink ]
	V interface{}
}

func (this *TextDocumentDeclaration_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DeclarationLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDeclaration_PartialResult_Or")
}

func (this *TextDocumentDeclaration_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DeclarationLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDeclaration_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDeclaration_Result_Or struct {
	// Or [ *Declaration []DeclarationLink *Null ]
	V interface{}
}

func (this *TextDocumentDeclaration_Result_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Declaration
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DeclarationLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDeclaration_Result_Or")
}

func (this *TextDocumentDeclaration_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Declaration); ok {
			break
		}

		if _, ok := this.V.([]DeclarationLink); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDeclaration_Result_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDefinition_PartialResult_Or struct {
	// Or [ []Location []DefinitionLink ]
	V interface{}
}

func (this *TextDocumentDefinition_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDefinition_PartialResult_Or")
}

func (this *TextDocumentDefinition_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDefinition_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDefinition_Result_Or struct {
	// Or [ *Definition []DefinitionLink *Null ]
	V interface{}
}

func (this *TextDocumentDefinition_Result_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Definition
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDefinition_Result_Or")
}

func (this *TextDocumentDefinition_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Definition); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentDefinition_Result_Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentDocumentSymbol_PartialResult_Or struct {
	// Or [ []SymbolInformation []DocumentSymbol ]
	V interface{}
}

func (this *TextDocumentDocumentSymbol_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []SymbolInformation
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DocumentSymbol
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDocumentSymbol_PartialResult_Or")
}

func (this *TextDocumentDocumentSymbol_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]DocumentSymbol); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDocumentSymbol_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDocumentSymbol_Result_Or struct {
	// Or [ []SymbolInformation []DocumentSymbol *Null ]
	V interface{}
}

func (this *TextDocumentDocumentSymbol_Result_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []SymbolInformation
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DocumentSymbol
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentDocumentSymbol_Result_Or")
}

func (this *TextDocumentDocumentSymbol_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]DocumentSymbol); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDocumentSymbol_Result_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentEdit_Edits_Element_Or struct {
	// Or [ *TextEdit *AnnotatedTextEdit ]
	V interface{}
}

func (this *TextDocumentEdit_Edits_Element_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *TextEdit
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *AnnotatedTextEdit
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentEdit_Edits_Element_Or")
}

func (this *TextDocumentEdit_Edits_Element_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextEdit); ok {
			break
		}

		if _, ok := this.V.(*AnnotatedTextEdit); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentEdit_Edits_Element_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentFilter_Or struct {
	// Or [ *TextDocumentFilter_Or_0 *TextDocumentFilter_Or_1
	// *TextDocumentFilter_Or_2 ]
	V interface{}
}

func (this *TextDocumentFilter_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *TextDocumentFilter_Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentFilter_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentFilter_Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentFilter_Or")
}

func (this *TextDocumentFilter_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentFilter_Or_0); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentFilter_Or_1); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentFilter_Or_2); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentFilter_Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentImplementation_PartialResult_Or struct {
	// Or [ []Location []DefinitionLink ]
	V interface{}
}

func (this *TextDocumentImplementation_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentImplementation_PartialResult_Or")
}

func (this *TextDocumentImplementation_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentImplementation_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentImplementation_Result_Or struct {
	// Or [ *Definition []DefinitionLink *Null ]
	V interface{}
}

func (this *TextDocumentImplementation_Result_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Definition
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentImplementation_Result_Or")
}

func (this *TextDocumentImplementation_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Definition); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentImplementation_Result_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSemanticTokensFullDelta_PartialResult_Or struct {
	// Or [ *SemanticTokensPartialResult *SemanticTokensDeltaPartialResult ]
	V interface{}
}

func (this *TextDocumentSemanticTokensFullDelta_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *SemanticTokensPartialResult
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensDeltaPartialResult
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"TextDocumentSemanticTokensFullDelta_PartialResult_Or",
	)
}

func (this *TextDocumentSemanticTokensFullDelta_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*SemanticTokensPartialResult); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensDeltaPartialResult); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentSemanticTokensFullDelta_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSemanticTokensFullDelta_Result_Or struct {
	// Or [ *SemanticTokens *SemanticTokensDelta *Null ]
	V interface{}
}

func (this *TextDocumentSemanticTokensFullDelta_Result_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *SemanticTokens
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensDelta
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentSemanticTokensFullDelta_Result_Or")
}

func (this *TextDocumentSemanticTokensFullDelta_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*SemanticTokens); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensDelta); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentSemanticTokensFullDelta_Result_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSyncOptions_Save_Or struct {
	// Or [ *Boolean *SaveOptions ]
	V interface{}
}

func (this *TextDocumentSyncOptions_Save_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SaveOptions
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentSyncOptions_Save_Or")
}

func (this *TextDocumentSyncOptions_Save_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SaveOptions); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentSyncOptions_Save_Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentTypeDefinition_PartialResult_Or struct {
	// Or [ []Location []DefinitionLink ]
	V interface{}
}

func (this *TextDocumentTypeDefinition_PartialResult_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp []Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentTypeDefinition_PartialResult_Or")
}

func (this *TextDocumentTypeDefinition_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentTypeDefinition_PartialResult_Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentTypeDefinition_Result_Or struct {
	// Or [ *Definition []DefinitionLink *Null ]
	V interface{}
}

func (this *TextDocumentTypeDefinition_Result_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *Definition
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []DefinitionLink
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentTypeDefinition_Result_Or")
}

func (this *TextDocumentTypeDefinition_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Definition); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentTypeDefinition_Result_Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceDocumentDiagnosticReport_Or struct {
	// Or [ *WorkspaceFullDocumentDiagnosticReport
	// *WorkspaceUnchangedDocumentDiagnosticReport ]
	V interface{}
}

func (this *WorkspaceDocumentDiagnosticReport_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *WorkspaceFullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *WorkspaceUnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceDocumentDiagnosticReport_Or")
}

func (this *WorkspaceDocumentDiagnosticReport_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*WorkspaceFullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceUnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceDocumentDiagnosticReport_Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceEdit_DocumentChanges_Element_Or struct {
	// Or [ *TextDocumentEdit *CreateFile *RenameFile *DeleteFile ]
	V interface{}
}

func (this *WorkspaceEdit_DocumentChanges_Element_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *TextDocumentEdit
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *CreateFile
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *RenameFile
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *DeleteFile
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceEdit_DocumentChanges_Element_Or")
}

func (this *WorkspaceEdit_DocumentChanges_Element_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentEdit); ok {
			break
		}

		if _, ok := this.V.(*CreateFile); ok {
			break
		}

		if _, ok := this.V.(*RenameFile); ok {
			break
		}

		if _, ok := this.V.(*DeleteFile); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceEdit_DocumentChanges_Element_Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceFoldersServerCapabilities_ChangeNotifications_Or struct {
	// Or [ *String *Boolean ]
	V interface{}
}

func (this *WorkspaceFoldersServerCapabilities_ChangeNotifications_Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"WorkspaceFoldersServerCapabilities_ChangeNotifications_Or",
	)
}

func (this *WorkspaceFoldersServerCapabilities_ChangeNotifications_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceFoldersServerCapabilities_ChangeNotifications_Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_Location_Or struct {
	// Or [ *Location *WorkspaceSymbol_Location_Or_1 ]
	V interface{}
}

func (this *WorkspaceSymbol_Location_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *WorkspaceSymbol_Location_Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceSymbol_Location_Or")
}

func (this *WorkspaceSymbol_Location_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceSymbol_Location_Or_1); ok {
			break
		}

		return nil, OrValidateFailed("WorkspaceSymbol_Location_Or")
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_PartialResult_Or struct {
	// Or [ []SymbolInformation []WorkspaceSymbol ]
	V interface{}
}

func (this *WorkspaceSymbol_PartialResult_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp []SymbolInformation
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []WorkspaceSymbol
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceSymbol_PartialResult_Or")
}

func (this *WorkspaceSymbol_PartialResult_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]WorkspaceSymbol); ok {
			break
		}

		return nil, OrValidateFailed("WorkspaceSymbol_PartialResult_Or")
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_Result_Or struct {
	// Or [ []SymbolInformation []WorkspaceSymbol *Null ]
	V interface{}
}

func (this *WorkspaceSymbol_Result_Or) UnmarshalJSON(data []byte) error {

	{

		var tmp []SymbolInformation
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp []WorkspaceSymbol
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *Null
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceSymbol_Result_Or")
}

func (this *WorkspaceSymbol_Result_Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]WorkspaceSymbol); ok {
			break
		}

		if _, ok := this.V.(*Null); ok {
			break
		}

		return nil, OrValidateFailed("WorkspaceSymbol_Result_Or")
	}
	return json.Marshal(this.V)
}
