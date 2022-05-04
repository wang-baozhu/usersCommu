package model

import "errors"

var (
	//注册使错误
	USER_EXIST_ERROR = errors.New("用户已存在") //400
	//登录时错误
	USER_NOT_EXIST_ERROR = errors.New("用户不存在") //500
	USER_PASSWORD_ERROR  = errors.New("密码错误")  //300
	//其他类型的错误返回状态码110
)
