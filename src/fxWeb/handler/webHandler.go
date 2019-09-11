package handler

import (
	"leiax00.com/fxWeb/conf"
)

type WebHandler struct {
	Conf *conf.Config
}

func NewWebHandler(conf *conf.Config) *WebHandler {
	return &WebHandler{Conf: conf}
}
