package export

import "github.com/black-desk/notels/internal/utils/logger"

var log = logger.Get("notels.export")

func Export(workspaceRootPath string, outputPath string) error {
	log.Debugw("Export called",
		"workspaceRootPath", workspaceRootPath,
		"outputPath", outputPath)
	return nil
}
