// Date:   Thu Jun 12 23:16:38 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package dal

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/comment/conf"
	"github.com/xiaoqixian/v2ex/backend/app/common/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
	Redis *redis.Client
)

func Init() {
	c := conf.GetConf()
	initMysql(&c.MySQL)
	initRedis(&c.Redis)
}

func initMysql(c *conf.MySQLConfig) {
	// connect to mysql db
	parseTime := "True"
	if !c.ParseTime {
		parseTime = "False"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", 
		c.User,
		c.Password,
		util.GetEnv("MYSQLADDR", "localhost"),
		c.Port,
		c.DBName,
		c.Charset,
		parseTime,
		c.Loc,
	)

	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func initRedis(c *conf.RedisConfig) {
	// connect to redis
	Redis = redis.NewClient(&redis.Options {
		Addr:     fmt.Sprintf("%s:6379", util.GetEnv("REDISADDR", "localhost")),
		Password: c.Password,
		DB:       c.DB,
	})
}
