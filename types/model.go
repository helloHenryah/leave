package types

type User struct {
	ID                uint   `gorm:"primarykey" json:"-"`
	Name              string `json:"name"`
	Sex               string `json:"sex"`
	StudentID         string `json:"student_id"`
	Phone             string `json:"phone"`
	ParentName        string `json:"parent_name"`
	ParentPhone       string `json:"parent_phone"`
	DormitoryID       string `json:"dormitory_id"`
	DormitoryName     string `json:"dormitory_name"`
	TeacherName       string `json:"teacher_name"`
	ClassName         string `json:"class_name"`
	Year              string `json:"year"`
	TeacherDepartment string `json:"teacher_department"`
}

type Submit struct {
	ID           uint   `gorm:"primarykey"`
	StudentID    string `json:"student_id" form:"student_id"`
	ApprovalTime string `json:"approval_time" form:"approval_time"`
	StartTime    string `json:"start_time" form:"start_time"`
	EndTime      string `json:"end_time" form:"end_time"`
	Lesson       string `json:"lesson" form:"lesson"`
	LeaveType    string `json:"leave_type" form:"leave_type"`
	Information  string `json:"information" form:"information"`
	TripMode     string `json:"trip_mode" form:"trip_mode"`
	WhereTo      string `json:"where_to" form:"where_to"`
	ReviewTime   string `json:"review_time" form:"review_time"`
}
