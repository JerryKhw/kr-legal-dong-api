package detail

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDetail godoc
// @Summary 상세 조회
// @Description 상세 조회
// @Param code path string true "code"
// @Success 200 {object} model.DataResponse{data=detail.GetDetail.detail} "success"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/detail/{code} [get]
func GetDetail(c *gin.Context) {
	code := c.Param("code")

	type detail struct {
		Code     string `json:"code" binding:"required"`
		SiCode   string `json:"siCode" binding:"required"`
		SiName   string `json:"siName" binding:"required"`
		GuCode   string `json:"guCode" binding:"required"`
		GuName   string `json:"guName" binding:"required"`
		DongCode string `json:"dongCode" binding:"required"`
		DongName string `json:"dongName" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Active   bool   `json:"active" binding:"required"`
	}

	var result detail

	err := db.DB.QueryRow("SELECT de.code, si.code, si.name, gu.code, gu.name, do.code, do.name, de.name, de.active FROM detail AS de INNER JOIN dong AS do ON de.dong_code = do.code INNER JOIN gu ON do.gu_code = gu.code INNER JOIN si ON gu.si_code = si.code WHERE de.code = ?", code).Scan(&result.Code, &result.SiCode, &result.SiName, &result.GuCode, &result.GuName, &result.DongCode, &result.DongName, &result.Name, &result.Active)

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
