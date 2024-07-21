package main

import (
	"gjob-admin/api"
	"gjob-admin/pkg/dao"
	"gjob-admin/pkg/middleware"
	"gjob-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	utils.InitConfig()
	// 初始化数据库
	dao.InitDB()
	// 初始化实例
	r := gin.Default()
	// 使用中间件
	r.Use(middleware.GinLogger, middleware.Cors())
	// 注册user模块路由
	api.Router(r)
	// project.Router(r)
	// api.Router(r)
	r.Run(utils.Config.GetString("server.base.listen_addr"))

}
