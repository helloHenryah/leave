package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leave/types"
	"math/rand/v2"
	"time"
)

func Submit(c *gin.Context) {
	var submit = new(types.Submit)
	if err := c.ShouldBind(&submit); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if submit.Lesson == "" {
		submit.Lesson = "无"
	}
	if submit.WhereTo == "" {
		submit.WhereTo = "学校"
	}
	db.Create(&submit)
	data := make(map[string]string)
	var user types.User
	err := db.Where(&types.User{StudentID: submit.StudentID}).Last(&user).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "用户不存在"})
		return
	}
	startTime, err := time.Parse("2006-01-02T15:04", submit.StartTime)
	if err != nil {
		fmt.Println(err, submit.StartTime)
		c.JSON(400, gin.H{"error": "开始时间格式错误"})
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04", submit.EndTime)
	if err != nil {
		fmt.Println(err, submit.EndTime)
		c.JSON(400, gin.H{"error": "结束时间格式错误"})
		return
	}
	if startTime.After(endTime) {
		c.JSON(400, gin.H{"error": "开始时间不能大于结束时间"})
		return
	}
	reviewTime, err := time.Parse("2006-01-02T15:04", submit.ReviewTime)
	if err != nil {
		fmt.Println(err, submit.ReviewTime)
		c.JSON(400, gin.H{"error": "审批时间格式错误"})
		return
	}
	today := time.Now().Format("2006-01-02")
	reviewDate := reviewTime.Format("2006-01-02")
	approvalTime, err := time.Parse("2006-01-02T15:04", submit.ApprovalTime)
	if err != nil {
		fmt.Println(err, submit.ApprovalTime)
		c.JSON(400, gin.H{"error": "编号时间格式错误"})
		return
	}
	if approvalTime.After(reviewTime) {
		c.JSON(400, gin.H{"error": "审批时间不能大于编号时间"})
		return
	}

	data["approval_id"] = fmt.Sprintf("%s%08d", approvalTime.Format("200601021504"), rand.IntN(59000000))
	if today == reviewDate {
		data["review_time"] = reviewTime.Format("15:04")
	} else {
		data["review_time"] = reviewTime.Format("2006-01-02 15:04")
	}

	data["name"] = user.Name
	data["student_id"] = submit.StudentID
	data["sex"] = user.Sex
	data["phone"] = user.Phone
	data["parent_name"] = user.ParentName
	data["parent_phone"] = user.ParentPhone
	data["dormitory_id"] = user.DormitoryID
	data["dormitory_name"] = user.DormitoryName
	data["teacher_id"] = user.TeacherID
	data["teacher_name"] = user.TeacherName
	data["img"] = user.Img
	data["start_time"] = startTime.Format("2006-01-02 15:04")
	data["end_time"] = endTime.Format("2006-01-02 15:04")
	data["during_time"] = fmt.Sprintf("%.2f", endTime.Sub(startTime).Minutes()/1440)
	data["lesson"] = submit.Lesson
	data["information"] = submit.Information
	if submit.LeaveType == "leave_out" {
		data["is_in_dormitory"] = "否"
		data["is_leave_school"] = "是"
		data["reason"] = "事假,校外办理"
		data["where_to"] = submit.WhereTo
		data["leave_type"] = "事假"
	} else if submit.LeaveType == "leave_in" {
		data["is_in_dormitory"] = "是"
		data["is_leave_school"] = "否"
		data["reason"] = "事假,校外办理"
		data["where_to"] = submit.WhereTo
		data["leave_type"] = "事假"
	} else if submit.LeaveType == "sick_in" {
		data["is_in_dormitory"] = "是"
		data["is_leave_school"] = "否"
		data["reason"] = "病假,校内办理"
		data["where_to"] = submit.WhereTo
		data["leave_type"] = "病假"
	} else if submit.LeaveType == "sick_out" {
		data["is_in_dormitory"] = "否"
		data["is_leave_school"] = "是"
		data["reason"] = "病假,校外办理"
		data["where_to"] = submit.WhereTo
		data["leave_type"] = "病假"
	}
	c.HTML(200, "index.html", data)
}
