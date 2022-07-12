package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shopping/config"
	"shopping/domain/user"
	"shopping/utils/api_helper"
	"shopping/utils/jwt_helper"
	"strconv"
	"time"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

// 实例化
func NewUserController(userService *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: userService,
		appConfig:   appConfig,
	}
}

// 根据给定的用户名和密码创建用户
func (this *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest

	// 检查参数
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrorInvalidBody)
		return
	}

	// 创建新用户
	newUser := user.NewUser(req.Username, req.Password)
	err = this.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateUserResponse{
			Username: req.Username,
		})
}

// 根据用户名和密码登陆
func (this *Controller) Login(g *gin.Context) {
	var req LoginRequest

	// 检查参数
	err := g.ShouldBind(&req)
	if err != nil {
		api_helper.HandleError(g, api_helper.ErrorInvalidBody)
		return
	}

	// 获取用户
	currentUser, err := this.userService.GetUser(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	decodedClaims := jwt_helper.VerifyToken(currentUser.Token, this.appConfig.SecretKey)
	if decodedClaims == nil {
		// 没有token 就需要重新生成
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(currentUser.ID), 10),
				"username": currentUser.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(24 * time.Hour).Unix(),
				"isAdmin":  currentUser.IsAdmin,
			})
		token := jwt_helper.GenerateToken(jwtClaims, this.appConfig.SecretKey)
		currentUser.Token = token

		// 更新用户信息
		err := this.userService.UpdateUser(&currentUser)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}

	g.JSON(
		http.StatusOK, LoginResponse{
			Username: currentUser.Username,
			UserId:   currentUser.ID,
			Token:    currentUser.Token,
		})
}

// 验证token
func (this *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwt_helper.VerifyToken(token, this.appConfig.SecretKey)
	g.JSON(http.StatusOK, decodedClaims)
}
