package ioc

import (
	"github.com/Duke1616/passkey/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	// 这个是假设你有一个独立的 Redis 的配置文件
	return redis.NewClient(&redis.Options{
		Addr: config.C().Redis.Addr,
	})
}
