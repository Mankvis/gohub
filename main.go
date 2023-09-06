package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务，默认端口为 8080，我们改为 8000
	err := r.Run(":3000")
	if err != nil {
		// 处理错误，端口被占用或者其他错误
		fmt.Println(err.Error())
	}
}
