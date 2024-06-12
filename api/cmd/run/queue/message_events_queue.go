package queue

import (
	"context"
	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"github.com/codelite7/momentum/api/common"
	message2 "github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"go.uber.org/zap"
)

const messageEventsQueueName = "messages"

func startMessageEventHandler(queue neoq.Neoq) error {
	h := handler.New(messageEventsQueueName, messageEventHandler)
	err := queue.Start(context.Background(), h)
	if err != nil {
		return err
	}

	return nil
}

func messageEventHandler(ctx context.Context) error {
	err := handleMessageEvent(ctx)
	if err != nil {
		common.Logger.Error("error handling mesage event", zap.Error(err))
		return err
	}
	return nil
}

func handleMessageEvent(ctx context.Context) error {
	common.Logger.Info("handling message event")
	job, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}

	messageId := job.Payload["id"].(string)
	common.Logger.Info("querying for message")
	message, err := common.EntClient.Message.Get(ctx, pulid.ID(messageId))
	if err != nil {
		//if ent.IsNotFound(err) {
		//	common.Logger.Error("message not found", zap.String("messageId", messageId))
		//	return nil
		//}
		return err
	}
	common.Logger.Info("got message")
	thread, err := message.Thread(ctx)
	if err != nil {
		return err
	}

	messages, err := common.GetMessageHistory(ctx, thread.ID)
	if err != nil {
		return err
	}
	chatMessages := common.ChatMessagesFromMessages(messages)
	sentBy, err := message.SentBy(ctx)
	if err != nil {
		return err
	}
	tenant, err := sentBy.QueryActiveTenant().First(ctx)
	if err != nil {
		return err
	}
	result, err := common.Prompt(chatMessages, thread.Provider)
	if err != nil {
		return err
	}
	if result != "" {
		_, err = common.EntClient.Message.Create().SetThread(thread).SetContent(result).SetTenant(tenant).SetSentBy(sentBy).SetMessageType(message2.MessageTypeAi).Save(ctx)
	}
	if err != nil {
		return err
	}
	return nil
}

func EnqueueMessageEvent(id pulid.ID) error {
	payload := map[string]interface{}{"id": id}
	job := &jobs.Job{
		Queue:   messageEventsQueueName,
		Payload: payload,
	}
	_, err := queue.Enqueue(context.Background(), job)
	if err != nil {
		return err
	}

	return nil
}
