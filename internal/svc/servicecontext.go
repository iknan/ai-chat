package svc

import (
	"ai_chat/internal/config"
	"ai_chat/internal/infra"
	"ai_chat/internal/infra/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	infra.InitInfra(&c)
	return &ServiceContext{
		Config: c,
		DB:     mysql.DB,
	}
}
