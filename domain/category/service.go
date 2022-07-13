package category

import (
	"mime/multipart"
	"shopping/utils/csv_helper"
	"shopping/utils/pagination"
)

type Service struct {
	repo Repository
}

// 实例化商品分类Service
func NewCategoryService(repo Repository) *Service {
	// 生成表
	repo.Migration()
	return &Service{
		repo: repo,
	}
}

// 创建分类
func (this *Service) Create(category *Category) error {
	cateList := this.repo.GetByName(category.Name)
	if len(cateList) > 0 {
		return ErrorCategoryExistWithName
	}

	err := this.repo.Create(category)
	return err
}

// 批量创建分类
func (this *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVar := range bulkCategory {
		categories = append(categories, NewCategory(categoryVar[0], categoryVar[1]))
	}
	count, err := this.repo.BulkCreate(categories)
	return count, nil
}

// 获得分页商品分类
func (this *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := this.repo.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
