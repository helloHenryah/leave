package api

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"html/template"
	"leave/types"
	"log"
	"strings"
	"time"
)

func (a *Api) Submit(c *gin.Context) {
	var submit = new(types.Submit)
	if err := c.Bind(&submit); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"error": err.Error()})
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
		log.Println(err)
		c.JSON(200, gin.H{"error": "用户不存在"})
		return
	}
	startTime, err := time.Parse("2006-01-02T15:04", submit.StartTime)
	if err != nil {
		log.Println(err, submit.StartTime)
		c.JSON(200, gin.H{"error": "开始时间格式错误"})
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04", submit.EndTime)
	if err != nil {
		log.Println(err, submit.EndTime)
		c.JSON(200, gin.H{"error": "结束时间格式错误"})
		return
	}

	if startTime.After(endTime) {
		log.Println("开始时间不能大于结束时间", submit.StartTime, submit.EndTime)
		c.JSON(200, gin.H{"error": "开始时间不能大于结束时间"})
		return
	}

	if endTime.Sub(startTime) > time.Hour*24 {
		log.Println("请假时间不能超过24小时", submit.StartTime, submit.EndTime)
		c.JSON(200, gin.H{"error": "请假时间不能超过24小时"})
		return
	}

	reviewTime, err := time.Parse("2006-01-02T15:04", submit.ReviewTime)
	if err != nil {
		log.Println(err, submit.ReviewTime)
		c.JSON(200, gin.H{"error": "审批时间格式错误"})
		return
	}
	approvalTime, err := time.Parse("2006-01-02T15:04", submit.ApprovalTime)
	if err != nil {
		log.Println(err, submit.ApprovalTime)
		c.JSON(200, gin.H{"error": "编号时间格式错误"})
		return
	}
	if approvalTime.After(reviewTime) {
		log.Println("编号时间不能大于审批时间", submit.ApprovalTime, submit.ReviewTime)
		c.JSON(200, gin.H{"error": "审批时间不能大于编号时间"})
		return
	}

	data["approval_time"] = approvalTime.Format("2006-01-02 15:04:05")
	data["review_time"] = reviewTime.Format("2006-01-02 15:04:05")
	data["name"] = user.Name
	data["student_id"] = submit.StudentID
	data["sex"] = user.Sex
	data["phone"] = user.Phone
	data["year"] = user.Year
	data["class_name"] = user.ClassName
	data["parent_name"] = user.ParentName
	data["parent_phone"] = user.ParentPhone
	data["dormitory_id"] = user.DormitoryID
	data["dormitory_name"] = user.DormitoryName
	data["teacher_name"] = user.TeacherName
	data["teacher_department"] = user.TeacherDepartment
	data["start_time"] = startTime.Format("2006-01-02 15:04")
	data["end_time"] = endTime.Format("2006-01-02 15:04")
	data["during_time"] = "1"
	data["lesson"] = submit.Lesson
	data["information"] = submit.Information
	if user.ClassName == "计算机科学与技术" {
		data["college"] = "智能科学与技术学院（网络空间安全学院）"
	} else {
		data["college"] = "信息工程学院"
	}

	if submit.LeaveType == "leave_out" {
		data["is_in_dormitory"] = "否"
		data["is_leave_school"] = "是"
		data["where_to"] = submit.WhereTo
	} else if submit.LeaveType == "leave_in" {
		data["is_in_dormitory"] = "是"
		data["is_leave_school"] = "否"
		data["where_to"] = "学校"
	}
	data["trip_mode"] = submit.TripMode
	//c.HTML(200, "index_out.html", data)
	if submit.Action == "html" {
		c.HTML(200, "index_out.html", data)
	} else if submit.Action == "image" {
		a.Template(c, data)
	}
}

func (a *Api) Template(c *gin.Context, data map[string]string) {

	// 手动读取 HTML 文件并渲染模板
	filePath := "templates/index_out.html" // 替换为实际的文件路径
	content, err := a.TemplateFS.ReadFile(filePath)
	if err != nil {
		log.Println("无法读取文件:", err)
		c.JSON(500, gin.H{"error": "无法读取文件"})
		return
	}

	tmpl, err := template.New("index").Parse(string(content))
	if err != nil {
		log.Println("解析模板失败:", err)
		c.JSON(500, gin.H{"error": "解析模板失败"})
		return
	}

	var htmlOutput strings.Builder
	if err := tmpl.Execute(&htmlOutput, data); err != nil {
		log.Println("渲染模板失败:", err)
		c.JSON(500, gin.H{"error": "渲染模板失败"})
		return
	}
	bytes := fullScreenshot(htmlOutput.String())
	//c.Header("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	//c.Header("Pragma", "no-cache")                                   // HTTP 1.0.
	//c.Header("Expires", "0")                                         // Proxies.
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprintf("%d", len(bytes)))
	c.Header("Content-Disposition", "attachment; filename=fullScreenshot.png")

	c.Writer.Write(bytes)
}

func fullScreenshot(htmlContent string) []byte {

	// 1. 创建有界面的浏览器选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		//chromedp.Flag("headless", false),
		chromedp.Flag("start-maximized", true),
		//chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"),
		chromedp.WindowSize(400, 2300),
	)
	//chromedp.Flag("headless", false),
	//chromedp.Flag("start-maximized", true),
	//chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"),

	// 2. 创建上下文
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, chromedp.Navigate("about:blank"),
		loadHTMLFromStringActionFunc(htmlContent),
		chromedp.Tasks{chromedp.FullScreenshot(&buf, 100)}); err != nil {
		log.Fatal(err)
	}
	//if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
	//	log.Fatal(err)
	//}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
	return buf

}

func loadHTMLFromStringActionFunc(content string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		ch := make(chan bool, 1)
		defer close(ch)

		go chromedp.ListenTarget(ctx, func(ev interface{}) {
			if _, ok := ev.(*page.EventLoadEventFired); ok {
				ch <- true
			}
		})

		frameTree, err := page.GetFrameTree().Do(ctx)
		if err != nil {
			return err
		}

		if err := page.SetDocumentContent(frameTree.Frame.ID, content).Do(ctx); err != nil {
			return err
		}

		select {
		case <-ch:
			return nil
		case <-ctx.Done():
			return context.DeadlineExceeded
		}
	}
}
