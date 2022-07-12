package user

import (
	"gorm.io/gorm"
	"shopping/utils/hash"
)

// 保存用户之前回调 如果密码没有被加密 加密密码和salt
func (user *User) BeforeSave(db *gorm.DB) (err error) {
	if user.Salt == "" {
		// 为salt创建一个随机字符串
		salt := hash.CreateSalt()
		// 创建hash加密密码
		password, err := hash.HashPassword(user.Password + salt)
		if err != nil {
			return err
		}
		user.Password = password
		user.Salt = salt
	}
	return
}
