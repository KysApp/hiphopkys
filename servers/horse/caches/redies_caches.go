package caches

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"hiphopkys/servers/horse/models"
	"sync"
	"time"
)

var (
	SyncOnce  sync.Once
	RedisPool *redis.Pool
)

func init() {
	SyncOnce.Do(func() {
		server := beego.AppConfig.String("cache::redis-server")
		password := beego.AppConfig.String("cache::redis-password")
		RedisPool = newPool(server, password)
	})
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
