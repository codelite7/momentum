package queue

import (
	"context"
	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/backends/postgres"
	"github.com/codelite7/momentum/api/config"
	"time"
)

var queue neoq.Neoq

func InitializeQueue() (err error) {
	queue, err = neoq.New(context.Background(),
		neoq.WithBackend(postgres.Backend),
		postgres.WithConnectionString(config.PostgresUri),
	)
	if err != nil {
		return err
	}
	err = InitializePromptQueue()
	if err != nil {
		return err
	}
	err = InitializeTestQueue()
	if err != nil {
		return err
	}
	//go queueThings() for testing queue things
	return nil
}

func queueThings() {
	count := 0
	for {
		err := QueueTest(context.Background(), TestQueueArgs{Count: count})
		if err != nil {
			panic(err)
		}
		count++
		time.Sleep(1 * time.Second)
	}
}
