package router

import (
	"github.com/gin-gonic/gin"
	"graphql/pkg/controller/graphql"
)

func InitRoute(route *gin.Engine) {
	route.GET("/graphal", graphql.Handler())
	route.POST("/graphal", graphql.Handler())
}
