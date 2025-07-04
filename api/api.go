package api

import (
	"embed"
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
	api.GET("/api/user", api.GetUserInfo)
	api.GET("/api/submit", api.GetSubmitInfo)
	api.POST("/api/register", api.Register)
	api.POST("/api/submit", api.Submit)

}
