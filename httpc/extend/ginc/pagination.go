package ginc

import (
	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/convc"
)

func GetPage(c *gin.Context) int {
	page := convc.Str(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convc.Str(c.Query("page_size")).MustInt()
	// todo
	if pageSize <= 0 {
		return 50
	}
	if pageSize > 200 {
		return 200
	}

	return pageSize
}
