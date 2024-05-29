package river

import (
	"context"
	"errors"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/riverqueue/river"
	"time"
)

type MessageArgs struct {
	ThreadId         pulid.ID  `json:"threadId"`
	MessageId        pulid.ID  `json:"messageId"`
	ResponseId       pulid.ID  `json:"responseId"`
	MessageCreatedAt time.Time `json:"messageCreatedAt"`
	MessageContent   string    `json:"messageContent"`
}

func (args MessageArgs) Kind() string { return "message" }

func (args MessageArgs) InsertOpts() river.InsertOpts {
	return river.InsertOpts{MaxAttempts: 10000000} // at 5 seconds per retry this is about 1.5 years
}

type MessageWorker struct {
	river.WorkerDefaults[MessageArgs]
}

func (w *MessageWorker) Work(ctx context.Context, job *river.Job[MessageArgs]) error {
	tx, err := common.EntClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	history, err := common.GetMessageHistory(ctx, tx, job.Args.ThreadId, job.Args.MessageCreatedAt)
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
	_, err = tx.Response.UpdateOneID(job.Args.ResponseId).SetContent(aiResponse).Save(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// message worker retries after 5 seconds
func (w *MessageWorker) NextRetry(job *river.Job[MessageArgs]) time.Time {
	return time.Now().Add(5 * time.Second)
}
