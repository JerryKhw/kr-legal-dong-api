package gu

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	si := rg.Group("/gu")
	{
		si.GET("", GetGu)
	}
}
