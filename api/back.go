package api

import (
	"github.com/gin-gonic/gin"
	"leave/types"
	"log"
	"strconv"
)

func (a *Api) GetUserInfo(c *gin.Context) {
	var user = new([]*types.User)
	if err := db.Find(&user).Error; err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user, "total": len(*user)})
}

func (a *Api) GetSubmitInfo(c *gin.Context) {
	page := c.Query("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		pageNum = 1
	}
	var submit []types.SubmitRespItem
	if err := db.Table("submits").
		Joins("LEFT JOIN users ON users.student_id = submits.student_id").
		Select("submits.*, users.name, users.sex, users.phone, users.parent_name, users.parent_phone, users.dormitory_id, users.teacher_name").
		Order("submits.id DESC").
		Offset((pageNum - 1) * 50).
		Limit(50).
		Scan(&submit).Error; err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": submit, "total": len(submit)})
}
