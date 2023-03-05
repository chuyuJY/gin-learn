package handler

import (
	"context"
	"gin-learn/gin-todolist/task/internal/repository"
	"gin-learn/gin-todolist/task/internal/service"
	"gin-learn/gin-todolist/task/pkg/e"
)

type TaskService struct {
	service.UnimplementedTaskServiceServer
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (ts *TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.SUCCESS
	resp.Msg = e.GetMsg(e.SUCCESS)
	if err = task.Create(req); err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	return resp, nil
}

func (ts *TaskService) TaskShow(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	var task repository.Task
	resp = new(service.TasksDetailResponse)
	taskList, err := task.Show(req)
	if err != nil {
		resp.Code = e.ERROR
	}
	resp.Code = e.SUCCESS
	resp.TaskDetail = repository.BuildTasks(taskList)
	return resp, nil
}

func (ts *TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task repository.Task
	resp = new(service.CommonResponse)
	if err = task.Update(req); err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Code = e.SUCCESS
	resp.Msg = e.GetMsg(e.SUCCESS)
	return resp, nil
}

func (ts *TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task repository.Task
	resp = new(service.CommonResponse)
	if err = task.Delete(req); err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Code = e.SUCCESS
	resp.Msg = e.GetMsg(e.SUCCESS)
	return resp, nil
}
