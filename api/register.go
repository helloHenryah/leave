package api

import (
	"github.com/gin-gonic/gin"
	"leave/types"
)

func Register(c *gin.Context) {
	var user = new(types.Register)
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = c.SaveUploadedFile(file, "upload/"+user.StudentID+".jpg")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var u = new(types.User)
	err = db.Where(&types.User{StudentID: user.StudentID}).First(u).Error
	if err != nil {
		u = &types.User{
			Name:        user.Name,
			Sex:         user.Sex,
			StudentID:   user.StudentID,
			Phone:       user.Phone,
			ParentName:  user.ParentName,
			ParentPhone: user.ParentPhone,
			DormitoryID: user.DormitoryID,
			TeacherID:   user.TeacherID,
			TeacherName: user.TeacherName,
			Img:         "upload/" + user.StudentID + ".jpg",
			Password:    user.Password,
		}
		if user.Sex == "男" {
			u.DormitoryName = "金川校区,北苑2号公寓"
		} else {
			u.DormitoryName = "金川校区,南苑A座公寓(南)"
		}

		err = db.Create(u).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"redirect": "/"})
	} else {
		err = db.Where(&types.User{StudentID: user.StudentID, Password: user.Password}).First(&u).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "用户名或密码错误"})
			return
		}
		u = &types.User{
			Model:       u.Model,
			Name:        user.Name,
			Sex:         user.Sex,
			Phone:       user.Phone,
			ParentName:  user.ParentName,
			ParentPhone: user.ParentPhone,
			DormitoryID: user.DormitoryID,
			TeacherID:   user.TeacherID,
			TeacherName: user.TeacherName,
			Img:         "upload/" + user.StudentID + ".jpg",
		}
		if user.Sex == "男" {
			u.DormitoryName = "金川校区,北苑2号公寓"
		} else {
			u.DormitoryName = "金川校区,南苑A座公寓(南)"
		}
		err = db.Updates(u).Error
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"redirect": "/"})
	}
}

func Login(c *gin.Context) {
	password := c.Query("password")
	studentID := c.Query("student_id")
	var user = new(types.User)
	err := db.Where(&types.User{StudentID: studentID, Password: password}).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "用户名或密码错误"})
		return
	}
	c.JSON(200, gin.H{"name": user.Name, "student_id": user.StudentID, "success": true})
}
