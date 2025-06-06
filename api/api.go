package api

import "github.com/gin-gonic/gin"

func InitApi(r *gin.Engine) {
	r.LoadHTMLFiles("templates/index.html",
		"templates/register.html",
		"templates/submit.html",
		"templates/login.html",
		"templates/table.html",
	)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "submit.html", nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	r.GET("/table", func(c *gin.Context) {
		c.HTML(200, "table.html", nil)
	})
	r.GET("/api/upload/:file", func(c *gin.Context) {
		c.File("upload/" + c.Param("file"))
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.GET("/api/img", func(c *gin.Context) {
		c.File("upload/img.jpg")
	})
	api := r.Group("/api")
	{
		api.GET("/user", GetUserInfo)
		api.GET("/submit", GetSubmitInfo)
		api.POST("/register", Register)
		api.POST("/login", Login)
		api.POST("/submit", Submit)
	}

}
