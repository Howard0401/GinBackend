package router

import (
	"VueGin/global"
	orderhandler "VueGin/handler/order"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(r *gin.RouterGroup) {
	methods := orderhandler.OrderHandler{
		OrderSrv: &service.OrderService{
			Repo: &repository.OrderRepository{
				DB: global.Global_DB,
			},
		}}

	order := r.Group("/order")
	{
		order.GET("/list", methods.OrderList)
		order.GET("/info/:id", methods.OrderInfo)
		order.POST("/add", methods.AddOrder)
		order.POST("/edit", methods.EditOrder)
		order.POST("/delete/:id", methods.DeleteOrder)
	}
}
