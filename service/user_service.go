package service

import (
	"errors"
	"go.uber.org/zap"
	"maker-checker/conf"
	"maker-checker/models"
	"maker-checker/repository"
	"maker-checker/utils"
	"time"
)

type UserService interface {
	GetMessages(messageId uint64) ([]*models.Message, error)
	UpdateMessage(request *models.UpdateMessageRequest) (*models.Message, error)
	CreateMessage(request *models.CreateMessageRequest) (*models.Message, error)
}

type userService struct {
	store repository.Store
	conf  *conf.GbeConfig
}

func NewUserService(
	store repository.Store,
	conf *conf.GbeConfig,
) UserService {
	return &userService{
		store: store,
		conf:  conf,
	}
}

// GetMessages Gets Message from database
func (u *userService) GetMessages(messageId uint64) ([]*models.Message, error) {
	return u.store.GetMessages(messageId)
}

// UpdateMessage Update Message Business logic and Update in the database
func (u *userService) UpdateMessage(request *models.UpdateMessageRequest) (*models.Message, error) {

	messages, err := u.store.GetMessages(request.RequestID)
	if err != nil {
		zap.L().Error("[UpdateMessage] Get Messages failed", zap.Error(err))
		return nil, err
	}
	if len(messages) == 0 {
		return nil, errors.New("message not found")
	}

	// Ensure the checker has the correct role
	checker, err := u.store.GetUser(request.UserID)
	if err != nil || checker.Role != models.CHECKER {
		zap.L().Error("[UpdateMessage] User role is not correct", zap.Error(err))
		return nil, errors.New("user is not checker cannot update the message")
	}

	// If already processed, return an error
	if messages[0].Status != "Pending" {
		zap.L().Error("[UpdateMessage] User role is not correct", zap.Error(err))
		return nil, errors.New("message status is not pending")
	}
	message := messages[0]
	message.Status = request.Status

	// approve case to send email only to notify recipient
	if request.Status == models.APPROVE {
		message.ApprovedBy = checker.Username

		//just need to send email to notify recipient by default it is false
		if u.conf.Email.IsEnabled {
			err := utils.SendEmail(u.conf.Email, message.Sender, message.Message)
			if err != nil {
				zap.L().Error("[UpdateMessage] Send Email failed", zap.Error(err))
				return nil, errors.New("failed to send email")
			}
		}

	} else if request.Status == models.REJECT {
		message.RejectedBy = checker.Username
	}

	//update message in the database
	err = u.store.UpdateMessage(*message)
	if err != nil {
		zap.L().Error("[UpdateMessage] update Message failed", zap.Error(err))
		return nil, errors.New("failed to update message")
	}
	return message, nil
}

// CreateMessage Create Message Business logic and insertion in the database
func (u *userService) CreateMessage(request *models.CreateMessageRequest) (*models.Message, error) {

	sender, err := u.store.GetUser(request.UserID)
	if err != nil || sender.Role != models.MAKER {
		zap.L().Error("[CreateMessage] User role is not correct", zap.Error(err))
		return nil, errors.New("user is not maker cannot create the message")
	}

	// Create the message request
	message := models.Message{
		Sender:    sender.Username,
		Recipient: request.Recipient,
		Message:   request.Message,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	err = u.store.CreateMessage(&message)
	if err != nil {
		zap.L().Error("[CreateMessage] update Message failed", zap.Error(err))
		return nil, errors.New("failed to create message")
	}

	return &message, err
}
