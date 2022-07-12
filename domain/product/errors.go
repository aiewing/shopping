package product

import "errors"

var (
	ErrorProductNotFound         = errors.New("商品没有找到")
	ErrorProductStockIsNotEnough = errors.New("商品库存不足")
)
