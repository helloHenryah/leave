package api

import (
	"context"
	"github.com/glebarez/sqlite"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"leave/types"
	"log"
	"strings"
)

var db *gorm.DB
var redisDB *redis.Client

func InitDb(dsn string, type_ string) (err error) {
	type_ = strings.ToLower(type_)
	log.Println(dsn, type_)
	switch type_ {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn))
	case "postgres", "pgsql":
		db, err = gorm.Open(postgres.Open(dsn))
	case "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open(dsn))
	default:
		if dsn != "" {
			db, err = gorm.Open(sqlite.Open(dsn))
		} else {
			db, err = gorm.Open(sqlite.Open("file.db"))
			log.Println("use default sqlite and file.db")
		}
	}
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	err = db.AutoMigrate(&types.User{}, &types.Submit{})
	if err != nil {
		panic("failed to migrate database" + err.Error())
	}
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	result, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect redis" + err.Error())
	}
	log.Println("redis:", result)
	return nil
}
