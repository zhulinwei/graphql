package main

import (
	"github.com/gin-gonic/gin"
	"graphql/pkg/router"
	"log"
)

func main () {
	route := gin.New()
	router.InitRoute(route)

	if err := route.Run(":8080"); err != nil {
		log.Fatalf("server run failed, err: %v", err.Error())
	}
}
