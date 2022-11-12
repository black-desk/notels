package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"

	"github.com/black-desk/notels/internal/utils/logger"
)

const (
	metaModelJsonFileName = "metaModel.json"
)

var log = logger.New("lspgen")

func main() {

	metaModel := loadMetaModel()

	genEnumerations(metaModel)

	genStructures(metaModel)

	genTypeAliases(metaModel)

	genClient(metaModel)

	genServer(metaModel)

	genExtra(metaModel)

	log.Info("done")
}

func loadMetaModel() *MetaModel {

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

	return &metaModel
}
