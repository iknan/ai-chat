package mysql

import (
	"ai_chat/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

var once sync.Once
var DB *gorm.DB

func InitMysql(c *config.Config) {
	var err error
	wg := &sync.WaitGroup{}
	if DB == nil {
		wg.Add(1)
		once.Do(func() {
			dsn := c.Mysql.User + ":" + c.Mysql.Pass + "@tcp(" + c.Mysql.Host + ":" + strconv.Itoa(c.Mysql.Port) + ")/" + c.Mysql.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
			DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(fmt.Sprintf("failed to connect mysql,err : %v\n", err))
			}
		})
		wg.Done()
		wg.Wait()
	}

	return
}
