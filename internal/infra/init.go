package infra

import (
	"ai_chat/internal/config"
	"ai_chat/internal/infra/mysql"
	"ai_chat/internal/infra/redis"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func InitInfra(c *config.Config) {
	mysql.InitMysql(c)
	redis.InitRedis(c)
	initLog(c)
}

func initLog(c *config.Config) {
	var LogCfg logx.LogConf
	_ = conf.FillDefault(&LogCfg)
	LogCfg.Mode = c.LogConf.Mode
	LogCfg.Level = c.LogConf.Level
	LogCfg.Path = c.LogConf.Path
	logc.MustSetup(LogCfg)
	logx.MustSetup(LogCfg)
	logx.AddWriter(logx.NewWriter(os.Stdout))
}
