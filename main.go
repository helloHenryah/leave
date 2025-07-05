package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"leave/api"
	"log"
	"os"
)

//go:embed templates/*
var TemplateFS embed.FS

func main() {

	dsn := os.Getenv("LEAVE_APP_DB_DSN")
	type_ := os.Getenv("LEAVE_APP_DB_TYPE")

	port := os.Getenv("LEAVE_APP_PORT")
	if type_ == "" {
		type_ = "sqlite"
		log.Println("use default db type sqlite")
	}
	if port == "" {
		port = "8080"
		log.Println("use default port 8080")
	}
	err := api.InitDb(dsn, type_)
	if err != nil {
		panic("init db err" + err.Error())
	}
	r := gin.Default()
	api.InitApi(r, TemplateFS)
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("init port err" + err.Error())
	}
}
