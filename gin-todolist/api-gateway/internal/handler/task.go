package handler

import (
	"context"
	"gin-learn/gin-todolist/api-gateway/internal/service"
	"gin-learn/gin-todolist/api-gateway/pkg/e"
	"gin-learn/gin-todolist/api-gateway/pkg/resp"
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskList(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.UserId = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := resp.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
		// Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.UserId = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := resp.Response{
		Data:   taskResp.Data,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
		// Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.UserId = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := resp.Response{
		Data:   taskResp.Data,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
		// Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteTask(ginCtx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.UserId = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskDelete(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := resp.Response{
		Data:   taskResp.Data,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
		// Error:  err.Error(),
	}
	ginCtx.JSON(http.StatusOK, r)
}
