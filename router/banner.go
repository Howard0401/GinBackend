package router

import (
	"VueGin/global"
	"VueGin/handler"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitBannerRouter(r *gin.RouterGroup) {
	methods := handler.BannerHandler{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: global.Global_DB,
			},
		}}
	banner := r.Group("/banner")
	{
		banner.GET("/list", methods.BannerList)
		banner.GET("/info", methods.BannerInfo)
		banner.POST("/add", methods.AddBanner)
		banner.POST("/edit", methods.EditBanner)     //坑點:要是id不存在，不會報錯，只會顯示nil why?
		banner.POST("/delete", methods.DeleteBanner) //坑點:要是id不存在，不會報錯，只會顯示nil why?
	}

}
