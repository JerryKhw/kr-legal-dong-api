package detail

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	si := rg.Group("/detail")
	{
		si.GET("", GetDetailList)
		si.GET(":code", GetDetail)
	}
}
