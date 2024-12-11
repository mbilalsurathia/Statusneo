package service

import (
	"maker-checker/conf"

	"maker-checker/repository"
	"maker-checker/repository/postgres"
)

type Container struct {
	GbeConfigService GbeConfigService
	LoggerService    LoggerService
	Store            repository.Store
	UserService      UserService
}

// NewServiceContainer Services initialize
func NewServiceContainer() *Container {
	gbeConfig := conf.GetConfig()
	gbeConfigService := NewGbeConfigService()

	postgresDB := postgres.SharedStore()

	//to add Users as a Maker and checker
	seedService := NewSeedService(postgresDB)
	loggerService := NewLoggerService(gbeConfig)
	userService := NewUserService(postgresDB, gbeConfig)

	// seed database for adding User
	seedService.Seed(gbeConfig.Env)

	return &Container{
		GbeConfigService: gbeConfigService,
		LoggerService:    loggerService,
		Store:            postgresDB,
		UserService:      userService,
	}
}
