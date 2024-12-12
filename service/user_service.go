package service

import (
	"errors"
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
		utils.LogError("[UpdateMessage] Get Messages failed", err)
		return nil, err
	}
	if len(messages) == 0 {
		return nil, errors.New("message not found")
	}

	checker, err := u.store.GetUser(request.UserID)
	if err != nil || checker.Role != models.CHECKER {
		utils.LogError("[UpdateMessage] User role is not correct", err)
		return nil, errors.New("user is not checker cannot update the message")
	}

	if messages[0].Status != "Pending" {
		utils.LogError("[UpdateMessage] Message status is not pending", err)
		return nil, errors.New("message status is not pending")
	}

	message := messages[0]
	message.Status = request.Status

	if request.Status == models.APPROVE {
		message.ApprovedBy = checker.Username
		if u.conf.Email.IsEnabled {
			go func() {
				if err := utils.SendEmail(u.conf.Email, message.Sender, message.Message); err != nil {
					utils.LogError("[UpdateMessage] Send Email failed", err)
				}
			}()
		}
	} else if request.Status == models.REJECT {
		message.RejectedBy = checker.Username
	}

	if err := u.store.UpdateMessage(*message); err != nil {
		utils.LogError("[UpdateMessage] Update Message failed", err)
		return nil, errors.New("failed to update message")
	}
	return message, nil
}

// CreateMessage Create Message Business logic and insertion in the database
func (u *userService) CreateMessage(request *models.CreateMessageRequest) (*models.Message, error) {

	sender, err := u.store.GetUser(request.UserID)
	if err != nil || sender.Role != models.MAKER {
		utils.LogError("[CreateMessage] User role is not correct", err)
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
		utils.LogError("[CreateMessage] update Message failed", err)
		return nil, errors.New("failed to create message")
	}

	return &message, err
}
