package router

import (
	"VueGin/global"
	producthandler "VueGin/handler/product"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitProductRouter(r *gin.RouterGroup) {
	methods := producthandler.ProductHandler{
		ProductSrv: &service.ProductService{
			Repo: &repository.ProductRepository{
				DB: global.Global_DB,
			},
		}}
	product := r.Group("/product")
	{
		product.GET("/list", methods.ProductList)
		product.GET("/info/:id", methods.ProductInfo)
		product.POST("/add", methods.AddProduct)
		product.POST("/edit", methods.EditProduct)
		product.POST("/delete/:id", methods.DeleteProduct)
	}

}
