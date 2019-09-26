package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

// @Summary Swagger API 文档测试
// @Produce  json
// @Param page query int true "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {string} json "{"code":200,"page":1,"pageSize":10}"
// @Failure 500 {string} json "{"errMessage":"内部错误"}"
// @Router /otc/im/swag [get]
func Swagger(c echo.Context) error {
	type HttpTestSwagger struct {
		Code     int `json:"code"`
		Page     int `form:"page" json:"page"`
		PageSize int `form:"pageSize" json:"pageSize"`
	}
	httpTestSwagger := HttpTestSwagger{}
	if err := c.Bind(&httpTestSwagger); err != nil {
		log.Error("bind err: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errMessage": "内存错误",
		})
	}
	log.Info("visit swagger")
	return c.JSON(http.StatusOK, &httpTestSwagger)
}
