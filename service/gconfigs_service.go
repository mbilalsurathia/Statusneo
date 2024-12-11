package service

import (
	"maker-checker/conf"
)

type GbeConfigService interface {
	GetConfig() *conf.GbeConfig
}

type gGbeConfigService struct {
}

func NewGbeConfigService() GbeConfigService {
	return &gGbeConfigService{}
}

func (c *gGbeConfigService) GetConfig() *conf.GbeConfig {

	return conf.GetConfig()

}
