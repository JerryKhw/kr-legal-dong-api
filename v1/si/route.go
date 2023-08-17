package si

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	si := rg.Group("/si")
	{
		si.GET("", GetSiList)
		si.GET(":code", GetSi)
	}
}
