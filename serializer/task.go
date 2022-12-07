package serializer

import "TodoList/model"

//序列化备忘录
type Task struct {
	ID      uint   `json:"id" example:"1"`         //备忘录id
	Title   string `json:"title" example:"学习"`     //标题
	Content string `json:"content" example:"GO编程"` //内容
	// View uint64 `json:"view" example"322"` //浏览量
	Status    int   `json:"status" example:"0"` //状态0未完成，1已完成
	CreatedAt int64 `json:"created_at"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
