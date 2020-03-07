package database

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var RD *redis.Pool

func InitRedis(){
	host:=beego.AppConfig.String("redisHost")
	pwd:=beego.AppConfig.String("redisPWD")
	MaxIdle:=beego.AppConfig.DefaultInt("redistMaxID",100)
	MaxActive:=beego.AppConfig.DefaultInt("redisMaxAct",100)
	RD=&redis.Pool{
		// Dial()方法返回一个连接，从在需要创建连接到的时候调用
		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",host,redis.DialPassword(pwd))
			if err!=nil{
				return nil, err
			}
			return c,nil
		},
		// 最大空闲连接数
		MaxIdle:MaxIdle,
		// 一个pool所能分配的最大的连接数目
		// 当设置成0的时候，该pool连接数没有限制
		MaxActive:MaxActive,
		// 空闲连接超时时间，超过超时时间的空闲连接会被关闭。
		IdleTimeout: 10*time.Minute,
	}
}
