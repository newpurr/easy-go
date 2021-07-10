package middleware

import (
	"fmt"
	"github.com/newpurr/easy-go/application"
	"time"

	"github.com/newpurr/easy-go/pkg/email"

	"github.com/newpurr/easy-go/pkg/domain"
	"github.com/newpurr/easy-go/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     application.EmailSetting.Host,
		Port:     application.EmailSetting.Port,
		IsSSL:    application.EmailSetting.IsSSL,
		UserName: application.EmailSetting.UserName,
		Password: application.EmailSetting.Password,
		From:     application.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				application.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				err := defailtMailer.SendMail(
					application.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					application.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				domain.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
