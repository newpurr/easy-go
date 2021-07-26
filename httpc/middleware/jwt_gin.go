package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/newpurr/easy-go/errcode"
	"github.com/newpurr/easy-go/http/extend/ginc"
	"github.com/newpurr/easy-go/util"

	"github.com/gin-gonic/gin"
)

func GinJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
			ecode = errcode.Success
		)

		token = tryToExtractToken(c)
		if token == "" {
			ecode = errcode.InvalidParams
			goto check
		}

		// todo
		_, err = util.ParseToken(token, "")
		if err == nil {
			c.Next()
			return
		}

		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			ecode = errcode.UnauthorizedTokenTimeout
		default:
			ecode = errcode.UnauthorizedTokenError
		}

	check:
		if ecode != errcode.Success {
			response := ginc.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
	}
}

func tryToExtractToken(c *gin.Context) string {
	var token string
	if s, exist := c.GetQuery("token"); exist {
		token = s
	} else {
		token = c.GetHeader("token")
	}
	return token
}
