package service

func (service *WebService) QueryRedisClientInfo() string {
	return service.dao.QueryRedisClientInfo()
}
