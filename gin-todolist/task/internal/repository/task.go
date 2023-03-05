package repository

import (
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
	"gin-learn/gin-todolist/task/internal/service"
)

type Task struct {
	TaskId    uint `gorm:"primarykey"` // id
	UserId    uint `gorm:"index"`      // 用户id
	Status    int  `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StratTime int64
	EndTime   int64
}

func (t *Task) Show(req *service.TaskRequest) (taskList []Task, err error) {
	err = DB.Model(Task{}).Where("user_id=?", req.UserId).Find(&taskList).Error
	return taskList, err
}

func (t *Task) Create(req *service.TaskRequest) error {
	task := &Task{
		UserId:    uint(req.UserId),
		Status:    int(req.Status),
		Title:     req.Title,
		Content:   req.Content,
		StratTime: int64(req.StartTime),
		EndTime:   int64(req.EndTime),
	}
	if err := DB.Create(&task).Error; err != nil {
		utils.LogrusObj.Error("Insert task failed, err: " + err.Error())
		return err
	}
	return nil
}

func (t *Task) Delete(req *service.TaskRequest) error {
	err := DB.Where("task_id=?", req.TaskId).Delete(Task{}).Error
	return err
}

func (t *Task) Update(req *service.TaskRequest) error {
	task := &Task{}
	err := DB.Where("task_id=?", req.TaskId).Find(task).Error
	if err != nil {
		return err
	}
	task.Title = req.Title
	task.Content = req.Content
	task.Status = int(req.Status)
	task.StratTime = int64(req.StartTime)
	task.EndTime = int64(req.EndTime)
	err = DB.Save(task).Error
	return err
}

// 序列化
func BuildTasks(tasks []Task) (tList []*service.TaskModel) {
	for _, task := range tasks {
		f := BuildTask(task)
		tList = append(tList, f)
	}
	return tList
}

func BuildTask(task Task) *service.TaskModel {
	return &service.TaskModel{
		TaskId:    uint32(task.TaskId),
		UserId:    uint32(task.UserId),
		Status:    uint32(task.Status),
		Title:     task.Title,
		Content:   task.Content,
		StartTime: uint32(task.StratTime),
		EndTime:   uint32(task.EndTime),
	}
}
