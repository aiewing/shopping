package cart

import (
	"gorm.io/gorm"
	"shopping/domain/user"
)

// 购物车结构体
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

// 实例化
func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}
