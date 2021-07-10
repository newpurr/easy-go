package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/pkg/domain"
	"github.com/newpurr/easy-go/pkg/errcode"
	"github.com/newpurr/easy-go/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := domain.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
