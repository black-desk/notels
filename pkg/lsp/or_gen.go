// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package protocol

import (
	"encoding/json"
	"fmt"
)

var OrValidateFailed = func(name string) error {
	return fmt.Errorf("or type \"%s\"validate failed", name)
}

type CancelParams_Id__Or struct {

	// *Integer
	// *String
	V interface{}
}

func (this *CancelParams_Id__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("CancelParams_Id__Or")
}

func (this *CancelParams_Id__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("CancelParams_Id__Or")
	}
	return json.Marshal(this.V)
}

type CompletionItem_Documentation__Or struct {

	// *String
	// *MarkupContent
	V interface{}
}

func (this *CompletionItem_Documentation__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("CompletionItem_Documentation__Or")
}

func (this *CompletionItem_Documentation__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("CompletionItem_Documentation__Or")
	}
	return json.Marshal(this.V)
}

type CompletionItem_TextEdit__Or struct {

	// *TextEdit
	// *InsertReplaceEdit
	V interface{}
}

func (this *CompletionItem_TextEdit__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("CompletionItem_TextEdit__Or")
}

func (this *CompletionItem_TextEdit__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextEdit); ok {
			break
		}

		if _, ok := this.V.(*InsertReplaceEdit); ok {
			break
		}

		return nil, OrValidateFailed("CompletionItem_TextEdit__Or")
	}
	return json.Marshal(this.V)
}

type CompletionList_ItemDefaults_EditRange__Or struct {
	V interface{}
}

func (this *CompletionList_ItemDefaults_EditRange__Or) UnmarshalJSON(
	data []byte,
) error {

	return OrValidateFailed("CompletionList_ItemDefaults_EditRange__Or")
}

func (this *CompletionList_ItemDefaults_EditRange__Or) MarshalJSON() ([]byte, error) {
	for {

		return nil, OrValidateFailed(
			"CompletionList_ItemDefaults_EditRange__Or",
		)
	}
	return json.Marshal(this.V)
}

type Declaration__Or struct {

	// *Location
	// []Location
	V interface{}
}

func (this *Declaration__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("Declaration__Or")
}

func (this *Declaration__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.([]Location); ok {
			break
		}

		return nil, OrValidateFailed("Declaration__Or")
	}
	return json.Marshal(this.V)
}

type Definition__Or struct {

	// *Location
	// []Location
	V interface{}
}

func (this *Definition__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("Definition__Or")
}

func (this *Definition__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.([]Location); ok {
			break
		}

		return nil, OrValidateFailed("Definition__Or")
	}
	return json.Marshal(this.V)
}

type Diagnostic_Code__Or struct {

	// *Integer
	// *String
	V interface{}
}

func (this *Diagnostic_Code__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("Diagnostic_Code__Or")
}

func (this *Diagnostic_Code__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("Diagnostic_Code__Or")
	}
	return json.Marshal(this.V)
}

type DidChangeConfigurationRegistrationOptions_Section__Or struct {

	// *String
	// []String
	V interface{}
}

func (this *DidChangeConfigurationRegistrationOptions_Section__Or) UnmarshalJSON(
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
		"DidChangeConfigurationRegistrationOptions_Section__Or",
	)
}

func (this *DidChangeConfigurationRegistrationOptions_Section__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.([]String); ok {
			break
		}

		return nil, OrValidateFailed(
			"DidChangeConfigurationRegistrationOptions_Section__Or",
		)
	}
	return json.Marshal(this.V)
}

type DocumentDiagnosticReportPartialResult_RelatedDocuments_Value__Or struct {

	// *FullDocumentDiagnosticReport
	// *UnchangedDocumentDiagnosticReport
	V interface{}
}

func (this *DocumentDiagnosticReportPartialResult_RelatedDocuments_Value__Or) UnmarshalJSON(
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
		"DocumentDiagnosticReportPartialResult_RelatedDocuments_Value__Or",
	)
}

func (this *DocumentDiagnosticReportPartialResult_RelatedDocuments_Value__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"DocumentDiagnosticReportPartialResult_RelatedDocuments_Value__Or",
		)
	}
	return json.Marshal(this.V)
}

type DocumentDiagnosticReport__Or struct {

	// *RelatedFullDocumentDiagnosticReport
	// *RelatedUnchangedDocumentDiagnosticReport
	V interface{}
}

func (this *DocumentDiagnosticReport__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("DocumentDiagnosticReport__Or")
}

func (this *DocumentDiagnosticReport__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*RelatedFullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*RelatedUnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed("DocumentDiagnosticReport__Or")
	}
	return json.Marshal(this.V)
}

