package utils

import (
	"VueGin/global"
	res "VueGin/resViewModel"

	"go.uber.org/zap"
)

func EntityLog(entity res.Entity) {
	global.Global_Logger.Debug("entity", zap.Reflect("entity", entity))
}
