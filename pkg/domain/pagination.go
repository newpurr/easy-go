package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/application"
	"github.com/newpurr/easy-go/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return application.AppSetting.DefaultPageSize
	}
	if pageSize > application.AppSetting.MaxPageSize {
		return application.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
