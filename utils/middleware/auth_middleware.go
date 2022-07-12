package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping/utils/jwt_helper"
)

// 管理员授权
func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authStr := context.GetHeader("Authorization")
		if authStr != "" {
			decodedClaims := jwt_helper.VerifyToken(authStr, secretKey)
			if decodedClaims != nil && decodedClaims.IsAdmin {
				context.Next()
				context.Abort()
				return
			}

			context.JSON(http.StatusForbidden, gin.H{"error": "你没有访问权限！"})
			context.Abort()
			return
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "你没有授权！"})
		}
		context.Abort()
		return
	}
}

// 用户授权
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authStr := context.GetHeader("Authorization")
		if authStr != "" {
			decodedClaims := jwt_helper.VerifyToken(authStr, secretKey)
			if decodedClaims != nil {
				context.Set("userId", decodedClaims.UserId)
				context.Next()
				context.Abort()
				return
			}

			context.JSON(http.StatusForbidden, gin.H{"error": "你没有访问权限！"})
			context.Abort()
			return
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "你没有授权！"})
		}
		context.Abort()
		return
	}
}
