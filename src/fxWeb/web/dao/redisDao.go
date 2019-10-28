package dao

import (
	"fmt"
)

func (dao *WebDao) QueryRedisClientInfo() string {
	val := dao.rc.Info("clients").Val()
	fmt.Println(val) // 1
	val = dao.rc.Info("clients").Val()
	fmt.Println(val) // 1 说明会自动关闭连接，多线程时才会多个连接
	return dao.rc.Info("clients").Val()
}
