package gu

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGu godoc
// @Summary 구 조회
// @Description 구 조회
// @Param code path string true "code"
// @Success 200 {object} model.DataResponse{data=gu.GetGu.gu} "success"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/gu/{code} [get]
func GetGu(c *gin.Context) {
	code := c.Param("code")

	type gu struct {
		Code   string `json:"code" binding:"required"`
		SiCode string `json:"siCode" binding:"required"`
		SiName string `json:"siName" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Active bool   `json:"active" binding:"required"`
	}

	var result gu

	err := db.DB.QueryRow("SELECT gu.code, si.code, si.name, gu.name, gu.active FROM gu INNER JOIN si ON gu.si_code = si.code WHERE gu.code = ?", code).Scan(&result.Code, &result.SiCode, &result.SiName, &result.Name, &result.Active)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_select",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    result,
	})
}
