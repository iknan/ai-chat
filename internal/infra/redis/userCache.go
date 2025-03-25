package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	ctx context.Context
	rdb *redis.Client
}

func NewUserCache(ctx context.Context, rdb *redis.Client) *UserCache {
	return &UserCache{
		ctx: context.Background(),
		rdb: rdb,
	}
}

func (u *UserCache) SetTokenCache(token string) (err error) {
	return u.rdb.Set(u.ctx, fmt.Sprintf(KeyUserToken, token), 1, TokenExpireTime).Err()
}

func (u *UserCache) ExistTokenCache(token string) (exist int64, err error) {
	return u.rdb.Exists(u.ctx, fmt.Sprintf(KeyUserToken, token)).Result()
}
