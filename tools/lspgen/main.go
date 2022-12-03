package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/black-desk/lib/go/logger"
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

	log.Infow("start", "working directory", wd)

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

	metaModel := MetaModel{}

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

	genEnumerations(&metaModel)
	genStruct(&metaModel)
	genTypeAliases(&metaModel)

	genClient(&metaModel)
	genServer(&metaModel)

	genAnd()
	genOr()
	genLiteral()
	genTuple()
	genStringLiteral()

	log.Info("done")
}
