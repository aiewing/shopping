package user

import "errors"

var (
	ErrorUserExistWithName = errors.New("用户名已经存在")
	ErrorUserNotFound      = errors.New("用户不存在")
	ErrorPasswordError     = errors.New("密码错误")
)
