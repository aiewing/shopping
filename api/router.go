package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"shopping/config"
	"shopping/domain/cart"
	"shopping/domain/category"
	"shopping/domain/order"
	"shopping/domain/product"
	"shopping/domain/user"
	"shopping/utils/database_handle"
	"shopping/utils/middleware"

	cartApi "shopping/api/cart"
	categoryApi "shopping/api/category"
	orderApi "shopping/api/order"
	productApi "shopping/api/product"
	userApi "shopping/api/user"
)

// Databases 结构体
type DataBases struct {
	categoryRepo    *category.Repository
	userRepo        *user.Repository
	productRepo     *product.Repository
	cartRepo        *cart.CartRepository
	cartItemRepo    *cart.CartItemRepository
	orderRepo       *order.OrderRepository
	orderedItemRepo *order.OrderedItemRepository
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}

// 根据配置文件创建数据库
func CreateDBs() *DataBases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}

	db := database_handle.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)
	return &DataBases{
		categoryRepo:    category.NewCategoryRepository(db),
		userRepo:        user.NewUserRepository(db),
		productRepo:     product.NewProductRepository(db),
		cartRepo:        cart.NewCartRepository(db),
		cartItemRepo:    cart.NewCartItemRepository(db),
		orderRepo:       order.NewOrderRepository(db),
		orderedItemRepo: order.NewOrderedItemRepository(db),
	}
}

// 注册所有控制器
func RegisterHandlers(engine *gin.Engine) {
	dbs := *CreateDBs()
	RegisterUserHandlers(engine, dbs)
	RegisterCategoryHandlers(engine, dbs)
	RegisterCartHandlers(engine, dbs)
	RegisterProductHandlers(engine, dbs)
	RegisterOrderHandlers(engine, dbs)
}

// 注册用户控制器
func RegisterUserHandlers(engine *gin.Engine, dbs DataBases) {
	userService := user.NewUserService(*dbs.userRepo)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := engine.Group("/user")
	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/login", userController.Login)
}

// 注册分类控制器
func RegisterCategoryHandlers(engine *gin.Engine, dbs DataBases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepo)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := engine.Group("/category")
	categoryGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateCategory)
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST(
		"/upload", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey),
		categoryController.BulkCreateCategory)
}

// 注册购物车控制器
func RegisterCartHandlers(engine *gin.Engine, dbs DataBases) {
	cartService := cart.NewService(*dbs.cartRepo, *dbs.cartItemRepo, *dbs.productRepo)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := engine.Group("/cart", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/", cartController.GetCart)
}

// 注册商品控制器
func RegisterProductHandlers(engine *gin.Engine, dbs DataBases) {
	productService := product.NewService(*dbs.productRepo)
	productController := productApi.NewProductController(*productService)
	productGroup := engine.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.CreateProduct)
	productGroup.DELETE(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProduct)
	productGroup.PATCH(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)

}

// 注册订单控制器
func RegisterOrderHandlers(engine *gin.Engine, dbs DataBases) {
	orderService := order.NewService(
		*dbs.orderRepo, *dbs.orderedItemRepo, *dbs.productRepo, *dbs.cartRepo,
		*dbs.cartItemRepo)
	productController := orderApi.NewOrderController(orderService)
	orderGroup := engine.Group("/order", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	orderGroup.POST("", productController.CompleteOrder)
	orderGroup.DELETE("", productController.CancelOrder)
	orderGroup.GET("", productController.GetOrders)
}
