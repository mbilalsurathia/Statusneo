package postgres

import (
	"maker-checker/models"
	"time"
)

func (s *Store) CreateUser(user *models.User) (*models.User, error) {
	if err := s.db.Model(&models.User{}).Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUser(userId string) (*models.User, error) {
	var user *models.User
	if err := s.db.Model(&models.User{}).Where("user_id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByUserName(username string) (*models.User, error) {
	var user *models.User
	if err := s.db.Model(&models.User{}).Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

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
