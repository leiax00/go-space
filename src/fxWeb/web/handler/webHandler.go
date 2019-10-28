package handler

import (
	"go.uber.org/fx"
	"leiax00.com/fxWeb/fw/conf"
	"leiax00.com/fxWeb/web/service"
)

type WebHandler struct {
	conf    *conf.Config
	service *service.WebService
}

type webHandlerIn struct {
	fx.In
	Conf    *conf.Config
	Service *service.WebService `optional:"true"` // optional:"true", 当无法注入时会设置为零值
}

func NewWebHandler(in webHandlerIn) *WebHandler {
	return &WebHandler{conf: in.Conf, service: in.Service}
}
