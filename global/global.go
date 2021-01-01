package global

import (
	settings "VueGin/config/settingModels"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Global_DB     *gorm.DB
	Global_Viper  *viper.Viper
	Global_Config *settings.ConfigSettings
	Global_Logger *zap.Logger
	Global_Tracer *opentracing.Tracer
)
