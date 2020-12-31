package handler

import (
	"VueGin/global"
	"VueGin/model"
	request "VueGin/model/requestType"
	"VueGin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	AuthSrv service.AuthSrv
}

func (h *AuthHandler) AuthByJWT(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "傳入JSON格式錯誤"})
		global.Global_Logger.Debug("msg", zap.Any("LoginFailed", err))
		return
	}

	if login.UserName == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "傳入帳號或密碼為空"})
		global.Global_Logger.Debug("msg", zap.String("LoginFailed", "前端傳入了空的帳號或密碼"))
		return
	}

	Login := model.User{NickName: login.UserName, Password: login.Password}

	token, err := h.AuthSrv.GetToken(Login)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"tocken": "couldn't retrieve tocken"})
		global.Global_Logger.Debug("msg", zap.Any("GetTokenFailed", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	global.Global_Logger.Debug("msg", zap.String("Login", "OK"))
}
