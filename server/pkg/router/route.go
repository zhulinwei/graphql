package router

import (
	"github.com/gin-gonic/gin"
	"graphql/pkg/controller/graphql"
)

func InitRoute(route *gin.Engine) {
	route.GET("/graphql", graphql.Handler())
	route.POST("/graphql", graphql.Handler())
}
