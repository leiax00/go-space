package bo

import (
	"github.com/go-redis/redis/v7"
	"net"
	"strconv"
	"strings"
	"time"
)

type RedisConf struct {
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	Password     string        `json:"password"`
	DialTimeout  time.Duration `json:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	PoolSize     int           `json:"pool_size"`
	PoolTimeout  time.Duration `json:"pool_timeout"`
}

func (conf *RedisConf) FillOpts(opts *redis.Options) {
	if strings.TrimSpace(conf.Host) != "" && 0 < conf.Port && conf.Port <= 65535 {
		opts.Addr = net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port))
	}
	if strings.TrimSpace(conf.Password) != "" {
		opts.Password = conf.Password
	}
	if conf.DialTimeout != 0 {
		opts.DialTimeout = conf.DialTimeout
		opts.PoolTimeout = conf.PoolTimeout + 1
	}
	if conf.ReadTimeout != 0 {
		opts.ReadTimeout = conf.ReadTimeout
	}
	if conf.WriteTimeout != 0 {
		opts.WriteTimeout = conf.WriteTimeout
	}
	if conf.PoolSize != 0 {
		opts.PoolSize = conf.PoolSize
	}
	if conf.PoolTimeout != 0 {
		opts.PoolTimeout = conf.PoolTimeout
	}
}
