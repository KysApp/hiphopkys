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

/**
 * 缓存中添加预约用户
 * @param  {[type]} model *models.AppointmentPlayerCacheModel) (string, string [description]
 * @return {(errorCode,desc)}       [description]
 */
func CachePushAppointmentUser(model *models.AppointmentPlayerCacheModel) (string, string) {
	conn := RedisPool.Get()
	defer conn.Close()
	max, err := beego.AppConfig.Int("room::capacity")
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	//获取数目
	reply, err := conn.Do("SCARD", model.AppointmentId)
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	count, err := redis.Int(reply, nil)
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	if count >= max {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_roomfull_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_roomfull_desc")
	}
	buffer, err := model.Marshal()
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	_, err = conn.Do("SADD", model.AppointmentId, buffer)
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	return "0", "操作成功"
}
