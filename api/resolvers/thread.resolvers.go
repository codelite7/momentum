package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
)

// CreateThread is the resolver for the createThread field.
func (r *mutationResolver) CreateThread(ctx context.Context, input ent.CreateThreadInput) (*ent.Thread, error) {
	//threadName, err := common.GetThreadName(input.Name)
	//if err != nil {
	//	return nil, gqlerror.Errorf(err.Error())
	//}
	//input.Name = threadName
	return ent.FromContext(ctx).Thread.Create().SetInput(input).Save(ctx)
}

// DeleteThread is the resolver for the deleteThread field.
func (r *mutationResolver) DeleteThread(ctx context.Context, id pulid.ID) (pulid.ID, error) {
	err := ent.FromContext(ctx).Thread.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Thread is the resolver for the thread field.
func (r *queryResolver) Thread(ctx context.Context, id pulid.ID) (*ent.Thread, error) {
	return r.client.Thread.Get(ctx, id)
}
