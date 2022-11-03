package protocol

import "fmt"

type URI string

type DocumentUri string

type Integer int64

type Uinteger uint64

type Decimal float64

type RegExp string

type String string

type Boolean bool


type interfaceUnmarshalJSONError []error

func (e*interfaceUnmarshalJSONError)Error()string{
        return fmt.Errorf("failed to unmarshal json to interface: %v", *e).Error()
}

