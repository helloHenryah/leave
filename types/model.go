package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Sex           string `json:"sex"`
	StudentID     string `json:"student_id"`
	Phone         string `json:"phone" `
	ParentName    string `json:"parentName"`
	ParentPhone   string `json:"parentPhone"`
	DormitoryID   string `json:"dormitoryID"`
	DormitoryName string `json:"dormitory_name"`
	Img           string `json:"img"`
	TeacherName   string `json:"teacherName"`
	TeacherID     string `json:"teacherID"`
	Password      string `json:"password"`
}
