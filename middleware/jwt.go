package middleware

import (
	jwt "VueGin/Utils/jwt"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"status": "帳號未登錄"})
			c.Abort()
			return
		}
		//這邊是把裝在utils包中的ParseToken方法：為何要這樣呢？
		// 1.這樣登入的路由也可以直接用utils包中的方法，而非引用middleware中的這個包，看起來可能會不直覺(已經透過middleware登入了，還要用middleware層中的方法創建新的JWT，倒不如放在utils層中)
		// Utils JWT included:
		// 1.返回JWT結構體，欄位含有toke NewJWT() return JWT struct, contains field SigInkey[]byte
		// 2.用Pointer Recever建構方法Cnreate, Parse方法

		claims, err := jwt.NewJWT().ParseToken(token)
		if err == errors.New("Token is expired") {
			c.JSON(http.StatusOK, gin.H{"status": "授權過期，解析出錯"})
			c.Abort()
			return
		}
		if err != nil {
			errMsg := fmt.Sprintf("解析出錯:%v", err)
			c.JSON(http.StatusOK, gin.H{"status": errMsg})
			c.Abort()
			return
		}
		//若使用者持續再使用則延長Token(因為是後臺管理，若是用戶的話，可能diable這部分好些)
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*30 //每次重新操作延時30分鐘
			newToken, err := jwt.NewJWT().CreateToken(*claims)
			if err != nil {
				fmt.Printf("CreateToken Failed: %v", err)
			}
			c.Header("token", newToken)
		}
		//TODO:參考開源專案可存進Redis

		c.Set("claims", claims)
		c.Next()
	}
}
