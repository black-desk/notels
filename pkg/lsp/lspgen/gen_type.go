package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/lspgen/internal/naming"
)

func genStructures(model *model.MetaModel) {
	fileName := "structures_gen.go"
	structuresGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer structuresGenFile.Close()
	structuresGenFile.Write(
		[]byte("// Code generated by \"lspgen\"; DO NOT EDIT\n"),
	)
	structuresGenFile.Write([]byte("package protocol\n"))
	for _, structure := range model.Structures {
		structuresGenFile.Write(
			[]byte(
				fmt.Sprintf(
					"/* %s */\ntype %s any\n",
					structure.Documentation,
					goMethodName(structure.Name),
				),
			),
		)
		// structuresGenFile.Write([]byte(fmt.Sprintf("func
		// lspgenUnmarshalJSONTo%s([]byte) (res *%s,err
		// error){return}\n", goMethodName(structure.Name),
		// goMethodName(structure.Name))))
		// structuresGenFile.Write([]byte(fmt.Sprintf("func
		// lspgenUnmarshalJSONTo%sSlice([]byte) (res *[]%s,err
		// error){return}\n", goMethodName(structure.Name),
		// goMethodName(structure.Name))))
	}
}

func genTypeAliases(model *model.MetaModel) {
	fileName := "typeAliases_gen.go"
	typeAliasesGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer typeAliasesGenFile.Close()
	typeAliasesGenFile.Write(
		[]byte("// Code generated by \"lspgen\"; DO NOT EDIT\n"),
	)
	typeAliasesGenFile.Write([]byte("package protocol\n"))
	for _, alias := range model.TypeAliases {
		typeAliasesGenFile.Write(
			[]byte(
				fmt.Sprintf(
					"/* %s */\ntype %s any\n",
					alias.Documentation,
					naming.MethodNameFromString(alias.Name),
				),
			),
		)
	}
}

func genExtra(model *model.MetaModel) {
	fileName := "extraType_gen.go"
	extraTypeGenFile, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatalw("failed to open file",
			"name", fileName,
			"error", err)
	}
	defer extraTypeGenFile.Close()
	extraTypeGenFile.Write(
		[]byte("// Code generated by \"lspgen\"; DO NOT EDIT\n"),
	)
	extraTypeGenFile.Write([]byte("package protocol\n"))

	types := map[string]struct{}{}

	for _, request := range model.Requests {
		if request.ErrorData != nil {
			types[goMethodName(request.Method)+"_ErrorData"] = struct{}{}
		}
		if request.PartialResult != nil {
			types[goMethodName(request.Method)+"_PartialResult"] = struct{}{}
		}
		if request.RegistrationOptions != nil {
			if request.RegistrationMethod != "" {

				types[goMethodName(request.RegistrationMethod)+"_Options"] = struct{}{}
			} else {
				types[goMethodName(request.Method)+"_RegistrationOptions"] = struct{}{}
			}
		}
		if request.Result != nil {
			types[goMethodName(request.Method)+"_Result"] = struct{}{}
		}
	}

	for _, notifi := range model.Notifications {
		if notifi.RegistrationOptions != nil {
			types[goMethodName(notifi.Method)+"_RegistrationOption"] = struct{}{}
		}

	}

	typesSlice := []string{}

	for t := range types {
		typesSlice = append(typesSlice, t)
	}

	sort.Strings(typesSlice)

	for _, t := range typesSlice {
		extraTypeGenFile.Write([]byte(fmt.Sprintf("type %s any\n", t)))
	}
}