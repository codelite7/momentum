package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/ent"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateThread is the resolver for the createThread field.
func (r *mutationResolver) CreateThread(ctx context.Context, input ent.CreateThreadInput) (*ent.Thread, error) {
	//userUuid, err := getUserUuid(ctx)
	//if err != nil {
	//	return nil, gqlerror.Errorf(err.Error())
	//}
	//input.CreatedByID = userUuid
	threadName, err := common.GetThreadName(input.Name)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	input.Name = threadName
	return ent.FromContext(ctx).Thread.Create().SetInput(input).Save(ctx)
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) BookmarkThread(ctx context.Context, id uuid.UUID) (*bool, error) {
	panic(fmt.Errorf("not implemented: BookmarkThread - bookmarkThread"))
}
func (r *mutationResolver) UnbookmarkThread(ctx context.Context, id uuid.UUID) (*bool, error) {
	panic(fmt.Errorf("not implemented: UnbookmarkThread - unbookmarkThread"))
}
