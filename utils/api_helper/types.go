package api_helper

import "errors"

// 响应结构体
type Response struct {
	Message string `json:"message"`
}

// 响应错误结构体
type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

// 自定义错误
var (
	ErrorInvalidBody = errors.New("参数错误")
)
