package main

var TypesToGenerate = map[string]*Type{}

func RegisterType(name string, t *Type) {
	if _, ok := TypesToGenerate[name]; !ok {
		TypesToGenerate[name] = t
	}
}

func RemoveType(name string) {
	if _, ok := TypesToGenerate[name]; ok {
		TypesToGenerate[name] = nil
	}
}
