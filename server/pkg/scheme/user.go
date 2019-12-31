package scheme

import (
	"github.com/graphql-go/graphql"
	"graphql/pkg/model"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "user",
	Description: "user model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryUser = graphql.Field{
	Name:"QueryUser",
	Description:"Query User",
	Type:graphql.NewList(userType),
	Args:graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type:graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
		id, _ := params.Args["id"].(int)
		name, _:= params.Args["name"].(string)
	},
}