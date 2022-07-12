package cart

import (
	"shopping/domain/product"
)

type Service struct {
	cartRepo    CartRepository
	itemRepo    CartItemRepository
	productRepo product.Repository
}

// 实例化service
func NewService(cartRepo CartRepository, itemRepo CartItemRepository, productRepo product.Repository) *Service {
	cartRepo.Migration()
	itemRepo.Migration()
	return &Service{
		cartRepo:    cartRepo,
		itemRepo:    itemRepo,
		productRepo: productRepo,
	}
}

// 添加item
func (this *Service) AddItem(userId uint, sku string, count int) error {
	if count <= 0 {
		return ErrorCountInvalid
	}

	// 查询商品
	currentProduct, err := this.productRepo.FindBySKU(sku)
	if err != nil {
		return err
	}

	// 查看库存是否充足
	if currentProduct.StockCount < count {
		return product.ErrorProductStockIsNotEnough
	}

	// 查询购物车
	currentCart, err := this.cartRepo.FindOrCreateByUserId(userId)
	if err != nil {
		return err
	}

	// 判断商品是否已经存在
	_, err = this.itemRepo.FindById(currentProduct.ID, currentCart.ID)
	if err == nil {
		return ErrorItemAlreadyExistedInCart
	}

	// 在购物车中加入商品
	err = this.itemRepo.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))
	return err
}

// 更新item
func (this *Service) UpdateItem(userId uint, sku string, count int) error {
	// 查询商品
	currentProduct, err := this.productRepo.FindBySKU(sku)
	if err != nil {
		return err
	}

	// 查询购物车
	currentCart, err := this.cartRepo.FindOrCreateByUserId(userId)
	if err != nil {
		return err
	}

	// 判断商品是否已经存在
	currentItem, err := this.itemRepo.FindById(currentProduct.ID, currentCart.ID)
	if err != nil {
		return ErrorItemNotExistInCart
	}

	// 查看库存是否充足
	if currentProduct.StockCount+currentItem.Count < count {
		return product.ErrorProductStockIsNotEnough
	}
	currentItem.Count = count

	// 更新商品信息
	err = this.itemRepo.Update(*currentItem)
	return err
}

// 获得item
func (this *Service) GetCartItems(userId uint) ([]CartItem, error) {
	// 查找购物车
	currentCart, err := this.cartRepo.FindOrCreateByUserId(userId)
	if err != nil {
		return nil, err
	}

	items, err := this.itemRepo.GetAllItems(currentCart.ID)
	return items, err
}
