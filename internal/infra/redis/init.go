package redis

import (
	"ai_chat/internal/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

var once sync.Once
var Rdb *redis.Client

func InitRedis(c *config.Config) {
	wg := &sync.WaitGroup{}
	if Rdb == nil {
		wg.Add(1)
		once.Do(func() {
			Rdb = redis.NewClient(&redis.Options{
				Addr:     c.Redis.Host,
				Password: c.Redis.Pass,
				DB:       c.Redis.Db,
			})

			_, err := Rdb.Ping(context.Background()).Result()
			if err != nil {
				panic(fmt.Sprintf("failed to connect redis,err : %v\n", err))
			}

			fmt.Printf("success connect redis,url : %s , DB : [%d]\n", c.Redis.Host, c.Redis.Db)
		})
		wg.Done()
		wg.Wait()
	}

	return
}
