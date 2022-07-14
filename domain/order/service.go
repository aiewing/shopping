package order

import (
	"shopping/domain/cart"
	"shopping/domain/product"
	"shopping/utils/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepo       OrderRepository
	orderedItemRepo OrderedItemRepository
	productRepo     product.Repository
	cartRepo        cart.CartRepository
	cartItemRepo    cart.CartItemRepository
}

// 实例化
func NewService(
	orderRepo OrderRepository,
	orderedItemRepo OrderedItemRepository,
	productRepo product.Repository,
	cartRepo cart.CartRepository,
	cartItemRepo cart.CartItemRepository,
) *Service {
	orderRepo.Migration()
	orderedItemRepo.Migration()

	return &Service{
		orderRepo:       orderRepo,
		orderedItemRepo: orderedItemRepo,
		productRepo:     productRepo,
		cartRepo:        cartRepo,
		cartItemRepo:    cartItemRepo,
	}

}

// 完成订单
func (this *Service) CreateOrder(userId uint) error {
	// 查询购物车
	currentCart, err := this.cartRepo.FindOrCreateByUserId(userId)
	if err != nil {
		return err
	}

	// 获取购物车中所有的内容
	cartItems, err := this.cartItemRepo.GetAllItems(currentCart.ID)
	if err != nil {
		return err
	}

	// 购物车为空
	if len(cartItems) == 0 {
		return ErrorEmptyCartFound
	}

	// 创建订单中的所有订单项
	orderedItems := make([]OrderedItem, 0)
	for _, item := range cartItems {
		orderedItems = append(orderedItems, *NewOrderedItem(item.Count, item.ProductID))
	}

	// 创建订单
	err = this.orderRepo.Create(NewOrder(userId, orderedItems))
	return err
}

// 取消订单
func (this *Service) CancelOrder(uid, oid uint) error {
	// 查找订单
	currentOrder, err := this.orderRepo.FindByOrderID(oid)
	if err != nil {
		return err
	}

	//
	if currentOrder.UserID != uid {
		return ErrorInvalidOrderID
	}

	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return ErrorCancelDurationPassed
	}

	// 更新订单
	currentOrder.IsCanceled = true
	err = this.orderRepo.Update(*currentOrder)

	return err
}

// 获得订单
func (this *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := this.orderRepo.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page
}
