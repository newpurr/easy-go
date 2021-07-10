package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/application"
)

type GinBootloader struct {
}

func NewGinBootloader() *GinBootloader {
	return &GinBootloader{}
}

func (sb GinBootloader) Boot() error {
	gin.SetMode(application.ServerSetting.RunMode)
	return nil
}
