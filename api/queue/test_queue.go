package queue

import (
	"context"
	"fmt"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"time"
)

const TestQueueName = "tests"

type TestQueueArgs struct {
	Count int
}

type TestQueueJob map[string]any

func InitializeTestQueue() error {
	h := handler.New(TestQueueName, handleTestQueueItem)
	err := queue.Start(context.Background(), h)
	if err != nil {
		return err
	}

	return nil
}

func handleTestQueueItem(ctx context.Context) (err error) {
	fmt.Printf("handling queue item")
	time.Sleep(5 * time.Second)
	job, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	fmt.Println(TestQueueJob(job.Payload).ToArgs().Count)
	return nil
}

func QueueTest(ctx context.Context, args TestQueueArgs) error {
	job := &jobs.Job{Queue: TestQueueName, Payload: args.ToJob()}
	_, err := queue.Enqueue(ctx, job)
	if err != nil {
		return err
	}

	return nil
}

func (p TestQueueArgs) ToJob() TestQueueJob {
	return map[string]any{
		"count": p.Count,
	}
}

func (p TestQueueJob) ToArgs() TestQueueArgs {
	return TestQueueArgs{Count: int(p["count"].(float64))}
}
