package cart

import (
	"gorm.io/gorm"
	"shopping/domain/product"
)

// Item
type CartItem struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// 创建Item
func NewCartItem(productId, cartId uint, count int) *CartItem {
	return &CartItem{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
