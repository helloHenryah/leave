package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io"
	"leave/api"
	"os"
)

//go:embed templates/*
var TemplateFS embed.FS

func main() {

	file, err := os.Open("setting.yaml")
	if err != nil {
		panic(err)
	}
	var m map[string]string
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &m)
	file.Close()

	api.InitDb(m["Dsn"])
	r := gin.Default()
	api.InitApi(r, TemplateFS)
	err = r.Run(fmt.Sprintf(":%s", m["Port"]))
	if err != nil {
		panic(err)
	}
}
