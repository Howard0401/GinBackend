package user

import (
	"VueGin/global"
	"VueGin/handler"
	"VueGin/repository"
	"VueGin/service"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	methods := handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: global.Global_DB,
			},
		}}
	user := r.Group("/user")
	{
		user.GET("/list", methods.UserList)
		user.GET("/info/:id", methods.UserInfo)
		user.POST("/add", methods.AddUser)
		user.POST("/edit", methods.EditUser)
		user.POST("/delete/:id", methods.DeleteUser)
	}

}
