package handler

import (
	"github.com/labstack/echo/v4"
	"leiax00.com/fxWeb/fxEcho"
	"net/http"
)

func (handler *WebHandler) HelloApp(c echo.Context) error {
	cc := c.(*fxEcho.FxEchoContext)
	cc.Pre()
	return c.JSON(http.StatusOK, map[string]string{"result": "Hello, FX-WEB！！"})
}
