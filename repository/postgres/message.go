package postgres

import (
	"maker-checker/models"
	"time"
)

func (s *Store) GetMessages(message uint64) ([]*models.Message, error) {
	var messages []*models.Message
	db := s.db.Model(&models.Message{})

	//if messageId in query parameters it will be return one message otherwise all the messages
	if message != 0 {
		db = db.Where("message_id = ?", message)
	}
	if err := db.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *Store) CreateMessage(message *models.Message) error {
	if err := s.db.Model(&models.Message{}).Create(&message).Error; err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateMessage(message models.Message) error {
	return s.db.Model(&models.Message{}).Where("message_id = ?", message.MessageId).Updates(map[string]interface{}{
		"status":      message.Status,
		"updated_at":  time.Now(),
		"approved_by": message.ApprovedBy,
		"rejected_by": message.RejectedBy,
	}).Error
}
