package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/domain/cart"
	"shopping/utils/api_helper"
)

type Controller struct {
	cartService *cart.Service
}

// 实例化
func NewCartController(service *cart.Service) *Controller {
	return &Controller{
		cartService: service,
	}
}

// 添加Item
func (this *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := this.cartService.AddItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "Item added to cart",
		})
}

// 更新Item
func (this *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)

	err := this.cartService.UpdateItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "updated",
		})
}

// 获得购物车商品列表
func (this *Controller) GetCart(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	result, err := this.cartService.GetCartItems(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, result)
}
