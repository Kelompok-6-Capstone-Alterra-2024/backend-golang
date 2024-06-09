package chat

import (
	"capstone/entities"
	chatEntities "capstone/entities/chat"
	"capstone/entities/doctor"
	"capstone/repositories/mysql/consultation"

	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

func (c *ChatRepo) GetAllChatByUserId(userId int) ([]chatEntities.Chat, error) {
	var consultationDB []consultation.Consultation

	if err := c.db.Where("user_id = ?", userId).Find(&consultationDB).Error; err != nil {
		return nil, err
	}

	var consultationIds []int
	for _, consultation := range consultationDB {
		consultationIds = append(consultationIds, int(consultation.ID))
	}

	var chatDB []Chat

	if err := c.db.Preload("Consultation").Preload("Consultation.Doctor").Where("consultation_id IN ?", consultationIds).Find(&chatDB).Error; err != nil {
		return nil, err
	}

	var mesageTemp []ChatMessage
	var latestMessage []ChatMessage

	for _, chat := range chatDB {
		err := c.db.Where("chat_id = ?", chat.ID).Order("created_at DESC").Limit(1).Find(&mesageTemp).Error
		if err != nil {
			return nil, err
		}
		latestMessage = append(latestMessage, mesageTemp...)
	}

	chatEnts := make([]chatEntities.Chat, len(chatDB))

	for i, chat := range chatDB {
		chatEnts[i].ID = chat.ID
		chatEnts[i].LatestMessageID = latestMessage[i].ID
		chatEnts[i].LatestMessageContent = latestMessage[i].Message
		chatEnts[i].LatestMessageTime = latestMessage[i].CreatedAt.Format("2006-01-02 15:04:05")
		chatEnts[i].Consultation.Doctor = &doctor.Doctor{
			ID:         chat.Consultation.Doctor.ID,
			Name:       chat.Consultation.Doctor.Name,
			Username:   chat.Consultation.Doctor.Username,
			ProfilePicture: chat.Consultation.Doctor.ProfilePicture,
			Specialist: chat.Consultation.Doctor.Specialist,
		}
		if chat.Consultation.Status == "pending" {
			chatEnts[i].Status = "process"
			chatEnts[i].Isrejected = false
		} else if chat.Consultation.Status == "accepted" && chat.Consultation.IsActive {
			chatEnts[i].Status = "active"
			chatEnts[i].Isrejected = false
		} else if chat.Consultation.Status == "accepted" && !chat.Consultation.IsActive {
			chatEnts[i].Status = "waiting"
			chatEnts[i].Isrejected = false
		} else if chat.Consultation.Status == "rejected" {
			chatEnts[i].Status = "complete"
			chatEnts[i].Isrejected = true
		}
	}

	return chatEnts, nil
}

func (c *ChatRepo) SendMessage(chatMessage chatEntities.ChatMessage) (chatEntities.ChatMessage, error) {
	var chatMessageDB ChatMessage
	chatMessageDB.ID = chatMessage.ID
	chatMessageDB.ChatId = chatMessage.ChatID
	chatMessageDB.Message = chatMessage.Message
	chatMessageDB.Role = chatMessage.Role

	if err := c.db.Create(&chatMessageDB).Error; err != nil {
		return chatEntities.ChatMessage{}, err
	}

	return chatEntities.ChatMessage{
		ID: chatMessageDB.ID,
		ChatID: chatMessageDB.ChatId,
		Message: chatMessageDB.Message,
		Role: chatMessageDB.Role,
		CreatedAt: chatMessageDB.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (c *ChatRepo) GetAllMessages(chatId int, lastMessageId int, metadata entities.Metadata) ([]chatEntities.ChatMessage, error) {
	var chatMessageDB []ChatMessage

	if lastMessageId == 0 {
		if err := c.db.Where("chat_id = ?", chatId).Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Find(&chatMessageDB).Error; err != nil {
			return nil, err
		}
	} else {
		if err := c.db.Where("chat_id = ?", chatId).Where("id > ?", lastMessageId).Limit(metadata.Limit).Offset((metadata.Page - 1) * metadata.Limit).Find(&chatMessageDB).Error; err != nil {
			return nil, err
		}
	}

	chatMessageEnts := make([]chatEntities.ChatMessage, len(chatMessageDB))

	for i, chatMessage := range chatMessageDB {
		chatMessageEnts[i].ID = chatMessage.ID
		chatMessageEnts[i].ChatID = chatMessage.ChatId
		chatMessageEnts[i].Message = chatMessage.Message
		chatMessageEnts[i].Role = chatMessage.Role
		chatMessageEnts[i].CreatedAt = chatMessage.CreatedAt.Format("2006-01-02 15:04:05")
	}

	return chatMessageEnts, nil
}