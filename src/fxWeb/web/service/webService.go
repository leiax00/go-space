package service

import (
	"leiax00.com/fxWeb/fw/conf"
	"leiax00.com/fxWeb/web/dao"
)

type WebService struct {
	conf *conf.Config
	dao  *dao.WebDao
}

func NewWebService(conf *conf.Config, dao *dao.WebDao) *WebService {
	return &WebService{conf: conf, dao: dao}
}
