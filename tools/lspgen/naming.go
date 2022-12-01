package main

import (
	"strings"
)

func MethodNameFromString(lspJsonRPCMethodName string) string {
	lspJsonRPCMethodName = strings.Replace(
		lspJsonRPCMethodName,
		"$",
		"Lsp",
		-1,
	)
	tmp := strings.Split(lspJsonRPCMethodName, "/")
	for i, s := range tmp {
		tmp[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	lspJsonRPCMethodName = strings.Join(tmp, "")
	return fix(lspJsonRPCMethodName)
}

func fix(name string) string {
	if name == "_InitializeParams" {
		return "XInitializeParams"
	}
	return name
}

func typeName(prefix string, current string, t *Type) string {
	if t.Kind == "reference" {
		return fix(t.Name)
	} else if t.Kind == "base" {
		return baseTypeName(t.Name)
	} else if t.Kind == "array" {
		if len(prefix) != 0 {
			return "[]" + typeName(
				prefix+"_"+current,
				"Element",
				t.Element,
			)
		} else {
			return "[]" + typeName(current, "Element", t.Element)

		}
	} else if t.Kind == "and" {
		RegisterAnd(prefix+"_"+current+"_And", t)
		return prefix + "_" + current + "_And"
	} else if t.Kind == "or" {
		if len(prefix) != 0 {
			RegisterOr(prefix+"_"+current+"_Or", t)
			return prefix + "_" + current + "_Or"

		} else {
			RegisterOr(current+"_Or", t)
			return current + "_Or"
		}
	} else if t.Kind == "map" {
		if len(prefix) != 0 {
			return "map[" + typeName(
				prefix+"_"+current,
				"Key",
				t.Key,
			) + "]" + typeValueName(prefix+"_"+current, "Value", t.Value)
		} else {
			return "map[" + typeName(
				current,
				"Key",
				t.Key,
			) + "]" + typeValueName(current, "Value", t.Value)
		}
	} else if t.Kind == "literal" {
		RegisterLiteral(prefix+"_"+current, t.Value.(*StructureLiteral))
		return prefix + "_" + current
	} else if t.Kind == "stringLiteral" {
                RegisterStringLiteral(prefix+"_"+current+"_StringLiteral",t.Value.(*StringLiteral))
		return prefix+"_"+current+"_StringLiteral"
	} else if t.Kind == "tuple" {
		RegisterTuple(prefix+"_"+current, t)
		return prefix + "_" + current
	}
	log.Fatalw("unexpected kind",
		"kind", t.Kind,
	)
	panic("")
}

func typeValueName(prefix string, current string, tv TypeValue) string {
	switch v := tv.(type) {
	case *Type:
		return typeName(prefix, current, v)
	// case *StringLiteral:
	// case *StructureLiteral:
	default:
		log.Fatalw("unexpected kind",
			"kind", tv,
		)
		panic("")
	}

}
