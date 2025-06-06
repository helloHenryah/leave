package types

type Register struct {
	Password    string `form:"password" json:"password"`
	Name        string `form:"name" json:"name"`
	StudentID   string `form:"student_id" json:"student_id"`
	Sex         string `form:"sex" json:"sex"`
	Phone       string `form:"phone" json:"phone"`
	ParentName  string `form:"parent_name" json:"parent_name"`
	ParentPhone string `form:"parent_phone" json:"parent_phone"`
	DormitoryID string `form:"dormitory_id" json:"dormitory_id"`
	TeacherID   string `form:"teacher_id" json:"teacher_id"`
	TeacherName string `form:"teacher_name" json:"teacher_name"`
}

type Submit struct {
	StudentID    string `json:"student_id" form:"student_id"`
	ApprovalTime string `json:"approval_time" form:"approval_time"`
	StartTime    string `json:"start_time" form:"start_time"`
	EndTime      string `json:"end_time" form:"end_time"`
	Lesson       string `json:"lesson" form:"lesson"`
	LeaveType    string `json:"leave_type" form:"leave_type"`
	Information  string `json:"information" form:"information"`
	WhereTo      string `json:"where_to" form:"where_to"`
	ReviewTime   string `json:"review_time" form:"review_time"`
}
