package cart

import (
	"gorm.io/gorm"
	"log"
)

// item
type CartItemRepository struct {
	db *gorm.DB
}

// 实例化item
func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{
		db: db,
	}
}

// 生成item表
func (this *CartItemRepository) Migration() {
	err := this.db.AutoMigrate(&CartItem{})
	if err != nil {
		log.Print(err)
	}
}

// 更新item
func (this *CartItemRepository) Update(item CartItem) error {
	return this.db.Save(&item).Error
}

// 根据商品id和购物车id查找item
func (this *CartItemRepository) FindById(productId uint, cartid uint) (*CartItem, error) {
	var item *CartItem
	err := this.db.Where(&CartItem{ProductID: productId, CartID: cartid}).First(&item).Error
	return item, err
}

// 常见item
func (this *CartItemRepository) Create(item *CartItem) error {
	err := this.db.Create(item).Error
	return err
}

// 返回购物车中所有item
func (this *CartItemRepository) GetAllItems(cartId uint) ([]CartItem, error) {
	var cartItems []CartItem
	err := this.db.Where(&CartItem{CartID: cartId}).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	for i, item := range cartItems {
		err := this.db.Model(item).Association("Product").Find(&cartItems[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return cartItems, nil
}
