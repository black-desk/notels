package types

import (
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/naming"
)

var typesToGenerate = map[string]*model.Type{}

func RegisterTypes(t *model.Type) {
	name := naming.TypeName(t)
	if _, ok := typesToGenerate[name]; !ok {
		typesToGenerate[name] = t
	}
}
