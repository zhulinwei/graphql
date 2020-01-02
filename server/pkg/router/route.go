package router

import (
	"github.com/gin-gonic/gin"
	"graphql/pkg/controller/graphql"
)

func InitRoute(route *gin.Engine) {
	// 用于访问GraphiQL客户端
	route.GET("/graphql", graphql.Handler())
	// 用于接受常规的CURD命令
	route.POST("/graphql", graphql.Handler())
}
