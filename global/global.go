package global

import (
	settings "VueGin/config/settingModels"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Global_DB     *gorm.DB
	Global_Viper  *viper.Viper
	Global_Config *settings.ConfigSettings
	Global_Logger *zap.Logger
	// Global_JWT
	// BannerHandler   handler.BannerHandler
	// CategoryHandler handler.CategoryHandler
	// OrderHandler    handler.OrderHandler
	// ProductHandler  handler.ProductHandler
	// UserHandler     handler.UserHandler
)
