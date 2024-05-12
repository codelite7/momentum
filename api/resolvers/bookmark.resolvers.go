package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/codelite7/momentum/api/ent"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateBookmark is the resolver for the createBookmark field.
func (r *mutationResolver) CreateBookmark(ctx context.Context, input ent.CreateBookmarkInput) (*ent.Bookmark, error) {
	userUuid, err := getUserUuid(ctx)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	input.UserID = userUuid
	return ent.FromContext(ctx).Bookmark.Create().SetInput(input).Save(ctx)
}

// DeleteBookmark is the resolver for the deleteBookmark field.
func (r *mutationResolver) DeleteBookmark(ctx context.Context, id uuid.UUID) (bool, error) {
	err := ent.FromContext(ctx).Bookmark.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}
