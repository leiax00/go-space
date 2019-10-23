package fwRedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go.uber.org/fx"
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
	return nil
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
