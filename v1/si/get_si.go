package si

import (
	"kr-legal-dong-api/db"
	"kr-legal-dong-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSi godoc
// @Summary 시 조회
// @Description 시 조회
// @Param request query si.GetSi.request true "query params"
// @Success 200 {object} model.DataResponse{data=[]si.GetSi.si} "success"
// @Router /v1/si [get]
func GetSi(c *gin.Context) {
	type request struct {
		Keyword *string `form:"keyword"`
		Active  *string `form:"active"`
	}

	type si struct {
		Code   string `json:"code" binding:"required"`
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
		where = append(where, "si.name LIKE ?")
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

		where = append(where, "si.active = ?")
		value = append(value, active)
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

	rows, err := db.DB.Query("SELECT si.code, si.name, si.active FROM si"+whereString, value...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed_select",
		})
	}

	data := []si{}

	for rows.Next() {
		var result si
		rows.Scan(&result.Code, &result.Name, &result.Active)
		data = append(data, result)
	}

	c.JSON(http.StatusOK, &model.DataResponse{
		Message: "success",
		Data:    data,
	})
}