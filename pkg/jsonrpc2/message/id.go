package message

import (
	"encoding/json"
	"fmt"

	. "github.com/black-desk/lib/go/errwrap"
)

type JsonRPCIDType uint

const (
	JsonRPCIDTypeNumber JsonRPCIDType = iota
	JsonRPCIDTypeString
	JsonRPCIDTypeNull
)

type ID struct {
	Type JsonRPCIDType
	Str  string
}

func (id *ID) IsStr() bool {
	return id.Type == JsonRPCIDTypeString
}

func (id *ID) IsNum() bool {
	return id.Type == JsonRPCIDTypeNumber
}

func (id *ID) IsNull() bool {
	return id.Type == JsonRPCIDTypeNull
}

func (id *ID) HashKey() string {
	return fmt.Sprintf("%d%s", id.Type, id.Str)
}

func (id *ID) MarshalJSON() ([]byte, error) {
	if id.IsNull() {
		return nil, nil
	}

	if id.IsStr() {
		ret, err := json.Marshal(id.Str)
		return ret, Trace(err)
	}

	if id.IsNum() {
		tmpNumber := json.Number(id.Str)
		ret, err := json.Marshal(tmpNumber)
		return ret, Trace(err)
	}

	panic("unknow JSON-RPC 2.0 ID type")
}

func (id *ID) UnmarshalJSON(raw []byte) error {
	if len(raw) <= 0 {
		return Trace(ErrInvalidID, "raw: %s", string(raw))
	}

	if raw[0] == 'n' { // null
		id = &ID{
			Type: JsonRPCIDTypeNull,
			Str:  "",
		}
		return nil
	}

	if raw[0] == '"' { // string
		err := json.Unmarshal(raw, &id.Str)
		if err != nil {
			return Trace(err)
		}

		id.Type = JsonRPCIDTypeString
		return nil
	}

	tmpNumber := json.Number(string(raw))
	if _, err := tmpNumber.Float64(); err != nil {
		return Annotate(err, "ID should be number / string / null")
	}

	id.Type = JsonRPCIDTypeNumber
	id.Str = tmpNumber.String()

	return nil
}

var _ json.Marshaler = &ID{}
var _ json.Unmarshaler = &ID{}
