package gu

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetGuList godoc
// @Summary 구 리스트 조회
// @Description 구 리스트 조회
// @Param request query gu.GetGuList.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]gu.GetGuList.gu} "success"
// @Failure 400 {object} model.DefaultResponse "bad_request"
// @Failure 500 {object} model.DefaultResponse "failed_select"
// @Router /v1/gu [get]
func GetGuList(c *gin.Context) {
	type request struct {
		SiCode  *string `form:"siCode"`
		Keyword *string `form:"keyword"`
		Active  *string `form:"active"`
	}

	type gu struct {
		Code   string `json:"code" binding:"required"`
		SiCode string `json:"siCode" binding:"required"`
		SiName string `json:"siName" binding:"required"`
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
		where = append(where, "gu.name LIKE ?")
		value = append(value, "%"+*req.Keyword+"%")
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

		where = append(where, "gu.active = ?")
		value = append(value, active)
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

	rows, err := db.DB.Query("SELECT gu.code, si.code, si.name, gu.name, gu.active FROM gu INNER JOIN si ON gu.si_code = si.code"+whereString, value...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_select",
		})
		c.Abort()
		return
	}

	data := []gu{}

	for rows.Next() {
		var result gu
		rows.Scan(&result.Code, &result.SiCode, &result.SiName, &result.Name, &result.Active)
		data = append(data, result)
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    data,
	})
}
