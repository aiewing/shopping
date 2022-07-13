package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/domain/category"
	"shopping/utils/api_helper"
	"shopping/utils/pagination"
)

type Controller struct {
	cateService *category.Service
}

// 实例化控制器
func NewCategoryController(service *category.Service) *Controller {
	return &Controller{
		cateService: service,
	}
}

// 根据给定的参数创建分类
func (this *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	if req.Name == "" || req.Desc == "" {
		api_helper.HandleError(g, api_helper.ErrorInvalidBody)
		return
	}

	newCategory := category.NewCategory(req.Name, req.Desc)
	err = this.cateService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusOK, api_helper.Response{
			Message: "Category created",
		})
}

// 根据给定的csv文件，批量创建分类
func (this *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, err := g.FormFile("file")
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	count, err := this.cateService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusOK, api_helper.Response{
			Message: fmt.Sprintf("%s uploaded! %d new categories created", fileHeader.Filename, count),
		})
}

// 获得分类列表
func (this *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = this.cateService.GetAll(page)
	g.JSON(http.StatusOK, page)
}
