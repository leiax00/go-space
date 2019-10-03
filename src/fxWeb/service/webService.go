package service

import (
	"leiax00.com/fxWeb/conf"
	"leiax00.com/fxWeb/dao"
)

type WebService struct {
	conf *conf.Config
	dao  *dao.WebDao
}

func NewWebService(conf *conf.Config, dao *dao.WebDao) *WebService {
	return &WebService{conf: conf, dao: dao}
}
