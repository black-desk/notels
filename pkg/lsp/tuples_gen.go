// Code generated from metaModel.json by "lspgen". DO NOT EDIT

package lsp

import (
	"encoding/json"
	"fmt"
)

var TupleValidateFailed = func(name string) error {
	return fmt.Errorf("tuple \"%s\"validate failed", name)
}

type ParameterInformation_Label_Or_1 struct {
	TupleItem0 Uinteger

	TupleItem1 Uinteger
}

func (this *ParameterInformation_Label_Or_1) UnmarshalJSON(data []byte) error {
	var msg = []json.RawMessage{}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	var tmpUnmarshal ParameterInformation_Label_Or_1

	if err := json.Unmarshal(msg[0], &tmpUnmarshal.TupleItem0); err != nil {
		return err
	}

	if err := json.Unmarshal(msg[1], &tmpUnmarshal.TupleItem1); err != nil {
		return err
	}

	*this = tmpUnmarshal
	return nil
}

func (this *ParameterInformation_Label_Or_1) MarshalJSON() ([]byte, error) {
	var msg = []json.RawMessage{}

	{
		bytes, err := json.Marshal(this.TupleItem0)
		if err != nil {
			return nil, err
		}
		msg = append(msg, bytes)
	}

	{
		bytes, err := json.Marshal(this.TupleItem1)
		if err != nil {
			return nil, err
		}
		msg = append(msg, bytes)
	}

	return json.Marshal(msg)
}
