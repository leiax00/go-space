package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (handler *WebHandler) HelloApp(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"result": "Hello, APP！！"})
}
