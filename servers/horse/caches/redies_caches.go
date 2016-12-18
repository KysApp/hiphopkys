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
	conn.Send("MULTI")
	conn.Do("SADD", model.AppointmentId, buffer)
	conn.Do("SET", model.UserId, model.AppointmentId)
	err = conn.Send("EXEC")
	if err != nil {
		return beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_code"), beego.AppConfig.String("errcode::cache_push_appointmentuser_error_inner_desc")
	}
	// setKey := beego.AppConfig.String("cache::key-appointment-id")
	// beego.BeeLogger.Debug("sefKey:%s", setKey)
	return "0", "success"
}

/**
 * 查询是否是预约玩家
 * @param {[type]} userId string) (bool, *models.AppointmentPlayerCacheModel [description]
 */
func CachePullAppointmentUser(userId string) (bool, *models.AppointmentPlayerCacheModel) {
	conn := RedisPool.Get()
	defer conn.Close()
	reply, err := conn.Do("GET", userId)
	if nil != err {
		beego.BeeLogger.Error("redies查找失败1:%s", err.Error())
		return false, nil
	}
	appointmentId, err := redis.String(reply, nil)
	if nil != err {
		beego.BeeLogger.Error("redies查找失败2:%s,%#v,userid:%s", err.Error(), reply, userId)
		return false, nil
	}

	reply, err = conn.Do("SMEMBERS", appointmentId)
	if nil != err {
		beego.BeeLogger.Error("redies查找失败3:%s,%#v,userid:%s,appointmentId:%s", err.Error(), reply, userId, appointmentId)
		return false, nil
	}

	byteBufferArray, err := redis.ByteSlices(reply, nil)
	if nil != err {
		beego.BeeLogger.Error("redies查找失败3:%s,%#v,userid:%s,appointmentId:%s", err.Error(), reply, userId, appointmentId)
		return false, nil
	}
	for _, buffer := range byteBufferArray {
		player := &models.AppointmentPlayerCacheModel{}
		if err := player.Unmarshal(buffer); err == nil {
			if player.UserId == userId {
				return true, player
			}
		} else {
			beego.BeeLogger.Error("redies查找失败6:%s", appointmentId)
		}
	}
	return false, nil
}
