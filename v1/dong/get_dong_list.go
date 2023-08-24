package dong

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"
	"strconv"

	koreanregexp "github.com/JerryKhw/korean-regexp"
	"github.com/gin-gonic/gin"
)

// GetDongList godoc
// @Summary 동 리스트 조회
// @Description 동 리스트 조회
// @Param request query dong.GetDongList.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]dong.GetDongList.dong} "success"
// @Failure 400 {object} model.DefaultResponse "bad_request"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/dong [get]
func GetDongList(c *gin.Context) {
	type request struct {
		SiCode    *string `form:"siCode"`
		GuCode    *string `form:"guCode"`
		Keyword   *string `form:"keyword"`
		Active    *string `form:"active"`
		UseRegExp *string `form:"useRegExp"`
	}

	type dong struct {
		Code   string `json:"code" binding:"required"`
		SiCode string `json:"siCode" binding:"required"`
		SiName string `json:"siName" binding:"required"`
		GuCode string `json:"guCode" binding:"required"`
		GuName string `json:"guName" binding:"required"`
		Name   string `json:"name" binding:"required"`
		Active bool   `json:"active" binding:"required"`
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
				where = append(where, "do.name REGEXP ?")
				value = append(value, koreanregexp.GetRegExp(*req.Keyword, koreanregexp.GetRegExpOptions{}).String())
			} else {
				where = append(where, "do.name Like ?")
				value = append(value, "%"+*req.Keyword+"%")
			}
		} else {
			where = append(where, "do.name Like ?")
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

		where = append(where, "do.active = ?")
		value = append(value, active)
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

	rows, err := db.DB.Query("SELECT do.code, si.code, si.name, gu.code, gu.name, do.name, do.active FROM dong AS do INNER JOIN gu ON do.gu_code = gu.code INNER JOIN si ON gu.si_code = si.code"+whereString, value...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_select",
		})
		c.Abort()
		return
	}

	data := []dong{}

	for rows.Next() {
		var result dong
		rows.Scan(&result.Code, &result.SiCode, &result.SiName, &result.GuCode, &result.GuName, &result.Name, &result.Active)
		data = append(data, result)
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    data,
	})
}
