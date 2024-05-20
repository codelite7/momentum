package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/codelite7/momentum/api/river"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input ent.CreateMessageInput) (*ent.Message, error) {
	// if the previous message doesn't have response content, don't allow creation of a new message
	userUuid, err := getUserUuid(ctx)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	input.SentByID = userUuid
	client := ent.FromContext(ctx)
	previousMessage, err := client.Message.Query().Where(message.HasThreadWith(thread.ID(input.ThreadID))).Order(message.ByCreatedAt(sql.OrderAsc())).First(ctx)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return nil, gqlerror.Errorf(err.Error())
	}
	if previousMessage != nil {
		previousResponse, err := previousMessage.Response(ctx)
		if err != nil {
			return nil, gqlerror.Errorf(err.Error())
		}
		if previousResponse == nil || previousResponse.Content == "" {
			return nil, gqlerror.Errorf("previous message response is empty, you must wait for a response before you can send another message")
		}
	}

	message, err := client.Message.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	agent, err := common.GetDefaultAgent(ctx, client)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	response, err := client.Response.Create().SetMessage(message).SetSentBy(agent).SetContent("").Save(ctx)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	_, err = river.RiverClient.Insert(ctx, river.MessageArgs{
		ThreadId:         input.ThreadID,
		MessageId:        message.ID,
		ResponseId:       response.ID,
		MessageContent:   message.Content,
		MessageCreatedAt: message.CreatedAt,
	}, nil)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return client.Message.Get(ctx, message.ID)
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type EntTx struct {
	client *ent.Client
}