type DocumentFilter__Or struct {

	// *TextDocumentFilter
	// *NotebookCellTextDocumentFilter
	V interface{}
}

func (this *DocumentFilter__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("DocumentFilter__Or")
}

func (this *DocumentFilter__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentFilter); ok {
			break
		}

		if _, ok := this.V.(*NotebookCellTextDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed("DocumentFilter__Or")
	}
	return json.Marshal(this.V)
}

type GlobPattern__Or struct {

	// *Pattern
	// *RelativePattern
	V interface{}
}

func (this *GlobPattern__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("GlobPattern__Or")
}

func (this *GlobPattern__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Pattern); ok {
			break
		}

		if _, ok := this.V.(*RelativePattern); ok {
			break
		}

		return nil, OrValidateFailed("GlobPattern__Or")
	}
	return json.Marshal(this.V)
}

type Hover_Contents__Or struct {

	// *MarkupContent
	// *MarkedString
	// []MarkedString
	V interface{}
}

func (this *Hover_Contents__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("Hover_Contents__Or")
}

func (this *Hover_Contents__Or) MarshalJSON() ([]byte, error) {
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

		return nil, OrValidateFailed("Hover_Contents__Or")
	}
	return json.Marshal(this.V)
}

type InlayHintLabelPart_Tooltip__Or struct {

	// *String
	// *MarkupContent
	V interface{}
}

func (this *InlayHintLabelPart_Tooltip__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("InlayHintLabelPart_Tooltip__Or")
}

func (this *InlayHintLabelPart_Tooltip__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("InlayHintLabelPart_Tooltip__Or")
	}
	return json.Marshal(this.V)
}

type InlayHint_Label__Or struct {

	// *String
	// []InlayHintLabelPart
	V interface{}
}

func (this *InlayHint_Label__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("InlayHint_Label__Or")
}

func (this *InlayHint_Label__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.([]InlayHintLabelPart); ok {
			break
		}

		return nil, OrValidateFailed("InlayHint_Label__Or")
	}
	return json.Marshal(this.V)
}

type InlayHint_Tooltip__Or struct {

	// *String
	// *MarkupContent
	V interface{}
}

func (this *InlayHint_Tooltip__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("InlayHint_Tooltip__Or")
}

func (this *InlayHint_Tooltip__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed("InlayHint_Tooltip__Or")
	}
	return json.Marshal(this.V)
}

type InlineValue__Or struct {

	// *InlineValueText
	// *InlineValueVariableLookup
	// *InlineValueEvaluatableExpression
	V interface{}
}

func (this *InlineValue__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("InlineValue__Or")
}

func (this *InlineValue__Or) MarshalJSON() ([]byte, error) {
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

		return nil, OrValidateFailed("InlineValue__Or")
	}
	return json.Marshal(this.V)
}

type LSPAny__Or struct {

	// *LSPObject
	// *LSPArray
	// *String
	// *Integer
	// *Uinteger
	// *Decimal
	// *Boolean
	// *Null
	V interface{}
}

func (this *LSPAny__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("LSPAny__Or")
}

func (this *LSPAny__Or) MarshalJSON() ([]byte, error) {
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

		return nil, OrValidateFailed("LSPAny__Or")
	}
	return json.Marshal(this.V)
}

type MarkedString__Or struct {

	// *String
	// *MarkedString__Or_1
	V interface{}
}

func (this *MarkedString__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *MarkedString__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("MarkedString__Or")
}

func (this *MarkedString__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkedString__Or_1); ok {
			break
		}

		return nil, OrValidateFailed("MarkedString__Or")
	}
	return json.Marshal(this.V)
}

type NotebookCellTextDocumentFilter_Notebook__Or struct {

	// *String
	// *NotebookDocumentFilter
	V interface{}
}

func (this *NotebookCellTextDocumentFilter_Notebook__Or) UnmarshalJSON(
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

	return OrValidateFailed("NotebookCellTextDocumentFilter_Notebook__Or")
}

func (this *NotebookCellTextDocumentFilter_Notebook__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookCellTextDocumentFilter_Notebook__Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentFilter__Or struct {

	// *NotebookDocumentFilter__Or_0
	// *NotebookDocumentFilter__Or_1
	// *NotebookDocumentFilter__Or_2
	V interface{}
}

func (this *NotebookDocumentFilter__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *NotebookDocumentFilter__Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentFilter__Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("NotebookDocumentFilter__Or")
}

func (this *NotebookDocumentFilter__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentFilter__Or_0); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter__Or_1); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentFilter__Or_2); ok {
			break
		}

		return nil, OrValidateFailed("NotebookDocumentFilter__Or")
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element__Or struct {

	// *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0
	// *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element__Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element__Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or struct {
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or) UnmarshalJSON(
	data []byte,
) error {

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or) MarshalJSON() ([]byte, error) {
	for {

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element__Or_0_Notebook__Or",
		)
	}
	return json.Marshal(this.V)
}

type NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or struct {
	V interface{}
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or) UnmarshalJSON(
	data []byte,
) error {

	return OrValidateFailed(
		"NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or",
	)
}

func (this *NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or) MarshalJSON() ([]byte, error) {
	for {

		return nil, OrValidateFailed(
			"NotebookDocumentSyncOptions_NotebookSelector_Element__Or_1_Notebook__Or",
		)
	}
	return json.Marshal(this.V)
}

type ParameterInformation_Documentation__Or struct {

	// *String
	// *MarkupContent
	V interface{}
}

func (this *ParameterInformation_Documentation__Or) UnmarshalJSON(
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

	return OrValidateFailed("ParameterInformation_Documentation__Or")
}

func (this *ParameterInformation_Documentation__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed(
			"ParameterInformation_Documentation__Or",
		)
	}
	return json.Marshal(this.V)
}

type ParameterInformation_Label__Or struct {

	// *String
	// *ParameterInformation_Label__Or_1_Tuple
	V interface{}
}

func (this *ParameterInformation_Label__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *String
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *ParameterInformation_Label__Or_1_Tuple
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("ParameterInformation_Label__Or")
}

func (this *ParameterInformation_Label__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*ParameterInformation_Label__Or_1_Tuple); ok {
			break
		}

		return nil, OrValidateFailed("ParameterInformation_Label__Or")
	}
	return json.Marshal(this.V)
}

type PrepareRenameResult__Or struct {

	// *Range
	// *PrepareRenameResult__Or_1
	// *PrepareRenameResult__Or_2
	V interface{}
}

func (this *PrepareRenameResult__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Range
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *PrepareRenameResult__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *PrepareRenameResult__Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("PrepareRenameResult__Or")
}

func (this *PrepareRenameResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Range); ok {
			break
		}

		if _, ok := this.V.(*PrepareRenameResult__Or_1); ok {
			break
		}

		if _, ok := this.V.(*PrepareRenameResult__Or_2); ok {
			break
		}

		return nil, OrValidateFailed("PrepareRenameResult__Or")
	}
	return json.Marshal(this.V)
}

type ProgressToken__Or struct {

	// *Integer
	// *String
	V interface{}
}

func (this *ProgressToken__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("ProgressToken__Or")
}

func (this *ProgressToken__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Integer); ok {
			break
		}

		if _, ok := this.V.(*String); ok {
			break
		}

		return nil, OrValidateFailed("ProgressToken__Or")
	}
	return json.Marshal(this.V)
}

type RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value__Or struct {

	// *FullDocumentDiagnosticReport
	// *UnchangedDocumentDiagnosticReport
	V interface{}
}

func (this *RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value__Or) UnmarshalJSON(
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
		"RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value__Or",
	)
}

func (this *RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"RelatedFullDocumentDiagnosticReport_RelatedDocuments_Value__Or",
		)
	}
	return json.Marshal(this.V)
}

type RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value__Or struct {

	// *FullDocumentDiagnosticReport
	// *UnchangedDocumentDiagnosticReport
	V interface{}
}

func (this *RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value__Or) UnmarshalJSON(
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
		"RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value__Or",
	)
}

func (this *RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*FullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*UnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments_Value__Or",
		)
	}
	return json.Marshal(this.V)
}

type RelativePattern_BaseUri__Or struct {

	// *WorkspaceFolder
	// *URI
	V interface{}
}

func (this *RelativePattern_BaseUri__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("RelativePattern_BaseUri__Or")
}

func (this *RelativePattern_BaseUri__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*WorkspaceFolder); ok {
			break
		}

		if _, ok := this.V.(*URI); ok {
			break
		}

		return nil, OrValidateFailed("RelativePattern_BaseUri__Or")
	}
	return json.Marshal(this.V)
}

type SemanticTokensClientCapabilities_Requests_Full__Or struct {

	// *Boolean
	// *SemanticTokensClientCapabilities_Requests_Full__Or_1
	V interface{}
}

func (this *SemanticTokensClientCapabilities_Requests_Full__Or) UnmarshalJSON(
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

		var tmp *SemanticTokensClientCapabilities_Requests_Full__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"SemanticTokensClientCapabilities_Requests_Full__Or",
	)
}

func (this *SemanticTokensClientCapabilities_Requests_Full__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensClientCapabilities_Requests_Full__Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"SemanticTokensClientCapabilities_Requests_Full__Or",
		)
	}
	return json.Marshal(this.V)
}

