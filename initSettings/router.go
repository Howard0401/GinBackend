package initSettings

import (
	"VueGin/router"

	"github.com/gin-gonic/gin"
)

//注意回傳型別
func Routers(r *gin.Engine) *gin.Engine {

	// r := gin.Default()
	//register Cors middleware
	// r.Use(middleware.Cors())
	// gin.SetMode(viper.GetString("mode"))

	PrivateGroup := r.Group("/api")
	PrivateGroup.Use()
	{
		router.InitBannerRouter(PrivateGroup)
		router.InitCategoryRouter(PrivateGroup)
		router.InitOrderRouter(PrivateGroup)
		router.InitProductRouter(PrivateGroup)
		router.InitUserRouter(PrivateGroup)
	}
	return r
}
