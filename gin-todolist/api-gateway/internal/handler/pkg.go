package handler

import (
	"errors"
	"gin-learn/gin-todolist/api-gateway/pkg/utils"
)

func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("User Service Err: " + err.Error())
		utils.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("Task Service Err: " + err.Error())
		utils.LogrusObj.Info(err)
		panic(err)
	}
}
