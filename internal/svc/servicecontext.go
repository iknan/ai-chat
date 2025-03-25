package svc

import (
	"ai_chat/internal/config"
	"ai_chat/internal/infra"
	"ai_chat/internal/infra/mysql"
	redisInfra "ai_chat/internal/infra/redis"
	"ai_chat/internal/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	AuthToken rest.Middleware
	DB        *gorm.DB
	Redis     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	infra.InitInfra(&c)
	return &ServiceContext{
		Config:    c,
		AuthToken: middleware.NewAuthTokenMiddleware(&c, redisInfra.Rdb).Handle,
		DB:        mysql.DB,
		Redis:     redisInfra.Rdb,
	}
}
