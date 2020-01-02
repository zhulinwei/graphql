package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"graphql/pkg/scheme"
	"log"
)

func Handler() gin.HandlerFunc {

	graphqlScheme, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: scheme.QueryType,
	})

	if err != nil {
		log.Fatalf("graphql build schema fail: %v", err.Error())
	}

	return func(context *gin.Context) {
		handler.New(&handler.Config{
			// 定义scheme
			Schema: &graphqlScheme,
			// 启用GraphiQL客户端
			GraphiQL: true,
		}).ServeHTTP(context.Writer, context.Request)
	}
}