package router

import (
	"VueGin/global"
	bannerhandler "VueGin/handler/banner"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

var methods bannerhandler.BannerHandler

//這樣寫是為了Unit Test方便
func GetMethods() {
	methods = bannerhandler.BannerHandler{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: global.Global_DB,
			},
		}}
}

func InitBannerRouter(r *gin.RouterGroup) {
	GetMethods()
	banner := r.Group("/banner")
	{
		banner.GET("/list", methods.BannerList)
		banner.GET("/info", methods.BannerInfo)
		banner.POST("/add", methods.AddBanner)
		banner.POST("/edit", methods.EditBanner)     //坑點:要是id不存在，不會報錯，只會顯示nil why?
		banner.POST("/delete", methods.DeleteBanner) //坑點:要是id不存在，不會報錯，只會顯示nil why?
	}

}
