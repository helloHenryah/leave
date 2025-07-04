package api

import (
	"github.com/gin-gonic/gin"
	"leave/types"
)

func (a *Api) Register(c *gin.Context) {
	var user = new(types.Register)
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var u = new(types.User)
	err := db.Where(&types.User{StudentID: user.StudentID}).First(u).Error
	if err != nil {
		u = &types.User{
			Name:              user.Name,
			Sex:               user.Sex,
			StudentID:         user.StudentID,
			Phone:             user.Phone,
			ParentName:        user.ParentName,
			ParentPhone:       user.ParentPhone,
			DormitoryID:       user.DormitoryID,
			TeacherName:       user.TeacherName,
			TeacherDepartment: user.TeacherDepartment,
			ClassName:         user.ClassName,
			Year:              user.Year,
		}
		if user.Sex == "男" {
			u.DormitoryName = "金川2号楼"
		} else {
			u.DormitoryName = "金川南苑A座公寓"
		}

		err = db.Create(u).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"redirect": "/", "student_id": user.StudentID})
	} else {
		u = &types.User{
			ID:                u.ID,
			Name:              user.Name,
			Sex:               user.Sex,
			StudentID:         user.StudentID,
			Phone:             user.Phone,
			ParentName:        user.ParentName,
			ParentPhone:       user.ParentPhone,
			DormitoryID:       user.DormitoryID,
			TeacherName:       user.TeacherName,
			ClassName:         user.ClassName,
			Year:              user.Year,
			TeacherDepartment: user.TeacherDepartment,
		}
		if user.Sex == "男" {
			u.DormitoryName = "金川2号楼"
		} else {
			u.DormitoryName = "金川南苑A座公寓"
		}
		err = db.Updates(u).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"redirect": "/", "student_id": user.StudentID})
	}
}
