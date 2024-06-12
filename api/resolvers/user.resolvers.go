package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/user"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*ent.User, error) {
	userInfo := common.GetUserIdFromContext(ctx)
	return r.client.User.Query().Where(user.ID(userInfo.UserId)).First(ctx)
}
