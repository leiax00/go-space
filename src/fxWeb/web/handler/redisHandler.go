package handler

import (
	"github.com/labstack/echo/v4"
	"leiax00.com/fxWeb/fw/bean/bo"
	"net/http"
)

func (handler *WebHandler) ModifyRedisConf(c echo.Context) error {
	conf := &bo.RedisConf{}
	if err := c.Bind(conf); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if pong, err := handler.service.ModifyRedisConf(conf); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, pong)
	}
}
func (handler *WebHandler) QueryRedisClientInfo(c echo.Context) error {
	info := handler.service.QueryRedisClientInfo()
	return c.String(http.StatusOK, info)
}
