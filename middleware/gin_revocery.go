package middleware

import (
	email "VueGin/Utils/email"
	"VueGin/global"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// func GinLogger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()
// 		path := c.Request.URL.Path
// 		query := c.Request.URL.RawQuery
// 		c.Next()

// 	}
// }

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logMail := email.NewEmail()
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					//這邊寫在前面的原因是，如果zaplogger沒有初始化好，報invalid memory address or nil pointer dereference錯誤時
					//還能優先回傳訊息
					err := logMail.SendEmail(
						global.Global_Config.SMTP.To,
						fmt.Sprintf("呼叫API報錯(calling api exception: exception time) 發生錯誤時間：%v", time.Now()),
						fmt.Sprintf("發生錯誤(error)；%v 錯誤資訊(info): %v", zap.Any("error", err).String, zap.String("request", string(httpRequest)).String))

					if err != nil {
						zap.L().Debug("Email", zap.Any("Send Email Error", err))
					}

					global.Global_Logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					//這邊寫在前面的原因是，如果zaplogger沒有初始化好，報invalid memory address or nil pointer dereference錯誤時
					//還能優先回傳訊息
					err := logMail.SendEmail(
						global.Global_Config.SMTP.To,
						fmt.Sprintf("呼叫API報錯(calling api exception: exception time) 發生錯誤時間：%v", time.Now()),
						fmt.Sprintf("發生錯誤(error)；%v 錯誤資訊(info): %v", err, zap.String("request", string(httpRequest)).String))

					if err != nil {
						zap.L().Debug("Email", zap.Any("Send Email Error", err))
					}

					global.Global_Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)

				} else {
					//這邊寫在前面的原因是，如果zaplogger沒有初始化好，報invalid memory address or nil pointer dereference錯誤時
					//還能優先回傳訊息
					err := logMail.SendEmail(
						global.Global_Config.SMTP.To,
						fmt.Sprintf("呼叫API報錯(calling api exception: exception time) 發生錯誤時間：%v", time.Now()),
						fmt.Sprintf("發生錯誤(error)；%v 錯誤資訊(info): %v", zap.Any("error", err).String, zap.String("request", string(httpRequest)).String))

					if err != nil {
						zap.L().Debug("Email", zap.Any("Send Email Error", err))
					}

					global.Global_Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
