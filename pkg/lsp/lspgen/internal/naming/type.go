package naming

import (
	"strings"

	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/extra"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
)

func TypeName(t *model.Type, alias string) (ret string) {

	defer func() {
		extra.RegisterType(ret, t)
	}()

	if t.Name != "" {
		return MethodNameFromString(t.Name)
	}

	if alias != "" {
		return alias
	}

	if t.Kind == "array" {
		return "Array" + TypeName(t.Element, "")
	}

	kind := ""

	if t.Kind == "or" {
		kind = "Or_"
	} else if t.Kind == "and" {
		kind = "And_"
	} else {
		log.Fatalw("unexpected kind",
			"kind", t.Kind,
		)
	}

	slice := []string{kind}

	for _, item := range t.Items {
		slice = append(slice, TypeName(&item, ""))
	}

	slice = append(slice, "") // extra _ at the end

	return strings.Join(slice, "_")
}
