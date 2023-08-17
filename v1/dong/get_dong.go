package dong

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDong godoc
// @Summary 동 조회
// @Description 동 조회
// @Success 200 {object} model.DataResponse{data=dong.GetDong.dong} "success"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/dong/{code} [get]
func GetDong(c *gin.Context) {
	code := c.Param("code")

	type dong struct {
		Code   string `json:"code" binding:"required"`
		SiCode string `json:"siCode" binding:"required"`
		SiName string `json:"siName" binding:"required"`
		GuCode string `json:"guCode" binding:"required"`
		GuName string `json:"guName" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Active bool   `json:"active" binding:"required"`
	}

	var result dong

	err := db.DB.QueryRow("SELECT do.code, si.code, si.name, gu.code, gu.name, do.name, do.active FROM dong AS do INNER JOIN gu ON do.gu_code = gu.code INNER JOIN si ON gu.si_code = si.code WHERE do.code = ?", code).Scan(&result.Code, &result.SiCode, &result.SiName, &result.GuCode, &result.GuName, &result.Name, &result.Active)

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
