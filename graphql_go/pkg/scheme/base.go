package scheme

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:"Query",
	Fields:graphql.Fields{
		"User": UserField,
	},
})
