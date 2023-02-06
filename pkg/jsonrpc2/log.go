package jsonrpc2

import (
	"github.com/black-desk/lib/go/logger"
	"github.com/black-desk/notels/pkg/jsonrpc2/internal/log"
	"go.uber.org/zap"
)

func EnableLog(newLogger *zap.SugaredLogger) {
	if newLogger == nil {
		newLogger = logger.Get("jsonrpc2")
		newLogger.Infof("logger for jsonrpc2 enabled")
	}

	log.LogMux.Lock()
	defer log.LogMux.Unlock()

	log.Log = newLogger

}

func DisableLog() {
	log.LogMux.Lock()
	defer log.LogMux.Unlock()

	log.Log.Infof("logger for jsonrpc2 disabled")
	log.Log = zap.NewNop().Sugar()
}
