package product

import "shopping/utils/pagination"

type Service struct {
	repo Repository
}

// 实例化
func NewService(repo Repository) *Service {
	repo.Migration()
	return &Service{
		repo: repo,
	}
}

// 获得所有商品分页
func (this *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := this.repo.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// 创建商品
func (this *Service) CreateProduct(name string, desc string, stockCount int, price float32, cid uint) error {
	newProduct := NewProduct(name, desc, stockCount, price, cid)
	err := this.repo.Create(newProduct)
	return err
}

// 删除商品
func (this *Service) DeleteProduct(sku string) error {
	err := this.repo.Delete(sku)
	return err
}

// 更新商品
func (this *Service) UpdateProduct(product Product) error {
	err := this.repo.Update(product)
	return err
}

// 查找商品
func (this *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := this.repo.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
