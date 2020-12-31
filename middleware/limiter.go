package middleware

import (
	limiter "VueGin/Utils/limiter"
	"net/http"

	"github.com/gin-gonic/gin"
)

//https://github.com/juju/ratelimit

func RateLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// TakeAvailable takes up to count immediately available tokens from the bucket. It returns the number of tokens removed, or zero if there are no available tokens.
			// 如果沒有可以分配的bucket就回覆 StatusTooManyRequests (429)
			count := bucket.TakeAvailable(1)
			if count == 0 {
				c.JSON(http.StatusTooManyRequests, gin.H{"code": "Too Many Requests"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
