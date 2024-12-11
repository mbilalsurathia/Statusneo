package repository

import (
	"maker-checker/models"
)

type Store interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(userId string) (*models.User, error)
	GetUserByUserName(userId string) (*models.User, error)
	GetMessages(messageId uint64) ([]*models.Message, error)
	CreateMessage(message *models.Message) error
	UpdateMessage(message models.Message) error
}
