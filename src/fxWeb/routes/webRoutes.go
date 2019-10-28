package routes

import (
	"github.com/labstack/echo/v4"
	"leiax00.com/fxWeb/web/handler"
)

func RegisterWebRoutes(handler *handler.WebHandler, e *echo.Echo) {
	group := e.Group("/web")
	group.GET("/hello", handler.HelloApp)
	group.GET("/db/table/create", handler.CreateTable)
	group.GET("/db/table/delete", handler.DeleteTable)
	group.GET("/redis/client", handler.QueryRedisClientInfo)
	e.Logger.Info("Success to register web routes!!!")
}
