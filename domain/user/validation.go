package user

import "regexp"

// 用户名正则表示式 最小8个字符 最大30个字符 用户名字母打头
var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_-]{4,16}$/")

// 密码正则表达式 最小8个字符 至少一个字符一个数字
var passwordRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,29}$")

func ValidateUserName(name string) bool {
	if len(name) > 0 {
		return true
	}
	return false
	//return usernameRegex.MatchString(name)
}

func ValidatePassword(password string) bool {
	if len(password) > 0 {
		return true
	}
	return false
	//return passwordRegex.MatchString(password)
}