type SemanticTokensClientCapabilities_Requests_Range__Or struct {

	// *Boolean
	// *SemanticTokensClientCapabilities_Requests_Range__Or_1
	V interface{}
}

func (this *SemanticTokensClientCapabilities_Requests_Range__Or) UnmarshalJSON(
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

		var tmp *SemanticTokensClientCapabilities_Requests_Range__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed(
		"SemanticTokensClientCapabilities_Requests_Range__Or",
	)
}

func (this *SemanticTokensClientCapabilities_Requests_Range__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensClientCapabilities_Requests_Range__Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"SemanticTokensClientCapabilities_Requests_Range__Or",
		)
	}
	return json.Marshal(this.V)
}

type SemanticTokensOptions_Full__Or struct {

	// *Boolean
	// *SemanticTokensOptions_Full__Or_1
	V interface{}
}

func (this *SemanticTokensOptions_Full__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensOptions_Full__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("SemanticTokensOptions_Full__Or")
}

func (this *SemanticTokensOptions_Full__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensOptions_Full__Or_1); ok {
			break
		}

		return nil, OrValidateFailed("SemanticTokensOptions_Full__Or")
	}
	return json.Marshal(this.V)
}

type SemanticTokensOptions_Range__Or struct {

	// *Boolean
	// *SemanticTokensOptions_Range__Or_1
	V interface{}
}

func (this *SemanticTokensOptions_Range__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Boolean
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *SemanticTokensOptions_Range__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("SemanticTokensOptions_Range__Or")
}

func (this *SemanticTokensOptions_Range__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensOptions_Range__Or_1); ok {
			break
		}

		return nil, OrValidateFailed("SemanticTokensOptions_Range__Or")
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_CallHierarchyProvider__Or struct {

	// *Boolean
	// *CallHierarchyOptions
	// *CallHierarchyRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_CallHierarchyProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_CallHierarchyProvider__Or")
}

