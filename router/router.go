package router

import (
	_ "VueGin/docs"
	"VueGin/global"
	"VueGin/middleware"

	"github.com/gin-gonic/gin"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	r.Use(middleware.Cors()) //register Cors middleware
	// }global.Global_Config.System.Timeout
	// r.Use(middleware.ContextTimeout(1 * time.Nanosecond))//先不用這個
	r.Use(middleware.RateLimiter(Limiter))
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := r.Group("")
	{
		InitLoginRouter(PublicGroup)
	}
	PrivateGroup := r.Group("/api")
	// PrivateGroup.Use(middleware.JWT()) //登入後需驗證JWT token
	PrivateGroup.Use()
	{
		InitBannerRouter(PrivateGroup)
		InitCategoryRouter(PrivateGroup)
		InitOrderRouter(PrivateGroup)
		InitProductRouter(PrivateGroup)
		InitUserRouter(PrivateGroup)
	}
	return r
}
