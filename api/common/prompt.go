package common

import (
	"context"
	"fmt"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/go-resty/resty/v2"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

func GetThreadName(messageContent string) (string, error) {
	prompt := fmt.Sprintf("Provide a concise 3 word summary of the following text:%s", messageContent)
	return Prompt([]ChatMessage{ChatMessage{
		Type: "human",
		Data: ChatMessageData{Content: prompt},
	}})
}

func Prompt(chatMessages []ChatMessage) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(chatMessages).
		Post(fmt.Sprintf("%s/perplexity", config.ApiLangchainBaseUrl))
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		err := tracerr.Errorf("unexpected status code received from langchain: %d", resp.StatusCode())
		Logger.Error("error prompting for response", zap.Error(err))
		return "", err
	}
	return resp.String(), nil
}
func GetMessageHistory(ctx context.Context, threadId pulid.ID) ([]*ent.Message, error) {
	return EntClient.Message.Query().Where(
		message.HasThreadWith(thread.ID(threadId)),
	).Order(message.ByCreatedAt()).All(ctx)
}

func ChatMessagesFromMessages(messages []*ent.Message) []ChatMessage {
	chatMessages := make([]ChatMessage, 0)
	for _, message := range messages {
		chatMessages = append(chatMessages, chatMessageFromMessage(message))
	}

	return chatMessages
}

type ChatMessage struct {
	Type string          `json:"type,omitempty"`
	Data ChatMessageData `json:"data,omitempty"`
}
type ChatMessageData struct {
	Content string `json:"content"`
}

func chatMessageFromMessage(messag *ent.Message) ChatMessage {
	return ChatMessage{Type: messag.MessageType.String(), Data: ChatMessageData{Content: messag.Content}}
}
