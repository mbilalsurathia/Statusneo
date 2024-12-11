package service

import (
	"maker-checker/models"
	"maker-checker/repository"
)

type SeedService interface {
	Seed(env string)
}

type seedService struct {
	store repository.Store
}

func NewSeedService(
	store repository.Store,

) SeedService {
	return &seedService{
		store: store,
	}
}

func (s *seedService) Seed(env string) {
	s.insertUsers()
}

func (s *seedService) insertUsers() {
	s.AddUser("user1", "Frank", models.MAKER)
	s.AddUser("user2", "Bob", models.CHECKER)

}

func (s *seedService) AddUser(userID, username, role string) {
	userDb, _ := s.store.GetUser(userID)

	if userDb == nil || userDb.Id == 0 {
		user := models.User{
			UserID:   userID,
			Username: username,
			Role:     role,
		}
		_, _ = s.store.CreateUser(&user)
	}
}
