package lsp

import (
	"encoding/json"
	"fmt"
)

type URI string

type DocumentUri string

type Integer int64

type Uinteger uint64

type Decimal float64

type RegExp string

type String string

type Boolean bool

type Null struct{}

func (this *Null) UnmarshalJSON(data []byte) error {
	var tmp *int
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp != nil {
		return fmt.Errorf("Null should always be null")
	}
	return nil
}

func (this *Null) MarshalJSON() ([]byte, error) {
	var tmp *int = nil
	return json.Marshal(tmp)
}

type interfaceUnmarshalJSONError []error

func (e *interfaceUnmarshalJSONError) Error() string {
	return fmt.Errorf("failed to unmarshal json to interface: %v", *e).
		Error()
}
