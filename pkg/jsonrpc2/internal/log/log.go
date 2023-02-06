package log

import (
	"sync"

	"go.uber.org/zap"
)

var LogMux = sync.Mutex{}
var Log = zap.NewNop().Sugar()
