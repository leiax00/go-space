package service

import (
	"leiax00.com/fxWeb/fw/bean/bo"
	"leiax00.com/fxWeb/fw/fwRedis"
)

func (service *WebService) ModifyRedisConf(conf *bo.RedisConf) (string, error) {
	client := fwRedis.GetClient()
	conf.FillOpts(client.Options())
	return client.Ping().Result()
}

func (service *WebService) QueryRedisClientInfo() string {
	return service.dao.QueryRedisClientInfo()
}
