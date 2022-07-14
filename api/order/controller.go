package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/domain/order"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"
)

type Controller struct {
	orderService *order.Service
}

// 实例订单
func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		orderService: orderService,
	}
}

// 完成订单
func (this *Controller) CreateOrder(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	err := this.orderService.CreateOrder(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Created",
		})
}

// 取消订单
func (this *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	userId := api_helper.GetUserId(g)
	err := this.orderService.CancelOrder(userId, req.OrderId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Canceled",
		})
}

// 获得订单列表
func (this *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = this.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)
}
