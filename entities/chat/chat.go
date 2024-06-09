package chat

import (
	"capstone/entities"
	"capstone/entities/consultation"
)

type Chat struct {
	ID           uint
	Consultation consultation.Consultation
	Status       string
	Isrejected   bool
	LatestMessageID uint
	LatestMessageContent string
	LatestMessageTime string
}

type ChatMessage struct {
	ID uint
	ChatID uint
	Message string
	Role string
	CreatedAt string
}

type RepositoryInterface interface {
	GetAllChatByUserId(userId int) ([]Chat, error)
	SendMessage(chatMessage ChatMessage) (ChatMessage, error)
	GetAllMessages(chatId int, lastMessageId int, metadata entities.Metadata) ([]ChatMessage, error)
}

type UseCaseInterface interface {
	GetAllChatByUserId(userId int) ([]Chat, error)
	SendMessage(chatMessage ChatMessage) (ChatMessage, error)
	GetAllMessages(chatId int, lastMessageId string, metadata entities.Metadata) ([]ChatMessage, error)
}