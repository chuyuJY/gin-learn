package handler

import (
	"context"
	"gin-learn/gin-todolist/api-gateway/internal/service"
	"gin-learn/gin-todolist/api-gateway/pkg/resp"
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
	"gin-learn/gin-todolist/user/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// gin.Key 中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := resp.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// gin.Key 中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.UserId))
	r := resp.Response{
		Data:   resp.TokenData{User: userResp.UserDetail, Token: token},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
