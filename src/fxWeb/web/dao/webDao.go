package dao

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"leiax00.com/fxWeb/fw/conf"
)

type WebDao struct {
	conf *conf.Config
	db   *gorm.DB
	rc   *redis.Client
}

func NewWebDao(conf *conf.Config, client *redis.Client) *WebDao {
	db, _ := gorm.Open("postgres", "host=localhost port=5432 user=leiax dbname=postgres password=")
	return &WebDao{conf: conf, db: db, rc: client}
}

func (dao WebDao) NewClient() *gorm.DB {
	//_ "github.com/jinzhu/gorm/dialects/postgres" gorm库依赖的pg驱动
	//_ "github.com/lib/pq" beego库依赖的pg驱动
	//orm.RegisterDataBase("postgres", "postgres", "leiax:@tcp(127.0.0.1:5432)/postgres?charset=utf-8", 30)
	//orm.RegisterModel(new(bo.User))
	//orm.RunSyncdb("user", false, true)
	//db, _ := gorm.Open("postgres", "host=localhost port=5432 user=leiax dbname=postgres password=")
	//user := bo.User{Name: "leiax01", Age: 18, Sex: 2}
	//db.CreateTable(&user)
	//db.Create(&user).NewRecord(user)
	//db.Close()
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=leiax dbname=postgres password=")
	if err != nil {
		log.Error(err.Error())
	}
	return db
}
