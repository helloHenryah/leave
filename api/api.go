package api

import (
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Api struct {
	*gin.Engine
	TemplateFS embed.FS
}

func InitApi(r *gin.Engine, fs embed.FS) {
	// 将 embed.FS 转换为 *template.Template
	tmpl := template.Must(template.New("").ParseFS(fs, "templates/*.html"))

	// 设置 Gin 使用嵌入的模板
	r.SetHTMLTemplate(tmpl)
	api := Api{
		Engine:     r,
		TemplateFS: fs,
	}
	api.GET("/", func(c *gin.Context) {
		c.HTML(200, "submit.html", nil)
	})
	api.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})
	api.GET("/table/user", func(c *gin.Context) {
		c.HTML(200, "table_user.html", nil)
	})
	api.GET("/table/submit", func(c *gin.Context) {
		c.HTML(200, "table_submit.html", nil)
	})

	api.GET("/html/:id", func(c *gin.Context) {
		c.HTML(200, "image.html", gin.H{"url": "/api/upload/" + c.Param("id")})
	})
	api.GET("/api/upload/:id", func(c *gin.Context) {
		result, err := redisDB.Get(context.Background(), fmt.Sprintf("image_%s", c.Param("id"))).Result()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.Data(200, "image/png", []byte(result))
	})
	api.GET("/api/download/:id", func(c *gin.Context) {
		result, err := redisDB.Get(context.Background(), fmt.Sprintf("image_%s", c.Param("id"))).Result()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.Header("Content-Disposition", "attachment; filename=病假条.png")
		c.Header("Content-Type", "application/octet-stream")
		c.Data(200, "image/png", []byte(result))
	})
	api.GET("/api/user", api.GetUserInfo)
	api.GET("/api/submit", api.GetSubmitInfo)
	api.POST("/api/register", api.Register)
	api.POST("/api/submit", api.Submit)

}
