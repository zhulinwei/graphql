package main

import (
	"github.com/99designs/gqlgen/handler"
	"graphql/pkg/controller"
	"graphql/pkg/util"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(util.NewExecutableSchema(util.Config{Resolvers: &controller.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
