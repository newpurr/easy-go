package middleware

import (
	"github.com/newpurr/easy-go/boot"
	"github.com/newpurr/easy-go/errcode"
	"github.com/newpurr/easy-go/event"
	"github.com/newpurr/easy-go/http/extend/ginc"

	"github.com/gin-gonic/gin"
)

func Recovery(bus event.Bus) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				bus.Publish("gin.Recovery", boot.EventContext{
					Param: err,
				})

				ginc.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