func (this *ServerCapabilities_CallHierarchyProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_CallHierarchyProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_CodeActionProvider__Or struct {

	// *Boolean
	// *CodeActionOptions
	V interface{}
}

func (this *ServerCapabilities_CodeActionProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_CodeActionProvider__Or")
}

func (this *ServerCapabilities_CodeActionProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*CodeActionOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_CodeActionProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ColorProvider__Or struct {

	// *Boolean
	// *DocumentColorOptions
	// *DocumentColorRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_ColorProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_ColorProvider__Or")
}

func (this *ServerCapabilities_ColorProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_ColorProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DeclarationProvider__Or struct {

	// *Boolean
	// *DeclarationOptions
	// *DeclarationRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_DeclarationProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_DeclarationProvider__Or")
}

func (this *ServerCapabilities_DeclarationProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_DeclarationProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DefinitionProvider__Or struct {

	// *Boolean
	// *DefinitionOptions
	V interface{}
}

func (this *ServerCapabilities_DefinitionProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_DefinitionProvider__Or")
}

func (this *ServerCapabilities_DefinitionProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DefinitionOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DefinitionProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DiagnosticProvider__Or struct {

	// *DiagnosticOptions
	// *DiagnosticRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_DiagnosticProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_DiagnosticProvider__Or")
}

func (this *ServerCapabilities_DiagnosticProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*DiagnosticOptions); ok {
			break
		}

		if _, ok := this.V.(*DiagnosticRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DiagnosticProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentFormattingProvider__Or struct {

	// *Boolean
	// *DocumentFormattingOptions
	V interface{}
}

func (this *ServerCapabilities_DocumentFormattingProvider__Or) UnmarshalJSON(
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
		"ServerCapabilities_DocumentFormattingProvider__Or",
	)
}

func (this *ServerCapabilities_DocumentFormattingProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentFormattingOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentFormattingProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentHighlightProvider__Or struct {

	// *Boolean
	// *DocumentHighlightOptions
	V interface{}
}

func (this *ServerCapabilities_DocumentHighlightProvider__Or) UnmarshalJSON(
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
		"ServerCapabilities_DocumentHighlightProvider__Or",
	)
}

func (this *ServerCapabilities_DocumentHighlightProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentHighlightOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentHighlightProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentRangeFormattingProvider__Or struct {

	// *Boolean
	// *DocumentRangeFormattingOptions
	V interface{}
}

func (this *ServerCapabilities_DocumentRangeFormattingProvider__Or) UnmarshalJSON(
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
		"ServerCapabilities_DocumentRangeFormattingProvider__Or",
	)
}

func (this *ServerCapabilities_DocumentRangeFormattingProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentRangeFormattingOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentRangeFormattingProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_DocumentSymbolProvider__Or struct {

	// *Boolean
	// *DocumentSymbolOptions
	V interface{}
}

func (this *ServerCapabilities_DocumentSymbolProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_DocumentSymbolProvider__Or")
}

func (this *ServerCapabilities_DocumentSymbolProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*DocumentSymbolOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_DocumentSymbolProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_FoldingRangeProvider__Or struct {

	// *Boolean
	// *FoldingRangeOptions
	// *FoldingRangeRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_FoldingRangeProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_FoldingRangeProvider__Or")
}

func (this *ServerCapabilities_FoldingRangeProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_FoldingRangeProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_HoverProvider__Or struct {

	// *Boolean
	// *HoverOptions
	V interface{}
}

func (this *ServerCapabilities_HoverProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_HoverProvider__Or")
}

func (this *ServerCapabilities_HoverProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*HoverOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_HoverProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ImplementationProvider__Or struct {

	// *Boolean
	// *ImplementationOptions
	// *ImplementationRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_ImplementationProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_ImplementationProvider__Or")
}

func (this *ServerCapabilities_ImplementationProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_ImplementationProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_InlayHintProvider__Or struct {

	// *Boolean
	// *InlayHintOptions
	// *InlayHintRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_InlayHintProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_InlayHintProvider__Or")
}

func (this *ServerCapabilities_InlayHintProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_InlayHintProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_InlineValueProvider__Or struct {

	// *Boolean
	// *InlineValueOptions
	// *InlineValueRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_InlineValueProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_InlineValueProvider__Or")
}

func (this *ServerCapabilities_InlineValueProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_InlineValueProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_LinkedEditingRangeProvider__Or struct {

	// *Boolean
	// *LinkedEditingRangeOptions
	// *LinkedEditingRangeRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_LinkedEditingRangeProvider__Or) UnmarshalJSON(
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
		"ServerCapabilities_LinkedEditingRangeProvider__Or",
	)
}

func (this *ServerCapabilities_LinkedEditingRangeProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_LinkedEditingRangeProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_MonikerProvider__Or struct {

	// *Boolean
	// *MonikerOptions
	// *MonikerRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_MonikerProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_MonikerProvider__Or")
}

func (this *ServerCapabilities_MonikerProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_MonikerProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_NotebookDocumentSync__Or struct {

	// *NotebookDocumentSyncOptions
	// *NotebookDocumentSyncRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_NotebookDocumentSync__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_NotebookDocumentSync__Or")
}

func (this *ServerCapabilities_NotebookDocumentSync__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*NotebookDocumentSyncOptions); ok {
			break
		}

		if _, ok := this.V.(*NotebookDocumentSyncRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_NotebookDocumentSync__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_ReferencesProvider__Or struct {

	// *Boolean
	// *ReferenceOptions
	V interface{}
}

func (this *ServerCapabilities_ReferencesProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_ReferencesProvider__Or")
}

func (this *ServerCapabilities_ReferencesProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*ReferenceOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_ReferencesProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_RenameProvider__Or struct {

	// *Boolean
	// *RenameOptions
	V interface{}
}

func (this *ServerCapabilities_RenameProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_RenameProvider__Or")
}

func (this *ServerCapabilities_RenameProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*RenameOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_RenameProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_SelectionRangeProvider__Or struct {

	// *Boolean
	// *SelectionRangeOptions
	// *SelectionRangeRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_SelectionRangeProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_SelectionRangeProvider__Or")
}

func (this *ServerCapabilities_SelectionRangeProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_SelectionRangeProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_SemanticTokensProvider__Or struct {

	// *SemanticTokensOptions
	// *SemanticTokensRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_SemanticTokensProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_SemanticTokensProvider__Or")
}

func (this *ServerCapabilities_SemanticTokensProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*SemanticTokensOptions); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensRegistrationOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_SemanticTokensProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TextDocumentSync__Or struct {

	// *TextDocumentSyncOptions
	// *TextDocumentSyncKind
	V interface{}
}

func (this *ServerCapabilities_TextDocumentSync__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_TextDocumentSync__Or")
}

func (this *ServerCapabilities_TextDocumentSync__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentSyncOptions); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentSyncKind); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_TextDocumentSync__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TypeDefinitionProvider__Or struct {

	// *Boolean
	// *TypeDefinitionOptions
	// *TypeDefinitionRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_TypeDefinitionProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_TypeDefinitionProvider__Or")
}

func (this *ServerCapabilities_TypeDefinitionProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_TypeDefinitionProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_TypeHierarchyProvider__Or struct {

	// *Boolean
	// *TypeHierarchyOptions
	// *TypeHierarchyRegistrationOptions
	V interface{}
}

func (this *ServerCapabilities_TypeHierarchyProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed("ServerCapabilities_TypeHierarchyProvider__Or")
}

func (this *ServerCapabilities_TypeHierarchyProvider__Or) MarshalJSON() ([]byte, error) {
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
			"ServerCapabilities_TypeHierarchyProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type ServerCapabilities_WorkspaceSymbolProvider__Or struct {

	// *Boolean
	// *WorkspaceSymbolOptions
	V interface{}
}

func (this *ServerCapabilities_WorkspaceSymbolProvider__Or) UnmarshalJSON(
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

	return OrValidateFailed(
		"ServerCapabilities_WorkspaceSymbolProvider__Or",
	)
}

func (this *ServerCapabilities_WorkspaceSymbolProvider__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceSymbolOptions); ok {
			break
		}

		return nil, OrValidateFailed(
			"ServerCapabilities_WorkspaceSymbolProvider__Or",
		)
	}
	return json.Marshal(this.V)
}

type SignatureInformation_Documentation__Or struct {

	// *String
	// *MarkupContent
	V interface{}
}

func (this *SignatureInformation_Documentation__Or) UnmarshalJSON(
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

	return OrValidateFailed("SignatureInformation_Documentation__Or")
}

func (this *SignatureInformation_Documentation__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*MarkupContent); ok {
			break
		}

		return nil, OrValidateFailed(
			"SignatureInformation_Documentation__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCodeAction_PartialResult_Element__Or struct {

	// *Command
	// *CodeAction
	V interface{}
}

func (this *TextDocumentCodeAction_PartialResult_Element__Or) UnmarshalJSON(
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
		"TextDocumentCodeAction_PartialResult_Element__Or",
	)
}

func (this *TextDocumentCodeAction_PartialResult_Element__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Command); ok {
			break
		}

		if _, ok := this.V.(*CodeAction); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentCodeAction_PartialResult_Element__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCodeAction_Result_Element__Or struct {

	// *Command
	// *CodeAction
	V interface{}
}

func (this *TextDocumentCodeAction_Result_Element__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentCodeAction_Result_Element__Or")
}

func (this *TextDocumentCodeAction_Result_Element__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Command); ok {
			break
		}

		if _, ok := this.V.(*CodeAction); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentCodeAction_Result_Element__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentCompletion_Result__Or struct {

	// []CompletionItem
	// *CompletionList
	// *Null
	V interface{}
}

func (this *TextDocumentCompletion_Result__Or) UnmarshalJSON(
	data []byte,
) error {

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

	return OrValidateFailed("TextDocumentCompletion_Result__Or")
}

func (this *TextDocumentCompletion_Result__Or) MarshalJSON() ([]byte, error) {
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

		return nil, OrValidateFailed(
			"TextDocumentCompletion_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentContentChangeEvent__Or struct {

	// *TextDocumentContentChangeEvent__Or_0
	// *TextDocumentContentChangeEvent__Or_1
	V interface{}
}

func (this *TextDocumentContentChangeEvent__Or) UnmarshalJSON(
	data []byte,
) error {

	{

		var tmp *TextDocumentContentChangeEvent__Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentContentChangeEvent__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentContentChangeEvent__Or")
}

func (this *TextDocumentContentChangeEvent__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentContentChangeEvent__Or_0); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentContentChangeEvent__Or_1); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentContentChangeEvent__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDeclaration_PartialResult__Or struct {

	// []Location
	// []DeclarationLink
	V interface{}
}

func (this *TextDocumentDeclaration_PartialResult__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDeclaration_PartialResult__Or")
}

func (this *TextDocumentDeclaration_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DeclarationLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDeclaration_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDeclaration_Result__Or struct {

	// *Declaration
	// []DeclarationLink
	// *Null
	V interface{}
}

func (this *TextDocumentDeclaration_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDeclaration_Result__Or")
}

func (this *TextDocumentDeclaration_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentDeclaration_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDefinition_PartialResult__Or struct {

	// []Location
	// []DefinitionLink
	V interface{}
}

func (this *TextDocumentDefinition_PartialResult__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDefinition_PartialResult__Or")
}

func (this *TextDocumentDefinition_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDefinition_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDefinition_Result__Or struct {

	// *Definition
	// []DefinitionLink
	// *Null
	V interface{}
}

func (this *TextDocumentDefinition_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDefinition_Result__Or")
}

func (this *TextDocumentDefinition_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentDefinition_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDocumentSymbol_PartialResult__Or struct {

	// []SymbolInformation
	// []DocumentSymbol
	V interface{}
}

func (this *TextDocumentDocumentSymbol_PartialResult__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDocumentSymbol_PartialResult__Or")
}

func (this *TextDocumentDocumentSymbol_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]DocumentSymbol); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentDocumentSymbol_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentDocumentSymbol_Result__Or struct {

	// []SymbolInformation
	// []DocumentSymbol
	// *Null
	V interface{}
}

func (this *TextDocumentDocumentSymbol_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentDocumentSymbol_Result__Or")
}

func (this *TextDocumentDocumentSymbol_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentDocumentSymbol_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentEdit_Edits_Element__Or struct {

	// *TextEdit
	// *AnnotatedTextEdit
	V interface{}
}

func (this *TextDocumentEdit_Edits_Element__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentEdit_Edits_Element__Or")
}

func (this *TextDocumentEdit_Edits_Element__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextEdit); ok {
			break
		}

		if _, ok := this.V.(*AnnotatedTextEdit); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentEdit_Edits_Element__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentFilter__Or struct {

	// *TextDocumentFilter__Or_0
	// *TextDocumentFilter__Or_1
	// *TextDocumentFilter__Or_2
	V interface{}
}

func (this *TextDocumentFilter__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *TextDocumentFilter__Or_0
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentFilter__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *TextDocumentFilter__Or_2
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("TextDocumentFilter__Or")
}

func (this *TextDocumentFilter__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*TextDocumentFilter__Or_0); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentFilter__Or_1); ok {
			break
		}

		if _, ok := this.V.(*TextDocumentFilter__Or_2); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentFilter__Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentImplementation_PartialResult__Or struct {

	// []Location
	// []DefinitionLink
	V interface{}
}

func (this *TextDocumentImplementation_PartialResult__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentImplementation_PartialResult__Or")
}

func (this *TextDocumentImplementation_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentImplementation_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentImplementation_Result__Or struct {

	// *Definition
	// []DefinitionLink
	// *Null
	V interface{}
}

func (this *TextDocumentImplementation_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentImplementation_Result__Or")
}

func (this *TextDocumentImplementation_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentImplementation_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSemanticTokensFullDelta_PartialResult__Or struct {

	// *SemanticTokensPartialResult
	// *SemanticTokensDeltaPartialResult
	V interface{}
}

func (this *TextDocumentSemanticTokensFullDelta_PartialResult__Or) UnmarshalJSON(
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
		"TextDocumentSemanticTokensFullDelta_PartialResult__Or",
	)
}

func (this *TextDocumentSemanticTokensFullDelta_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*SemanticTokensPartialResult); ok {
			break
		}

		if _, ok := this.V.(*SemanticTokensDeltaPartialResult); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentSemanticTokensFullDelta_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSemanticTokensFullDelta_Result__Or struct {

	// *SemanticTokens
	// *SemanticTokensDelta
	// *Null
	V interface{}
}

func (this *TextDocumentSemanticTokensFullDelta_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed(
		"TextDocumentSemanticTokensFullDelta_Result__Or",
	)
}

func (this *TextDocumentSemanticTokensFullDelta_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentSemanticTokensFullDelta_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentSyncOptions_Save__Or struct {

	// *Boolean
	// *SaveOptions
	V interface{}
}

func (this *TextDocumentSyncOptions_Save__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("TextDocumentSyncOptions_Save__Or")
}

func (this *TextDocumentSyncOptions_Save__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		if _, ok := this.V.(*SaveOptions); ok {
			break
		}

		return nil, OrValidateFailed("TextDocumentSyncOptions_Save__Or")
	}
	return json.Marshal(this.V)
}

type TextDocumentTypeDefinition_PartialResult__Or struct {

	// []Location
	// []DefinitionLink
	V interface{}
}

func (this *TextDocumentTypeDefinition_PartialResult__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentTypeDefinition_PartialResult__Or")
}

func (this *TextDocumentTypeDefinition_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]Location); ok {
			break
		}

		if _, ok := this.V.([]DefinitionLink); ok {
			break
		}

		return nil, OrValidateFailed(
			"TextDocumentTypeDefinition_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type TextDocumentTypeDefinition_Result__Or struct {

	// *Definition
	// []DefinitionLink
	// *Null
	V interface{}
}

func (this *TextDocumentTypeDefinition_Result__Or) UnmarshalJSON(
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

	return OrValidateFailed("TextDocumentTypeDefinition_Result__Or")
}

func (this *TextDocumentTypeDefinition_Result__Or) MarshalJSON() ([]byte, error) {
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
			"TextDocumentTypeDefinition_Result__Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceDocumentDiagnosticReport__Or struct {

	// *WorkspaceFullDocumentDiagnosticReport
	// *WorkspaceUnchangedDocumentDiagnosticReport
	V interface{}
}

func (this *WorkspaceDocumentDiagnosticReport__Or) UnmarshalJSON(
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

	return OrValidateFailed("WorkspaceDocumentDiagnosticReport__Or")
}

func (this *WorkspaceDocumentDiagnosticReport__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*WorkspaceFullDocumentDiagnosticReport); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceUnchangedDocumentDiagnosticReport); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceDocumentDiagnosticReport__Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceEdit_DocumentChanges_Element__Or struct {

	// *TextDocumentEdit
	// *CreateFile
	// *RenameFile
	// *DeleteFile
	V interface{}
}

func (this *WorkspaceEdit_DocumentChanges_Element__Or) UnmarshalJSON(
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

	return OrValidateFailed("WorkspaceEdit_DocumentChanges_Element__Or")
}

func (this *WorkspaceEdit_DocumentChanges_Element__Or) MarshalJSON() ([]byte, error) {
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
			"WorkspaceEdit_DocumentChanges_Element__Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceFoldersServerCapabilities_ChangeNotifications__Or struct {

	// *String
	// *Boolean
	V interface{}
}

func (this *WorkspaceFoldersServerCapabilities_ChangeNotifications__Or) UnmarshalJSON(
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
		"WorkspaceFoldersServerCapabilities_ChangeNotifications__Or",
	)
}

func (this *WorkspaceFoldersServerCapabilities_ChangeNotifications__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*String); ok {
			break
		}

		if _, ok := this.V.(*Boolean); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceFoldersServerCapabilities_ChangeNotifications__Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_Location__Or struct {

	// *Location
	// *WorkspaceSymbol_Location__Or_1
	V interface{}
}

func (this *WorkspaceSymbol_Location__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *Location
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *WorkspaceSymbol_Location__Or_1
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceSymbol_Location__Or")
}

func (this *WorkspaceSymbol_Location__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*Location); ok {
			break
		}

		if _, ok := this.V.(*WorkspaceSymbol_Location__Or_1); ok {
			break
		}

		return nil, OrValidateFailed("WorkspaceSymbol_Location__Or")
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_PartialResult__Or struct {

	// []SymbolInformation
	// []WorkspaceSymbol
	V interface{}
}

func (this *WorkspaceSymbol_PartialResult__Or) UnmarshalJSON(
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

		var tmp []WorkspaceSymbol
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("WorkspaceSymbol_PartialResult__Or")
}

func (this *WorkspaceSymbol_PartialResult__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.([]SymbolInformation); ok {
			break
		}

		if _, ok := this.V.([]WorkspaceSymbol); ok {
			break
		}

		return nil, OrValidateFailed(
			"WorkspaceSymbol_PartialResult__Or",
		)
	}
	return json.Marshal(this.V)
}

type WorkspaceSymbol_Result__Or struct {

	// []SymbolInformation
	// []WorkspaceSymbol
	// *Null
	V interface{}
}

func (this *WorkspaceSymbol_Result__Or) UnmarshalJSON(data []byte) error {

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

	return OrValidateFailed("WorkspaceSymbol_Result__Or")
}

func (this *WorkspaceSymbol_Result__Or) MarshalJSON() ([]byte, error) {
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

		return nil, OrValidateFailed("WorkspaceSymbol_Result__Or")
	}
	return json.Marshal(this.V)
}

type XInitializeParams_Trace__Or struct {

	// *string
	// *string
	// *string
	// *string
	V interface{}
}

func (this *XInitializeParams_Trace__Or) UnmarshalJSON(data []byte) error {

	{

		var tmp *string
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *string
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *string
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	{

		var tmp *string
		if err := json.Unmarshal(data, &tmp); err == nil {
			this.V = tmp
			return nil
		}
	}

	return OrValidateFailed("XInitializeParams_Trace__Or")
}

func (this *XInitializeParams_Trace__Or) MarshalJSON() ([]byte, error) {
	for {

		if _, ok := this.V.(*string); ok {
			break
		}

		if _, ok := this.V.(*string); ok {
			break
		}

		if _, ok := this.V.(*string); ok {
			break
		}

		if _, ok := this.V.(*string); ok {
			break
		}

		return nil, OrValidateFailed("XInitializeParams_Trace__Or")
	}
	return json.Marshal(this.V)
}
