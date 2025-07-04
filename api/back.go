package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leave/types"
)

func (a *Api) GetUserInfo(c *gin.Context) {
	var user = new([]*types.User)
	if err := db.Find(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user, "total": len(*user)})
}

func (a *Api) GetSubmitInfo(c *gin.Context) {
	var submit []types.SubmitRespItem
	if err := db.Table("submits").
		Joins("LEFT JOIN users ON users.student_id = submits.student_id").
		Select("submits.*, users.name,users.student_id, users.sex, users.phone, users.parent_name, users.parent_phone, users.dormitory_id, users.teacher_id, users.teacher_name").
		Order("submits.id DESC").
		Scan(&submit).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": submit, "total": len(submit)})
}
