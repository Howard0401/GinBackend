package global

import (
	"VueGin/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Global_DB     *gorm.DB
	Global_Viper  *viper.Viper
	Global_Config config.Config
	// Global_JWT
	// BannerHandler   handler.BannerHandler
	// CategoryHandler handler.CategoryHandler
	// OrderHandler    handler.OrderHandler
	// ProductHandler  handler.ProductHandler
	// UserHandler     handler.UserHandler
)
