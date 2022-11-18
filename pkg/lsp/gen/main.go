package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/black-desk/notels/internal/utils/logger"
	"github.com/black-desk/notels/pkg/lsp/gen/internal/model"
	"github.com/black-desk/notels/pkg/lsp/gen/internal/types"
)

const (
	metaModelJsonFileName = "metaModel.json"
)

var log = logger.Get("lspgen")

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalw("failed to getwd",
			"error", err)
	}

	log.Infow("codegen start", "working directory", wd)

	metaModelJsonFile, err := os.Open(metaModelJsonFileName)
	if err != nil {
		log.Fatalw("failed to open meta model json file",
			"path", filepath.Join(wd, metaModelJsonFileName),
			"error", err)
	}

	content, err := io.ReadAll(metaModelJsonFile)
	if err != nil {
		log.Fatalw("failed to read meta model file",
			"error", err)
	}

	metaModel := model.MetaModel{}

	err = json.Unmarshal(content, &metaModel)
	if err != nil {
		log.Fatalw("failed to unmarshal meta model",
			"error", err)
	}

	log.Debugw("meta model parsed",
		"version", metaModel.MetaData.Version,
		"enumerations", len(metaModel.Enumerations),
		"structures", len(metaModel.Structures),
		"typeAliases", len(metaModel.TypeAliases),
		"notifications", len(metaModel.Notifications),
		"requests", len(metaModel.Requests),
	)

	types.GenEnumerations(&metaModel)

	genStructures(&metaModel)

	genTypeAliases(&metaModel)

	genClient(&metaModel)

	genServer(&metaModel)

	genExtra(&metaModel)

	log.Info("done")
}

func goMethodName(json string) string {
	json = strings.Replace(json, "$", "Lsp", -1)
	tmp := strings.Split(json, "/")
	for i, s := range tmp {
		tmp[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	json = strings.Join(tmp, "")
	return json
}
