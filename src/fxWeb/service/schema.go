package service

func (service WebService) CreateTable() error {
	return service.dao.CreateTable()
}

func (service WebService) DeleteTable() error {
	return service.dao.DeleteTable()
}
