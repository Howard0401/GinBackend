package initSettings

import (
	"VueGin/middleware"
	"VueGin/router"

	"github.com/gin-gonic/gin"
)

//注意回傳型別
func Routers(r *gin.Engine) *gin.Engine {

	//middleware
	//其實，呼叫Default時，就已預設使用Logger(), Recovery()這兩個middleware
	//這邊使用zaplog設定，抽換原先Gin預設的log
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	//register Cors middleware
	r.Use(middleware.Cors())

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
