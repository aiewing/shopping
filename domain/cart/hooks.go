package cart

import "gorm.io/gorm"

// 如果计数为零 则删除商品
func (item *CartItem) AfterUpdate(db *gorm.DB) error {
	if item.Count <= 0 {
		return db.Unscoped().Delete(&item).Error
	}
	return nil
}
