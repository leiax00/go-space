package fwRedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go.uber.org/fx"
	"time"
)

func NewFwRedisModule() fx.Option {
	return fx.Options(
		fx.Provide(newRedisConf),
		fx.Provide(newRedisClient),
		fx.Populate(&fwRedisClient),
		fx.Invoke(registerRedisHook),
	)
}

func newRedisConf() *redis.Options {
	timeout := 20 * time.Second
	return &redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		PoolSize:     10,
		PoolTimeout:  timeout + 1,
	}
}

func newRedisClient(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

func registerRedisHook(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			fmt.Println("Close redis client....")
			return fwRedisClient.Close()
		},
	})
}
