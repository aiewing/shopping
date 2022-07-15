package order

import (
	"gorm.io/gorm"
	"log"
)

type OrderRepository struct {
	db *gorm.DB
}

// 实例化
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

// 创建表
func (this *OrderRepository) Migration() {
	err := this.db.AutoMigrate(&Order{})
	if err != nil {
		log.Print(err)
	}
}

// 根据订单id查找
func (this *OrderRepository) FindByOrderID(oid uint) (*Order, error) {
	var currentOrder *Order
	err := this.db.Where("IsCanceled = ?", false).Where("ID", oid).First(&currentOrder).Error
	return currentOrder, err
}

// 更新订单
func (this *OrderRepository) Update(newOrder Order) error {
	err := this.db.Save(&newOrder).Error
	return err
}

// 创建订单
func (this *OrderRepository) Create(ci *Order) error {
	err := this.db.Create(ci).Error
	return err
}

// 获得所有订单
func (this *OrderRepository) GetAll(pageIndex, pageSize int, uid uint) ([]Order, int) {
	var orders []Order
	var count int64

	this.db.Where("IsCanceled = ?", 0).Where(
		"UserID", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count)
	for i, order := range orders {
		this.db.Where("OrderID = ?", order.ID).Find(&orders[i].OrderedItems)
		// 计算总价格
		var totalPrice float32 = 0.0
		for j, item := range orders[i].OrderedItems {
			this.db.Where("ID = ?", item.ProductID).First(&orders[i].OrderedItems[j].Product)
			totalPrice += orders[i].OrderedItems[j].Product.Price * float32(orders[i].OrderedItems[j].Count)
		}
		orders[i].TotalPrice = totalPrice
	}
	return orders, int(count)
}
