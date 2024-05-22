package queue

import (
	"context"
	"errors"
	"fmt"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"github.com/codelite7/momentum/api/common"
	"github.com/google/uuid"
	"time"
)

const PromptQueueName = "prompts"

type PromptQueueArgs struct {
	ThreadId         uuid.UUID `json:"threadId"`
	MessageId        uuid.UUID `json:"messageId"`
	ResponseId       uuid.UUID `json:"responseId"`
	MessageCreatedAt time.Time `json:"messageCreatedAt"`
	MessageContent   string    `json:"messageContent"`
}

type PromptQueueJob map[string]any

func InitializePromptQueue() error {
	h := handler.New(PromptQueueName, handlePromptQueueItem)
	err := queue.Start(context.Background(), h)
	if err != nil {
		return err
	}

	return nil
}

func handlePromptQueueItem(ctx context.Context) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.New(fmt.Sprint("recovered from panic: ", r))
		}
	}()
	job, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	payload := job.Payload
	queueJob := PromptQueueJob(payload)
	args, err := queueJob.ToArgs()
	if err != nil {
		return err
	}
	tx, err := common.EntClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	history, err := common.GetMessageHistory(ctx, tx, args.ThreadId, args.MessageCreatedAt)
	if len(history) == 0 {
		return errors.New("history is empty")
	}
	if err != nil {
		return err
	}
	chatMessages, err := common.ChatMessagesFromMessages(history)
	if err != nil {
		return err
	}
	aiResponse, err := common.Prompt(chatMessages)
	if err != nil {
		return err
	}
	if aiResponse == "" {
		return err
	}
	_, err = tx.Response.UpdateOneID(args.ResponseId).SetContent(aiResponse).Save(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func QueuePrompt(ctx context.Context, args PromptQueueArgs) error {
	job := &jobs.Job{Queue: PromptQueueName, Payload: args.ToJob()}
	_, err := queue.Enqueue(ctx, job)
	if err != nil {
		return err
	}

	return nil
}

func (p PromptQueueArgs) ToJob() PromptQueueJob {
	return map[string]any{
		"threadId":         p.ThreadId,
		"messageId":        p.MessageId,
		"responseId":       p.ResponseId,
		"messageCreatedAt": p.MessageCreatedAt,
		"messageContent":   p.MessageContent,
	}
}

func (p PromptQueueJob) ToArgs() (PromptQueueArgs, error) {
	threadId := p["threadId"].(string)
	messageId := p["messageId"].(string)
	responseId := p["responseId"].(string)
	messageCreatedAt := p["messageCreatedAt"].(string)
	messageContent := p["messageContent"].(string)
	args := PromptQueueArgs{}
	var err error
	args.ThreadId, err = uuid.Parse(threadId)
	if err != nil {
		return args, err
	}
	args.MessageId, err = uuid.Parse(messageId)
	if err != nil {
		return args, err
	}
	args.ResponseId, err = uuid.Parse(responseId)
	if err != nil {
		return args, err
	}
	args.ThreadId, err = uuid.Parse(threadId)
	if err != nil {
		return args, err
	}
	args.MessageCreatedAt, err = time.Parse(time.RFC3339, messageCreatedAt)
	if err != nil {
		return args, err
	}
	args.MessageContent = messageContent
	return args, nil
}
