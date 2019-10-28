package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *WebHandler) QueryRedisClientInfo(c echo.Context) error {
	info := handler.service.QueryRedisClientInfo()
	return c.String(http.StatusOK, info)
}
