package si

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSi godoc
// @Summary 시 조회
// @Description 시 조회
// @Success 200 {object} model.DataResponse{data=si.GetSi.si} "success"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/si/{code} [get]
func GetSi(c *gin.Context) {
	code := c.Param("code")

	type si struct {
		Code   string `json:"code" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Active bool   `json:"active" binding:"required"`
	}

	var result si

	err := db.DB.QueryRow("SELECT si.code, si.name, si.active FROM si WHERE si.code = ?", code).Scan(&result.Code, &result.Name, &result.Active)

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
