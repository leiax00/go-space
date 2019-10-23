package fwRedis

import "github.com/go-redis/redis/v7"

var fwRedisClient *redis.Client

func GetClient() *redis.Client {
	return fwRedisClient
}
