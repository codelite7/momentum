package river

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/riverqueue/river"
)

type MessageArgs struct {
	MessageId      uuid.UUID `json:"messageId"`
	ResponseId     uuid.UUID `json:"responseId"`
	MessageContent string    `json:"messageContent"`
}

func (MessageArgs) Kind() string { return "message" }

type MessageWorker struct {
	river.WorkerDefaults[MessageArgs]
}

func (w *MessageWorker) Work(ctx context.Context, job *river.Job[MessageArgs]) error {
	fmt.Println("Handling Message Job")
	return nil
}
