package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"leave/api"
	"os"
)

type config struct {
	Port int    `yaml:"port"`
	Dsn  string `yaml:"dsn"`
}

var c config

func init() {
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &c)
}

func main() {
	api.InitDb(c.Dsn)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api.InitApi(r)
	err := r.Run(fmt.Sprintf(":%d", c.Port))
	if err != nil {
		panic(err)
	}
}
