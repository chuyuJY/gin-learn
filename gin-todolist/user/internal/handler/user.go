package handler

import (
	"context"
	"gin-learn/gin-todolist/user/internal/repository"
	"gin-learn/gin-todolist/user/internal/service"
	"gin-learn/gin-todolist/user/pkg/e"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	user := &repository.User{}
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	if err := user.ShowUserInfo(req); err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	user := &repository.User{}
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success
	if err := user.CreateUser(req); err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserLogout(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	resp = new(service.UserDetailResponse)
	return resp, nil
}
