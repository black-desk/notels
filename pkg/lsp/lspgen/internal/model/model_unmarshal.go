package model

import (
	"encoding/json"
)

// UnmarshalJSON implements json.Unmarshaler
func (n *Notification) UnmarshalJSON(data []byte) error {
	type unmarshalNotification struct {
		Documentation       string          `json:"documentation"`
		MessageDirection    string          `json:"messageDirection"`
		Method              string          `json:"method"`
		Params              json.RawMessage `json:"params"`
		Proposed            bool            `json:"proposed"`
		RegistrationMethod  string          `json:"registrationMethod"`
		RegistrationOptions *Type           `json:"registrationOptions"`
		Since               string          `json:"since"`
	}

	var tmp unmarshalNotification

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	n.Documentation = tmp.Documentation
	n.MessageDirection = tmp.MessageDirection
	n.Method = tmp.Method

	var tmpType Type
	if tmp.Params != nil {

		if err := json.Unmarshal(tmp.Params, &tmpType); err == nil {
			n.Params = []Type{tmpType}
		} else {
			err = json.Unmarshal(tmp.Params, &n.Params)
			if err != nil {
				return err
			}
		}
	}

	n.Proposed = tmp.Proposed
	n.RegistrationMethod = tmp.RegistrationMethod
	n.RegistrationOptions = tmp.RegistrationOptions
	n.Since = tmp.Since

	return nil
}

var _ json.Unmarshaler = &Notification{}

// UnmarshalJSON implements json.Unmarshaler
func (r *Request) UnmarshalJSON(data []byte) error {

	type unmarshalRequest struct {
		Documentation       string          `json:"documentation"`
		ErrorData           *Type           `json:"errorData"`
		MessageDirection    string          `json:"messageDirection"`
		Method              string          `json:"method"`
		Params              json.RawMessage `json:"params"`
		PartialResult       *Type           `json:"partialResult"`
		Proposed            bool            `json:"proposed"`
		RegistrationMethod  string          `json:"registrationMethod"`
		RegistrationOptions *Type           `json:"registrationOptions"`
		Result              *Type           `json:"result"`
		Since               string          `json:"since"`
	}

	var tmp unmarshalRequest

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	r.Documentation = tmp.Documentation
	r.ErrorData = tmp.ErrorData
	r.MessageDirection = tmp.MessageDirection
	r.Method = tmp.Method

	if tmp.Params != nil {
		var tmpType Type
		if err := json.Unmarshal(tmp.Params, &tmpType); err == nil {
			r.Params = []Type{tmpType}
		} else {
			err = json.Unmarshal(tmp.Params, &r.Params)
			if err != nil {
				return err
			}
		}
	}

	r.PartialResult = tmp.PartialResult
	r.Proposed = tmp.Proposed
	r.RegistrationMethod = tmp.RegistrationMethod
	r.RegistrationOptions = tmp.RegistrationOptions
	r.Result = tmp.Result
	r.Since = tmp.Since
	return nil

}

var _ json.Unmarshaler = &Request{}

// UnmarshalJSON implements json.Unmarshaler
func (t *Type) UnmarshalJSON(data []byte) error {
	type unmarshalType struct {
		Kind    string          `json:"kind"`
		Name    string          `json:"name"`
		Element *Type           `json:"element"`
		Key     *Type           `json:"key"`
		Value   json.RawMessage `json:"value"`
		Items   []Type          `json:"items"`
	}

	var tmpUnmarshalType unmarshalType

	err := json.Unmarshal(data, &tmpUnmarshalType)
	if err != nil {
		return err
	}

	t.Kind = tmpUnmarshalType.Kind
	t.Name = tmpUnmarshalType.Name
	t.Element = tmpUnmarshalType.Element
	t.Key = tmpUnmarshalType.Key
	t.Value = toTypeValue(tmpUnmarshalType.Value)
	t.Items = tmpUnmarshalType.Items
	return nil
}

var _ json.Unmarshaler = &Type{}

// isTypeValue implements TypeValue
func (*Type) isTypeValue() {
	panic("unimplemented")
}

var _ TypeValue = &Type{}
var _ TypeValue = &StructureLiteral{}

func _() {
	var s String
	var _ TypeValue = &s
}

func toTypeValue(data json.RawMessage) TypeValue {
	if data == nil {
		return nil
	}
	var tmpType Type
	err := json.Unmarshal(data, &tmpType)
	if err == nil {
		return &tmpType
	}
	var tmpStructureLiteral StructureLiteral
	err = json.Unmarshal(data, &tmpStructureLiteral)
	if err == nil {
		return &tmpStructureLiteral
	}
	var tmpString String
	err = json.Unmarshal(data, &tmpString)
	if err == nil {
		return &tmpString
	}
	panic("all failed: " + string(data))
}

// isTypeValue implements TypeValue
func (*String) isTypeValue() {
	panic("unimplemented")
}

// isTypeValue implements TypeValue
func (*StructureLiteral) isTypeValue() {
	panic("unimplemented")
}
