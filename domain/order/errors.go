package order

import "errors"

var (
	ErrorEmptyCartFound       = errors.New("购物车是空的")
	ErrorInvalidOrderID       = errors.New("无效订单")
	ErrorCancelDurationPassed = errors.New("已通过取消持续时间")
	ErrorNotEnoughStock       = errors.New("没有足够库存")
)
