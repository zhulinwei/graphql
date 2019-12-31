package scheme

import (
	"github.com/graphql-go/graphql"
	"graphql/pkg/model"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:"Query",
	Fields:graphql.Fields{
		"User": &graphql.Field{
			Type:UserType,
			Args:graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (result interface{}, err error) {
				return model.User{}, nil
			},
		},
	},
})
