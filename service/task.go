package service

import (
	"TodoList/model"
	"TodoList/serializer"
	"fmt"
	"strconv"
	"time"
)

// 创建一个结构体，接收前端的参数
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未完成，1是已完成
}

// 查询时需要一个空的结构体
type ShowTaskService struct {
}

// 页数码，一页显示几条，做个分页
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`   //页数码
	PageSize int `json:"page_size" form:"page_size"` //每页最多条数
}

// 更新
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

// 模糊查询
type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`   //页数码
	PageSize int    `json:"page_size" form:"page_size"` //每页最多条数
}

// 删除备忘录
type DeleteTaskService struct {
}

// 创建备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	code := 200
	task := model.Task{
		// User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    service.Status,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500 //失败
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// 查询备忘录的方法
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功",
	}
}

// 列表返回用户所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	fmt.Println("页数码====", service.PageNum)
	// service.PageNum = 2
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	fmt.Println("页数码====", service.PageNum)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 列表返回用户所有备忘录
func (service *ListTaskService) List1(uid uint, PageNum, PageSize string) serializer.Response {
	var tasks []model.Task
	count := 0
	fmt.Println("num&size======", PageNum, PageSize)
	PageSize1, _ := strconv.Atoi(PageSize)
	PageNum1, _ := strconv.Atoi(PageNum)

	if PageSize1 == 0 {
		PageSize1 = 10
	}
	fmt.Println("num1&size1======", PageNum1, PageSize1)

	fmt.Println("页数码====", PageNum1)
	// service.PageNum = 2
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(PageSize1).
		Offset((PageNum1 - 1) * PageSize1).Find(&tasks)
	fmt.Println("页数码====", PageNum1)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 更新
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	code := 200
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		code = 500 //失败
		return serializer.Response{
			Status: code,
			Msg:    "修改失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "修改成功",
	}
}

// 模糊查询
func (service *SearchTaskService) Search(uid uint) serializer.Response {

	var tasks []model.Task
	code := 200
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10 //每页显示10条
	}

	fmt.Println("info====", service.Info)
	fmt.Println("PageNum====", service.PageNum)
	fmt.Println("PageSize====", service.PageSize)
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		code = 500 //失败
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 删除
func (delete *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
