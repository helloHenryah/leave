package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leave/types"
)

func GetUserInfo(c *gin.Context) {
	var user = new([]*types.User)
	if err := db.Find(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user, "total": len(*user)})
}

func GetSubmitInfo(c *gin.Context) {
	var submit = new([]*types.Submit)
	if err := db.Find(&submit).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": submit, "total": len(*submit)})
}
