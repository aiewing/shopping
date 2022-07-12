package product

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生产表
func (this *Repository) Migration() {
	err := this.db.AutoMigrate(&Product{})
	if err != nil {
		log.Print(err)
	}
}

// 更新
func (this *Repository) Update(updateProduct Product) error {
	savedProduct, err := this.FindBySKU(updateProduct.SKU)
	if err != nil {
		return err
	}

	err = this.db.Model(&savedProduct).Updates(updateProduct).Error
	return err
}

// 创建
func (this *Repository) Create(product *Product) error {
	result := this.db.Create(product)
	return result.Error
}

// 查询所有商品
func (this *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64

	this.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

// 根据sku删除
func (this *Repository) Delete(sku string) error {
	product, err := this.FindBySKU(sku)
	if err != nil {
		return err
	}
	product.IsDeleted = true

	err = this.db.Save(product).Error
	return err
}

// 根据sku查询
func (this *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	err := this.db.Where("IsDeleted = ?", 0).Where(Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, ErrorProductNotFound
	}
	return product, nil
}

// 搜索返回分页结果
func (this *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	this.db.Where("IsDeleted = ?", false).Where(
		"Name LIKE ? OR SKU Like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}
