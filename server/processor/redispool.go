package processor

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//初始化连接池
func InitPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) *redis.Pool {

	myPool := &redis.Pool{
		//最大空闲连接数
		MaxIdle: maxIdle,
		//最大连 接数，0表示没有限制
		MaxActive: maxActive,
		//最大空闲时间，当被使用过的连接过了这个时间仍没有被使用，就会被放到连接池中
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	return myPool

}
