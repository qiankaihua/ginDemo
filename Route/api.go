package Route

import (
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Http"
)

// 该文件编写 api 路由转发
func AddApiRoute() {
	Http.Router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "success",
			"version": "v1.0",
		})
	})
}