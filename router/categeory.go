package router

import (
	"VueGin/global"
	categoryhandler "VueGin/handler/category"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(r *gin.RouterGroup) {

	methods := categoryhandler.CategoryHandler{
		CategorySrv: &service.CategoryService{
			Repo: &repository.CategoryRepository{
				DB: global.Global_DB,
			},
		}}
	category := r.Group("/category")
	{
		category.GET("/list", methods.CategoryList)
		category.GET("/list4backend", methods.CategoryListForBackend)
		category.GET("/info/:id", methods.CategoryInfo) //查第三級分類
		category.POST("/add", methods.AddCategory)
		category.POST("/edit", methods.EditCategory)
		category.POST("/delete/:id", methods.DeleteCategory) //soft deleted
	}
}
