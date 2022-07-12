package user

import "gorm.io/gorm"

// 用户模型
type User struct {
	gorm.Model

	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDeleted bool
	IsAdmin   bool
}

// 新建用户实例
func NewUser(username, password string) *User {
	return &User{
		Username:  username,
		Password:  password,
		IsDeleted: false,
		IsAdmin:   false,
	}
}
