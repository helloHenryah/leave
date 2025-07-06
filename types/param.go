package types

type Register struct {
	Name              string `form:"name" json:"name"`
	StudentID         string `form:"student_id" json:"student_id"`
	Sex               string `form:"sex" json:"sex"`
	Phone             string `form:"phone" json:"phone"`
	ParentName        string `form:"parent_name" json:"parent_name"`
	ParentPhone       string `form:"parent_phone" json:"parent_phone"`
	DormitoryID       string `form:"dormitory_id" json:"dormitory_id"`
	TeacherName       string `form:"teacher_name" json:"teacher_name"`
	ClassName         string `form:"class_name" json:"class_name"`
	Year              string `form:"year" json:"year"`
	TeacherDepartment string `form:"teacher_department" json:"teacher_department"`
}

type SubmitRespItem struct {
	Submit
	StudentID   string `json:"student_id" gorm:"column:student_id"`
	Name        string `json:"name" gorm:"column:name"`
	Sex         string `json:"sex" gorm:"column:sex"`
	Phone       string `json:"phone" gorm:"column:phone"`
	ParentName  string `json:"parent_name" gorm:"column:parent_name"`
	ParentPhone string `json:"parent_phone" gorm:"column:parent_phone"`
	DormitoryID string `json:"dormitory_id" gorm:"column:dormitory_id"`
	TeacherName string `json:"teacher_name" gorm:"column:teacher_name"`
}
