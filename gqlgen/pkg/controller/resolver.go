package controller

import (
	"context"
	"graphql/pkg/model"
	"graphql/pkg/util"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() util.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	userId := 123
	username := "tony"
	return &model.User{
		ID: &userId,
		Name: &username,
	}, nil
}
