package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/errcode"
	"github.com/newpurr/easy-go/http/extend/ginc"
	"github.com/newpurr/easy-go/limiter"
)

func RateLimiter(l limiter.Interface) gin.HandlerFunc {
	getKey := func(c *gin.Context) string {
		return c.FullPath()
	}

	return func(c *gin.Context) {
		key := getKey(c)
		bucket, ok := l.GetBucket(key)
		if !ok {
			c.Next()
			return
		}

		count := bucket.TakeAvailable(1)
		if count == 0 {
			response := ginc.NewResponse(c)
			response.ToErrorResponse(errcode.TooManyRequests)
			c.Abort()
			return
		}
	}
}
