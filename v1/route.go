package v1

import (
	"kr-legal-dong-api/v1/detail"
	"kr-legal-dong-api/v1/dong"
	"kr-legal-dong-api/v1/gu"
	"kr-legal-dong-api/v1/si"

	"github.com/gin-gonic/gin"
)

func SetRoute(eg *gin.Engine) {
	v1 := eg.Group("/v1")

	si.SetRoute(v1)
	gu.SetRoute(v1)
	dong.SetRoute(v1)
	detail.SetRoute(v1)
}
