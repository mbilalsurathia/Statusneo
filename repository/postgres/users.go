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
