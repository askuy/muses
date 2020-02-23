package redis

import (
	"github.com/BurntSushi/toml"
	"github.com/i2eco/muses/pkg/cache/redis"
	"github.com/i2eco/muses/pkg/common"
	"github.com/i2eco/muses/pkg/logger"
	"github.com/i2eco/muses/pkg/token"
)

var defaultCallerStore = &callerStore{}

type callerStore struct {
	Name string
	cfg Cfg
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	clientName := c.cfg.Muses.Token.JWT.client
	loggerName := c.cfg.Muses.Token.JWT.logger
	client := redis.Caller(clientName)
	log := logger.Caller(loggerName)

	// 初始化token里面的Accessor
	token.Accessor = initTokenAccessor(log, client)
	return nil
}

func Register() common.Caller{
	return defaultCallerStore
}
