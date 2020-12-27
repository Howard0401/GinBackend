package handler

import (
	"VueGin/model"
	request "VueGin/model/requestType"
	"VueGin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthSrv service.AuthSrv
}

func (h *AuthHandler) AuthByJWT(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "帳號密碼輸入錯誤"})
	}
	Login := model.User{NickName: login.UserName, Password: login.Password}
	tocken, err := h.AuthSrv.GetToken(Login)

	c.JSON(http.StatusOK, gin.H{"tocken": tocken})
}
