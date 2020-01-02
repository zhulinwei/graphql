package scheme

import (
	"github.com/graphql-go/graphql"
	"graphql/pkg/model"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "user",
	Description: "user info",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "user id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "username",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "user email",
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "user phone",
		},
		"status": &graphql.Field{
			Type:        UserStatusEnumType,
			Description: "user status",
		},
	},
})

var UserStatusEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserStatusEnum",
	Description: "user status enum",
	Values: graphql.EnumValueConfigMap{
		"EnableUser": &graphql.EnumValueConfig{
			Value:       model.EnableStatus,
			Description: "user enable",
		},
		"DisableUser": &graphql.EnumValueConfig{
			Value:       model.DisableStatus,
			Description: "user disable",
		},
	},
})