package format

import (
	"VueGin/global"
	res "VueGin/model/res_view_model"

	"go.uber.org/zap"
)

func EntityLog(entity res.Entity) {
	global.Global_Logger.Debug("entity", zap.Reflect("entity", entity))
}
