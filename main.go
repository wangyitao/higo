package main

import (
	"higo/conf"
	"higo/server"
)

// @title higo Example API
// @version 1.0
// @description This is a awesome web server.

// @host localhost:3000
// @BasePath /api/v1/
func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
