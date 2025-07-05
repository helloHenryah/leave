package api

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"leave/types"
	"strings"
)

var db *gorm.DB

func InitDb(dsn string, type_ string) (err error) {
	type_ = strings.ToLower(type_)
	fmt.Println(dsn, type_)
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
			fmt.Println("use default sqlite and file.db")
		}
	}
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	err = db.AutoMigrate(&types.User{}, &types.Submit{})
	if err != nil {
		panic("failed to migrate database" + err.Error())
	}
	return nil
}
