package dong

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(rg *gin.RouterGroup) {
	si := rg.Group("/dong")
	{
		si.GET("", GetDongList)
		si.GET(":code", GetDong)
	}
}
