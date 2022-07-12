package cart

import "errors"

var (
	ErrorItemAlreadyExistedInCart = errors.New("商品已经存在")
	ErrorCountInvalid             = errors.New("数量不能是负数")

	ErrorItemNotExistInCart = errors.New("购物车商品不存在")
)
