package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 错误处理
func HandleError(g *gin.Context, err error) {
	g.JSON(
		http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	g.Abort()
}
