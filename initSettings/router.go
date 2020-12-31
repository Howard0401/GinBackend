package initSettings

import (
	"VueGin/global"
	"VueGin/middleware"
	"VueGin/router"

	"github.com/gin-gonic/gin"
)

//注意回傳型別
func Routers(r *gin.Engine) *gin.Engine {

	//middleware
	r.Use(middleware.AppInfo()) //版本號 version

	//其實，呼叫Default時，就已預設使用Logger(), Recovery()這兩個middleware
	//這邊使用zaplog設定，抽換掉原先Gin預設的log
	if global.Global_Config.System.Mode == "debug" {
		r.Use(middleware.GinLogger())
		r.Use(middleware.GinRecovery(true))
	}
	r.Use(middleware.ContextTimeout(global.Global_Config.System.Timeout))
	r.Use(middleware.RateLimiter(Limiter))
	r.Use(middleware.Cors()) //register Cors middleware

	PublicGroup := r.Group("")
	{
		router.InitLoginRouter(PublicGroup)
	}
	PrivateGroup := r.Group("/api")
	// PrivateGroup.Use(middleware.JWT()) //登入後需驗證JWT token
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
