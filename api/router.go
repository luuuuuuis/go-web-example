package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 创建登录路由
	r.GET("/hello", hello)

}
