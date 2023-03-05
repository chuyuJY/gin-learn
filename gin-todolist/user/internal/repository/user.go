package repository

import (
	"errors"
	"gin-learn/gin-todolist/user/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const (
	PasswordCost = 12 // 加密难度
)

// CheckUserExist 检查用户是否存在
func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		return false
	}
	return true
}

// CreateUser 新建用户
func (user *User) CreateUser(req *service.UserRequest) (err error) {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count > 0 {
		err = errors.New("UserName is exist")
		return
	}
	user = &User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	// 密码摘要
	if err = user.HashPassword(req.Password); err != nil {
		return
	}
	if err = DB.Create(user).Error; err != nil {
		return
	}
	return nil
}

// SetPassword 得到密码摘要
func (user *User) HashPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 检验密码摘要
func (user *User) CheckPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)); err != nil {
		return false
	}
	return true
}

// ShowUserInfo 获取用户信息
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("用户不存在")
}

// BuildUser 将 User 封装
func BuildUser(user *User) *service.UserModel {
	userModel := &service.UserModel{
		UserId:   uint32(user.UserId),
		UserName: user.UserName,
		NickName: user.NickName,
	}
	return userModel
}
