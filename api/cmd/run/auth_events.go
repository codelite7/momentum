package run

import (
	"context"
	"errors"
	"fmt"
	"github.com/codelite7/momentum/api/cmd/run/queue"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	"github.com/workos/workos-go/v4/pkg/events"
	"time"
)

func HandleAuthEvents() {
	for {
		err := PollForEvents()
		if err != nil {
			fmt.Printf("PollForEvents error: %v\n", err)
		}
		time.Sleep(config.WorkosPollInterval)
	}
}

func PollForEvents() error {
	eventTypes := []string{
		"user.created",
	}

	cursor, err := common.EntClient.WorkosEventCursor.Query().First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			cursor, err = common.EntClient.WorkosEventCursor.Create().Save(context.Background())
			if err != nil {
				return errors.New("error creating new workos event cursor: " + err.Error())
			}
		} else {
			return errors.New("failed to get workos event cursor: " + err.Error())
		}
	}
	opts := events.ListEventsOpts{
		Events: eventTypes,
	}
	if cursor.UserCreatedCursor != "" {
		opts.After = cursor.UserCreatedCursor
	}
	events, err := events.ListEvents(context.Background(), opts)
	if err != nil {
		return errors.New("error polling for events: " + err.Error())
	}
	for _, event := range events.Data {
		err = handleEvent(event)
		if err != nil {
			return errors.New("error handling event: " + err.Error())
		}
		cursor, err = cursor.Update().SetUserCreatedCursor(event.ID).Save(context.Background())
		if err != nil {
			return errors.New("error updating cursor: " + err.Error())
		}
	}
	return nil
}

func handleEvent(event events.Event) (err error) {
	switch event.Event {
	case "user.created":
		err = queue.EnqueueWorkosUserCreatedEvent(event)
	default:
		err = errors.New(fmt.Sprintf("unknown event: %s", event.Event))
	}
	return
}
