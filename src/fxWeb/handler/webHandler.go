package handler

import (
	"leiax00.com/fxWeb/conf"
	"leiax00.com/fxWeb/service"
)

type WebHandler struct {
	conf    *conf.Config
	service *service.WebService
}

func NewWebHandler(conf *conf.Config, service *service.WebService) *WebHandler {
	return &WebHandler{conf: conf, service: service}
}
