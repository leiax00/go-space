package dao

import "leiax00.com/fxWeb/bean/bo"

func (dao *WebDao) CreateTable() error {
	dao.db.CreateTable(&bo.User{})
	return nil
}

func (dao *WebDao) DeleteTable() error {
	dao.db.DropTableIfExists(&bo.User{})
	return nil
}
