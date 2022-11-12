package main

import (
	"github.com/goccy/go-json"
)

type MetaModel struct {
	MetaData struct {
		Version string `json:"version"`
	} `json:"metaData"`

	Enumerations  []Enumeration  `json:"enumerations"`
	Notifications []Notification `json:"notifications"`
	Requests      []Request      `json:"requests"`
	Structures    []Structure    `json:"structures"`
	TypeAliases   []TypeAlias    `json:"typeAliases"`
}

type Enumeration struct {
	Documentation         string `json:"documentation"`
	Name                  string `json:"name"`
	Proposed              bool   `json:"proposed"`
	Since                 string `json:"since"`
	SupportsCustiomValues bool   `json:"supportsCustomValues"`
	Type                  struct {
		Name string `json:"name"`
	} `json:"type"`
	Values []struct {
		Documentation string          `json:"documentation"`
		Name          string          `json:"name"`
		Proposed      bool            `json:"proposed"`
		Since         string          `json:"since"`
		Value         json.RawMessage `json:"value"` // number or string
	} `json:"values"`
}

type Notification struct {
	Documentation       string `json:"documentation"`
	MessageDirection    string `json:"messageDirection"`
	Method              string `json:"method"`
	Params              []Type `json:"params"`
	Proposed            bool   `json:"proposed"`
	RegistrationMethod  string `json:"registrationMethod"`
	RegistrationOptions *Type   `json:"registrationOptions"`
	Since               string `json:"since"`
}

type Request struct {
	Documentation       string `json:"documentation"`
	ErrorData           *Type   `json:"errorData"`
	MessageDirection    string `json:"messageDirection"`
	Method              string `json:"method"`
	Params              []Type `json:"params"`
	PartialResult       *Type   `json:"partialResult"`
	Proposed            bool   `json:"proposed"`
	RegistrationMethod  string `json:"registrationMethod"`
	RegistrationOptions *Type   `json:"registrationOptions"`
	Result              *Type   `json:"result"`
	Since               string `json:"since"`
}

type Structure struct {
	Documentation string     `json:"documentation"`
	Extends       []Type     `json:"extends"`
	Mixins        []Type     `json:"mixins"`
	Name          string     `json:"name"`
	Properties    []Property `json:"properties"`
	Proposed      bool       `json:"proposed"`
	Since         string     `json:"since"`
}

type TypeAlias struct {
	Documentation string `json:"documentation"`
	Name          string `json:"name"`
	Proposed      bool   `json:"proposed"`
	Since         string `json:"since"`
	Type          Type   `json:"type"`
}

type Type struct {
	Kind    string    `json:"kind"`
	Name    string    `json:"name"`
	Element *Type     `json:"element"`
	Key     *Type     `json:"key"`
	Value   TypeValue `json:"value"`
	Items   []Type    `json:"items"`
}

type TypeValue interface {
	isTypeValue()
}

type StructureLiteral struct {
	Properties []Property `json:"properties"`
	Proposed   bool       `json:"proposed"`
	Since      string     `json:"since"`
}

type String string

type Property struct {
	Documentation string `json:"documentation"`
	Name          string `json:"name"`
	Optional      bool   `json:"optional"`
	Proposed      bool   `json:"proposed"`
	Since         string `json:"since"`
	Type          Type   `json:"type"`
}
