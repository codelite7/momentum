package common

import (
	"context"
	"fmt"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/message"
	"github.com/codelite7/momentum/api/ent/thread"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

func GetThreadName(messageContent string) (string, error) {
	prompt := fmt.Sprintf("Provide a concise 3 word summary of the following text:%s", messageContent)
	return Prompt([]*ChatMessage{&ChatMessage{
		Type: "human",
		Data: ChatMessageData{Content: prompt},
	}})
}

func Prompt(chatMessages []*ChatMessage) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(chatMessages).
		Post("http://localhost:6543/perplexity")
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
func GetMessageHistory(ctx context.Context, tx *ent.Tx, threadId uuid.UUID, lastMessageCreatedAt time.Time) ([]*ent.Message, error) {
	return tx.Message.Query().Where(
		message.HasThreadWith(thread.ID(threadId)),
		message.CreatedAtLTE(lastMessageCreatedAt),
	).All(ctx)
}

func ChatMessagesFromMessages(messages []*ent.Message) ([]*ChatMessage, error) {
	chatMessages := make([]*ChatMessage, 0)
	for _, message := range messages {
		humanChatMessage, err := chatMessageFromMessage(message)
		if err != nil {
			return nil, err
		}
		chatMessages = append(chatMessages, humanChatMessage)
		response, err := message.Response(context.Background())
		if err != nil {
			return nil, err
		}
		if response != nil && response.Content != "" {
			aiMessage, err := chatMessageFromResponse(response)
			if err != nil {
				return nil, err
			}
			chatMessages = append(chatMessages, aiMessage)
		}
	}
	return chatMessages, nil
}

type ChatMessage struct {
	Type string          `json:"type,omitempty"`
	Data ChatMessageData `json:"data,omitempty"`
}
type ChatMessageData struct {
	Content string `json:"content"`
}

func chatMessageFromMessage(message *ent.Message) (*ChatMessage, error) {
	chatMessage := &ChatMessage{Data: ChatMessageData{Content: message.Content}}
	chatMessage.Type = "human"

	return chatMessage, nil
}
func chatMessageFromResponse(response *ent.Response) (*ChatMessage, error) {
	chatMessage := &ChatMessage{Data: ChatMessageData{Content: response.Content}}
	chatMessage.Type = "ai"

	return chatMessage, nil
}
