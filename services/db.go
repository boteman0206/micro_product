package services

import (
	"errors"
	"fmt"
	"micro_product/config"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine      *xorm.Engine
	redisHandle *redis.Client
)

func SetupDB() {
	redisHandle = GetRedisConn()

	engine = NewDbConn()
}

func CloseDB() {
	engine.Close()
	redisHandle.Close()
}

//获取redis集群客户端
func GetRedisConn() *redis.Client {
	var addr, pwd string
	var db int
	if config.ConfigRes.Redis != nil {
		// 后期加配置
		redisini := config.ConfigRes.Redis
		redisAddr := fmt.Sprintf("%s:%d", redisini.Addr, redisini.Port)
		addr = redisAddr
		pwd = redisini.Pwd
		db = redisini.Db
	} else {
		addr = "127.0.0.1:6379"
		pwd = "admin1234"
		db = 0
	}

	if redisHandle != nil {
		_, err := redisHandle.Ping().Result()
		if err == nil {
			return redisHandle
		} else {
			redisHandle = redis.NewClient(&redis.Options{
				Addr:         addr,
				Password:     pwd,
				DB:           db,
				MinIdleConns: 28,
				IdleTimeout:  30,
				PoolSize:     512,
				MaxConnAge:   30 * time.Second,
			})
			return redisHandle
		}
	}

	redisHandle = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	return redisHandle
}

// NewDbConn 连接池，请勿关闭
func NewDbConn(dataSourceName ...string) *xorm.Engine {
	var err error

	if engine != nil {
		if err = engine.DB().Ping(); err == nil {
			return engine
		}
		return engine
	}

	var mySqlStr string
	if len(config.ConfigRes.Mysql.MysqlStr) > 0 {
		mySqlStr = config.ConfigRes.Mysql.MysqlStr
	} else {
		mySqlStr = "root:admin123456@(127.0.0.1:3306)/micro_product?charset=utf8mb4"
	}

	engine, err = xorm.NewEngine("mysql", mySqlStr)
	if err != nil {
		panic(GetDBError(err))
	}
	engine.ShowSQL()
	if config.IsDebug {
		engine.ShowSQL()
		engine.ShowExecTime()
	}

	//空闲关闭时间
	engine.SetConnMaxLifetime(120 * time.Second)
	//最大空闲连接
	engine.SetMaxIdleConns(10)
	//最大连接数
	engine.SetMaxOpenConns(1000)

	engine.SetTZLocation(time.Local)

	return engine
}

func GetDBError(err error) error {
	return errors.New("数据库操作失败：" + err.Error())
}
