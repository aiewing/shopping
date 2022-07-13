package category

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// 创建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成商品分类表
func (this *Repository) Migration() {
	err := this.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

// 创建商品分类
func (this *Repository) Create(cate *Category) error {
	result := this.db.Create(cate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 通过名称查询商品分类
func (this *Repository) GetByName(name string) []Category {
	var categories []Category
	this.db.Where("Name = ?", name).Find(&categories)
	return categories
}

// 批量创建商品分类
func (this *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := this.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// 获得分页商品分类
func (this *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64

	this.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)

	return categories, int(count)
}
