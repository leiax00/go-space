package web

import (
	"go.uber.org/fx"
	"leiax00.com/fxWeb/routes"
	"leiax00.com/fxWeb/web/dao"
	"leiax00.com/fxWeb/web/handler"
	"leiax00.com/fxWeb/web/service"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			handler.NewWebHandler,
			service.NewWebService,
			dao.NewWebDao,
		),
		fx.Invoke(
			routes.RegisterWebRoutes,
		),
	)
}
