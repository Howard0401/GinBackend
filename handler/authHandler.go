package handler

import (
	"VueGin/global"
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
		c.JSON(http.StatusBadRequest, gin.H{"status": "帳號密碼輸入錯誤"})
		global.Global_Logger.Debug("msg")
	}
	Login := model.User{NickName: login.UserName, Password: login.Password}
	token, err := h.AuthSrv.GetToken(Login)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"tocken": "couldn't retrieve tocken"})
		global.Global_Logger.Debug("msg")
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
	global.Global_Logger.Debug("msg")
}
