package api

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"leave/types"
)

var db *gorm.DB

func InitDb(dsn string) {
	DB, err := gorm.Open(postgres.Open(dsn))
	fmt.Println("connecting to postgres")
	if err != nil {
		DB, err = gorm.Open(sqlite.Open("file.db"))
		if err != nil {
			panic("failed to connect database")
		}
		panic("failed to connect database")
	}

	db = DB
	err = db.AutoMigrate(&types.User{}, &types.Submit{})
	if err != nil {
		panic("failed to migrate database" + err.Error())
	}
}
