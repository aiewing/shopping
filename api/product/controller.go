package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/domain/product"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"
)

type Controller struct {
	productService product.Service
}

// 实例化
func NewProductController(service product.Service) *Controller {
	return &Controller{
		productService: service,
	}
}

// 获得商品列表（分页）
func (this *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	queryText := g.Query("qt")
	if queryText != "" {
		page = this.productService.SearchProduct(queryText, page)
	} else {
		page = this.productService.GetAll(page)
	}
	g.JSON(http.StatusOK, page)

}

// 创建商品
func (this *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := this.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Product Created",
		})
}

// 根据sku删除商品
func (this *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := this.productService.DeleteProduct(req.SKU)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusOK, api_helper.Response{
			Message: "Product Deleted",
		})
}

// 根据sku更新商品
func (this *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := this.productService.UpdateProduct(req.ToProduct())
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusOK, CreateProductResponse{
			Message: "Product Updated",
		})
}
