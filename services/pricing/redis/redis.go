package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
)

var Client = redis.NewClient(&redis.Options{
	Addr: config.Default.RedisAddr,
})
