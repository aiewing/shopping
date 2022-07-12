package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 保存商品之前生成商品sku
func (this *Product) BeforeSave(db *gorm.DB) {
	this.SKU = uuid.New().String()
}
