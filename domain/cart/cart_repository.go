package cart

import (
	"gorm.io/gorm"
	"log"
)

type CartRepository struct {
	db *gorm.DB
}

// 实例化
func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

// 创建表
func (this *CartRepository) Migration() {
	err := this.db.AutoMigrate(&Cart{})
	if err != nil {
		log.Print(err)
	}
}

// 更新
func (this *CartRepository) Update(cart Cart) error {
	return this.db.Save(cart).Error
}

// 根据用户id查找或创建购物车
func (this *CartRepository) FindOrCreateByUserId(userId uint) (*Cart, error) {
	var cart *Cart
	err := this.db.Where(Cart{UserID: userId}).Attrs(NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// 根据用户id查找购物车
func (this *CartRepository) FindByUserId(userId uint) (*Cart, error) {
	var cart *Cart
	err := this.db.Where(Cart{UserID: userId}).Attrs(NewCart(userId)).First(&cart).Error
	return cart, err
}
