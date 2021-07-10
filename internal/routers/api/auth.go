package api

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/internal/service"
	"github.com/newpurr/easy-go/pkg/domain"
	"github.com/newpurr/easy-go/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := domain.NewResponse(c)
	valid, errs := domain.BindAndValid(c, &param)
	if !valid {
		application.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		application.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := domain.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		application.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
