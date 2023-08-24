package detail

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"
	"strconv"

	koreanregexp "github.com/JerryKhw/korean-regexp"
	"github.com/gin-gonic/gin"
)

// GetDetailList godoc
// @Summary 상세 리스트 조회
// @Description 상세 리스트 조회
// @Param request query detail.GetDetailList.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]detail.GetDetailList.detail} "success"
// @Failure 400 {object} model.DefaultResponse "bad_request"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/detail [get]
func GetDetailList(c *gin.Context) {
	type request struct {
		SiCode    *string `form:"siCode"`
		GuCode    *string `form:"guCode"`
		DongCode  *string `form:"dongCode"`
		Keyword   *string `form:"keyword"`
		Active    *string `form:"active"`
		UseRegExp *string `form:"useRegExp"`
	}

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

	req := &request{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "bad_request",
		})
		c.Abort()
		return
	}

	where := []string{}
	value := []any{}

	if req.Keyword != nil {
		if req.UseRegExp != nil {
			useRegExp, err := strconv.ParseBool(*req.UseRegExp)

			if err != nil {
				c.JSON(http.StatusBadRequest, &model.DefaultResponse{
					Message: "bad_request",
				})
				c.Abort()
				return
			}

			if useRegExp {
				where = append(where, "de.name REGEXP ?")
				value = append(value, koreanregexp.GetRegExp(*req.Keyword, koreanregexp.GetRegExpOptions{}).String())
			} else {
				where = append(where, "de.name Like ?")
				value = append(value, "%"+*req.Keyword+"%")
			}
		} else {
			where = append(where, "de.name Like ?")
			value = append(value, "%"+*req.Keyword+"%")
		}
	}

	if req.Active != nil {
		active, err := strconv.ParseBool(*req.Active)

		if err != nil {
			c.JSON(http.StatusBadRequest, &model.DefaultResponse{
				Message: "bad_request",
			})
			c.Abort()
			return
		}

		where = append(where, "de.active = ?")
		value = append(value, active)
	}

	if req.DongCode != nil {
		where = append(where, "do.code = ?")
		value = append(value, req.DongCode)
	}

	if req.GuCode != nil {
		where = append(where, "gu.code = ?")
		value = append(value, req.GuCode)
	}

	if req.SiCode != nil {
		where = append(where, "si.code = ?")
		value = append(value, req.SiCode)
	}

	whereString := ""

	for index, data := range where {
		if index == 0 {
			whereString += " WHERE "
		} else {
			whereString += " AND "
		}
		whereString += data
	}

	rows, err := db.DB.Query("SELECT de.code, si.code, si.name, gu.code, gu.name, do.code, do.name, de.name, de.active FROM detail AS de INNER JOIN dong AS do ON de.dong_code = do.code INNER JOIN gu ON do.gu_code = gu.code INNER JOIN si ON gu.si_code = si.code"+whereString, value...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_select",
		})
		c.Abort()
		return
	}

	data := []detail{}

	for rows.Next() {
		var result detail
		rows.Scan(&result.Code, &result.SiCode, &result.SiName, &result.GuCode, &result.GuName, &result.DongCode, &result.DongName, &result.Name, &result.Active)
		data = append(data, result)
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    data,
	})
}
