package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

type Integer int64

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
