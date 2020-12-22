package server

import (
	"VueGin/global"
	"VueGin/handler"
	"VueGin/initSettings"
	"VueGin/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	BannerHandler   handler.BannerHandler
	CategoryHandler handler.CategoryHandler
	OrderHandler    handler.OrderHandler
	ProductHandler  handler.ProductHandler
	UserHandler     handler.UserHandler
)

func RunServer() {
	//ver1. 使用gin設定路由、啟動API
	r := gin.Default()
	r.Use(middleware.Cors())

	initSettings.Routers(r)
	gin.SetMode(viper.GetString("mode"))

	port := global.Global_Viper.GetString("port")
	r.Run(port)
	// if err := r.Run(port); err != nil {
	// 	log.Fatalf("r.Run(port) failed:%v", err)
	// }

	//TODO: ver2. 使用gin設定路由、使用原生Http監聽，並用middleware設定CasbinHandler攔截訊息
	// r := gin.Default()
	// r.Use(middleware.Cors())
	// gin.SetMode(viper.GetString("mode"))
	// //要用原生http的話，要再用一個CasbinHandler做middleware
	// Router := initSettings.Routers(r)
	// address := fmt.Sprintf(":%s", viper.GetString("port"))

	// s := initHTTPServer(address, Router)
	// err := s.ListenAndServe()
	// if err != nil {
	// 	log.Fatalf(" s.ListenAndServe() failed, err: %v", err)
	// }
}
