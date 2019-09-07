package basic

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"leiax00.com/fxWeb/handler"
)

var EchoObj *echo.Echo

func NewEchoModule() fx.Option {
	EchoObj = echo.New()
	return fx.Provide(EchoObj)
}

func RegisterRoutes(handler *handler.WebHandler, e *echo.Echo) {
	group := e.Group("web")
	group.GET("/hello", handler.HelloApp)
}
