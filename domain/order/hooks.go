package order

import (
	"gorm.io/gorm"
	"shopping/domain/cart"
	"shopping/domain/product"
)

// 创建之前 查找购物车并删除里面的项目
func (this *Order) BeforeCreate(db *gorm.DB) error {

	// 查找购物车
	var currentCart cart.Cart
	err := db.Where("UserID = ?", this.UserID).First(&currentCart).Error
	if err != nil {
		return err
	}

	// 删除购物车中的item
	err = db.Where("CartID = ?", currentCart.ID).Unscoped().Delete(&cart.CartItem{}).Error
	if err != nil {
		return err
	}

	// 删除购物车
	err = db.Unscoped().Delete(&currentCart).Error
	return err
}

// 保存之后 更新产品库存
func (this *OrderedItem) BeforeSave(db *gorm.DB) (err error) {
	var currentProduct product.Product
	var currentOrderedItem OrderedItem

	// 通过产品ID查找产品
	err = db.Where("ID = ?", this.ProductID).First(&currentProduct).Error
	if err != nil {
		return err
	}

	// 原来购物车中的产品数
	reservedStockCount := 0
	err = db.Where("ID = ?", this.ID).First(&currentOrderedItem).Error
	if err == nil {
		reservedStockCount = currentOrderedItem.Count
	}
	newStockCount := currentProduct.StockCount + reservedStockCount - this.Count
	if newStockCount < 0 {
		return ErrorNotEnoughStock
	}

	// 更新产品库存
	err = db.Model(&currentProduct).Update("StockCount", newStockCount).Error
	if err != nil {
		return err
	}

	// 如果数量为0 删除当前项
	if this.Count == 0 {
		err := db.Unscoped().Delete(currentOrderedItem).Error
		return err
	}
	return err
}

// 如果订单被取消 金额将会返回产品库存
func (this *Order) BeforeUpdate(db *gorm.DB) (err error) {
	if this.IsCanceled {
		var orderedItems []OrderedItem

		// 查询订单中所有的订单项
		err = db.Where("OrderID = ?", this.ID).Find(&orderedItems).Error
		if err != nil {
			return err
		}

		for _, item := range orderedItems {
			var currentProduct product.Product
			// 通过产品ID 查找产品
			err = db.Where("ID = ?", item.ProductID).First(&currentProduct).Error
			if err != nil {
				return err
			}

			// 更新库存
			newStockCount := currentProduct.StockCount + item.Count
			err = db.Model(&currentProduct).Update("StockCount", newStockCount).Error
			if err != nil {
				return err
			}

			// 更新订单项为已经取消
			err = db.Model(&item).Update("IsCanceled", true).Error
			if err != nil {
				return err
			}
		}
	}
	return
}
