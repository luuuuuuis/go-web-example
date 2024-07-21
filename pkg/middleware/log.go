/*
@File    :   log.go
@Time    :   2024/04/10 23:19:25
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package middleware

import (
	"gjob-admin/pkg/utils"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

// GinLogger 用于替换gin框架的Logger中间件，不传参数
func GinLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	c.Next() // 执行视图函数
	// 视图函数执行完成，统计时间，记录日志
	cost := time.Since(start)
	if int(cost) >= 10000000000 {
		utils.Logger.Error().
			Err(errors.New("请求响应超时")).
			Stack().
			Int("status", 500).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("query", query).
			Str("ip", c.ClientIP()).
			Str("user-agent", c.Request.UserAgent()).
			Msg("请求响应超时")
	} else {
		utils.Logger.Info().
			Int("status", 200).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("query", query).
			Str("ip", c.ClientIP()).
			Str("user-agent", c.Request.UserAgent()).
			Msg("请求成功")
	}
}
