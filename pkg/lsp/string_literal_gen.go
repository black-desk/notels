// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
	"encoding/json"
	"fmt"
)

var StringLiteralValidateFailed = func(name string) error {
	return fmt.Errorf("literal structure \"%s\"validate failed", name)
}

type CreateFile_Kind_StringLiteral struct {
	V String // create
}

func (this *CreateFile_Kind_StringLiteral) UnmarshalJSON(data []byte) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "create" {
		return StringLiteralValidateFailed(
			"CreateFile_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *CreateFile_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "create" {
		return nil, StringLiteralValidateFailed(
			"CreateFile_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type DeleteFile_Kind_StringLiteral struct {
	V String // delete
}

func (this *DeleteFile_Kind_StringLiteral) UnmarshalJSON(data []byte) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "delete" {
		return StringLiteralValidateFailed(
			"DeleteFile_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *DeleteFile_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "delete" {
		return nil, StringLiteralValidateFailed(
			"DeleteFile_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type FullDocumentDiagnosticReport_Kind_StringLiteral struct {
	V String // full
}

func (this *FullDocumentDiagnosticReport_Kind_StringLiteral) UnmarshalJSON(
	data []byte,
) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "full" {
		return StringLiteralValidateFailed(
			"FullDocumentDiagnosticReport_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *FullDocumentDiagnosticReport_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "full" {
		return nil, StringLiteralValidateFailed(
			"FullDocumentDiagnosticReport_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type RenameFile_Kind_StringLiteral struct {
	V String // rename
}

func (this *RenameFile_Kind_StringLiteral) UnmarshalJSON(data []byte) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "rename" {
		return StringLiteralValidateFailed(
			"RenameFile_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *RenameFile_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "rename" {
		return nil, StringLiteralValidateFailed(
			"RenameFile_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type UnchangedDocumentDiagnosticReport_Kind_StringLiteral struct {
	V String // unchanged
}

func (this *UnchangedDocumentDiagnosticReport_Kind_StringLiteral) UnmarshalJSON(
	data []byte,
) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "unchanged" {
		return StringLiteralValidateFailed(
			"UnchangedDocumentDiagnosticReport_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *UnchangedDocumentDiagnosticReport_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "unchanged" {
		return nil, StringLiteralValidateFailed(
			"UnchangedDocumentDiagnosticReport_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type WorkDoneProgressBegin_Kind_StringLiteral struct {
	V String // begin
}

func (this *WorkDoneProgressBegin_Kind_StringLiteral) UnmarshalJSON(
	data []byte,
) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "begin" {
		return StringLiteralValidateFailed(
			"WorkDoneProgressBegin_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *WorkDoneProgressBegin_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "begin" {
		return nil, StringLiteralValidateFailed(
			"WorkDoneProgressBegin_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type WorkDoneProgressEnd_Kind_StringLiteral struct {
	V String // end
}

func (this *WorkDoneProgressEnd_Kind_StringLiteral) UnmarshalJSON(
	data []byte,
) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "end" {
		return StringLiteralValidateFailed(
			"WorkDoneProgressEnd_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *WorkDoneProgressEnd_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "end" {
		return nil, StringLiteralValidateFailed(
			"WorkDoneProgressEnd_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}

type WorkDoneProgressReport_Kind_StringLiteral struct {
	V String // report
}

func (this *WorkDoneProgressReport_Kind_StringLiteral) UnmarshalJSON(
	data []byte,
) error {
	var tmp string
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp != "report" {
		return StringLiteralValidateFailed(
			"WorkDoneProgressReport_Kind_StringLiteral",
		)
	}
	this.V = String(tmp)
	return nil
}

func (this *WorkDoneProgressReport_Kind_StringLiteral) MarshalJSON() ([]byte, error) {
	if this.V != "report" {
		return nil, StringLiteralValidateFailed(
			"WorkDoneProgressReport_Kind_StringLiteral",
		)
	}
	return json.Marshal(&this.V)
}
