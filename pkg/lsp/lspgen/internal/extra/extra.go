package extra

import (
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
)

var TypesToGenerate = map[string]*model.Type{}

func RegisterType(name string, t *model.Type) {
	if _, ok := TypesToGenerate[name]; !ok {
		TypesToGenerate[name] = t
	}
}

func RemoveType(name string) {
	if _, ok := TypesToGenerate[name]; ok {
		TypesToGenerate[name] = nil
	}
}
