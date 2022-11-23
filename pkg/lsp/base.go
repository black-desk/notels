package lsp

import (
	"github.com/black-desk/notels/pkg/jsonrpc2"
)

type URI string

type DocumentUri string

type Integer int64

type Uinteger uint64

type Decimal float64

type RegExp string

type String = jsonrpc2.String

type Boolean = jsonrpc2.Boolean

type Null = jsonrpc2.Null
